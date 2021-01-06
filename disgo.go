package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

const (
	IMP = iota
	IMPA
	MARK2
	BRA
	IMM
	ZP
	ZPX
	ZPY
	INDX
	INDY
	IND
	MARK3
	ABS
	ABSX
	ABSY
	IND16
	IND1X
)

const (
    I_ADC = iota
    I_AND
    I_ASL
    I_BCC
    I_BCS
    I_BEQ
    I_BIT
    I_BMI
    I_BNE
    I_BPL
    I_BRA
    I_BRK
    I_BVC
    I_BVS
    I_CLC
    I_CLD
    I_CLI
    I_CLV
    I_CMP
    I_CPX
    I_CPY
    I_DEC
    I_DEX
    I_DEY
    I_EOR
    I_INC
    I_INX
    I_INY
    I_JMP
    I_JSR
    I_LDA
    I_LDX
    I_LDY
    I_LSR
    I_NOP
    I_ORA
    I_PHA
    I_PHP
    I_PHX
    I_PHY
    I_PLA
    I_PLP
    I_PLX
    I_PLY
    I_ROL
    I_ROR
    I_RTI
    I_RTS
    I_SBC
    I_SEC
    I_SED
    I_SEI
    I_STA
    I_STP
    I_STX
    I_STY
    I_STZ
    I_TAX
    I_TAY
    I_TRB
    I_TSB
    I_TSX
    I_TXA
    I_TXS
    I_TYA
    I_WAI
    I_XXX
)

var opstring = [...] string {
"ADC",
"AND",
"ASL",
"BCC",
"BCS",
"BEQ",
"BIT",
"BMI",
"BNE",
"BPL",
"BRA",
"BRK",
"BVC",
"BVS",
"CLC",
"CLD",
"CLI",
"CLV",
"CMP",
"CPX",
"CPY",
"DEC",
"DEX",
"DEY",
"EOR",
"INC",
"INX",
"INY",
"JMP",
"JSR",
"LDA",
"LDX",
"LDY",
"LSR",
"NOP",
"ORA",
"PHA",
"PHP",
"PHX",
"PHY",
"PLA",
"PLP",
"PLX",
"PLY",
"ROL",
"ROR",
"RTI",
"RTS",
"SBC",
"SEC",
"SED",
"SEI",
"STA",
"STP",
"STX",
"STY",
"STZ",
"TAX",
"TAY",
"TRB",
"TSB",
"TSX",
"TXA",
"TXS",
"TYA",
"WAI",
"???" }

var opname = [...] int {
	I_BRK, I_ORA, I_XXX, I_XXX, I_TSB, I_ORA, I_ASL, I_XXX, I_PHP, I_ORA, I_ASL, I_XXX, I_TSB, I_ORA, I_ASL, I_XXX,  // 00
		I_BPL, I_ORA, I_ORA, I_XXX, I_TRB, I_ORA, I_ASL, I_XXX, I_CLC, I_ORA, I_INC, I_XXX, I_TRB, I_ORA, I_ASL, I_XXX,  // 10
		I_JSR, I_AND, I_XXX, I_XXX, I_BIT, I_AND, I_ROL, I_XXX, I_PLP, I_AND, I_ROL, I_XXX, I_BIT, I_AND, I_ROL, I_XXX,  // 20
		I_BMI, I_AND, I_AND, I_XXX, I_BIT, I_AND, I_ROL, I_XXX, I_SEC, I_AND, I_DEC, I_XXX, I_BIT, I_AND, I_ROL, I_XXX,  // 30
		I_RTI, I_EOR, I_XXX, I_XXX, I_XXX, I_EOR, I_LSR, I_XXX, I_PHA, I_EOR, I_LSR, I_XXX, I_JMP, I_EOR, I_LSR, I_XXX,  // 40
		I_BVC, I_EOR, I_EOR, I_XXX, I_XXX, I_EOR, I_LSR, I_XXX, I_CLI, I_EOR, I_PHY, I_XXX, I_XXX, I_EOR, I_LSR, I_XXX,  // 50
		I_RTS, I_ADC, I_XXX, I_XXX, I_STZ, I_ADC, I_ROR, I_XXX, I_PLA, I_ADC, I_ROR, I_XXX, I_JMP, I_ADC, I_ROR, I_XXX,  // 60
		I_BVS, I_ADC, I_ADC, I_XXX, I_STZ, I_ADC, I_ROR, I_XXX, I_SEI, I_ADC, I_PLY, I_XXX, I_JMP, I_ADC, I_ROR, I_XXX,  // 70
		I_BRA, I_STA, I_XXX, I_XXX, I_STY, I_STA, I_STX, I_XXX, I_DEY, I_BIT, I_TXA, I_XXX, I_STY, I_STA, I_STX, I_XXX,  // 80
		I_BCC, I_STA, I_STA, I_XXX, I_STY, I_STA, I_STX, I_XXX, I_TYA, I_STA, I_TXS, I_XXX, I_STZ, I_STA, I_STZ, I_XXX,  // 90
		I_LDY, I_LDA, I_LDX, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX, I_TAY, I_LDA, I_TAX, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX,  // A0
		I_BCS, I_LDA, I_LDA, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX, I_CLV, I_LDA, I_TSX, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX,  // B0
		I_CPY, I_CMP, I_XXX, I_XXX, I_CPY, I_CMP, I_DEC, I_XXX, I_INY, I_CMP, I_DEX, I_WAI, I_CPY, I_CMP, I_DEC, I_XXX,  // C0
		I_BNE, I_CMP, I_CMP, I_XXX, I_XXX, I_CMP, I_DEC, I_XXX, I_CLD, I_CMP, I_PHX, I_STP, I_XXX, I_CMP, I_DEC, I_XXX,  // D0
		I_CPX, I_SBC, I_XXX, I_XXX, I_CPX, I_SBC, I_INC, I_XXX, I_INX, I_SBC, I_NOP, I_XXX, I_CPX, I_SBC, I_INC, I_XXX,  // E0
		I_BEQ, I_SBC, I_XXX, I_XXX, I_XXX, I_SBC, I_INC, I_XXX, I_SED, I_SBC, I_PLX, I_XXX, I_XXX, I_SBC, I_INC, I_XXX } // F0

var opmode = [...] int {
	IMP, INDX,  IMP, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMPA,  IMP,  ABS,    ABS,   ABS,  IMP,  // 00
		BRA, INDY,  IND, IMP,  ZP,   ZPX,   ZPX,    IMP,   IMP,  ABSY,  IMPA,  IMP,  ABS,    ABSX,  ABSX, IMP,  // 10
		ABS, INDX,  IMP, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMPA,  IMP,  ABS,    ABS,   ABS,  IMP,  // 20
		BRA, INDY,  IND, IMP,  ZPX,  ZPX,   ZPX,    IMP,   IMP,  ABSY,  IMPA,  IMP,  ABSX,   ABSX,  ABSX, IMP,  // 30
		IMP, INDX,  IMP, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMPA,  IMP,  ABS,    ABS,   ABS,  IMP,  // 40
		BRA, INDY,  IND, IMP,  ZP,   ZPX,   ZPX,    IMP,   IMP,  ABSY,  IMP,   IMP,  ABS,    ABSX,  ABSX, IMP,  // 50
		IMP, INDX,  IMP, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMPA,  IMP,  IND16,  ABS,   ABS,  IMP,  // 60
		BRA, INDY,  IND, IMP,  ZPX,  ZPX,   ZPX,    IMP,   IMP,  ABSY,  IMP,   IMP,  IND1X,  ABSX,  ABSX, IMP,  // 70
		BRA, INDX,  IMP, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMP,   IMP,  ABS,    ABS,   ABS,  IMP,  // 80
		BRA, INDY,  IND, IMP,  ZPX,  ZPX,   ZPY,    IMP,   IMP,  ABSY,  IMP,   IMP,  ABS,    ABSX,  ABSX, IMP,  // 90
		IMM, INDX,  IMM, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMP,   IMP,  ABS,    ABS,   ABS,  IMP,  // A0
		BRA, INDY,  IND, IMP,  ZPX,  ZPX,   ZPY,    IMP,   IMP,  ABSY,  IMP,   IMP,  ABSX,   ABSX,  ABSY, IMP,  // B0
		IMM, INDX,  IMP, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMP,   IMP,  ABS,    ABS,   ABS,  IMP,  // C0
		BRA, INDY,  IND, IMP,  ZP,   ZPX,   ZPX,    IMP,   IMP,  ABSY,  IMP,   IMP,  ABS,    ABSX,  ABSX, IMP,  // D0
		IMM, INDX,  IMP, IMP,  ZP,   ZP,     ZP,    IMP,   IMP,  IMM,   IMP,   IMP,  ABS,    ABS,   ABS,  IMP,  // E0
		BRA, INDY,  IMP, IMP,  ZP,   ZPX,   ZPX,    IMP,   IMP,  ABSY,  IMP,   IMP,  ABS,    ABSX,  ABSX, IMP } // F0 

func main() {

    if len(os.Args) != 2 {
		fmt.Println("Usage: disgo <filename> <address>")
        return
    }

    f, err := os.Open(os.Args[1])

    if err != nil {
        fmt.Println(err)
        return
    }

	defer f.Close()
	data, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err)
		return
	}

	var totalBytes = len(data)
	var currentOffset = 0
	const baseAddress = 0x2000

	for currentOffset < totalBytes {

		var thisByte byte = data[currentOffset]
		var mode = opmode[thisByte]
		var bytesRequired = 0

		switch {
			case (mode>MARK3):
			bytesRequired=2
			case (mode>MARK2):
			bytesRequired=1
		}

		var name = opname[thisByte]

		if totalBytes-currentOffset < bytesRequired {
			panic("todo")
		}

		var outputStr = fmt.Sprintf("%02X ", baseAddress+currentOffset)

		for i := 0; i < 3; i++ {
			if i <= bytesRequired {
				outputStr += fmt.Sprintf("%02X ", data[currentOffset+i])
            } else {
				outputStr += "   "
			}
		}

		outputStr += fmt.Sprintf("%s", opstring[name])
		currentOffset++

		switch (mode) {
		//case IMP:
		//case IMPA:
		case BRA:
			branchRange := int(int8(data[currentOffset]+1))
			branchTarget := (currentOffset + branchRange) + baseAddress
			outputStr += fmt.Sprintf(" &%04X", branchTarget);
			currentOffset++
		case IMM:
			outputStr += fmt.Sprintf(" #&%02X", data[currentOffset])
			currentOffset++;
			break;
		case ZP:
			outputStr += fmt.Sprintf(" &%02X", data[currentOffset])
			currentOffset++;
			break;
		case ZPX:
			outputStr += fmt.Sprintf(" %02X,X", data[currentOffset])
			currentOffset++;
		case ZPY:
			outputStr += fmt.Sprintf(" %02X,Y", data[currentOffset])
			currentOffset++;
		case IND:
			outputStr += fmt.Sprintf(" (%02X)", data[currentOffset])
			currentOffset++;
		case INDX:
			outputStr += fmt.Sprintf(" (&%02X,X)", data[currentOffset])
			currentOffset++;
		case INDY:
			outputStr += fmt.Sprintf(" (&%02X),Y", data[currentOffset])
			currentOffset++;
		case ABS:
			outputStr += fmt.Sprintf(" &%02X%02X", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2;
		case ABSX:
			outputStr += fmt.Sprintf(" &%02X%02X,X", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2;
		case ABSY:
			outputStr += fmt.Sprintf(" &%02X%02X,Y", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2;
		case IND16:
			outputStr += fmt.Sprintf(" (&%02X%02X)", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2;
		case IND1X:
			outputStr += fmt.Sprintf(" (&%02X%02X,X)", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2;
		}

		fmt.Println(outputStr)
	}
}
