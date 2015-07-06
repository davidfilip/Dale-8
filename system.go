package main

const (
	cpuFrequency     = 60
	numRegisters     = 16
	ramCapacity      = 4096
	programAreaStart = 0x200
	screenWidth      = 64
	screenHeight     = 32
	pixelSize        = 10
)

var (
	memory [ramCapacity]byte
	gfx    [screenWidth * screenHeight]byte
	stack  [16]byte
	V      [numRegisters]byte
	I      uint16
	PC     uint16
	SP     byte
	key    [16]byte

	DelayTimer byte
	SoundTimer byte
)
