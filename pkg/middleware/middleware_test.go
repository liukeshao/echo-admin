package middleware

import (
	"github.com/liukeshao/echo-admin/pkg/services"
	"os"
	"testing"

	"github.com/liukeshao/echo-admin/config"
	"github.com/liukeshao/echo-admin/ent"
	"github.com/liukeshao/echo-admin/pkg/tests"
)

var (
	c   *services.Container
	usr *ent.User
)

func TestMain(m *testing.M) {
	// Set the environment to test
	config.SwitchEnvironment(config.EnvTest)

	// Create a new container
	c = services.NewContainer()

	// Create a user
	var err error
	if usr, err = tests.CreateUser(c.ORM); err != nil {
		panic(err)
	}

	// Run tests
	exitVal := m.Run()

	// Shutdown the container
	if err = c.Shutdown(); err != nil {
		panic(err)
	}

	os.Exit(exitVal)
}
