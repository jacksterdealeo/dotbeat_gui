package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"math"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Width  = 480
	Height = 480
	scale  = 64
)

var (
	clockCircle  *ebiten.Image
	clockHand    *ebiten.Image
	thisMoment           = time.Now()
	handRotation float64 = 0
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	now := time.Now()
	midnight := time.Date(now.Year(), now.Month(), now.Day(), -1, 0, 0, 0, time.UTC)
	secondsFromMidnight := float64(now.Sub(midnight).Seconds())
	dotbeatTime := math.Mod((secondsFromMidnight / 86.4), 1000)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("d%d.%s\n@%0.2f\n", midnight.Day(), midnight.Month().String(), dotbeatTime))

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(clockCircle, op)

	op = &ebiten.DrawImageOptions{}
	handRotation = dotbeatTime*(0.0062831853) + (3.1415926536)
	op.GeoM.Rotate(handRotation)
	op.GeoM.Scale(0.25, 0.25)
	op.GeoM.Translate(float64(screen.Bounds().Dx()/2), float64(screen.Bounds().Dy()/2))

	// degrees = radians * (180/pi)
	// radians = degrees * (pi/180)
	screen.DrawImage(clockHand, op)

	// TODO: pin the hands directly to the center.
	op = &ebiten.DrawImageOptions{}
	handRotation *= 100
	op.GeoM.Rotate(handRotation)
	op.GeoM.Scale(0.5, 0.40)
	op.GeoM.Translate(float64(screen.Bounds().Dx()/2), float64(screen.Bounds().Dy()/2))
	screen.DrawImage(clockHand, op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}

func main() {

	circleFile, err := os.ReadFile("circle.png")
	if err != nil {
		log.Fatal(err)
	}
	lineFile, err := os.ReadFile("line.png")
	if err != nil {
		log.Fatal(err)
	}
	// Decode an image from the image file's byte slice.
	img_circleFile, _, err := image.Decode(bytes.NewReader(circleFile))
	if err != nil {
		log.Fatal(err)
	}
	clockCircle = ebiten.NewImageFromImage(img_circleFile)

	img_lineFile, _, err := image.Decode(bytes.NewReader(lineFile))
	if err != nil {
		log.Fatal(err)
	}
	clockHand = ebiten.NewImageFromImage(img_lineFile)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Dotbeat Clock")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
