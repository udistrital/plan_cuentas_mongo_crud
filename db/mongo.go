package db

import (
	"os"
	"time"

	"github.com/astaxie/beego"
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
	// This comporbation is for test cases.
	if mongoHost == "" {
		mongoHost = os.Getenv("FINANCIERA_MONGO_CRUD_DB_URL")
	}
	mongoUser := beego.AppConfig.String("mongo_user")
	if mongoUser == "" {
		mongoUser = os.Getenv("FINANCIERA_MONGO_CRUD_DB_USER")
	}
	mongoPassword := beego.AppConfig.String("mongo_pass")
	if mongoPassword == "" {
		mongoPassword = os.Getenv("FINANCIERA_MONGO_CRUD_DB_PASS")
	}
	mongoDatabase := beego.AppConfig.String("mongo_db_connect")
	if mongoDatabase == "" {
		mongoDatabase = os.Getenv("FINANCIERA_MONGO_CRUD_DB_NAME")
	}
	mongoAuthDb := beego.AppConfig.String("mongo_db_auth")
	if mongoAuthDb == "" {
		mongoAuthDb = os.Getenv("FINANCIERA_MONGO_CRUD_DB_AUTH")
	}
	info := &mgo.DialInfo{
		Addrs:    []string{mongoHost},
		Timeout:  60 * time.Second,
		Database: mongoDatabase,
		Username: mongoUser,
		Password: mongoPassword,
		Source:   mongoDatabase,
	}

	session, err := mgo.DialWithInfo(info)
	if err != nil {
		panic(err)
	} else {
		session.SetMode(mgo.Monotonic, true)
	}

	return session, err
}
