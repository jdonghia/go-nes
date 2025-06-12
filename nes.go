package main

import (
	"fmt"
	"time"
)

type CPU struct {
	PC      uint16
	A       uint8
	X       uint8
	RAM     [0xffff]uint8
	state   int    // Phase of current instruction
	opcode  uint8  // Current opcode
	addrLo  uint8  // Low byte temp
	addrHi  uint8  // High byte temp
	address uint16 // Computed effective address
}

func (cpu *CPU) step() {
	if cpu.state == 0 {
		cpu.opcode = cpu.RAM[cpu.PC]
		fmt.Printf("Clock: PC=0x%04X OPCODE=0x%02X\n", cpu.PC, cpu.opcode)
		cpu.PC++
		switch cpu.opcode {
		case 0xA5:
			cpu.state = 1
		case 0xAD:
			cpu.state = 10
		case 0x8D:
			cpu.state = 20
		default:
			fmt.Printf("Unknown opcode 0x%02X\n", cpu.opcode)
			cpu.state = -1
		}
		return
	}

	switch cpu.opcode {
	case 0xA5:
		cpu.ldaZeroPage()
	case 0xAD:
		cpu.ldaAbsolute()
	case 0x8D:
		cpu.staAbsolute()
	}
}

func (cpu *CPU) ldaZeroPage() {
	switch cpu.state {
	case 1:
		cpu.addrLo = cpu.RAM[cpu.PC]
		fmt.Printf("Clock: PC=0x%04X ZERO_PAGE_ADDR=0x%02X\n", cpu.PC, cpu.addrLo)
		cpu.PC++
		cpu.state = 2
	case 2:
		addr := uint16(cpu.addrLo)
		value := cpu.RAM[addr]
		cpu.A = value
		fmt.Printf("Clock: ADDR=0x%04X VALUE=0x%02X A=0x%02X\n", addr, value, cpu.A)
		cpu.state = 0
	}
}

func (cpu *CPU) ldaAbsolute() {
	switch cpu.state {
	case 10:
		cpu.addrLo = cpu.RAM[cpu.PC]
		fmt.Printf("Clock: PC=0x%04X ADDR_LOW=0x%02X\n", cpu.PC, cpu.addrLo)
		cpu.PC++
		cpu.state = 11
	case 11:
		cpu.addrHi = cpu.RAM[cpu.PC]
		fmt.Printf("Clock: PC=0x%04X ADDR_HIGH=0x%02X\n", cpu.PC, cpu.addrHi)
		cpu.PC++
		cpu.state = 12
	case 12:
		cpu.address = uint16(cpu.addrHi)<<8 | uint16(cpu.addrLo)
		value := cpu.RAM[cpu.address]
		cpu.A = value
		fmt.Printf("Clock: ADDR=0x%04X VALUE=0x%02X A=0x%02X\n", cpu.address, value, cpu.A)
		cpu.state = 0
	}
}

func (cpu *CPU) staAbsolute() {
	switch cpu.state {
	case 20:
		cpu.addrLo = cpu.RAM[cpu.PC]
		fmt.Printf("Clock: PC=0x%04X ADDR_LOW=0x%02X\n", cpu.PC, cpu.addrLo)
		cpu.PC++
		cpu.state = 21
	case 21:
		cpu.addrHi = cpu.RAM[cpu.PC]
		fmt.Printf("Clock: PC=0x%04X ADDR_HIGH=0x%02X\n", cpu.PC, cpu.addrHi)
		cpu.PC++
		cpu.state = 22
	case 22:
		cpu.address = uint16(cpu.addrHi)<<8 | uint16(cpu.addrLo)
		cpu.RAM[cpu.address] = cpu.A
		fmt.Printf("Clock: STORE A=0x%02X -> ADDR=0x%04X\n", cpu.A, cpu.address)
		cpu.state = 0
	}
}

func main() {
	cpu := CPU{
		PC: 0x0000,
		A:  0x42, // preset accumulator
		RAM: [0xffff]uint8{
			0x0000: 0xAD, // LDA Absolute
			0x0001: 0x34,
			0x0002: 0x12,
			0x1234: 0x99,

			// 0x0003: 0xA5, // LDA Zero Page
			// 0x0004: 0x10,
			// 0x0010: 0x55,

			0x0003: 0x8D, // STA Absolute
			0x0004: 0x00,
			0x0005: 0x50, // store to 0x2000
		},
	}

	for i := 0; i < 12; i++ {
		cpu.step()
		time.Sleep(1 * time.Second)
	}

	// Confirm result
	fmt.Printf("\nRAM[0x5000] = 0x%02X (should be 0x%02X)\n", cpu.RAM[0x5000], cpu.A)
}
