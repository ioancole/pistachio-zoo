package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadAnimals() ([]Animal, error) {
	path := "./animals.json"

	data, err := os.ReadFile(path)

	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", path, err)
	}

	var animals []Animal
	if err := json.Unmarshal(data, &animals); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	return animals, nil
}
