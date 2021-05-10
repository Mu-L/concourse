package runtime_test

import (
	"context"
	"errors"
	"reflect"
	"strings"

	"github.com/concourse/concourse/worker/runtime"
	"github.com/concourse/concourse/worker/runtime/iptables/iptablesfakes"
	"github.com/concourse/concourse/worker/runtime/libcontainerd/libcontainerdfakes"
	"github.com/concourse/concourse/worker/runtime/runtimefakes"
	"github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type CNINetworkSuite struct {
	suite.Suite
	*require.Assertions

	network  runtime.Network
	cni      *runtimefakes.FakeCNI
	store    *runtimefakes.FakeFileStore
	iptables *iptablesfakes.FakeIptables
}

func (s *CNINetworkSuite) SetupTest() {
	var err error

	s.store = new(runtimefakes.FakeFileStore)
	s.cni = new(runtimefakes.FakeCNI)
	s.iptables = new(iptablesfakes.FakeIptables)

	s.network, err = runtime.NewCNINetwork(
		runtime.WithCNIFileStore(s.store),
		runtime.WithCNIClient(s.cni),
		runtime.WithIptables(s.iptables),
	)
	s.NoError(err)
}

func (s *CNINetworkSuite) TestNewCNINetworkWithInvalidConfigDoesntFail() {
	// CNI defers the actual interpretation of the network configuration to
	// the plugins.
	//
	_, err := runtime.NewCNINetwork(
		runtime.WithCNINetworkConfig(runtime.CNINetworkConfig{
			Subnet: "_____________",
		}),
		runtime.WithIptables(s.iptables),
	)
	s.NoError(err)
}

func (s *CNINetworkSuite) TestSetupMountsEmptyHandle() {
	_, err := s.network.SetupMounts("")
	s.EqualError(err, "empty handle")
}

func (s *CNINetworkSuite) TestSetupMountsFailToCreateHosts() {
	s.store.CreateReturnsOnCall(0, "", errors.New("create-hosts-err"))

	_, err := s.network.SetupMounts("handle")
	s.EqualError(errors.Unwrap(err), "create-hosts-err")

	s.Equal(1, s.store.CreateCallCount())
	fname, _ := s.store.CreateArgsForCall(0)

	s.Equal("handle/hosts", fname)
}

func (s *CNINetworkSuite) TestSetupMountsFailToCreateResolvConf() {
	s.store.CreateReturnsOnCall(1, "", errors.New("create-resolvconf-err"))

	_, err := s.network.SetupMounts("handle")
	s.EqualError(errors.Unwrap(err), "create-resolvconf-err")

	s.Equal(2, s.store.CreateCallCount())
	fname, _ := s.store.CreateArgsForCall(1)

	s.Equal("handle/resolv.conf", fname)
}

func (s *CNINetworkSuite) TestSetupMountsReturnsMountpoints() {
	s.store.CreateReturnsOnCall(0, "/tmp/handle/etc/hosts", nil)
	s.store.CreateReturnsOnCall(1, "/tmp/handle/etc/resolv.conf", nil)

	mounts, err := s.network.SetupMounts("some-handle")
	s.NoError(err)

	s.Len(mounts, 2)
	s.Equal(mounts, []specs.Mount{
		{
			Destination: "/etc/hosts",
			Type:        "bind",
			Source:      "/tmp/handle/etc/hosts",
			Options:     []string{"bind", "rw"},
		},
		{
			Destination: "/etc/resolv.conf",
			Type:        "bind",
			Source:      "/tmp/handle/etc/resolv.conf",
			Options:     []string{"bind", "rw"},
		},
	})
}

func (s *CNINetworkSuite) TestSetupMountsCallsStoreWithNameServers() {
	network, err := runtime.NewCNINetwork(
		runtime.WithCNIFileStore(s.store),
		runtime.WithNameServers([]string{"6.6.7.7", "1.2.3.4"}),
		runtime.WithIptables(s.iptables),
	)
	s.NoError(err)

	_, err = network.SetupMounts("some-handle")
	s.NoError(err)

	_, resolvConfContents := s.store.CreateArgsForCall(1)
	s.Equal(resolvConfContents, []byte("nameserver 6.6.7.7\nnameserver 1.2.3.4\n"))
}

func (s *CNINetworkSuite) TestSetupMountsCallsStoreWithoutNameServers() {
	network, err := runtime.NewCNINetwork(
		runtime.WithCNIFileStore(s.store),
		runtime.WithIptables(s.iptables),
	)
	s.NoError(err)

	_, err = network.SetupMounts("some-handle")
	s.NoError(err)

	actualResolvContents, err := runtime.ParseHostResolveConf("/etc/resolv.conf")
	s.NoError(err)

	contents := strings.Join(actualResolvContents, "\n") + "\n"

	_, resolvConfContents := s.store.CreateArgsForCall(1)
	s.Equal(resolvConfContents, []byte(contents))
}

func (s *CNINetworkSuite) TestSetupHostNetwork() {
	hostIp, err := runtime.GetHostIp()
	s.NoError(err)

	ts := map[string]struct {
		cniNetworkSetup   func() (runtime.Network, error)
		expectedTableName string
		expectedChainName string
		expectedRuleSpec  []string
	}{
		"flushes the CONCOURSE-OPERATOR chain": {
			cniNetworkSetup: func() (runtime.Network, error) {
				return runtime.NewCNINetwork(
					runtime.WithIptables(s.iptables),
				)
			},
			expectedTableName: "filter",
			expectedChainName: "CONCOURSE-OPERATOR",
		},
		"adds rule to CONCOURSE-OPERATOR chain for accepting established connections": {
			cniNetworkSetup: func() (runtime.Network, error) {
				return runtime.NewCNINetwork(
					runtime.WithIptables(s.iptables),
				)
			},
			expectedTableName: "filter",
			expectedChainName: "CONCOURSE-OPERATOR",
			expectedRuleSpec:  []string{"-m", "conntrack", "--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT"},
		},
		"adds rule to CONCOURSE-OPERATOR chain to reject IP 1.1.1.1": {
			cniNetworkSetup: func() (runtime.Network, error) {
				return runtime.NewCNINetwork(
					runtime.WithRestrictedNetworks([]string{"1.1.1.1", "8.8.8.8"}),
					runtime.WithIptables(s.iptables),
				)
			},
			expectedTableName: "filter",
			expectedChainName: "CONCOURSE-OPERATOR",
			expectedRuleSpec:  []string{"-d", "1.1.1.1", "-j", "REJECT"},
		},
		"adds rule to CONCOURSE-OPERATOR chain to reject IP 8.8.8.8": {
			cniNetworkSetup: func() (runtime.Network, error) {
				return runtime.NewCNINetwork(
					runtime.WithRestrictedNetworks([]string{"1.1.1.1", "8.8.8.8"}),
					runtime.WithIptables(s.iptables),
				)
			},
			expectedTableName: "filter",
			expectedChainName: "CONCOURSE-OPERATOR",
			expectedRuleSpec:  []string{"-d", "8.8.8.8", "-j", "REJECT"},
		},
		"flushes the INPUT chain": {
			cniNetworkSetup: func() (runtime.Network, error) {
				return runtime.NewCNINetwork(
					runtime.WithIptables(s.iptables),
				)
			},
			expectedTableName: "filter",
			expectedChainName: "INPUT",
		},
		"adds rule to INPUT chain to block host access by default": {
			cniNetworkSetup: func() (runtime.Network, error) {
				return runtime.NewCNINetwork(
					runtime.WithIptables(s.iptables),
				)
			},
			expectedTableName: "filter",
			expectedChainName: "INPUT",
			expectedRuleSpec:  []string{"-i", "concourse0", "-d", hostIp, "-j", "DROP"},
		},
	}

	for desc, c := range ts {
		network, err := c.cniNetworkSetup()
		s.NoError(err)
		err = network.SetupHostNetwork()
		s.NoError(err)

		foundExpected := false
		if c.expectedRuleSpec == nil {
			numOfCalls := s.iptables.CreateChainOrFlushIfExistsCallCount()
			for i := 0; i < numOfCalls; i++ {
				tablename, chainName := s.iptables.CreateChainOrFlushIfExistsArgsForCall(i)
				if tablename == c.expectedTableName && chainName == c.expectedChainName {
					foundExpected = true
					break
				}
			}
		} else {
			numOfCalls := s.iptables.AppendRuleCallCount()
			for i := 0; i < numOfCalls; i++ {
				tablename, chainName, rulespec := s.iptables.AppendRuleArgsForCall(i)
				if tablename == c.expectedTableName && chainName == c.expectedChainName && reflect.DeepEqual(rulespec, c.expectedRuleSpec) {
					foundExpected = true
					break
				}
			}

		}

		s.Equal(foundExpected, true, desc)
	}
}

func (s *CNINetworkSuite) TestAddNilTask() {
	err := s.network.Add(context.Background(), nil)
	s.EqualError(err, "nil task")
}

func (s *CNINetworkSuite) TestAddSetupErrors() {
	s.cni.SetupReturns(nil, errors.New("setup-err"))
	task := new(libcontainerdfakes.FakeTask)

	err := s.network.Add(context.Background(), task)
	s.EqualError(errors.Unwrap(err), "setup-err")
}

func (s *CNINetworkSuite) TestAdd() {
	task := new(libcontainerdfakes.FakeTask)
	task.PidReturns(123)
	task.IDReturns("id")

	err := s.network.Add(context.Background(), task)
	s.NoError(err)

	s.Equal(1, s.cni.SetupCallCount())
	_, id, netns, _ := s.cni.SetupArgsForCall(0)
	s.Equal("id", id)
	s.Equal("/proc/123/ns/net", netns)
}

func (s *CNINetworkSuite) TestRemoveNilTask() {
	err := s.network.Remove(context.Background(), nil)
	s.EqualError(err, "nil task")
}

func (s *CNINetworkSuite) TestRemoveSetupErrors() {
	s.cni.RemoveReturns(errors.New("remove-err"))
	task := new(libcontainerdfakes.FakeTask)

	err := s.network.Remove(context.Background(), task)
	s.EqualError(errors.Unwrap(err), "remove-err")
}

func (s *CNINetworkSuite) TestRemove() {
	task := new(libcontainerdfakes.FakeTask)
	task.PidReturns(123)
	task.IDReturns("id")

	err := s.network.Remove(context.Background(), task)
	s.NoError(err)

	s.Equal(1, s.cni.RemoveCallCount())
	_, id, netns, _ := s.cni.RemoveArgsForCall(0)
	s.Equal("id", id)
	s.Equal("/proc/123/ns/net", netns)
}
