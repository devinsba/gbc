package gbc

import (
	"time"

	"github.com/devinsba/gbc-go/gbc/cpu"
)

type GameboyColor struct {
	cartridge Cartridge
	bootRom   []byte
}

func (gb *GameboyColor) WithBootRom(bootRom []byte) {
	gb.bootRom = bootRom
}

func (gb *GameboyColor) InsertCartridge(cartridge *Cartridge) {
	gb.cartridge = *cartridge

	logger().Debugf("Game name %s", gb.cartridge.GetName())
	logger().Debugf("Cartridge flags [GBC: %x] [SGB: %x]", gb.cartridge.GetCGBFlag(), gb.cartridge.GetSGBFlag())
}

func (gb *GameboyColor) Start() {
	cpu := cpu.InitGameboyCpu(gb.bootRom, gb.cartridge.getRom())

	go cpu.Run()

	for {
		time.Sleep(1000 * time.Millisecond)
	}
}
