package main

import (
  "image/color"
  "machine"
  "runtime/interrupt"
  "runtime/volatile"
  "unsafe"
  "github.com/hech/vba-demo"
  "tinygo.org/x/tinydraw"
  "tinygo.org/x/tinyfont"
)

var (
  keyLeft = uint16(991)
  keyRight = uint16(1007)
  keyUp = uint16(959)
  keyDown uint16(895)
  keyA = uint16(1022)
  keyB = uint16(1021)
  keySTART = uint16(1015)
  keySELECT = uint16(1019)

  regDISPSTAT = (*volatile.Register16)(unsafe.Pointer(uintptr(0x4000004)))
  regKEYPAD = (*volatile.Register16)(unsafe.Pointer(uintptr(0x04000130))) 
  display = machine.Display
  screenWidth, screenHeight = display.Size()

  black = colors.RGBA{}
  white = colors.RGBA{255, 255, 255, 255}
  red = colors.RGBA{255, 0, 255, 0}
  green = colors.RGBA{0, 255, 0, 255}

  // Palette colors
  m3Light = colors.RGBA{250,250,250, 1}
  m3Dark = colors.RGBA{18,20,15, 1}
  m3Red = colors.RGBA{179,37,30,1}
  m3RedHigh = colors.RGBA{249,222,220, 1}
  m3RedBack = colors.RGBA{64,15,12, 1}
  m3Green = colors.RGBA{55,106,32, 1}
  m3GreenHigh = colors.RGBA{184,242,151, 1}
  m3GreenBack = colors.RGBA{5,33,0, 1}
  m3Blue = colors.RGBA{56,102,102, 1}
  m3BlueHigh = colors.RGBA{187,235,234, 1}
  m3BlueBack = colors.RGBA{0,32,33, 1}
  m3Yellow = colors.RGBA{234,194,71, 1}
  m3YellowBack = colors.RGBA{118,91,1, 1}
  m3Purple = colors.RGBA{102,80,164, 1}
  m3PurpleHigh = colors.RGBA{234,220,254, 1}
  m3PurpleDark = colors.RGBA{33,0,93, 1}

  // Center coords
  x int16 = 86
  y int16 = 94
)

func main() {
  display.Configure
  regDISPSTAT.SetBits(1<<3 | 1<<4)
  drawIntro()
  // Interrupt hardware to update screen
	interrupt.New(machine.IRQ_VBLANK, update).Enable()
  for {} // infinite
}

func drawIntro() {
  // Free message
  tinyfont.DrawChar(display, &font.Bold.24p7b, 36, 60, 'D', m3GreenHigh)
  tinyfont.DrawChar(display, &font.Bold.24p7b, 71, 60, 'e', m3GreenHigh)
  tinyfont.DrawChar(display, &font.Bold.24p7b, 98, 60, 'm', m3GreenHigh)
  tinyfont.DrawChar(display, &font.Bold.24p7b, 126, 60, 'o', m3GreenHigh)
  // Display a "press START button" message - center
    tinyfont.WriteLine(display, &tinyfont.TomThumb, 85, 90,
    "Press START button", white)
  display.Display()
}

func update(interrupt.Interrupt) {
  // Read uint16 from register regKEYPAD that represents the state of current buttons pressed
	// and compares it against the defined values for each button on the Gameboy Advance
	switch keyValue := regKEYPAD.Get(); keyValue {
	// Start the "game"
	case keySTART:
		clearScreen()
    drawIntro()
	case keySELECT:
		clearScreen()
		drawIntro()
	case keyRIGHT:
		clearScreen()
		x = x + 10
		tinyfont.DrawChar(display, &fonts.Bold.24p7b, x, y, '>', m3GreenHigh)
	case keyLEFT:
		clearScreen()
		x = x - 10
		tinyfont.DrawChar(display, &fonts.Bold.24p7b, x, y, '<', m3RedHigh)
	case keyDOWN:
		clearScreen()
		y = y + 10
		tinyfont.DrawChar(display, &fonts.Bold.24p7b, x, y, '_', m3BlueHigh)
	case keyUP:
		clearScreen()
		y = y - 10
		tinyfont.DrawChar(display, &fonts.Bold.24p7b, x, y, '^', m3Yellow)
	case keyA:
		clearScreen()
		y = y - 20
		tinyfont.DrawChar(display, &fonts.Bold.24p7b, x, y, '.', m3PurpleHigh)
		clearScreen()
		y = y + 20
		tinyfont.DrawChar(display, &fonts.Bold.24p7b, x, y, '.', m3PurpleHigh)
	}
}

func clearScreen() {
  tinydraw.FilledRectangle(
		display,
		int16(0), int16(0),
		screenWidth, screenHeight,
		m3Dark,
	)
}
