package skycmd

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/go-viper/mapstructure/v2"
	flags "github.com/jessevdk/go-flags"
	"sigs.k8s.io/yaml"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/flag/v2"
)

var connectors []*Connector

func RegisterConnector(connector *Connector) {
	connectors = append(connectors, connector)
}

func GetConnectors() []*Connector {
	return connectors
}

func WireConnectors(group *flags.Group) {
	for _, c := range connectors {
		description := fmt.Sprintf("%s (%s)", group.ShortDescription, c.Name())
		connGroup, _ := group.AddGroup(description, description, c.config)
		connGroup.Namespace = c.ID()
	}
}

func WireTeamConnectors(group *flags.Group) {
	for _, c := range connectors {
		description := fmt.Sprintf("%s (%s)", group.ShortDescription, c.Name())
		connTeamGroup, _ := group.AddGroup(description, description, c.teamConfig)
		connTeamGroup.Namespace = c.ID()
	}
}

type AuthFlags struct {
	SecureCookies     bool              `long:"cookie-secure" description:"Force sending secure flag on http cookies"`
	Expiration        time.Duration     `long:"auth-duration" default:"24h" description:"Length of time for which tokens are valid. Afterwards, users will have to log back in."`
	SigningKey        *flag.PrivateKey  `long:"session-signing-key" required:"true" description:"File containing an RSA private key, used to sign auth tokens."`
	PasswordConnector string            `long:"password-connector" default:"local" choice:"local" choice:"ldap" description:"Connector to use when authenticating via 'fly login -u ... -p ...'"`
	LocalUsers        map[string]string `long:"add-local-user" description:"List of username:password combinations for all your local users. The password can be bcrypted - if so, it must have a minimum cost of 10." value-name:"USERNAME:PASSWORD"`
	Clients           map[string]string `long:"add-client" description:"List of client_id:client_secret combinations" value-name:"CLIENT_ID:CLIENT_SECRET"`
}

type AuthTeamFlags struct {
	LocalUsers []string  `long:"local-user" description:"A whitelisted local concourse user. These are the users you've added at web startup with the --add-local-user flag." value-name:"USERNAME"`
	Config     flag.File `short:"c" long:"config" description:"Configuration file for specifying team params"`
}

func (flag *AuthTeamFlags) Format() (atc.TeamAuth, error) {

	if path := flag.Config.Path(); path != "" {
		return flag.formatFromFile()

	}
	return flag.formatFromFlags()

}

// When formatting from a configuration file we iterate over each connector
// type and create a new instance of the TeamConfig object for each connector.
// These connectors all have their own unique configuration so we need to use
// mapstructure to decode the generic result into a known struct.

// e.g.
// The github connector has configuration for: users, teams, orgs
// The cloudfoundry conncetor has configuration for: users, orgs, spaces

func (flag *AuthTeamFlags) formatFromFile() (atc.TeamAuth, error) {

	content, err := os.ReadFile(flag.Config.Path())
	if err != nil {
		return nil, err
	}

	var data struct {
		Roles []map[string]any `json:"roles"`
	}
	if err = yaml.Unmarshal(content, &data); err != nil {
		return nil, err
	}

	auth := atc.TeamAuth{}

	for _, role := range data.Roles {
		roleName := role["name"].(string)

		users := []string{}
		groups := []string{}

		for _, connector := range connectors {
			config, ok := role[connector.ID()]
			if !ok {
				continue
			}

			teamConfig, err := connector.newTeamConfig()
			if err != nil {
				return nil, err
			}

			err = mapstructure.Decode(config, &teamConfig)
			if err != nil {
				return nil, err
			}

			for _, user := range teamConfig.GetUsers() {
				if user != "" {
					users = append(users, connector.ID()+":"+strings.ToLower(user))
				}
			}

			for _, group := range teamConfig.GetGroups() {
				if group != "" {
					groups = append(groups, connector.ID()+":"+strings.ToLower(group))
				}
			}
		}

		if conf, ok := role["local"].(map[string]any); ok {
			for _, user := range conf["users"].([]any) {
				if user != "" {
					users = append(users, "local:"+strings.ToLower(user.(string)))
				}
			}
		}

		if len(users) == 0 && len(groups) == 0 {
			continue
		}

		auth[roleName] = map[string][]string{
			"users":  users,
			"groups": groups,
		}
	}

	if err := auth.Validate(); err != nil {
		return nil, err
	}

	return auth, nil
}

// When formatting team config from the command line flags, the connector's
// TeamConfig has already been populated by the flags library. All we need to
// do is grab the teamConfig object and extract the users and groups.

func (flag *AuthTeamFlags) formatFromFlags() (atc.TeamAuth, error) {

	users := []string{}
	groups := []string{}

	for _, connector := range connectors {

		teamConfig := connector.teamConfig

		for _, user := range teamConfig.GetUsers() {
			if user != "" {
				users = append(users, connector.ID()+":"+strings.ToLower(user))
			}
		}

		for _, group := range teamConfig.GetGroups() {
			if group != "" {
				groups = append(groups, connector.ID()+":"+strings.ToLower(group))
			}
		}
	}

	for _, user := range flag.LocalUsers {
		if user != "" {
			users = append(users, "local:"+strings.ToLower(user))
		}
	}

	if len(users) == 0 && len(groups) == 0 {
		return nil, atc.ErrAuthConfigInvalid
	}

	return atc.TeamAuth{
		"owner": map[string][]string{
			"users":  users,
			"groups": groups,
		},
	}, nil
}

type Config interface {
	Name() string
	Serialize(redirectURI string) ([]byte, error)
}

type TeamConfig interface {
	GetUsers() []string
	GetGroups() []string
}

type Connector struct {
	id         string
	config     Config
	teamConfig TeamConfig
}

func (con *Connector) ID() string {
	return con.id
}

func (con *Connector) Name() string {
	return con.config.Name()
}

func (con *Connector) Serialize(redirectURI string) ([]byte, error) {
	return con.config.Serialize(redirectURI)
}

func (con *Connector) newTeamConfig() (TeamConfig, error) {

	typeof := reflect.TypeOf(con.teamConfig)
	if typeof.Kind() == reflect.Ptr {
		typeof = typeof.Elem()
	}

	valueof := reflect.ValueOf(con.teamConfig)
	if valueof.Kind() == reflect.Ptr {
		valueof = valueof.Elem()
	}

	instance := reflect.New(typeof).Interface()
	res, ok := instance.(TeamConfig)
	if !ok {
		return nil, errors.New("Invalid TeamConfig")
	}

	return res, nil
}

type skyDisplayUserIdGenerator struct {
	mapConnectorUserid map[string]string
}

func NewSkyDisplayUserIdGenerator(config map[string]string) (atc.DisplayUserIdGenerator, error) {
	generator := &skyDisplayUserIdGenerator{
		mapConnectorUserid: map[string]string{},
	}

	for connectorId, fieldName := range config {
		valid := false
		if connectorId == "local" {
			valid = true
		} else {
			for _, connector := range GetConnectors() {
				if connectorId == connector.ID() {
					valid = true
					break
				}
			}
		}
		if !valid {
			return nil, fmt.Errorf("invalid connector: %s", connectorId)
		}

		switch fieldName {
		case "user_id", "name", "username", "email":
		default:
			return nil, fmt.Errorf("invalid user field %s of connector %s", fieldName, connectorId)
		}

		generator.mapConnectorUserid[connectorId] = fieldName
	}
	return generator, nil
}

func (g *skyDisplayUserIdGenerator) DisplayUserId(connector, userid, username, preferredUsername, email string) string {
	if fieldName, ok := g.mapConnectorUserid[connector]; ok {
		switch fieldName {
		case "user_id":
			return userid
		case "name":
			return username
		case "username":
			return preferredUsername
		case "email":
			return email
		}
	}

	// For unconfigured connector, applies a default rule.
	if preferredUsername != "" {
		return preferredUsername
	} else if userid != "" {
		return userid
	} else {
		return username
	}
}
