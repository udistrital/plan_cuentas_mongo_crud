package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/udistrital/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		AdicionFuenteToApropiacion := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "ad_fuente_apropiacion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}
		ReduccionFuenteToApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "rd_fuente_apropiacion",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}
		parameters := []interface{}{
			AdicionFuenteToApropiacion,
			ReduccionFuenteToApropiacion,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).InsertMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		AdicionFuenteToApropiacion := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "ad_fuente",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}
		ReduccionFuenteToApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "rd_fuente",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}
		parameters := []interface{}{
			AdicionFuenteToApropiacion,
			ReduccionFuenteToApropiacion,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).DeleteMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	})
}
