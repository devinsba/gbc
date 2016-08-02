package main

import (
	"github.com/ventu-io/slog"
	"github.com/ventu-io/slog/basic"
	"github.com/ventu-io/slf"
	"github.com/devinsba/gbc-go/gbc"
	"io/ioutil"
)

func main() {
	appender := basic.New()
	appender.SetTemplate(basic.StandardTermTemplate)


	logger := slog.New()
	logger.SetLevel(slf.LevelDebug)
	logger.AddEntryHandler(appender)
	slf.Set(logger)

	rom, err := ioutil.ReadFile("/Users/devinsba/Downloads/PokemonCrystal.gbc")
	if err != nil {
		panic("Couldn't read ROM")
	}
	cart := gbc.NewCartridge(rom)

	gameboy := new(gbc.GameboyColor)

	logger.WithContext("main").Debugf("%v", AssetNames())
	boot, err := Asset("data/gbc_bios.bin")
	if err == nil {
		gameboy.WithBootRom(boot)
	}

	gameboy.InsertCartridge(cart)
	gameboy.Start()
}
