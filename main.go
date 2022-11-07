package main

import (
	"os"

	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/routers"
	"gitlab.com/meta-node/mail/server"
)

func main() {
	finish := make(chan bool)
	go runHttpsServer()
	go server.MailServer()

	<-finish
}

func runHttpsServer() {
	dbConn := database.InstanceDB()
	database.Migration(dbConn)
	r := routers.InitRouter()
	os.LookupEnv("PORT")
	// port := config.GetConfig().Server.HTTP.Port
	port, ok := os.LookupEnv("PORT")

	if ok == false {
		port = "3000"
	}
	r.Run(port)
}
