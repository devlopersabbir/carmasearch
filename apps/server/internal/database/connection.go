package database

import (
	"context"
	"log"

	"github.com/carmasearch/carma-server/arch/redis"
	"github.com/carmasearch/carma-server/internal/config"
	"github.com/carmasearch/carma-server/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection(env *config.Config, ctx context.Context) (*gorm.DB, redis.Store) {
	dsn := env.GeneratePGConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	// auto migrations
	err = migrations.Automigrate(db)
	if err != nil {
		log.Fatalf("Couldn't migrate database: %s", err)
	}

	store := redis.NewStore(ctx, &env.Redis)
	store.Connect()

	// Elastic Search Connection
	ESClientConnection()
	ESCreateIndexIfNotExist()

	return db, store
}
