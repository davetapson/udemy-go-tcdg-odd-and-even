package main

import "fmt"

//In this assignment you'll get some practice with slices and for loops.
//
//Here are the directions:
//
//Create a new project folder to house this new project and create a 'main.go' file inside of it.
//In the main.go file, define a function called 'main'.  Remember that the 'main' function will be called automatically.
//In the main function, create a slice of ints from 0 through 10
//Iterate through the slice with a for loop.  For each element, check to see if the number is even or odd.
//If the value is even, print out "even".  If it is odd, print out "odd"
//Run your code from the terminal by changing into your project directory then running 'go run main.go'.
//Your output should look like this:
//
//0 is even
//1 is odd
//2 is even
//3 is odd
//4 is even
//5 is odd
//6 is even
//7 is odd
//8 is even
//9 is odd
//10 is even

// I used modulus to determine odd and even.
// I matched colour of text in your answer, but I suspect that may be a function of your IDE

func main(){

	x := []int{}
	for i := range x {
		x[i] = i
	}

	fmt.Printf(x)

	s := []int{0,1,2,3,4,5,6,7,8,9,10}

	for i := range s{
		if s[i]%2 == 0{
			fmt.Printf("%v \x1b[31;1m is \x1b[30;1m even\n", s[i])
		} else {
			fmt.Printf("%v \x1b[31;1m is \x1b[30;1m odd\n", s[i])
		}
	}
}