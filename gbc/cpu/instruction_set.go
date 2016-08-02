package cpu

type instruction struct {
	opcode          byte
	instructionSize uint16
	cpuCycles       int
	method          InstructionMethod
}

type InstructionMethod func(GameboyCpu, instruction) uint16

var instructionSet = map[byte]instruction{
	0x00: instruction{0x00, 1, 4, nop},
	0x01: instruction{0x01, 3, 12, loadImmediateTo16BitReg},
	0x02: instruction{0x02, 2, 8, loadAToAddress},
	0x0F: instruction{0x0F, 1, 1, rotateRightCarryA},
	0x11: instruction{0x11, 3, 12, loadImmediateTo16BitReg},
	0x20: instruction{0x20, 2, 8, jumpImmediateConditional},
	0x28: instruction{0x28, 2, 8, jumpImmediateConditional},
	0x30: instruction{0x30, 2, 8, jumpImmediateConditional},
	0x31: instruction{0x31, 3, 12, loadImmediateTo16BitReg},
	0x38: instruction{0x38, 2, 8, jumpImmediateConditional},
	0x3E: instruction{0x3E, 2, 8, loadValueIntoA},
	0xAF: instruction{0xAF, 2, 4, xor},
	0xC3: instruction{0xC3, 3, 12, jumpDirect},
	0xE0: instruction{0xE0, 2, 12, loadAToHighRam},
	0xF3: instruction{0xF3, 1, 4, disableInterupts},
	0xFE: instruction{0xFE, 2, 8, compareA},
}
