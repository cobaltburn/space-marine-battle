package main

import (
	"image/color"
	"math"
	"math/rand"
)

type army int
type unitType int

const (
	SpaceMarine army = iota
	Tyranid
)

const (
	Termagant unitType = iota
	Leaper
	Tactical
	Sniper
)

const (
	TermagantRange    = 6
	TermagantHp       = 5
	TermagantDMG      = 3
	TermagantWS       = 4
	TermagantMovement = 3

	LeaperRange    = 2
	LeaperHp       = 7
	LeaperDMG      = 4
	LeaperWS       = 2
	LeaperMovement = 5

	TacticalRange    = 12
	TacticalHp       = 9
	TacticalDMG      = 5
	TacticalWS       = 3
	TacticalMovement = 3

	SniperRange    = 18
	SniperHp       = 6
	SniperDMG      = 7
	SniperWS       = 2
	SniperMovement = 3
)

var (
	termColor   = color.RGBA{218, 36, 246, 255}
	leaperColor = color.RGBA{190, 48, 213, 255}
	assultColor = color.RGBA{48, 125, 213, 255}
	sniperColor = color.RGBA{0, 84, 255, 255}
)

type Unit struct {
	army army
	unit unitType
	x    int
	y    int
	hp   int
}

func (unit Unit) GetPoint() (int, int) {
	return unit.x, unit.y
}

func (unit Unit) GetColor() color.Color {
	switch u := unit.unit; {
	case u == Termagant:
		return termColor
	case u == Leaper:
		return leaperColor
	case u == Tactical:
		return assultColor
	case u == Sniper:
		return sniperColor
	default:
		panic("invalid unit type")
	}
}

func (unit *Unit) Move(targets []Unit, allies []Unit) {
	tar := unit.findTarget(targets)
	if unit.distance(*tar) > unit.AttackRange() {
		step := unit.stepOptions(*tar, append(targets, allies...))
		unit.x += step[0]
		unit.y += step[1]
	}
}

func (unit Unit) findTarget(targets []Unit) *Unit {
	idx := 0
	for i, target := range targets {
		dist := unit.distance(target)
		curDist := unit.distance(targets[idx])
		if curDist > dist {
			idx = i
		}
	}
	return &targets[idx]
}

func (unit Unit) distance(target Unit) float64 {
	x1, y1 := unit.GetPoint()
	x2, y2 := target.GetPoint()
	x, y := float64(x2-x1), float64(y2-y1)
	sqrX, sqrY := math.Pow(x, 2.0), math.Pow(y, 2.0)
	return math.Sqrt(sqrX + sqrY)
}

func (unit Unit) stepOptions(target Unit, units []Unit) []int {
	spaces := make([][]int, 0)
	for x := -unit.Movement(); x <= unit.Movement(); x++ {
		for y := -unit.Movement(); y <= unit.Movement(); y++ {
			spaces = append(spaces, []int{x, y})
		}
	}
	current := unit.distance(target)
	open := []int{0, 0}
	for _, s := range spaces {
		distance := Unit{x: unit.x + s[0], y: unit.y + s[1]}.distance(target)
		if current > distance && checkOpen(unit.x+s[0], unit.y+s[1], units) {
			current = distance
			open = s
		}
	}
	return open
}

func (unit Unit) AttackRange() float64 {
	switch {
	case unit.unit == Termagant:
		return TermagantRange
	case unit.unit == Leaper:
		return LeaperRange
	case unit.unit == Tactical:
		return TacticalRange
	case unit.unit == Sniper:
		return SniperRange
	default:
		panic("invalid unit type")
	}
}

func (unit Unit) Damage() int {
	switch {
	case unit.unit == Termagant:
		return TermagantDMG
	case unit.unit == Leaper:
		return LeaperDMG
	case unit.unit == Tactical:
		return TacticalDMG
	case unit.unit == Sniper:
		return SniperDMG
	default:
		panic("invalid unit type")
	}
}

func (unit Unit) WeaponSkill() int {
	switch {
	case unit.unit == Termagant:
		return TermagantWS
	case unit.unit == Leaper:
		return LeaperWS
	case unit.unit == Tactical:
		return TacticalWS
	case unit.unit == Sniper:
		return SniperWS
	default:
		panic("invalid unit type")
	}
}

func (unit Unit) Movement() int {
	switch {
	case unit.unit == Termagant:
		return TermagantMovement
	case unit.unit == Leaper:
		return LeaperMovement
	case unit.unit == Tactical:
		return TacticalMovement
	case unit.unit == Sniper:
		return SniperMovement
	default:
		panic("invalid unit type")
	}
}

func (unit Unit) Attack(targets []Unit) {
	attack := rand.Intn(6) + 1
	if !(attack >= unit.WeaponSkill()) {
		return
	}
	tar := unit.findTarget(targets)
	dist := unit.distance(*tar)
	if dist <= unit.AttackRange() {
		dmg := rand.Intn(unit.Damage()) + 1
		(*tar).hp -= dmg
	}
}

func checkOpen(x, y int, units []Unit) bool {
	for _, u := range units {
		if u.x == x && u.y == y {
			return false
		}
	}
	return true
}

func GenerateUnit(unit unitType, x, y int) Unit {
	switch {
	case unit == Termagant:
		return Unit{
			Tyranid,
			unit,
			x,
			y,
			TermagantHp,
		}
	case unit == Leaper:
		return Unit{
			Tyranid,
			unit,
			x,
			y,
			LeaperHp,
		}
	case unit == Tactical:
		return Unit{
			SpaceMarine,
			unit,
			x,
			y,
			TacticalHp,
		}
	case unit == Sniper:
		return Unit{
			SpaceMarine,
			unit,
			x,
			y,
			SniperHp,
		}
	default:
		panic("invalid unit type")
	}
}
