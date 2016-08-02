package cpu

import (
	"time"
	"encoding/binary"
)

func InitGameboyCpu(bootRom []byte, gameRom []byte) *GameboyCpu {
	gameboyCpu := GameboyCpu{rom: []byte{}}
	gameboyCpu.setPC(0x0000)
	if bootRom != nil && len(bootRom) > 0 {
		gameboyCpu.rom = append(gameboyCpu.rom, bootRom...)
	}
	if bootRom != nil {
		gameboyCpu.rom = append(gameboyCpu.rom, gameRom[len(bootRom):]...)
	} else {
		gameboyCpu.rom = append(gameboyCpu.rom, gameRom...)
		gameboyCpu.setPC(0x0100)
		gameboyCpu.setBC(0x0013)
	}

	return &gameboyCpu
}

func (cpu *GameboyCpu) Run() {
	for {
		logger().Debugf("PC %x INST %x", cpu.getPC(), cpu.rom[cpu.getPC()])
		inst := instructionSet[cpu.rom[cpu.getPC()]]
		cpu.setPC(inst.method(*cpu, inst))
		//cpu.pc = exec(cpu.rom, cpu.pc)
		time.Sleep(1000 * time.Millisecond)
	}
}


type GameboyCpu struct{
	// TODO set up different mem areas. pretty sure I shouldn't be overwriting bytes in the game ROM
	rom []byte
}

var cpuRegisters []byte = []byte{
	0x00, 0x00, //A F
	0x00, 0x00, //B C
	0x00, 0x00, //D E
	0x00, 0x00, //H L
	0x00,0x00,  //SP
	0x00,0x00,  //PC
}

const (
	offset_A uint16 = 0
	offset_F uint16 = 1
	offset_B uint16 = 2
	offset_C uint16 = 3
	offset_SP uint16 = 8
	offset_PC uint16 = 10
)

func (cpu *GameboyCpu) setPC(pc uint16) {
	temp := []byte{0x00,0x00}
	binary.LittleEndian.PutUint16(temp, pc)
	cpuRegisters[offset_PC] = temp[0]
	cpuRegisters[offset_PC + 1] = temp[1]
}

func (cpu *GameboyCpu) getPC() uint16 {
	addressLow := cpuRegisters[offset_PC]
	addressHigh := cpuRegisters[offset_PC + 1]
	return binary.LittleEndian.Uint16([]byte{addressLow, addressHigh})
}

func (cpu *GameboyCpu) setSP(pc uint16) {
	temp := []byte{0x00,0x00}
	binary.LittleEndian.PutUint16(temp, pc)
	cpuRegisters[offset_SP] = temp[0]
	cpuRegisters[offset_SP + 1] = temp[1]
}

func (cpu *GameboyCpu) getSP() uint16 {
	addressLow := cpuRegisters[offset_SP]
	addressHigh := cpuRegisters[offset_SP + 1]
	return binary.LittleEndian.Uint16([]byte{addressLow, addressHigh})
}

func (cpu *GameboyCpu) setA(a byte) {
	cpuRegisters[offset_A] = a
}

func (cpu *GameboyCpu) getA() byte  {
	return cpuRegisters[offset_A]
}

func (cpu *GameboyCpu) setB(b byte) {
	cpuRegisters[offset_B] = b
}

func (cpu *GameboyCpu) getB() byte  {
	return cpuRegisters[offset_B]
}

func (cpu *GameboyCpu) setC(c byte) {
	cpuRegisters[offset_C] = c
}

func (cpu *GameboyCpu) getC() byte  {
	return cpuRegisters[offset_C]
}

func (cpu *GameboyCpu) setBC(bc uint16) {
	temp := []byte{0x00,0x00}
	binary.LittleEndian.PutUint16(temp, bc)
	cpuRegisters[offset_B] = temp[0]
	cpuRegisters[offset_C] = temp[1]
}

func (cpu *GameboyCpu) getBC() uint16  {
	addressLow := cpuRegisters[offset_B]
	addressHigh := cpuRegisters[offset_C]
	return binary.LittleEndian.Uint16([]byte{addressLow, addressHigh})
}

func (cpu *GameboyCpu) setFlagZ(val bool) {
	cpuRegisters[offset_F] = setFlag(cpuRegisters[offset_F], val, 7)
}

func (cpu *GameboyCpu) getFlagZ() bool {
	return cpuRegisters[offset_F] >> 7 & 0x01 == 1
}

func (cpu *GameboyCpu) setFlagN(val bool) {
	cpuRegisters[offset_F] = setFlag(cpuRegisters[offset_F], val, 6)
}

func (cpu *GameboyCpu) getFlagN() bool {
	return cpuRegisters[offset_F] >> 6 & 0x01 == 1
}

func (cpu *GameboyCpu) setFlagH(val bool) {
	cpuRegisters[offset_F] = setFlag(cpuRegisters[offset_F], val, 5)
}

func (cpu *GameboyCpu) getFlagH() bool {
	return cpuRegisters[offset_F] >> 5 & 0x01 == 1
}

func (cpu *GameboyCpu) setFlagC(val bool) {
	cpuRegisters[offset_F] = setFlag(cpuRegisters[offset_F], val, 4)
}

func (cpu *GameboyCpu) getFlagC() bool {
	return cpuRegisters[offset_F] >> 4 & 0x01 == 1
}

func setFlag(flags byte, val bool, leftShift uint) byte {
	if val {
		return flags | byte(1 << leftShift)
	} else {
		return flags & (0xFF ^ byte(1 << leftShift))
	}
}