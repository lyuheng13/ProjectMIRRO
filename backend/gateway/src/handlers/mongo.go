package handlers

import (
	"fmt"
	"os"

	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var Session *mgo.Session
var Databases *mgo.Database
var MgoError error

const (
	MONGO_HOST = ""
	MONGO_PORT = ""
	MONGO_DB   = ""
	MONGO_USER = ""
	MONGO_PWD  = ""
)

//Init two mongodb and settings
func (ch *Context) Init() {
	mongoes := ch.MongoDBs
	mongoes.Mongo[0] = ch.ConnectMongo("mongoone", "27137", "user", "niupi", "@NIUPI123")
	mongoes.Mongo[1] = ch.ConnectMongo("mongotwo", "27138", "ml", "niupi", "@NIUPI123")
}

//Take MongoDB host, port, dbname, username, and password. Return a mongodb connection
func (ch *Context) ConnectMongo(MONGO_HOST string, MONGO_PORT string, MONGO_DB string, MONGO_USER string, MONGO_PWD string) *mgo.Database {
	Session, MgoError = mgo.Dial(fmt.Sprintf("%s:%s", MONGO_HOST, MONGO_PORT))
	if MgoError != nil {
		fmt.Println("Wrong mongdb link！")
		os.Exit(1)
	}

	Databases = Session.DB(MONGO_DB)

	MgoError = Databases.Login(MONGO_USER, MONGO_PWD)
	if MgoError != nil {
		fmt.Println("Authentication failed！")
		os.Exit(1)
	}

	return Databases
	// defer Session.Close()
}
