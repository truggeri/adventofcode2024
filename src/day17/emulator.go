package day17

import (
	"strconv"
)

type instruction uint8

const (
	ADV instruction = iota
	BXL
	BST
	JNZ
	BXC
	OUT
	BDV
	CDV
)

type computer struct {
	regA, regB, regC uint
	ip               uint
	instructions     []instruction
	outBuffer        []string
}

func (c *computer) tick() bool {
	instr := c.instructions[c.ip]
	operand := c.nextOperand()
	advance := true

	switch instr {
	case ADV:
		result := c.regA / (1 << operand)
		c.regA = result
		break
	case BXL:
		result := c.regB ^ uint(operand)
		c.regB = result
		break
	case BST:
		result := uint(operand) % 8
		c.regB = result
		break
	case JNZ:
		if c.regA == 0 {
			break
		}
		c.ip = operand
		advance = false
		break
	case BXC:
		result := c.regB ^ c.regC
		c.regB = result
		break
	case OUT:
		result := operand % 8
		c.outBuffer = append(c.outBuffer, strconv.Itoa(int(result)))
		break
	case BDV:
		result := c.regA / (1 << operand)
		c.regB = result
		break
	case CDV:
		result := c.regA / (1 << operand)
		c.regC = result
		break
	}
	if advance {
		c.ip += 2
	}

	return int(c.ip) >= len(c.instructions)
}

func (c *computer) nextOperand() uint {
	i := c.instructions[c.ip]
	o := c.instructions[c.ip+1]
	comboOperand := i == ADV || i == BST || i == OUT || i == BDV || i == CDV
	if !comboOperand {
		return uint(o)
	}

	if o <= 3 {
		return uint(o)
	} else if o == 4 {
		return c.regA
	} else if o == 5 {
		return c.regB
	} else if o == 6 {
		return c.regC
	}
	panic("Combo operand of 7 provided")
}
