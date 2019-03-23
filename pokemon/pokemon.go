package pokemon

import (
	"errors"

	"github.com/kjfsm/pokemon-go/pokedex"
)

// Pokemon 1匹のポケモンを表す
type Pokemon struct {
	no        int
	name      string
	types     []string
	level     int
	form      string
	ability   string
	nature    string
	gender    string
	happiness int
	BaseStats [6]int
	iv        individualValue
	ev        effortValue
}

// EffortValue 努力値
type effortValue [6]int

// IndividualValue 個体値
type individualValue [6]int

// Config 初期設定
type Config struct {
	Name    string
	Form    string
	Ability string
	Nature  string
	Gender  string
}

// New ポケモンの情報を指定してPokemonを作る
func New(c *Config) (Pokemon, error) {
	if c.Name == "" {
		return Pokemon{}, errors.New("名前が指定されていません")
	}
	if c.Ability == "" {
		return Pokemon{}, errors.New("とくせいが指定されていません")
	}
	pokedex, err := pokedex.GetPokedex()
	if err != nil {
		return Pokemon{}, err
	}
	for _, pokemon := range pokedex {
		if pokemon.Name == c.Name {
			if pokemon.Form != c.Form {
				continue
			}
			abilities := append(pokemon.Abilities, pokemon.HiddenAbilities...)
			for _, ability := range abilities {
				if ability == c.Ability {
					return Pokemon{
						no:        pokemon.No,
						name:      pokemon.Name,
						types:     pokemon.Types,
						level:     50,
						form:      c.Form,
						ability:   c.Ability,
						nature:    c.Nature,
						gender:    c.Gender,
						happiness: 0,
						BaseStats: [6]int{
							pokemon.Stats.Hp,
							pokemon.Stats.Attack,
							pokemon.Stats.Defence,
							pokemon.Stats.SpAttack,
							pokemon.Stats.SpDefence,
							pokemon.Stats.Speed,
						},
					}, nil
				}
			}
		}
	}
	return Pokemon{}, errors.New("見つかりませんでした")
}

// GetName 名前を返す
func (p *Pokemon) GetName() string {
	return p.name
}

// SetLevel レベルをセットする
func (p *Pokemon) SetLevel(level int) error {
	if level < 1 {
		return errors.New("レベルが低すぎます")
	}
	if level > 100 {
		return errors.New("レベルが高すぎます")
	}
	p.level = level
	return nil
}

// GetLevel レベルを返す
func (p *Pokemon) GetLevel() int {
	return p.level
}

// SetIV 個体値のセット
func (p *Pokemon) SetIV(iv [6]int) error {
	for _, value := range iv {
		if value < 0 || value > 31 {
			return errors.New("値が異常です")
		}
	}
	p.iv = iv
	return nil
}

// SetEV 努力値のセット
func (p *Pokemon) SetEV(ev [6]int) error {
	for _, value := range ev {
		if value < 0 || value > 252 {
			return errors.New("値が異常です")
		}
	}
	sum := 0
	for _, x := range ev {
		sum += x
	}
	if sum > 510 {
		return errors.New("努力値の合計値が510を超えています")
	}
	p.ev = ev
	return nil
}

// CalcStatus 実数値の計算
// BaseStatus 種族値
// IV 個体値
// EV 努力値
func (p *Pokemon) CalcStatus() [6]int {
	return [6]int{
		int((p.BaseStats[0]*2+p.iv[0]+int(p.ev[0]/4))*p.level/100) + 10 + p.level,
		int(int((p.BaseStats[1]*2+p.iv[1]+int(p.ev[1]/4))*p.level/100) + 5),
		int(int((p.BaseStats[2]*2+p.iv[2]+int(p.ev[2]/4))*p.level/100) + 5),
		int(int((p.BaseStats[3]*2+p.iv[3]+int(p.ev[3]/4))*p.level/100) + 5),
		int(int((p.BaseStats[4]*2+p.iv[4]+int(p.ev[4]/4))*p.level/100) + 5),
		int(int((p.BaseStats[5]*2+p.iv[5]+int(p.ev[5]/4))*p.level/100) + 5),
	}
}

// CalcIV 個体値を計算する
func (p *Pokemon) CalcIV() (max [6]int, min [6]int) {
	return [6]int{31, 31, 31, 31, 31, 31}, [6]int{31, 30, 30, 30, 31, 30}
}
