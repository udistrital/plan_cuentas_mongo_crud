package migrations

import (
	"context"

	"github.com/udistrital/plan_cuentas_mongo_crud/models"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		root3 := models.ArbolRubroParameter{
			Activo:          true,
			Tipo:            "raiz",
			Valor:           "3",
			UnidadEjecutora: "0",
		}
		root2 := models.ArbolRubroParameter{
			Activo:          true,
			Tipo:            "raiz",
			Valor:           "2",
			UnidadEjecutora: "0",
		}

		parameters := []interface{}{
			root2,
			root3,
		}
		_, err := db.Collection(models.ArbolRubroParameterCollection).InsertMany(context.TODO(), parameters)
		if err != nil {
			return err
		}
		return nil
	}, func(db *mongo.Database) error {
		_, err := db.Collection(models.ArbolRubroParameterCollection).DeleteOne(context.TODO(), models.ArbolRubroParameter{
			Activo:          true,
			Tipo:            "raiz",
			Valor:           "3",
			UnidadEjecutora: "0",
		})
		if err != nil {
			return err
		}
		return nil
	})
}
