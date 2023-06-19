package main

import (
	"bufio"
	"fmt"
	"os"

	brainfact "github.com/ever-eduardo/brain-fact"
)

const BFFileExtension string = ".bf"
const Welcome string = "Welcome to Brainfact REPL! 🌎"
const Description string = "A brainfuck interpreter written in Go."
const Help string = "Type .help for assistance"

func main() {
	clearScreen()
	if len(os.Args) > 1 {

	} else {
		var code string
		running := true
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("\n\n\n   %s\n   %s\n   %s\n\n\n", Welcome, Description, Help)
		for running {
			fmt.Printf("%v", brainfact.Prompt)
			if scanner.Scan() {
				code = scanner.Text()
				switch code {
				case ".help":
					printHelp()
				case ".clear":
					clearScreen()
					fmt.Printf("\n\n\n   %s\n   %s\n   %s\n\n\n", Welcome, Description, Help)
				case ".exit":
					running = false
				default:
					println()
					if err := brainfact.Run(code); err != nil {
						fmt.Printf("\n\n\n Error: %s\n\n\n", err)
					}
					println()
				}
			}
		}
		clearScreen()
		fmt.Printf("\n\n\n   Leaving the Brainfact REPL! 🌎\n\n\n")
	}
}

func clearScreen() {
	fmt.Printf("\u001B[H")
	fmt.Printf("\u001B[2J")
}

func printHelp() {
	clearScreen()
	fmt.Printf("\n\n\n   Brainfact REPL! 🌎\n\n\n")
	fmt.Printf("   Command list:\n")
	fmt.Println("   .help  - Show this message")
	fmt.Println("   .exit  - Quit the REPL")
	fmt.Println("   .clear - Clear the terminal")
	fmt.Printf("\n\n")
}
