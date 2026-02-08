package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and press ENTER when ready."

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//rand generates a number between 0 and whatever is passed as a parameter
	//we add 1 to it because we want the number used to be at least 1, and not
	//greater than 10
	var firstNumber = rand.Intn(9) + 1
	var secondNumber = rand.Intn(9) + 1
	var subtraction = rand.Intn(9) + 1
	var answer = firstNumber * secondNumber - subtraction

	playGuessNumber(firstNumber,secondNumber,subtraction,answer)
}

func playGuessNumber(firstNumber,secondNumber,subtraction,answer int){
	reader := bufio.NewReader(os.Stdin)
	
	// display a welcome/instruction
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	//take  them through the games
	fmt.Println("Multiply your number by",firstNumber,prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by",secondNumber,prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of",prompt)
	reader.ReadString('\n')
	
	fmt.Println("Now subtract", subtraction,prompt)
	reader.ReadString('\n')

	//give them the answer
	fmt.Println("The answer is", answer)
}