package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

var pos = []string{" ", " ", " ", " ", " ", " ", " ", " ", " "}
var played = []int{}
var player int = 1
var plays int = 0

var winCombo1 = []int{0, 3, 6, 0, 1, 2, 0, 2}
var winCombo2 = []int{1, 4, 7, 3, 4, 5, 4, 4}
var winCombo3 = []int{2, 5, 8, 6, 7, 8, 8, 6}

func main() {
	showTable()
	play()
}

func showTable() {
	clearTerminal()
	fmt.Printf("  %s | %s | %s   \n", pos[0], pos[1], pos[2])
	fmt.Println(" ----------- ")
	fmt.Printf("  %s | %s | %s   \n", pos[3], pos[4], pos[5])
	fmt.Println(" ----------- ")
	fmt.Printf("  %s | %s | %s   \n", pos[6], pos[7], pos[8])
	fmt.Println("")
}

func play() {
	fmt.Printf("Player %d, what position do you want to play?\n", player)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter position(1-9): ")
	text, _ := reader.ReadString('\n')
	conv, err := strconv.Atoi(string(text[0]))
	if err != nil || conv == 0 {
		showTable()
		fmt.Println("ERROR: Must enter a number from 1 - 9!")
		play()
	}
	if contains(conv) {
		fmt.Println("ERROR: You already tried that position... Try again!")
		play()
	}
	if player == 1 {
		pos[conv-1] = "x"
		player = 2
	} else {
		pos[conv-1] = "o"
		player = 1
	}
	plays++
	played = append(played, conv)
	verifyWin()
}

func verifyWin() {
	for i := 0; i < 8; i++ {
		if pos[winCombo1[i]] == "x" && pos[winCombo2[i]] == "x" && pos[winCombo3[i]] == "x" {
			showTable()
			fmt.Println("Player 1 wins!!")
			os.Exit(0)
		} else if pos[winCombo1[i]] == "o" && pos[winCombo2[i]] == "o" && pos[winCombo3[i]] == "o" {
			showTable()
			fmt.Println("Player 2 wins!!")
			os.Exit(0)
		} else if plays > 8 {
			showTable()
			fmt.Println("Game Tied!!")
			os.Exit(0)
		}
	}
	showTable()
	play()
}

// Clears the terminal ( command for MAC or LINUX )
func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

// Verifies if the position picked has not already been chosen
func contains(val int) bool {
	for _, d := range played {
		if d == val {
			return true
		}
	}
	return false
}
