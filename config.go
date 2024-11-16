package main

import (
	"image/color"
	"time"

	"golang.org/x/exp/rand"
)

const (
	screenWidth      = 640
	screenHeight     = 480
	paddleWidth      = 10
	paddleHeight     = 80
	ballRadius       = 5
	paddleSpeed      = 5
	initialBallSpeed = 300
	minBallSpeedY    = 1
	ballAcceleration = 50
)

var (
	// Random
	random = rand.New(rand.NewSource(uint64(time.Now().UnixNano())))

	// Texts
	titleText       = "Go Ping Pong"
	startText       = "Press SPACE to Start"
	controlText     = "Left: W/S | Right: Up/Down | Reset: R"
	pointTextFormat = "Point for %s! Press SPACE to continue."
	serveText       = "Press SPACE to serve. The ball must touch the paddle."
	leftPlayerText  = "Left Player"
	rightPlayerText = "Right Player"

	// Initial config
	leftScore, rightScore  int
	ballSpeedX, ballSpeedY = initialBallSpeed, initialBallSpeed
	gameStarted            bool
	showPointScreen        bool
	showStartScreen        = true
	waitingForServe        = true
	serveLeft              = random.Intn(2) == 0
	spaceKeyWasPressed     bool

	leftPaddleY  = (screenHeight - paddleHeight) / 2
	rightPaddleY = (screenHeight - paddleHeight) / 2
	ballX, ballY = screenWidth/2 - ballRadius/2, screenHeight/2 - ballRadius/2

	previousTime time.Time

	// Colors
	backgroundColor  = color.RGBA{35, 140, 35, 255}   // Gren table
	leftPaddleColor  = color.RGBA{0, 0, 255, 255}     // Blue paddle
	rightPaddleColor = color.RGBA{255, 0, 0, 255}     // Red paddle
	ballColor        = color.RGBA{255, 255, 255, 255} // White ball
	lineColor        = color.RGBA{255, 255, 255, 255} // White line

)
