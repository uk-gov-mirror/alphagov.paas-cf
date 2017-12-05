package acceptance_test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os/exec"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/config"
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
	"github.com/cloudfoundry-incubator/cf-test-helpers/workflowhelpers"
)

const (
	BYTE     = int64(1)
	KILOBYTE = 1024 * BYTE
	MEGABYTE = 1024 * KILOBYTE
	GIGABYTE = 1024 * MEGABYTE
	TERABYTE = 1024 * GIGABYTE

	DEFAULT_MEMORY_LIMIT = "256M"
	DB_CREATE_TIMEOUT    = 30 * time.Minute
)

var (
	testConfig *config.Config
	httpClient *http.Client
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)

	testConfig = config.LoadConfig()
	testConfig.NamePrefix = "ACC"

	httpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: testConfig.SkipSSLValidation},
		},
	}

	testContext := workflowhelpers.NewTestSuiteSetup(testConfig)

	BeforeSuite(func() {
		testContext.Setup()
		// FIXME this should be removed once these services are generally available.
		org := testContext.GetOrganizationName()
		workflowhelpers.AsUser(testContext.AdminUserContext(), testContext.ShortTimeout(), func() {
			enableServiceAccess := cf.Cf("enable-service-access", "mongodb", "-o", org).Wait(testConfig.DefaultTimeoutDuration())
			Expect(enableServiceAccess).To(Exit(0))
			Expect(enableServiceAccess).To(Say("OK"))

			enableServiceAccess = cf.Cf("enable-service-access", "elasticsearch", "-o", org).Wait(testConfig.DefaultTimeoutDuration())
			Expect(enableServiceAccess).To(Exit(0))
			Expect(enableServiceAccess).To(Say("OK"))

			enableServiceAccess = cf.Cf("enable-service-access", "compose-redis", "-o", org).Wait(testConfig.DefaultTimeoutDuration())
			Expect(enableServiceAccess).To(Exit(0))
			Expect(enableServiceAccess).To(Say("OK"))

			enableServiceAccess = cf.Cf("enable-service-access", "redis", "-o", org).Wait(testConfig.DefaultTimeoutDuration())
			Expect(enableServiceAccess).To(Exit(0))
			Expect(enableServiceAccess).To(Say("OK"))
		})
	})

	AfterSuite(func() {
		testContext.Teardown()
	})

	componentName := "Custom-Acceptance-Tests"
	if testConfig.ArtifactsDirectory != "" {
		helpers.EnableCFTrace(testConfig, componentName)
	}

	RunSpecs(t, componentName)
}

// quietCf is an equivelent of cf.Cf that doesn't send the output to
// GinkgoWriter. Used when you don't want the output, even in verbose mode (eg
// when polling the API)
func quietCf(program string, args ...string) *Session {
	command, err := Start(exec.Command(program, args...), nil, nil)
	Expect(err).NotTo(HaveOccurred())
	return command
}

func pollForServiceCreationCompletion(dbInstanceName string) {
	fmt.Fprint(GinkgoWriter, "Polling for service creation to complete")
	Eventually(func() *Buffer {
		fmt.Fprint(GinkgoWriter, ".")
		command := quietCf("cf", "service", dbInstanceName).Wait(testConfig.DefaultTimeoutDuration())
		Expect(command).To(Exit(0))
		return command.Out
	}, DB_CREATE_TIMEOUT, 15*time.Second).Should(Say("create succeeded"))
	fmt.Fprint(GinkgoWriter, "done\n")
}

func pollForServiceDeletionCompletion(dbInstanceName string) {
	fmt.Fprint(GinkgoWriter, "Polling for service destruction to complete")
	Eventually(func() *Buffer {
		fmt.Fprint(GinkgoWriter, ".")
		command := quietCf("cf", "services").Wait(testConfig.DefaultTimeoutDuration())
		Expect(command).To(Exit(0))
		return command.Out
	}, DB_CREATE_TIMEOUT, 15*time.Second).ShouldNot(Say(dbInstanceName))
	fmt.Fprint(GinkgoWriter, "done\n")
}