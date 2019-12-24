package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/udistrital/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {

		AnulacionCdpToRp := models.MovimientoParameter{
			Multiplicador:       -1,
			TipoMovimientoHijo:  "anul_cdp_suspension",
			TipoMovimientoPadre: "cdp_suspension",
			Initial:             true,
		}

		AnulacionCdpToApropiacion := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "anul_cdp_suspension",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			WithOutChangeState:   true,
		}

		CdpModificacionApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "cdp_suspension",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			WithOutChangeState:   true,
			Initial:              true,
		}

		parameters := []interface{}{
			AnulacionCdpToRp,
			AnulacionCdpToApropiacion,
			CdpModificacionApropiacion,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).InsertMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {

		AnulacionCdpToRp := models.MovimientoParameter{
			Multiplicador:       -1,
			TipoMovimientoHijo:  "anul_cdp_suspension",
			TipoMovimientoPadre: "cdp_suspension",
			Initial:             true,
		}

		AnulacionCdpToApropiacion := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "anul_cdp_suspension",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			WithOutChangeState:   true,
		}

		CdpModificacionApropiacion := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "cdp_suspension",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			WithOutChangeState:   true,
		}

		parameters := []interface{}{
			AnulacionCdpToRp,
			AnulacionCdpToApropiacion,
			CdpModificacionApropiacion,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).DeleteMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	})
}
