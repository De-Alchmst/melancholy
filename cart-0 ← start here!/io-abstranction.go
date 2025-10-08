package main

import "cart/w4"

// We will do what's called a "pro gamer move"
import "math/rand"

type InputKey int
const (
	KeyUp InputKey = iota
	KeyRight
	KeyDown
	KeyLeft
	KeyX
	KeyZ
	KayAny
	KeyAction
	KeyMovement
)

var (
	prevPressedKeys map[InputKey]bool = make(map[InputKey]bool)
	randomSeed int64 = 37 // the most random of numbers
)

func Held(key InputKey) bool {
	var gamepad = *w4.GAMEPAD1
	// we are about to do what's called a pro gamer move
	randomSeed += int64(gamepad)

	// Yandere dev in shambels
	switch key {
	case KeyUp:
		return gamepad&w4.BUTTON_UP != 0
	case KeyRight:
		return gamepad&w4.BUTTON_RIGHT != 0
	case KeyDown:
		return gamepad&w4.BUTTON_DOWN != 0
	case KeyLeft:
		return gamepad&w4.BUTTON_LEFT != 0
	case KeyX:
		return gamepad&w4.BUTTON_1 != 0
	case KeyZ:
		return gamepad&w4.BUTTON_2 != 0
	case KayAny:
		return gamepad != 0
	case KeyAction:
		return gamepad&(w4.BUTTON_1|w4.BUTTON_2) != 0
	case KeyMovement:
		return gamepad&(w4.BUTTON_UP|w4.BUTTON_DOWN|w4.BUTTON_LEFT|w4.BUTTON_RIGHT) != 0
	default:
		return false
	}
	// probably...
}


func Pressed(key InputKey) bool {
	return Held(key) && !prevPressedKeys[key]
}


func UpdatePressed() {
	for _, key := range [...]InputKey{KeyUp, KeyRight, KeyDown, KeyLeft, KeyX, KeyZ, KayAny, KeyAction, KeyMovement} {
		prevPressedKeys[key] = Held(key)
	}
}


func GetRandomN(n int) int {
	Rnd := rand.New(rand.NewSource(randomSeed))
	randomSeed += int64(rand.Int())
	return Rnd.Intn(n)
}
