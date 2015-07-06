package main

import (
	"github.com/go-gl/glfw/v3.1/glfw"
)

func initGraphics() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(screenWidth*pixelSize, screenHeight*pixelSize, "Dale: Chip-8 Emulator", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
