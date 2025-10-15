package game

import (
	"spacegame/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image            *ebiten.Image
	position         Vector
	game             *Game
	laserLoadingTime *Timer
}

func NewPlayer(game *Game) *Player {
	image := assets.PlayerSprite
	bounds := image.Bounds()
	halfWidth := float64(bounds.Dx()) / 2

	position := Vector{
		X: (screenWidth / 2) - halfWidth,
		Y: 500,
	}

	return &Player{
		image:            image,
		game:             game,
		position:         position,
		laserLoadingTime: NewTimer(12),
	}
}

func (p *Player) Update() {
	speed := 7.0

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= speed
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += speed
	}

	p.laserLoadingTime.Update()
	if ebiten.IsKeyPressed(ebiten.KeySpace) && p.laserLoadingTime.IsReady() {
		p.laserLoadingTime.Reset()
		bounds := p.image.Bounds()
		halfWidth := float64(bounds.Dx()) / 2
		halfHeight := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			p.position.X + halfWidth,
			p.position.Y - halfHeight/2,
		}

		laser := NewLaser(spawnPos)
		p.game.AddLasers(laser)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.image, op)
}
