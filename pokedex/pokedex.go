package pokedex

import (
	"github.com/thamaji/req"
)

// Pokedex ポケモン図鑑
type Pokedex []pokedata

// Pokedata ポケモン1匹のデータ
type pokedata struct {
	No              int      `json:"no"`
	Name            string   `json:"name"`
	Form            string   `json:"form"`
	IsMegaEvolution bool     `json:"isMegaEvolution"`
	Evolutions      []int    `json:"evolutions"`
	Types           []string `json:"types"`
	Abilities       []string `json:"abilities"`
	HiddenAbilities []string `json:"hiddenAbilities"`
	Stats           stats    `json:"stats"`
}

// Stats 種族値
type stats struct {
	Hp        int `json:"hp"`
	Attack    int `json:"attack"`
	Defence   int `json:"defence"`
	SpAttack  int `json:"spAttack"`
	SpDefence int `json:"spDefence"`
	Speed     int `json:"speed"`
}

// GetPokedex JSONから受け取る
func GetPokedex() (Pokedex, error) {
	var pokedex Pokedex
	url := "https://raw.githubusercontent.com/kotofurumiya/pokemon_data/master/data/pokemon_data.json"
	request := req.Get(url, nil)
	if err := request.FetchJSON(&pokedex); err != nil {
		return nil, err
	}
	return pokedex, nil
}
