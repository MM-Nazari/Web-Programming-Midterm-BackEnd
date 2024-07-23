package Migration

import (
	"WebMidterm/Controller"
	"WebMidterm/Interface"
	"WebMidterm/Repository"
	"database/sql"
	"log"
)

func Db() Controller.BasketController {
	db, err := sql.Open("sqlite3", "Basket.db")
	if err != nil {
		log.Fatal(err)
	}

	var basketRepository Interface.Repository
	basketRepository = Repository.NewSQLiteRepository(db)

	if err := basketRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	h := Controller.BasketController{Repo: basketRepository}
	return h
}
