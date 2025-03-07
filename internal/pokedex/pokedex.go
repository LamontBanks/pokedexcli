package pokedex

type Pokedex struct {
	Pokemon map[string]PokedexEntry
}

type PokedexEntry struct {
	NumberCaught int
}

func NewPokedex() Pokedex {
	return Pokedex{
		Pokemon: map[string]PokedexEntry{},
	}
}

func (dex *Pokedex) AddCaughtPokemon(name string) {
	pokemon, exist := dex.Pokemon[name]
	if exist {
		pokemon.NumberCaught += 1
	} else {
		dex.Pokemon[name] = PokedexEntry{
			NumberCaught: 1,
		}
	}
}
