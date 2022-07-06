package db

import (
	"os"
	"time"

	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

// Cursor devuelve un apuntador con la conexión a la bd y la colección especifica
func Cursor(session *mgo.Session, Collection string) *mgo.Collection {
	mongoDB := beego.AppConfig.String("mongo::db")
	c := session.DB(mongoDB).C(Collection)
	return c
}

// GetSession crea una sesión con las credenciales
func GetSession() (*mgo.Session, error) {

	mongoHost := beego.AppConfig.String("mongo::host")
	// This comporbation is for test cases.
	if mongoHost == "" {
		mongoHost = os.Getenv("FINANCIERA_MONGO_CRUD_DB_URL")
	}
	mongoPort := beego.AppConfig.String("mongo::port")
	if mongoPort == "" {
		mongoPort = os.Getenv("FINANCIERA_MONGO_CRUD_DB_PORT")
	}
	mongoUser := beego.AppConfig.String("mongo::user")
	if mongoUser == "" {
		mongoUser = os.Getenv("FINANCIERA_MONGO_CRUD_DB_USER")
	}
	mongoPassword := beego.AppConfig.String("mongo::pass")
	if mongoPassword == "" {
		mongoPassword = os.Getenv("FINANCIERA_MONGO_CRUD_DB_PASS")
	}
	mongoDatabase := beego.AppConfig.String("mongo::db")
	if mongoDatabase == "" {
		mongoDatabase = os.Getenv("FINANCIERA_MONGO_CRUD_DB_NAME")
	}
	mongoAuthDb := beego.AppConfig.String("mongo::auth")
	if mongoAuthDb == "" {
		mongoAuthDb = os.Getenv("FINANCIERA_MONGO_CRUD_DB_AUTH")
	}
	info := &mgo.DialInfo{
		Addrs:    []string{mongoHost + ":" + mongoPort},
		Timeout:  60 * time.Second,
		Database: mongoDatabase,
		Username: mongoUser,
		Password: mongoPassword,
		Source:   mongoAuthDb,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	} else {
		session.SetMode(mgo.Monotonic, true)
	}

	return session, err
}
