package bootstrap

import (
	migrations "github.com/gobardofw/console/migration"
	"github.com/gobardofw/database"
	"github.com/jmoiron/sqlx"
)

// SetupDatabase driver
func SetupDatabase() {
	conf := app.Config()
	if db, err := database.NewMySQLConnector(
		conf.String("mysql.host", "127.0.0.1"),
		conf.String("mysql.username", "root"),
		conf.String("mysql.password", ""),
		conf.String("mysql.database", "// {{.name}}"),
	); err != nil {
		_container.Register("--APP-DB", db)
	} else {
		panic("failed to init database " + err.Error())
	}

	_cli.AddCommand(migrations.MigrationCommand(func(driver string) *sqlx.DB {
		return app.Database(driver)
	}, "--APP-DB", "./database/migrations", "./database/seeds"))
}

// Database get database driver
// leave name empty to resolve default
func (app *AppDriver) Database(names ...string) *sqlx.DB {
	name := "--APP-DB"
	if len(names) > 0 {
		name = names[0]
	}
	if dep, exists := app.Container.Resolve(name); exists {
		if res, ok := dep.(*sqlx.DB); ok {
			return res
		}
	}
	return nil
}
