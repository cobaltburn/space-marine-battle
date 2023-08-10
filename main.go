package main

import (
	"log"
	"math/rand"
	"reflect"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	HEIGHT = 44
	WIDTH  = 60
)

const (
	TotalTatical = 10
	TotalSniper  = 5

	TotalTermagants = 25
	TotalLeapers    = 5
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("space marine battle")
	marines, tyranids := placeCharacters()
	if err := ebiten.RunGame(&Game{marines: marines, tyranids: tyranids, height: HEIGHT, width: WIDTH}); err != nil {
		log.Fatal(err)
	}
}

func placeCharacters() ([]Unit, []Unit) {
	spaceMarines := make([]Unit, 0)
	points := make([][]int, 0)
	for len(points) < TotalTatical+TotalSniper {
		x := rand.Intn(WIDTH) / 8
		y := rand.Intn(HEIGHT)
		pt := []int{x, y}
		if !containsPoint(points, pt) {
			points = append(points, pt)
			var marine Unit
			if len(spaceMarines) < TotalTatical {
				marine = GenerateUnit(Tactical, pt[0], pt[1])
			} else {
				marine = GenerateUnit(Sniper, pt[0], pt[1])
			}
			spaceMarines = append(spaceMarines, marine)
		}
	}
	points = points[:0]
	tyranids := make([]Unit, 0)
	for len(points) < TotalTermagants+TotalLeapers {
		x := HEIGHT - (rand.Intn(WIDTH) / 8)
		y := rand.Intn(HEIGHT)
		pt := []int{x, y}
		if !containsPoint(points, pt) {
			points = append(points, pt)
			var tyranid Unit
			if len(tyranids) < TotalTermagants {
				tyranid = GenerateUnit(Termagant, pt[0], pt[1])
			} else {
				tyranid = GenerateUnit(Leaper, pt[0], pt[1])
			}
			tyranids = append(tyranids, tyranid)
		}
	}
	return spaceMarines, tyranids
}

func containsPoint(pts [][]int, pt []int) bool {
	for _, p := range pts {
		if reflect.DeepEqual(pt, p) {
			return true
		}
	}
	return false
}
