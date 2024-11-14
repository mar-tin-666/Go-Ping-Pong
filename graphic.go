package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) Draw(screen *ebiten.Image) {
	// Drawing the table and grid lines
	screen.Fill(backgroundColor)
	lineImg := ebiten.NewImage(1, screenHeight)
	lineImg.Fill(lineColor)
	lineOpts := &ebiten.DrawImageOptions{}
	lineOpts.GeoM.Translate(float64(screenWidth)/2, 0)
	screen.DrawImage(lineImg, lineOpts)

	// Drawing palettes
	leftPaddleImg := ebiten.NewImage(paddleWidth, paddleHeight)
	leftPaddleImg.Fill(leftPaddleColor)
	leftPaddleOpts := &ebiten.DrawImageOptions{}
	leftPaddleOpts.GeoM.Translate(0, float64(leftPaddleY))
	screen.DrawImage(leftPaddleImg, leftPaddleOpts)

	rightPaddleImg := ebiten.NewImage(paddleWidth, paddleHeight)
	rightPaddleImg.Fill(rightPaddleColor)
	rightPaddleOpts := &ebiten.DrawImageOptions{}
	rightPaddleOpts.GeoM.Translate(float64(screenWidth-paddleWidth), float64(rightPaddleY))
	screen.DrawImage(rightPaddleImg, rightPaddleOpts)

	// Drawing ball
	vector.DrawFilledCircle(screen, float32(ballX), float32(ballY), float32(ballRadius), ballColor, true)

	// Displaying scores
	scoreText := fmt.Sprintf("%d : %d", leftScore, rightScore)
	scoreX := (screenWidth - len(scoreText)*6) / 2
	ebitenutil.DebugPrintAt(screen, scoreText, scoreX, 20)

	// Welcome screen, point message, and serve message
	if showStartScreen {
		ebitenutil.DebugPrintAt(screen, titleText, screenWidth/2-len(titleText)*3, screenHeight/2-20)
		ebitenutil.DebugPrintAt(screen, startText, screenWidth/2-len(startText)*3, screenHeight/2+10)
		ebitenutil.DebugPrintAt(screen, controlText, screenWidth/2-len(controlText)*3, screenHeight/2+40)
	} else if showPointScreen {
		playerText := leftPlayerText
		if serveLeft {
			playerText = rightPlayerText
		}
		pointText := fmt.Sprintf(pointTextFormat, playerText)
		ebitenutil.DebugPrintAt(screen, pointText, screenWidth/2-len(pointText)*3, screenHeight/2)
	} else if waitingForServe {
		ebitenutil.DebugPrintAt(screen, serveText, screenWidth/2-len(serveText)*3, screenHeight/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
