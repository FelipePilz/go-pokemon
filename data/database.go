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

type Type struct {
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

func FindTypes() ([]Type, error) {
	var types []Type

	rows, err := db.Query("SELECT * FROM TYPE")
	if err != nil {
		return nil, fmt.Errorf("FindTypes: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var t Type
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return nil, fmt.Errorf("FindTypes: %v", err)
		}
		types = append(types, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindTypes: %v", err)
	}

	return types, nil
}

func FindPokemons() ([]Pokemon, error) {
	var pokemons []Pokemon

	rows, err := db.Query("SELECT * FROM POKEMON")
	if err != nil {
		return nil, fmt.Errorf("FindPokemons: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var pokemon Pokemon
		if err := rows.Scan(&pokemon.ID, &pokemon.Name); err != nil {
			return nil, fmt.Errorf("FindPokemons: %v", err)
		}
		pokemons = append(pokemons, pokemon)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindPokemons: %v", err)
	}

	return pokemons, nil
}

func FindPokemonsByType(typeName string) ([]Pokemon, error) {
	var pokemons []Pokemon

	rows, err := db.Query("SELECT p.* FROM POKEMON p "+
		"INNER JOIN POKEMON_HAS_TYPE pt ON pt.pokemon_id = p.id "+
		"INNER JOIN TYPE t ON t.id = pt.type_id "+
		"WHERE t.name = ?", typeName)

	if err != nil {
		return nil, fmt.Errorf("FindPokemonsByType %q: %v", typeName, err)
	}

	defer rows.Close()

	for rows.Next() {
		var pokemon Pokemon
		if err := rows.Scan(&pokemon.ID, &pokemon.Name); err != nil {
			return nil, fmt.Errorf("FindPokemonsByType %q: %v", typeName, err)
		}
		pokemons = append(pokemons, pokemon)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("FindPokemonsByType %q: %v", typeName, err)
	}

	return pokemons, nil
}
