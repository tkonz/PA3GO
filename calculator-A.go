package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	right_float := 0.0
	left_float := 0.0
	right_int := 0
	left_int := 0
	op := operatorStack.Pop().(byte)

	r := operandStack.Top()

	if reflect.ValueOf(r).Kind() == reflect.Float64 {
		right_float = operandStack.Pop().(float64)
		r = right_float
	} else {
		right_int = operandStack.Pop().(int)
		r = right_int
	}

	l := operandStack.Top()

	if reflect.ValueOf(l).Kind() == reflect.Float64 {
		left_float = operandStack.Pop().(float64)
		l = left_float
	} else {
		left_int = operandStack.Pop().(int)
		l = left_int
	}

	if reflect.ValueOf(r).Kind() == reflect.ValueOf(l).Kind() {
		if reflect.ValueOf(r).Kind() == reflect.Int {
			switch op {
			case '+':
				operandStack.Push(left_int + right_int)
			case '-':
				operandStack.Push(left_int - right_int)
			case '*':
				operandStack.Push(left_int * right_int)
			case '/':
				operandStack.Push(left_int / right_int)
			default:
				panic("illegal operator")
			}
		} else {
			switch op {
			case '+':
				operandStack.Push(left_float + right_float)
			case '-':
				operandStack.Push(left_float - right_float)
			case '*':
				operandStack.Push(left_float * right_float)
			case '/':
				operandStack.Push(left_float / right_float)
			default:
				panic("illegal operator")
			}
		}
	} else if reflect.ValueOf(r).Kind() == reflect.Int {
		right_float = float64(right_int)
		switch op {
		case '+':
			operandStack.Push(left_float + right_float)
		case '-':
			operandStack.Push(left_float - right_float)
		case '*':
			operandStack.Push(left_float * right_float)
		case '/':
			operandStack.Push(left_float / right_float)
		default:
			panic("illegal operator")
		}
	} else {
		left_float = float64(left_int)
		switch op {
		case '+':
			operandStack.Push(left_float + right_float)
		case '-':
			operandStack.Push(left_float - right_float)
		case '*':
			operandStack.Push(left_float * right_float)
		case '/':
			operandStack.Push(left_float / right_float)
		default:
			panic("illegal operator")
		}
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
			v := int(0)
			fmt.Println("number is currently: ", line[i])
			for {
				v = v*10 + int(line[i]-'0')
				fmt.Println("number is currently: ", v)
				i++
				pre_dec_flag = true //flag for mixed number is on
				fmt.Println("pre_dec_flag = ", pre_dec_flag)
				if i == len(line) || !('0' <= line[i] && line[i] <= '9') {
					break
				}
			}
			//operandIntStack.Push(v)
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
			var f int
			d = float64(tempV) / float64(tempV1)
			fmt.Println("d is a decimal value = ", d)
			var e float64 = float64(d)
			if pre_dec_flag == true {
				f = operandStack.Pop().(int) // pop off int, no need for int
				e += float64(f)              // pop int add to decimal
				fmt.Println("decimal number = ", e)
			}
			operandStack.Push(e) // push e on float stack
			fmt.Println("e = ", e)

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
	if reflect.ValueOf(operandStack.Top()).Kind() == reflect.Int {
		r := operandStack.Pop()
		fmt.Println(r)
	} else {
		r := operandStack.Pop().(float64)
		fmt.Println(r)
	}
}
