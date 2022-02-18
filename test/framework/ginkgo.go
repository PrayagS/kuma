package framework

import (
	"github.com/onsi/ginkgo"
	ginkgo_config "github.com/onsi/ginkgo/config"
)

var suiteFailed = false

func ShouldSkipCleanup() bool {
	return ginkgo.CurrentGinkgoTestDescription().Failed && ginkgo_config.GinkgoConfig.FailFast
}

func ShouldSkipSuiteCleanup() bool {
	return suiteFailed && ginkgo_config.GinkgoConfig.FailFast
}

// E2EAfterEachMarkIfFailed will set global variable suiteFailed to true,
// if any spec would fail. It's necessary as if we want to prevent cleanup which
// is defined in AfterEachSuite blocks Ginkgo v1 is not setting ginkgo.CurrentGinkgoTestDescription().Failed
// to true, so we have to rely on the variable set in this function
// ref. https://github.com/onsi/ginkgo/issues/361
func E2EAfterEachMarkIfFailed() {
	ginkgo.AfterEach(func() {
		suiteFailed = suiteFailed || ginkgo.CurrentGinkgoTestDescription().Failed
	})
}

func E2EAfterEach(fn func()) {
	ginkgo.AfterEach(func() {
		if ShouldSkipCleanup() {
			return
		}
		fn()
	})
}

func E2EAfterSuite(fn func()) {
	ginkgo.AfterSuite(func() {
		if ShouldSkipSuiteCleanup() {
			return
		}
		fn()
	})
}

func E2EBeforeSuite(fn func()) {
	ginkgo.BeforeSuite(fn)
}
