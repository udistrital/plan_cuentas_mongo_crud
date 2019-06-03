package db

import (
	"time"

	"github.com/globalsign/mgo"
)

// Cursor devuelve un apuntador con la conexión a la bd y la colección especifica
func Cursor(session *mgo.Session, Collection string) *mgo.Collection {
	mongoDB := beego.AppConfig.String("mongo_db")
	c := session.DB(mongoDB).C(Collection)
	return c
}

// GetSession crea una sesión con las credenciales
func GetSession() (*mgo.Session, error) {

	mongoHost := beego.AppConfig.String("mongo_host")
	mongoUser := beego.AppConfig.String("mongo_user")
	mongoPassword := beego.AppConfig.String("mongo_pass")
	mongoDatabase := beego.AppConfig.String("mongo_db_connect")

	info := &mgo.DialInfo{
		Addrs:    []string{mongoHost},
		Timeout:  60 * time.Second,
		Database: mongoDatabase,
		Username: mongoUser,
		Password: mongoPassword,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	} else {
		session.SetMode(mgo.Monotonic, true)
	}

	return session, err
}
