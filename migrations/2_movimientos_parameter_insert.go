package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/udistrital/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		anulateSTatus := "anulado"
		AnulacionCdpToRp := models.MovimientoParameter{
			Multiplicador:          -1,
			TipoMovimientoHijo:     "anul_cdp",
			TipoMovimientoPadre:    "cdp",
			Initial:                true,
			NoBalanceLeftStateName: &anulateSTatus,
		}
		AnulacionRpToRp := models.MovimientoParameter{
			Multiplicador:       -1,
			TipoMovimientoHijo:  "anul_rp",
			TipoMovimientoPadre: "rp",
			Initial:             true,
		}
		AnulacionRpToCDP := models.MovimientoParameter{
			Multiplicador:       1,
			TipoMovimientoHijo:  "anul_rp",
			TipoMovimientoPadre: "cdp",
		}
		CreacionRp := models.MovimientoParameter{
			Multiplicador:       -1,
			TipoMovimientoHijo:  "rp",
			TipoMovimientoPadre: "cdp",
			Initial:             true,
		}
		CreacionCdp := models.MovimientoParameter{
			Multiplicador:        -1,
			TipoMovimientoHijo:   "cdp",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			Initial:              true,
			WithOutChangeState:   true,
		}
		AnulacionCdpToApropiacion := models.MovimientoParameter{
			Multiplicador:        1,
			TipoMovimientoHijo:   "anul_cdp",
			TipoMovimientoPadre:  "apropiacion",
			FatherCollectionName: "arbol_rubro_apropiacion",
			WithOutChangeState:   true,
		}

		parameters := []interface{}{
			AnulacionCdpToRp,
			AnulacionRpToRp,
			AnulacionRpToCDP,
			CreacionRp,
			CreacionCdp,
			AnulacionCdpToApropiacion,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).InsertMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		AnulacionCdp := models.MovimientoParameter{
			Multiplicador:       -1,
			TipoMovimientoHijo:  "anulacion_cdp",
			TipoMovimientoPadre: "cdp",
		}
		AnulacionRp := models.MovimientoParameter{
			Multiplicador:       -1,
			TipoMovimientoHijo:  "anulacion_rp",
			TipoMovimientoPadre: "rp",
			Initial:             true,
		}
		CreacionRp := models.MovimientoParameter{
			Multiplicador:       -1,
			TipoMovimientoHijo:  "rp",
			TipoMovimientoPadre: "cdp",
			Initial:             true,
		}

		parameters := []interface{}{
			AnulacionCdp,
			AnulacionRp,
			CreacionRp,
		}
		_, err := db.Collection(models.MovimientoParameterCollection).DeleteMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	})
}
