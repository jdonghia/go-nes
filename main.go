package main

import (
	"fmt"
	"time"
)

type Nes struct {
	cpu Cpu
	bus Bus
	RAM [0xffff]uint8
}

type Bus struct {
	write func(RAM *[0xffff]uint8, data uint8, addr uint16)
	read  func(RAM *[0xffff]uint8, addr uint16) uint8
}

type Cpu struct {
	pc uint16
	sp uint16
	a  uint8
	x  uint8
	y  uint8
	sr uint8
}

const (
	FLAGS = iota
	C     = (1 << 0)
	Z     = (1 << 1)
	I     = (1 << 2)
	D     = (1 << 3)
	B     = (1 << 4)
	U     = (1 << 5)
	V     = (1 << 6)
	N     = (1 << 7)
)

func write(RAM *[0xffff]uint8, data uint8, addr uint16) {
	RAM[addr] = data
}

func read(RAM *[0xffff]uint8, addr uint16) uint8 {
	return RAM[addr]
}

func main() {
	nes := Nes{
		cpu: Cpu{},
		bus: Bus{
			write: write,
			read:  read,
		},
		RAM: [0xffff]uint8{}}

	// nes.bus.write(&nes.RAM, 0x42, 0x1234)

	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			clock(&nes)
		}
	}

}

func clock(nes *Nes) {
	nes.cpu.a++
	fmt.Printf("0x%04x\n", nes.cpu.a)
	println(nes.cpu.a)
}

//Adressing Modes
func IMP() uint8 { return 0 }
func IMM() uint8 { return 0 }
func ZP0() uint8 { return 0 }
func ZPX() uint8 { return 0 }
func ZPY() uint8 { return 0 }
func REL() uint8 { return 0 }
func ABS() uint8 { return 0 }
func ABX() uint8 { return 0 }
func ABY() uint8 { return 0 }
func IND() uint8 { return 0 }
func IZX() uint8 { return 0 }
func IZY() uint8 { return 0 }

// Instructions
func ADC() uint8 { return 0 }
func AND() uint8 { return 0 }
func ASL() uint8 { return 0 }
func BCC() uint8 { return 0 }
func BCS() uint8 { return 0 }
func BEQ() uint8 { return 0 }
func BIT() uint8 { return 0 }
func BMI() uint8 { return 0 }
func BNE() uint8 { return 0 }
func BPL() uint8 { return 0 }
func BRK() uint8 { return 0 }
func BVC() uint8 { return 0 }
func BVS() uint8 { return 0 }
func CLC() uint8 { return 0 }
func CLD() uint8 { return 0 }
func CLI() uint8 { return 0 }
func CLV() uint8 { return 0 }
func CMP() uint8 { return 0 }
func CPX() uint8 { return 0 }
func CPY() uint8 { return 0 }
func DEC() uint8 { return 0 }
func DEX() uint8 { return 0 }
func DEY() uint8 { return 0 }
func EOR() uint8 { return 0 }
func INC() uint8 { return 0 }
func INX() uint8 { return 0 }
func INY() uint8 { return 0 }
func JMP() uint8 { return 0 }
func JSR() uint8 { return 0 }
func LDA() uint8 { return 0 }
func LDX() uint8 { return 0 }
func LDY() uint8 { return 0 }
func LSR() uint8 { return 0 }
func NOP() uint8 { return 0 }
func ORA() uint8 { return 0 }
func PHA() uint8 { return 0 }
func PHP() uint8 { return 0 }
func PLA() uint8 { return 0 }
func PLP() uint8 { return 0 }
func ROL() uint8 { return 0 }
func ROR() uint8 { return 0 }
func RTI() uint8 { return 0 }
func RTS() uint8 { return 0 }
func SBC() uint8 { return 0 }
func SEC() uint8 { return 0 }
func SED() uint8 { return 0 }
func SEI() uint8 { return 0 }
func STA() uint8 { return 0 }
func STX() uint8 { return 0 }
func STY() uint8 { return 0 }
func TAX() uint8 { return 0 }
func TAY() uint8 { return 0 }
func TSX() uint8 { return 0 }
func TXA() uint8 { return 0 }
func TXS() uint8 { return 0 }
func TYA() uint8 { return 0 }

func XXX() uint8 { return 0 }
