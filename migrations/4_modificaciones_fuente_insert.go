package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/udistrital/mongo-migrate"
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

		TrasladoOrigenFuente := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "tr_fuente",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		TrasladoDestinoFuente := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "tr_fuente_destino",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		parameters := []interface{}{
			AdicionFuente,
			TrasladoOrigenFuente,
			TrasladoDestinoFuente,
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

		TrasladoOrigenFuente := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "tr_fuente",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		TrasladoDestinoFuente := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "tr_fuente_destino",
			TipoMovimientoPadre:  "fuente",
			FatherCollectionName: "fuente_financiamiento",
			Initial:              true,
			WithOutChangeState:   true,
		}

		parameters := []interface{}{
			AdicionFuente,
			TrasladoOrigenFuente,
			TrasladoDestinoFuente,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).DeleteMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	})
}
