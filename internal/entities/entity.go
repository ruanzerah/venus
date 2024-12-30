package entities

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity interface {
	Attack()
	Heal()
	Stats()
	Walk()
}

type CharacterVocation string

const (
	Mage    CharacterVocation = "Mage"
	Warrior CharacterVocation = "Warrior"
	Druid   CharacterVocation = "Druid"
	Hunter  CharacterVocation = "Hunter"
)

type Player struct {
	Health   float32
	Mana     float32
	Rage     float32
	Vocation CharacterVocation
	Model    rl.Model
	Vec      rl.Vector3
}

func (p *Player) Render() {
}

func (p *Player) Walk() {
}

func (p *Player) Stats() {
}

func (p *Player) Heal() {
}

func (p *Player) Attack() {
}
