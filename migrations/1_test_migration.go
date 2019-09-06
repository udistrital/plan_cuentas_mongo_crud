package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/udistrital/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		_, err := db.Collection(models.MovimientoParameterCollection).InsertOne(context.TODO(), models.MovimientoParameter{
			Multiplicador:       1,
			TipoMovimientoHijo:  "test",
			TipoMovimientoPadre: "test",
		})
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(models.MovimientoParameterCollection).DeleteOne(context.TODO(), models.MovimientoParameter{
			Multiplicador:       1,
			TipoMovimientoHijo:  "test",
			TipoMovimientoPadre: "test",
		})
		if err != nil {
			return err
		}
		return nil
	})
}
