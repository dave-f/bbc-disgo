package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

var opstring = [...]string{
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
	"???"}

var opname = [...]int{
	I_BRK, I_ORA, I_XXX, I_XXX, I_TSB, I_ORA, I_ASL, I_XXX, I_PHP, I_ORA, I_ASL, I_XXX, I_TSB, I_ORA, I_ASL, I_XXX, // 00
	I_BPL, I_ORA, I_ORA, I_XXX, I_TRB, I_ORA, I_ASL, I_XXX, I_CLC, I_ORA, I_INC, I_XXX, I_TRB, I_ORA, I_ASL, I_XXX, // 10
	I_JSR, I_AND, I_XXX, I_XXX, I_BIT, I_AND, I_ROL, I_XXX, I_PLP, I_AND, I_ROL, I_XXX, I_BIT, I_AND, I_ROL, I_XXX, // 20
	I_BMI, I_AND, I_AND, I_XXX, I_BIT, I_AND, I_ROL, I_XXX, I_SEC, I_AND, I_DEC, I_XXX, I_BIT, I_AND, I_ROL, I_XXX, // 30
	I_RTI, I_EOR, I_XXX, I_XXX, I_XXX, I_EOR, I_LSR, I_XXX, I_PHA, I_EOR, I_LSR, I_XXX, I_JMP, I_EOR, I_LSR, I_XXX, // 40
	I_BVC, I_EOR, I_EOR, I_XXX, I_XXX, I_EOR, I_LSR, I_XXX, I_CLI, I_EOR, I_PHY, I_XXX, I_XXX, I_EOR, I_LSR, I_XXX, // 50
	I_RTS, I_ADC, I_XXX, I_XXX, I_STZ, I_ADC, I_ROR, I_XXX, I_PLA, I_ADC, I_ROR, I_XXX, I_JMP, I_ADC, I_ROR, I_XXX, // 60
	I_BVS, I_ADC, I_ADC, I_XXX, I_STZ, I_ADC, I_ROR, I_XXX, I_SEI, I_ADC, I_PLY, I_XXX, I_JMP, I_ADC, I_ROR, I_XXX, // 70
	I_BRA, I_STA, I_XXX, I_XXX, I_STY, I_STA, I_STX, I_XXX, I_DEY, I_BIT, I_TXA, I_XXX, I_STY, I_STA, I_STX, I_XXX, // 80
	I_BCC, I_STA, I_STA, I_XXX, I_STY, I_STA, I_STX, I_XXX, I_TYA, I_STA, I_TXS, I_XXX, I_STZ, I_STA, I_STZ, I_XXX, // 90
	I_LDY, I_LDA, I_LDX, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX, I_TAY, I_LDA, I_TAX, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX, // A0
	I_BCS, I_LDA, I_LDA, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX, I_CLV, I_LDA, I_TSX, I_XXX, I_LDY, I_LDA, I_LDX, I_XXX, // B0
	I_CPY, I_CMP, I_XXX, I_XXX, I_CPY, I_CMP, I_DEC, I_XXX, I_INY, I_CMP, I_DEX, I_WAI, I_CPY, I_CMP, I_DEC, I_XXX, // C0
	I_BNE, I_CMP, I_CMP, I_XXX, I_XXX, I_CMP, I_DEC, I_XXX, I_CLD, I_CMP, I_PHX, I_STP, I_XXX, I_CMP, I_DEC, I_XXX, // D0
	I_CPX, I_SBC, I_XXX, I_XXX, I_CPX, I_SBC, I_INC, I_XXX, I_INX, I_SBC, I_NOP, I_XXX, I_CPX, I_SBC, I_INC, I_XXX, // E0
	I_BEQ, I_SBC, I_SBC, I_XXX, I_XXX, I_SBC, I_INC, I_XXX, I_SED, I_SBC, I_PLX, I_XXX, I_XXX, I_SBC, I_INC, I_XXX} // F0

var opmode = [...]int{
	IMP, INDX, IMP, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMPA, IMP, ABS, ABS, ABS, IMP, // 00
	BRA, INDY, IND, IMP, ZP, ZPX, ZPX, IMP, IMP, ABSY, IMPA, IMP, ABS, ABSX, ABSX, IMP, // 10
	ABS, INDX, IMP, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMPA, IMP, ABS, ABS, ABS, IMP, // 20
	BRA, INDY, IND, IMP, ZPX, ZPX, ZPX, IMP, IMP, ABSY, IMPA, IMP, ABSX, ABSX, ABSX, IMP, // 30
	IMP, INDX, IMP, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMPA, IMP, ABS, ABS, ABS, IMP, // 40
	BRA, INDY, IND, IMP, ZP, ZPX, ZPX, IMP, IMP, ABSY, IMP, IMP, ABS, ABSX, ABSX, IMP, // 50
	IMP, INDX, IMP, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMPA, IMP, IND16, ABS, ABS, IMP, // 60
	BRA, INDY, IND, IMP, ZPX, ZPX, ZPX, IMP, IMP, ABSY, IMP, IMP, IND1X, ABSX, ABSX, IMP, // 70
	BRA, INDX, IMP, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMP, IMP, ABS, ABS, ABS, IMP, // 80
	BRA, INDY, IND, IMP, ZPX, ZPX, ZPY, IMP, IMP, ABSY, IMP, IMP, ABS, ABSX, ABSX, IMP, // 90
	IMM, INDX, IMM, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMP, IMP, ABS, ABS, ABS, IMP, // A0
	BRA, INDY, IND, IMP, ZPX, ZPX, ZPY, IMP, IMP, ABSY, IMP, IMP, ABSX, ABSX, ABSY, IMP, // B0
	IMM, INDX, IMP, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMP, IMP, ABS, ABS, ABS, IMP, // C0
	BRA, INDY, IND, IMP, ZP, ZPX, ZPX, IMP, IMP, ABSY, IMP, IMP, ABS, ABSX, ABSX, IMP, // D0
	IMM, INDX, IMP, IMP, ZP, ZP, ZP, IMP, IMP, IMM, IMP, IMP, ABS, ABS, ABS, IMP, // E0
	BRA, INDY, IND, IMP, ZP, ZPX, ZPX, IMP, IMP, ABSY, IMP, IMP, ABS, ABSX, ABSX, IMP} // F0

const (
	DATA = iota
	CODE
)

type CodePoint struct {
	offset   int // actual offset in file
	address  int // address
	bytetype int // code (1) or data (0)
}

// Parse the control file.  Control files currently have 4 commands:
// load <file> - the source file
// base <address> - set the base address
// data <address>,<count> - mark a region as data (can be used multiple times)
// save <file> - the target file
func parseControlFile(controlfilename string) (string, string, [][]CodePoint, error) {

	var parsedLoadFile string
	var parsedSaveFile string
	var parsedBase int
	f, err := os.Open(controlfilename)

	if err != nil {
		return "", "", nil, err
	}

	defer f.Close()

	type DataBlock struct {
		address int
		length  int
	}

	dataBlocks := make([]DataBlock, 0, 5)
	s := bufio.NewScanner(f)

	for s.Scan() {
		l := strings.TrimSpace(s.Text())
		if strings.HasPrefix(l, "load") {
			parsedLoadFile = strings.TrimSpace(strings.TrimPrefix(l, "load"))
		} else if strings.HasPrefix(l, "base") {
			base, err := strconv.ParseInt(strings.TrimSpace(strings.TrimPrefix(l, "base")), 0, 0)
			if err != nil {
				return "", "", nil, err
			}
			parsedBase = int(base)
		} else if strings.HasPrefix(l, "data") {
			cmd := strings.Split(strings.TrimPrefix(l, "data"), ",")
			if len(cmd) != 2 {
				return "", "", nil, errors.New("bad data command")
			}
			parsedAddress, err := strconv.ParseInt(strings.TrimSpace(cmd[0]), 0, 0)
			if err != nil {
				return "", "", nil, err
			}
			parsedLength, err := strconv.ParseInt(strings.TrimSpace(cmd[1]), 0, 0)
			if err != nil {
				return "", "", nil, err
			}
			var newBlock DataBlock
			newBlock.address = int(parsedAddress)
			newBlock.length = int(parsedLength)
			dataBlocks = append(dataBlocks, newBlock)
		} else if strings.HasPrefix(l, "save") {
			parsedSaveFile = strings.TrimSpace(strings.TrimPrefix(l, "save"))
		}
	}

	fi, err := os.Stat(parsedLoadFile)

	if err != nil {
		return "", "", nil, err
	}

	filesize := int(fi.Size())
	d := make([]CodePoint, filesize)
	r := make([][]CodePoint, 0, 5)

	// Assume all bytes are code to begin with
	for i := 0; i < filesize; i++ {
		d[i].offset = i
		d[i].address = parsedBase + i
		d[i].bytetype = CODE
	}

	// Now go through data segments above and set up the byte types
	for _, v := range dataBlocks {
		//fmt.Printf("Data block at 0x%x, length %d\n", v.address, v.length)
		fileOffs := v.address - parsedBase // file offset
		for i, j := fileOffs, 0; j < v.length; i, j = i+1, j+1 {
			d[i].bytetype = DATA
		}
	}

	// Finally, create the slices from the data types
	lastByteType := -1
	idx := -1

	for _, v := range d {
		if v.bytetype != lastByteType {
			// fmt.Printf("New slice at 0x%x\n", v.address)
			idx++
			r = append(r, make([]CodePoint, 0, filesize))
		}
		// add to current slice
		r[idx] = append(r[idx], v)
		lastByteType = v.bytetype
	}

	return parsedLoadFile, parsedSaveFile, r, nil
}

type Comment struct {
	address int
	comment string
}

// Go through the target file, saving all comments and their address
func saveComments(targetFilename string, commentCol int) ([]Comment, error) {
	f, err := os.Open(targetFilename)

	if err != nil {
		return nil, err
	}

	defer f.Close()

	comments := make([]Comment, 0, 10)
	s := bufio.NewScanner(f)

	for s.Scan() {
		l := s.Text()
		if len(l) > commentCol {
			p := strings.SplitN(l, " ", 2)
			if (l[commentCol] == ';') && (len(p) == 2) {
				a, err := strconv.ParseInt(p[0], 16, 0)
				if err == nil {
					c := strings.SplitN(l, ";", 2)
					fmt.Println("Comment at", p[0], "->", c[1])
					var newComment Comment
					newComment.address = int(a)
					newComment.comment = c[1]
					comments = append(comments, newComment)
				} else {
					fmt.Println("Error parsing comments", err)
				}
			}
		}
	}

	return comments, nil
}

func getCommentForAddress(address int, comments []Comment) (bool, string) {
	for _, i := range comments {
		if i.address == address {
			return true, i.comment
		}
	}
	return false, ""
}

func dis(data []byte, baseAddress int, asmType int, writer io.Writer, applyComments bool, comments []Comment, commentColumn int) {
	var currentOffset = 0
	var totalBytes = len(data)
	last := false

	for currentOffset < totalBytes {
		var thisByte byte = data[currentOffset]
		var mode = opmode[thisByte]
		var bytesRequired = 0

		switch {
		case (mode > MARK3):
			bytesRequired = 2
		case (mode > MARK2):
			bytesRequired = 1
		}

		var name int

		if asmType == DATA {
			bytesRequired = 0
			mode = IMP
			name = len(opstring) - 1
		} else {
			name = opname[thisByte]
		}

		var outputStr = fmt.Sprintf("%02X ", baseAddress+currentOffset)
		commentExists, commentText := getCommentForAddress(baseAddress+currentOffset, comments)

		// If we do not have enough bytes in this slice to disassemble this instruction, drop out here
		if (currentOffset + bytesRequired) >= totalBytes {
			bytesRequired = totalBytes - (currentOffset + bytesRequired)
			last = true
		}

		for i := 0; i < 3; i++ {
			if i <= bytesRequired {
				outputStr += fmt.Sprintf("%02X ", data[currentOffset+i])
			} else {
				outputStr += "   "
			}
		}

		if last {
			outputStr += "???"
			fmt.Println(outputStr)
			break
		}

		outputStr += opstring[name]
		currentOffset++

		switch mode {
		//case IMP:
		case IMPA:
			outputStr += " A"
		case BRA:
			branchRange := int(int8(data[currentOffset] + 1))
			branchTarget := (currentOffset + branchRange) + baseAddress
			outputStr += fmt.Sprintf(" &%04X", branchTarget)
			currentOffset++
		case IMM:
			outputStr += fmt.Sprintf(" #&%02X", data[currentOffset])
			currentOffset++
		case ZP:
			outputStr += fmt.Sprintf(" &%02X", data[currentOffset])
			currentOffset++
		case ZPX:
			outputStr += fmt.Sprintf(" &%02X,X", data[currentOffset])
			currentOffset++
		case ZPY:
			outputStr += fmt.Sprintf(" &%02X,Y", data[currentOffset])
			currentOffset++
		case IND:
			outputStr += fmt.Sprintf(" (&%02X)", data[currentOffset])
			currentOffset++
		case INDX:
			outputStr += fmt.Sprintf(" (&%02X,X)", data[currentOffset])
			currentOffset++
		case INDY:
			outputStr += fmt.Sprintf(" (&%02X),Y", data[currentOffset])
			currentOffset++
		case ABS:
			outputStr += fmt.Sprintf(" &%02X%02X", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2
		case ABSX:
			outputStr += fmt.Sprintf(" &%02X%02X,X", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2
		case ABSY:
			outputStr += fmt.Sprintf(" &%02X%02X,Y", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2
		case IND16:
			outputStr += fmt.Sprintf(" (&%02X%02X)", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2
		case IND1X:
			outputStr += fmt.Sprintf(" (&%02X%02X,X)", data[currentOffset+1], data[currentOffset+0])
			currentOffset += 2
		}

		if commentExists {
			formatStr := fmt.Sprintf("%%-%ds", commentColumn)
			outputStr = fmt.Sprintf(formatStr, outputStr)
			outputStr += ";"
			outputStr += commentText
		}

		_, err := io.WriteString(writer, outputStr+"\n")

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

func main() {

	wipeComments := flag.Bool("wipe", false, "Wipe comments")
	commentColumn := flag.Int("column", 28, "Column number for comments")
	_ = flag.Bool("stdout", false, "Output to stdout rather than the specifed file (TODO)")

	usageFunc := func() {
		fmt.Println("Usage: disgo <control file>")
		fmt.Printf("Options:\n")
		flag.PrintDefaults()
	}

	flag.Usage = usageFunc
	flag.Parse()

	if flag.NArg() != 1 {
		usageFunc()
		return
	}

	sourceFilename, targetFilename, p, err := parseControlFile(flag.Arg(0))

	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Open(sourceFilename)

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

	var comments []Comment

	if *commentColumn < 28 {
		fmt.Println("Comment column too low; defaulting to 28")
		*commentColumn = 28
	}

	if !*wipeComments {
		comments, err = saveComments(targetFilename, *commentColumn)
		if err != nil {
			fmt.Println("Cannot read comments (possibly new target file?)")
		} else {
			fmt.Println("Comments found", len(comments))
		}
	}

	of, err := os.Create(targetFilename)

	if err == nil {

		defer of.Close()

		for _, v := range p {
			thisSlice := data[v[0].offset : 1+v[len(v)-1].offset]
			var thisSliceType string
			if v[0].bytetype == 0 {
				thisSliceType = "data"
			} else if v[0].bytetype == 1 {
				thisSliceType = "code"
			} else {
				thisSliceType = "undefined"
			}
			fmt.Printf("Disassembling slice length %d address 0x%x type %s\n", len(thisSlice), v[0].address, thisSliceType)
			dis(thisSlice, v[0].address, v[0].bytetype, of, !*wipeComments, comments, *commentColumn)
		}
	} else {
		fmt.Println(err)
		return
	}

	fmt.Println("OK")
}
