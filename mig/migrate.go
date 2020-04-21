package mig

import (
	"database/sql"
	"fmt"
	"github.com/aibotsoft/micro/config"
	"github.com/golang-migrate/migrate/v4"
	//_ "github.com/denisenkom/go-mssqldb"
	//_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	stub "github.com/golang-migrate/migrate/v4/database/sqlserver"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"time"
)

func MigrateUp(cfg *config.Config, log *zap.SugaredLogger, db *sql.DB) error {
	start := time.Now()
	log.Info("begin db migration..")
	instance, err := stub.WithInstance(db, &stub.Config{SchemaName: "dbo", DatabaseName: cfg.Mssql.Database})
	if err != nil {
		return errors.Wrap(err, "WithInstance error")
	}
	version, dirty, err := instance.Version()
	if err != nil {
		return errors.Wrap(err, "instance.Version error")
	}
	log.Info("current db version: ", version, ", is dirty: ", dirty)
	if dirty {
		return errors.New("db is dirty")
	}
	sourceURL := fmt.Sprintf("github://%s:%s@%s/%s/migrations", cfg.Migrate.User, cfg.Migrate.Token, cfg.Migrate.User, cfg.Service.Name)

	migrateInstance, err := migrate.NewWithDatabaseInstance(sourceURL, "sqlserver", instance)
	if err != nil {
		return errors.Wrap(err, "migrate.NewWithDatabaseInstance error")
	}

	err = migrateInstance.Up()
	switch {
	case err == migrate.ErrNoChange:
		log.Infow("migration done, no change", "time", time.Since(start))
		return nil
	case err != nil:
		return errors.Wrap(err, "migrateInstance.Up error")
	}
	newVersion, dirty, err := migrateInstance.Version()
	log.Infow("migration done", "db_version: ", newVersion, "is dirty", dirty, "time", time.Since(start))
	return nil
}
