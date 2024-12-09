package services

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strings"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/labstack/echo/v4"
	"github.com/liukeshao/echo-admin/config"
	"github.com/liukeshao/echo-admin/ent"
	_ "github.com/mattn/go-sqlite3"
	// Require by ent
	_ "github.com/liukeshao/echo-admin/ent/runtime"
)

// Container contains all services used by the application and provides an easy way to handle dependency
// injection including within tests
type Container struct {
	Web      *echo.Echo
	Config   *config.Config
	Cache    *CacheClient
	Database *sql.DB
	ORM      *ent.Client

	Auth    *AuthService
	User    *UserService
	Org     *OrgService
	Cluster *ClusterService
}

// NewContainer creates and initializes a new Container
func NewContainer() *Container {
	c := new(Container)
	c.initConfig()
	c.initWeb()
	c.initCache()
	c.initDatabase()
	c.initORM()
	c.initServices()
	return c
}

// Shutdown shuts the Container down and disconnects all connections.
// If the task runner was started, cancel the context to shut it down prior to calling this.
func (c *Container) Shutdown() error {
	if err := c.ORM.Close(); err != nil {
		return err
	}
	if err := c.Database.Close(); err != nil {
		return err
	}
	c.Cache.Close()

	return nil
}

// initConfig initializes configuration
func (c *Container) initConfig() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	c.Config = &cfg

	// Configure logging
	switch cfg.App.Environment {
	case config.EnvProduction:
		slog.SetLogLoggerLevel(slog.LevelInfo)
	default:
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}
}

// initWeb initializes the web framework
func (c *Container) initWeb() {
	c.Web = echo.New()
	c.Web.HideBanner = true
}

// initCache initializes the cache
func (c *Container) initCache() {
	store, err := newInMemoryCache(c.Config.Cache.Capacity)
	if err != nil {
		panic(err)
	}

	c.Cache = NewCacheClient(store)
}

// initDatabase initializes the database
func (c *Container) initDatabase() {
	var err error
	var connection string

	switch c.Config.App.Environment {
	case config.EnvTest:
		// TODO: Drop/recreate the DB, if this isn't in memory?
		connection = c.Config.Database.TestConnection
	default:
		connection = c.Config.Database.Connection
	}

	c.Database, err = openDB(c.Config.Database.Driver, connection)
	if err != nil {
		panic(err)
	}
}

// initORM initializes the ORM
func (c *Container) initORM() {
	drv := entsql.OpenDB(c.Config.Database.Driver, c.Database)
	c.ORM = ent.NewClient(ent.Driver(drv))
	if c.Config.App.Environment != config.EnvProduction {
		c.ORM = c.ORM.Debug()
	}

	// Run the auto migration tool.
	if err := c.ORM.Schema.Create(context.Background()); err != nil {
		panic(err)
	}
}

func (c *Container) initServices() {
	c.User = NewUserService(c.ORM)
	c.Auth = NewAuthService(c.Config)
	c.Org = NewOrgService(c.ORM)
	c.Cluster = NewClusterService(c.ORM)
}

// openDB opens a database connection
func openDB(driver, connection string) (*sql.DB, error) {
	// Helper to automatically create the directories that the specified sqlite file
	// should reside in, if one
	if driver == "sqlite3" {
		d := strings.Split(connection, "/")

		if len(d) > 1 {
			path := strings.Join(d[:len(d)-1], "/")

			if err := os.MkdirAll(path, 0755); err != nil {
				return nil, err
			}
		}
	}

	return sql.Open(driver, connection)
}
