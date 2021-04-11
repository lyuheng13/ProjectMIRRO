package handlers

import "ProjectMIRRO/backend/gateway/model"

//Middleware to stores: mysql connection
type Context struct {
	UserSql *model.MySQLStore
}
