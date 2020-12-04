package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		AdicionFuente := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "ad_fuente",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		ReduccionFuente := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "rd_fuente",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		parameters := []interface{}{
			AdicionFuente,
			ReduccionFuente,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).InsertMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		AdicionFuente := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "ad_fuente",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		ReduccionFuente := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "rd_fuente",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		parameters := []interface{}{
			AdicionFuente,
			ReduccionFuente,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).DeleteMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	})
}
