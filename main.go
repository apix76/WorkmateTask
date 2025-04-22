package main

import (
	"WorkmateTask/conf"
	"WorkmateTask/db"
	"WorkmateTask/rest"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	con, err := conf.NewConf()
	if err != nil {
		log.Fatal(err)
	}

	db := db.NewDbAccess(con.RdbUrl)
	rest.Server(con, db)
}
