package worldofwish

import (
	"fmt"
	"math/rand"
	"time"
)

type Human interface {
	NewHuman(name string, vie int32, degats uint32) *Races_t
}

type Orc interface {
	NewOrc(name string, vie int32, degats uint32) *Races_t
}

type Races interface {
	GetVie() int32
	GetDegats() int32
	GetName() string
	SetVie(vie int32)
	SetDegats(degats int32)
	SetName(name string)
	Attack(enemy *Races_t)
	Human
	Orc
}

type Races_t struct {
	vie    int32
	degats int32
	name   string
}

func New() *Races_t {
	return &Races_t{}
}

func (rs *Races_t) NewHuman(name string, vie int32, degats int32) *Races_t {
	rs.vie = vie
	rs.degats = degats
	rs.name = name
	return rs
}

func (rs *Races_t) NewOrc(name string, vie int32, degats int32) *Races_t {
	rs.vie = vie
	rs.degats = degats
	rs.name = name
	return rs
}

func (rs *Races_t) GetVie() int32 {
	return rs.vie
}

func (rs *Races_t) GetDegats() int32 {
	return rs.degats
}

func (rs *Races_t) GetName() string {
	return rs.name
}

func (rs *Races_t) SetVie(vie int32) {
	rs.vie = vie
}

func (rs *Races_t) SetDegats(degats int32) {
	rs.degats = degats
}

func (rs *Races_t) SetName(name string) {
	rs.name = name
}

func random(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := min + rand.Intn(max-min+1)

	return randomNumber
}

func (rs *Races_t) Attack(enemy *Races_t) {

	lucky := random(100, 1)

	if lucky <= 50 { // 50% attack is failed
		fmt.Println("Attaque failed !!!")
		return
	} else {
		if enemy.GetVie() <= 0 {
			enemy.SetVie(0)
			return
		} else {
=			enemy.SetVie(enemy.GetVie() - rs.GetDegats())
			fmt.Println(rs.GetName(), "attack", enemy.GetName())
			return
		}

	}
}
