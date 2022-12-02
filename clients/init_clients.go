package clients

import "sample-gin-server/clients/db"

var GameDB db.Database
var URLDB db.URLDatabase

func Init() {
	GameDB = db.Init()
	URLDB = db.InitURLDatabase()
}
