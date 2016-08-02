package cpu

import "encoding/binary"

func compareA(cpu GameboyCpu, inst instruction) uint16 {
	cpu.setFlagN(true)
	if cpu.getA() == cpu.rom[cpu.getPC()+1] {
		cpu.setFlagZ(true)
	}
	// TODO figure out Flags H and C for subtraction i guess?

	return cpu.getPC() + inst.instructionSize
}

func nop(cpu GameboyCpu, inst instruction) uint16 {
	return cpu.getPC() + inst.instructionSize
}

func jumpDirect(cpu GameboyCpu, inst instruction) uint16 {
	addressLow := cpu.rom[cpu.getPC()+1]
	addressHigh := cpu.rom[cpu.getPC()+2]
	newPc := binary.LittleEndian.Uint16([]byte{addressLow, addressHigh})
	return newPc
}

func jumpImmediateConditional(cpu GameboyCpu, inst instruction) uint16 {
	switch inst.opcode {
	case 0x20:
	case 0x28:
		logger().Debugf("JUMP (+ %x) if Z set", cpu.rom[cpu.getPC()+1])
		if cpu.getFlagZ() {
			logger().Debug("Z is set")
			return cpu.getPC() + uint16(cpu.rom[cpu.getPC()+1])
		}
	case 0x30:
	case 0x38:
	}
	return cpu.getPC() + inst.instructionSize
}

func xor(cpu GameboyCpu, inst instruction) uint16 {
	switch inst.opcode {
	case 0xAF:
		cpu.setA(0)
		cpu.setFlagZ(true)
		cpu.setFlagN(false)
		cpu.setFlagH(false)
		cpu.setFlagC(false)
	}
	return cpu.getPC() + inst.instructionSize
}

func loadAToAddress(cpu GameboyCpu, inst instruction) uint16 {
	switch inst.opcode {
	case 0x02:
		addr := cpu.getBC() + 0xFF00
		logger().Debugf("Setting mem %x to A [%x]", addr, cpu.getA())
		cpu.rom[addr] = cpu.getA()
	}
	return cpu.getPC() + inst.instructionSize
}

func loadAToHighRam(cpu GameboyCpu, inst instruction) uint16 {
	switch inst.opcode {
	case 0xE0:
		var memLocation int = 0xFF00 + int(cpu.rom[cpu.getPC()+1])
		cpu.rom[memLocation] = cpu.getA()
	}
	return cpu.getPC() + inst.instructionSize
}

func loadImmediateTo16BitReg(cpu GameboyCpu, inst instruction) uint16 {
	addressLow := cpu.rom[cpu.getPC()+1]
	addressHigh := cpu.rom[cpu.getPC()+2]
	newVal := binary.LittleEndian.Uint16([]byte{addressLow, addressHigh})
	switch inst.opcode {
	case 0x01:
		cpu.setBC(newVal)
		break
	case 0x31:
		logger().Debugf("SP to become 0x%x", newVal)
		cpu.setSP(newVal)
		break
	}
	return cpu.getPC() + inst.instructionSize
}

func loadValueIntoA(cpu GameboyCpu, inst instruction) uint16 {
	switch inst.opcode {
	case 0x3E:
		cpu.setA(cpu.rom[cpu.getPC()+1])
	}
	return cpu.getPC() + inst.instructionSize
}

func disableInterupts(cpu GameboyCpu, inst instruction) uint16 {
	// TODO interupts
	return cpu.getPC() + inst.instructionSize
}

func rotateRightCarryA(cpu GameboyCpu, inst instruction) uint16 {
	a := cpu.getA()
	bit0 := a&1 == 1
	a = a >> 1
	cpu.setA(a)
	if a == 0 {
		cpu.setFlagZ(true)
	} else {
		// TODO maybe?
		//cpu.setFlagZ(false)
	}
	cpu.setFlagN(false)
	cpu.setFlagH(false)
	cpu.setFlagC(bit0)

	return cpu.getPC() + inst.instructionSize
}
