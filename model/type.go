package model

import (
	"fmt"
)

type Type struct {
	ID   int64
	Name string
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
