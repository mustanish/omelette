package tests_test

import (
	"os"
	"testing"

	"github.com/mustanish/omelette/app/config"
	"github.com/mustanish/omelette/app/connectors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var _ = BeforeSuite(func() {
	os.Setenv("ENV", "testing")
	config, _ := config.LoadConfig()
	connectors.Initialize(config)
	os.Unsetenv("ENV")
	//log.Println("First One")
})

var _ = AfterSuite(func() {
	connectors.Drop()
	//log.Println("First Two")
})
