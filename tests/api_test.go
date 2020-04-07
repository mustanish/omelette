package tests_test

import (
	"testing"

	"github.com/mustanish/omelette/app/connectors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var _ = BeforeSuite(func() {
	connectors.InitializeDB()
})

var _ = AfterSuite(func() {
})
