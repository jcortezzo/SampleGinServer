package clients

import "sample-gin-server/clients/db"

var GameDB db.Database

func Init() {
	GameDB = db.Init()
}
