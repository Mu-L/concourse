package testflight_test

import (
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("Reconfiguring a resource", func() {
	It("picks up the new configuration immediately", func() {
		guid1, err := uuid.NewRandom()
		Expect(err).ToNot(HaveOccurred())

		guid2, err := uuid.NewRandom()
		Expect(err).ToNot(HaveOccurred())

		setAndUnpausePipeline(
			"fixtures/reconfiguring.yml",
			"-v", "force_version="+guid1.String(),
		)

		watch := fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
		Expect(watch).To(gbytes.Say(guid1.String()))

		setAndUnpausePipeline(
			"fixtures/reconfiguring.yml",
			"-v", "force_version="+guid2.String(),
		)

		watch = fly("trigger-job", "-j", inPipeline("some-passing-job"), "-w")
		Expect(watch).To(gbytes.Say(guid2.String()))
	})
})
