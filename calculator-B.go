package main

import (
	"bufio"
	"fmt"
	"os"
	_ "reflect"
	"stack"
)

var operatorStack = stack.NewStack()
var operandStack = stack.NewStack()

func precedence(op byte) uint8 {
	switch op {
	case '+', '-':
		return 0
	case '*', '/':
		return 1
	case ')':
		return 2
	case '(':
		return 3
	default:
		panic("illegal operator")
	}
}

func apply() {

	op := operatorStack.Pop().(byte)
	fmt.Println(op, " is operator")
	right := operandStack.Pop().(float64)
	fmt.Println(right, " is right hand side")
	left := operandStack.Pop().(float64)
	fmt.Println(left, " is left hand side")

	switch op {
	case '+':
		operandStack.Push(left + right)
	case '-':
		operandStack.Push(left - right)
	case '*':
		operandStack.Push(left * right)
	case '/':
		operandStack.Push(left / right)
	default:
		panic("illegal operator")
	}
}

func main() {
	// Read a from Stdin.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	//
	//flag for space error
	num_flag := false
	pre_dec_flag := false

	for i := 0; i < len(line); {
		switch line[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			v := float64(0)
			fmt.Println("number is currently: ", line[i])
			for {
				v = v*10 + float64(line[i]-'0')
				fmt.Println("number is currently: ", v)
				i++
				pre_dec_flag = true //flag for mixed number is on
				fmt.Println("pre_dec_flag = ", pre_dec_flag)
				if i == len(line) || !('0' <= line[i] && line[i] <= '9') {
					break
				}
			}
			operandStack.Push(v)
			fmt.Println("top of operand stack is ", v)
			fmt.Println(v, " is pushed onto number stack")
			num_flag = true // turn on num flag

		case '.':
			i++ // to get to number portion after .
			tempV := 0
			tempV1 := 1
			lenDec := 0 // keep track of length of decimal

			for {
				tempV = tempV*10 + int(line[i]-'0')
				fmt.Println("number is currently: ", tempV)
				i++
				lenDec++
				fmt.Println("length of dec is currently: ", lenDec)

				if i == len(line) || !('0' <= line[i] && line[i] <= '9') {
					break
				}
			}
			fmt.Println("Made it out of for loop")
			// evaluate decimal
			for index := 0; index < lenDec; {
				tempV1 *= 10
				fmt.Println("created tempV1 = ", tempV1)
				index++
			}
			var d float64
			d = float64(tempV) / float64(tempV1)
			fmt.Println("d is a decimal value = ", d)
			var e float64 = float64(d)
			if pre_dec_flag == true {
				e += operandStack.Pop().(float64)
				fmt.Println("decimal number = ", e)
			}
			operandStack.Push(e)
		case '+', '-', '*', '/':
			num_flag = false // turn off num flag
			for !operatorStack.IsEmpty() && precedence(operatorStack.Top().(byte)) < 3 &&
				precedence(operatorStack.Top().(byte)) >= precedence(line[i]) {
				apply()
			}
			fmt.Println("current spot in string is ", line[i])
			fmt.Println("+ - * / is pushed onto operator stack")
			fmt.Println("num_flag after +, -, *, /", num_flag)
			operatorStack.Push(line[i])
			i++
		case '(':
			num_flag = false // turn off num flag
			fmt.Println("num_flag after ( is ", num_flag)
			operatorStack.Push(line[i])
			fmt.Println("left paren is pushed onto operator stack")
			i++
		case ')':
			num_flag = false // turn off num flag
			fmt.Println("num_flag after ) is ", num_flag)
			for precedence(operatorStack.Top().(byte)) < 3 {
				apply()
			}

			// then pop the left paren
			operatorStack.Pop()
			i++
		case ' ':
			fmt.Println("num_flag after a space is ", num_flag)
			if num_flag == true && '0' <= line[i+1] && line[i+1] <= '9' {
				panic("space error")
			}
			fmt.Println("current spot in string is ", i)
			i++

		default:
			panic("illegal character")
		}
	}
	for !operatorStack.IsEmpty() {
		apply()
	}
	r := operandStack.Pop().(float64)
	fmt.Println(r)
}
