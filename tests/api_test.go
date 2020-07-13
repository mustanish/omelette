package tests_test

import (
	"omelette/app/connectors"
	"omelette/app/routes"
	"omelette/config"
	"os"
	"testing"

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
	os.Setenv("DATABASE_URL", "http://localhost:8529")
	cfg, _ := config.LoadConfig()
	connectors.Initialize(cfg)
	routes.InitializeRouter()
})

var _ = AfterSuite(func() {
	connectors.Drop()
	os.Unsetenv("ENV")
	os.Unsetenv("DATABASE_URL")
})
