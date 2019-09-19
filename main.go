package main

import (
	"image/color"

	"github.com/oakmound/oak"
	// "github.com/oakmound/oak/collision"
	"github.com/oakmound/oak/entities"
	"github.com/oakmound/oak/event"
	"github.com/oakmound/oak/render"
	"github.com/oakmound/oak/scene"
)

var (
	score1 = 0
	score2 = 0
)

const (
	windowWidth = 640
	windowHeight = 480
	// playerLabel collision.Label = 1
)

func main() {
	const sceneName = "game_scene"
	oak.Add(sceneName,
		func(prevScene string, data interface{}) {
			newPlayer()
		}, func() bool { return true },
		func() (string, *scene.Result) { return sceneName, nil })
	oak.Init(sceneName)
}

func newPlayer() {
	const width = 20
	const height = 60
	var gray = color.RGBA{190, 190, 190, 255}
	var xPos float64 = (windowWidth - width) / 2
	var yPos float64 = windowHeight - height - 20

	playerCharacter := entities.NewMoving(xPos, yPos, width, height, render.NewColorBox(width, height, gray), nil, 0, 0)
	playerCharacter.Speed.SetY(4)
	playerCharacter.Speed.SetX(4)
	render.Draw(playerCharacter.R, 1)
	// playerCharacter.Space.UpdateLabel(playerLabel)
	playerCharacter.Bind(inputBindings("UpArrow", "DownArrow", "LeftArrow", "RightArrow"), event.Enter)
	playerCharacter.SetPos(xPos, yPos)
}

func inputBindings(up string, down string, left string, right string) func(int, interface{}) int {
	return func(id int, nothing interface{}) int {
		playerCharacter := event.GetEntity(id).(*entities.Moving)
		playerCharacter.Delta.SetY(0)
		playerCharacter.Delta.SetX(0)

		if oak.IsDown(up) {
			playerCharacter.Delta.SetY(-playerCharacter.Speed.Y())
		} else if oak.IsDown(down) {
			playerCharacter.Delta.SetY(playerCharacter.Speed.Y())
		}
		playerCharacter.ShiftY(playerCharacter.Delta.Y())
		if playerCharacter.Y() < 0 || playerCharacter.Y() > (windowHeight-playerCharacter.H) {
			playerCharacter.ShiftY(-playerCharacter.Delta.Y())
		}

		if oak.IsDown(left) {
			playerCharacter.Delta.SetX(-playerCharacter.Speed.X())
		} else if oak.IsDown(right) {
			playerCharacter.Delta.SetX(playerCharacter.Speed.X())
		}
		playerCharacter.ShiftX(playerCharacter.Delta.X())
		if playerCharacter.X() < 0 || playerCharacter.X() > (windowWidth-playerCharacter.W) {
			playerCharacter.ShiftX(-playerCharacter.Delta.X())
		}

		return 0
	}
}