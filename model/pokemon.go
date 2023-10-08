package model

import (
	"fmt"
)

type Pokemon struct {
	ID   int64
	Name string
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
