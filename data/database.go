package server

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

type Pokemon struct {
	ID   int64
	Name string
}

func Start() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "pokemons",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Database connection success!")
}

func FindPokemons() ([]Pokemon, error) {
	var pokemons []Pokemon

	rows, err := db.Query("SELECT * FROM POKEMON")
	if err != nil {
		return nil, fmt.Errorf("findPokemons: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var pokemon Pokemon
		if err := rows.Scan(&pokemon.ID, &pokemon.Name); err != nil {
			return nil, fmt.Errorf("findPokemons: %v", err)
		}
		pokemons = append(pokemons, pokemon)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findPokemons: %v", err)
	}

	return pokemons, nil
}
