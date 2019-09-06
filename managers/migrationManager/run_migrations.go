package migrationmanager

import (
	"context"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	migrate "github.com/miguelramirez93/mongo-migrate"

	// migrtions ... import the migration file.
	_ "github.com/udistrital/plan_cuentas_mongo_crud/migrations"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RunMigrations ... Migrate all files in migrations package.
func RunMigrations() (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", beego.AppConfig.String("mongo_user"), beego.AppConfig.String("mongo_pass"), beego.AppConfig.String("mongo_host"), "27017")
	opt := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database(beego.AppConfig.String("mongo_db"))
	migrate.SetDatabase(db)
	if err := migrate.Up(migrate.AllAvailable); err != nil {
		return nil, err
	}
	return db, nil
}
