package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func (g *Game) Update() error {
	// Time control for animation
	currentTime := time.Now()
	deltaTime := currentTime.Sub(previousTime).Seconds()
	previousTime = currentTime

	// Screen selector
	if ebiten.IsKeyPressed(ebiten.KeySpace) && !spaceKeyWasPressed {
		spaceKeyWasPressed = true
	} else if !ebiten.IsKeyPressed(ebiten.KeySpace) && spaceKeyWasPressed {
		spaceKeyWasPressed = false
		if showStartScreen {
			showStartScreen = false
			resetPaddlesAndBall()
		} else if showPointScreen {
			showPointScreen = false
			waitingForServe = true
			resetPaddlesAndBall()
		} else if waitingForServe {
			if paddleTouchesBall() {
				waitingForServe = false
				gameStarted = true
				randomizeBallDirection()
			}
		}
	}

	//R to press to reset game
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		gameStarted = false
		showStartScreen = true
		waitingForServe = true
	}

	//Paddle control
	if !showStartScreen && !showPointScreen {
		if ebiten.IsKeyPressed(ebiten.KeyW) && leftPaddleY > 0 {
			leftPaddleY -= paddleSpeed
		} else if ebiten.IsKeyPressed(ebiten.KeyS) && leftPaddleY < screenHeight-paddleHeight {
			leftPaddleY += paddleSpeed
		}
		if ebiten.IsKeyPressed(ebiten.KeyUp) && rightPaddleY > 0 {
			rightPaddleY -= paddleSpeed
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) && rightPaddleY < screenHeight-paddleHeight {
			rightPaddleY += paddleSpeed
		}
	}

	// Game logic
	if gameStarted {
		// Ball movement
		ballX += int(float64(ballSpeedX) * deltaTime)
		ballY += int(float64(ballSpeedY) * deltaTime)

		// Collision with walls
		if ballY <= 0 || ballY+ballRadius >= screenHeight {
			ballSpeedY = -ballSpeedY
		}

		// Collision with the left paddle
		if ballX <= paddleWidth && ballY+ballRadius >= leftPaddleY && ballY <= leftPaddleY+paddleHeight {
			ballSpeedX = int(math.Abs(float64(ballSpeedX)))
			adjustBallSpeed(leftPaddleY)
		}

		// Collision with the right paddle
		if ballX+ballRadius >= screenWidth-paddleWidth && ballY+ballRadius >= rightPaddleY && ballY <= rightPaddleY+paddleHeight {
			ballSpeedX = -int(math.Abs(float64(ballSpeedX)))
			adjustBallSpeed(rightPaddleY)
		}

		// Scoring
		if ballX < 0 {
			rightScore++
			serveLeft = true
			showPointScreen = true
			gameStarted = false
		} else if ballX > screenWidth {
			leftScore++
			serveLeft = false
			showPointScreen = true
			gameStarted = false
		}
	}

	return nil
}

// Game reset
func resetPaddlesAndBall() {
	leftPaddleY, rightPaddleY = (screenHeight-paddleHeight)/2, (screenHeight-paddleHeight)/2
	resetBallForServe()
}

// Setting the ball for serve
func resetBallForServe() {
	if !serveLeft {
		ballX = paddleWidth + ballRadius
		ballY = leftPaddleY + paddleHeight/2
	} else {
		ballX = screenWidth - paddleWidth - ballRadius
		ballY = rightPaddleY + paddleHeight/2
	}
	ballSpeedX, ballSpeedY = initialBallSpeed, initialBallSpeed
	gameStarted = false
}

// Checking if the ball is touching the paddle
func paddleTouchesBall() bool {
	if serveLeft {
		return ballX+ballRadius >= screenWidth-paddleWidth && ballY+ballRadius >= rightPaddleY && ballY-ballRadius <= rightPaddleY+paddleHeight
	}
	return ballX-ballRadius <= paddleWidth && ballY+ballRadius >= leftPaddleY && ballY-ballRadius <= leftPaddleY+paddleHeight
}

// Increase ball speed
func adjustBallSpeed(paddleY int) {
	distFromCenter := (ballY) - (paddleY + paddleHeight/2)
	ballSpeedY += int(distFromCenter / 10)

	if ballSpeedX < 0 {
		ballSpeedX -= 1
	} else {
		ballSpeedX += 1
	}

	if math.Abs(float64(ballSpeedY)) < minBallSpeedY {
		ballSpeedY = minBallSpeedY * int(math.Copysign(1, float64(ballSpeedY)))
	}
}

// Randomize ball direction
func randomizeBallDirection() {
	ballSpeedX = initialBallSpeed * (2*random.Intn(2) - 1)
	ballSpeedY = initialBallSpeed * (2*random.Intn(2) - 1)
}
