package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"

	"github.com/bbrainstormer/bootdev-pokedex/internal/commands"
	"github.com/bbrainstormer/bootdev-pokedex/internal/lib"
)

func main() {
	// Handle interrupt
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		fmt.Println("\nProgram interrupted.")
		os.Exit(0)
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("Error found: %v\n", err)
		}
		input := scanner.Text()
		clean_input := lib.CleanInput(input)
		if len(clean_input) == 0 {
			continue
		}
		word := clean_input[0]
		cmd, exists := commands.GetCommand(word)
		if exists {
			err := cmd.Callback(clean_input[1:])
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
