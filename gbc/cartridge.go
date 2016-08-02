package gbc

import "bytes"

type Cartridge struct {
	rom []byte
}

func NewCartridge(rom []byte) *Cartridge {
	cart := new(Cartridge)
	cart.rom = rom

	return cart
}

func (cart *Cartridge) getRom() []byte {
	return cart.rom
}

func (cart *Cartridge) GetName() string {
	nameBytes := cart.rom[0x134:0x142]
	i := bytes.IndexByte(nameBytes, 0)

	return string(nameBytes[:i])
}

func (cart *Cartridge) GetCGBFlag() int {
	return int(cart.rom[0x143])
}

func (cart *Cartridge) GetSGBFlag() int {
	return int(cart.rom[0x146])
}
