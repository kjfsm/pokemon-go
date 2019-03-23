package pokemon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	_, err := New(&Config{Name: "フシギダネ", Ability: "しんりょく"})
	assert.Nil(t, err)

	_, err = New(&Config{})
	assert.NotNil(t, err)

	_, err = New(&Config{Name: "デオキシス", Form: "アタックフォルム", Ability: "プレッシャー"})
	assert.Nil(t, err)

	_, err = New(&Config{Name: "アネﾞデパミﾞ", Ability: "？？？"})
	assert.NotNil(t, err)
}

func TestGetName(t *testing.T) {
	hushi, err := New(&Config{Name: "フシギダネ", Ability: "しんりょく"})
	assert.Nil(t, err)
	assert.Equal(t, hushi.GetName(), "フシギダネ")
}

func TestSetLevel(t *testing.T) {
	hushi, _ := New(&Config{Name: "フシギダネ", Ability: "しんりょく"})
	err := hushi.SetLevel(1)
	assert.Nil(t, err)
	err = hushi.SetLevel(100)
	assert.Nil(t, err)
	assert.Equal(t, 100, hushi.level)
	err = hushi.SetLevel(0)
	assert.NotNil(t, err)
	err = hushi.SetLevel(101)
	assert.NotNil(t, err)
}

func TestGetLevel(t *testing.T) {
	hushi, _ := New(&Config{Name: "フシギダネ"})
	hushi.SetLevel(50)
	assert.Equal(t, 50, hushi.GetLevel())
}

func TestSetIV(t *testing.T) {
	hushi, _ := New(&Config{Name: "フシギダネ"})
	err := hushi.SetIV([6]int{2, 2, 3, 4, 5, 6})
	assert.Nil(t, err)

	err = hushi.SetIV([6]int{2, 2, 3, 4, 5, 32})
	assert.NotNil(t, err)
}

func TestSetEV(t *testing.T) {
	hushi, _ := New(&Config{Name: "フシギダネ"})
	err := hushi.SetEV([6]int{1, 2, 3, 4, 5, 6})
	assert.Nil(t, err)

	err = hushi.SetEV([6]int{253, 0, 0, 0, 0, 0})
	assert.NotNil(t, err)

	err = hushi.SetEV([6]int{252, 252, 252, 252, 252, 252})
	assert.NotNil(t, err)
}

func TestCalcStatus(t *testing.T) {
	happinas, err := New(&Config{Name: "ハピナス", Ability: "しぜんかいふく"})
	assert.Nil(t, err)
	happinas.SetLevel(50)
	happinas.SetIV([6]int{31, 31, 31, 31, 31, 31})
	happinas.SetEV([6]int{252, 0, 0, 0, 252, 0})
	status := happinas.CalcStatus()
	assert.Equal(t, status, [6]int{362, 30, 30, 95, 187, 75})
}

func TestCalcSecretStatus(t *testing.T) {
	happinas, err := New(&Config{Name: "ハピナス", Ability: "しぜんかいふく"})
	assert.Nil(t, err)
	happinas.SetLevel(50)
	happinas.SetEV([6]int{252, 0, 0, 0, 252, 0})
	max, min := happinas.CalcIV()
	assert.Equal(t, max, [6]int{31, 31, 31, 31, 31, 31})
	assert.Equal(t, min, [6]int{31, 30, 30, 30, 31, 30})

}
