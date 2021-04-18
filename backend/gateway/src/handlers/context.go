package handlers

import (
	"ProjectMIRRO/backend/gateway/model"

	"gopkg.in/mgo.v2"
)

//Middleware to stores: mysql connection
type Context struct {
	UserSql  *model.MySQLStore
	User     *User
	MongoDBs *Mongos
}

//Current user information
type User struct {
	Userid    string `json:"userid"`
	UserEmail string `json:"useremail"`
}

//Mongo Connections
//Mongo1 is for user information on the personal profile
//Mongo2 is for overall information for ML
type Mongos struct {
	Mongo []*mgo.Database
}
