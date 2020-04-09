package tests_test

import (
	"os"
	"testing"

	"github.com/mustanish/omelette/app/config"
	"github.com/mustanish/omelette/app/connectors"
	"github.com/mustanish/omelette/app/routes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var accessToken string
var refreshToken string

var _ = BeforeSuite(func() {
	os.Setenv("ENV", "testing")
	cfg, _ := config.LoadConfig()
	connectors.Initialize(cfg)
	routes.InitializeRouter()
})

var _ = AfterSuite(func() {
	connectors.Drop()
	os.Unsetenv("ENV")
})
