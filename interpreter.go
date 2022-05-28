package brainfact

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Global variables configurables.
var (
	TapeSize uint = math.MaxInt16
	Prompt        = "brain-fact > "
)

// interpreter compiles and executes brainfuck code.
type interpreter struct {
	code     string
	compiler *compiler
	tape     []byte
	scanner  *bufio.Scanner
}

// Run creates a new Interpreter and runs the code passed as argument. It returns nil if no error happened.
func Run(code string) error {
	i := &interpreter{
		code:     code,
		compiler: newCompiler(code),
		tape:     make([]byte, TapeSize),
		scanner:  bufio.NewScanner(os.Stdin),
	}
	i.compiler.run()
	if i.compiler.err != nil {
		return i.compiler.err
	}
	return i.run()
}

// run runs the brainfuck interpreter.
func (i *interpreter) run() error {
	tape := i.tape
	bytecode := i.compiler.bytecode
	limit := uint(len(bytecode))
	var pointer uint
	var pc uint
	var op byte
	for {
		op = bytecode[pc]
		switch op {
		case opLeft:
			if pointer == 0 {
				return fmt.Errorf("tape underflow")
			}
			pointer--
			pc++
		case opRight:
			if pointer == limit-1 {
				return fmt.Errorf("tape overflow")
			}
			pointer++
			pc++
		case opAdd:
			tape[pointer]++
			pc++
		case opSub:
			tape[pointer]--
			pc++
		case opLB:
			if tape[pointer] == 0 {
				pc = uint(binary.BigEndian.Uint16(bytecode[pc+1:]))
			} else {
				pc += 3
			}
		case opRB:
			if tape[pointer] != 0 {
				pc = uint(binary.BigEndian.Uint16(bytecode[pc+1:]))
			} else {
				pc += 3
			}
		case opComma:
			fmt.Print(Prompt)
		scanLoop:
			for i.scanner.Scan() {
				if result, err := strconv.ParseUint(i.scanner.Text(), 0, 64); err == nil {
					tape[pointer] = byte(result % 256)
					break scanLoop
				} else {
					break scanLoop
				}
			}
			pc++
		case opDot:
			fmt.Print(string(tape[pointer]))
			pc++
		case opEND:
			return nil
		default:
			return fmt.Errorf("unknown bytecode %v", op)
		}
	}
}
