package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Integration", func() {
	var (
		session *gexec.Session
		err     error
		port    string
	)

	Context("", func() {
		BeforeEach(func() {
			port = strconv.Itoa(8080)
			session, err = gexec.Start(
				exec.Command(buildPath),
				GinkgoWriter,
				GinkgoWriter,
			)
			Expect(err).NotTo(HaveOccurred())

			time.Sleep(8 * time.Second)
		})

		AfterEach(func() {
			session.Kill()
		})

		Describe("/", func() {
			It("returns HTML we expect", func() {
				response, err := http.Get("http://test@127.0.0.1:" + port)
				Expect(err).NotTo(HaveOccurred())
				Expect(response).To(BeSuccessful())

				defer response.Body.Close()
				_, err = ioutil.ReadAll(response.Body)
				Expect(err).NotTo(HaveOccurred())
			})
		})
	})

})

func BeSuccessful() types.GomegaMatcher {
	return &SuccessfulHTTPResponseMatcher{}
}

type SuccessfulHTTPResponseMatcher struct{}

func (matcher *SuccessfulHTTPResponseMatcher) Match(actual interface{}) (success bool, err error) {
	response, ok := actual.(*http.Response)
	if !ok {
		return false, fmt.Errorf("SuccessfulHTTPResponseMatcher matcher expects an http.Response")
	}

	return (response.StatusCode >= 200) && (response.StatusCode < 400), nil
}

func (matcher *SuccessfulHTTPResponseMatcher) FailureMessage(actual interface{}) (message string) {
	response := actual.(*http.Response)

	return fmt.Sprintf("Expected Status Code\n\t%d\nto be successful (2XX or 3XX)", response.StatusCode)
}

func (matcher *SuccessfulHTTPResponseMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	response := actual.(*http.Response)

	return fmt.Sprintf("Expected Status Code\n\t%d\nto not be successful (1XX, 4XX, 5XX)", response.StatusCode)
}
