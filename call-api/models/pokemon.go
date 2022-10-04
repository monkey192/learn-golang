package models

import (
	"fmt"
)

type Response struct {
	Name    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

// A Pokemon Struct to map every pokemon to.
type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

func (poke Pokemon) Info() {
	fmt.Printf("Entrynumber: %d - Name: %s\n", poke.EntryNo, poke.Species.Name)
}
