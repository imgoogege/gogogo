// calculator.go
// 	example: calculate 3 + 4 = 7 as input: 3 ENTER 4 ENTER + ENTER --> result = 7,

package main

import (
	"./stack/stack"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	calc1 := new(stack.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error !")
			os.Exit(1)
		}
		token = token[0 : len(token)-2] // remove "\r\n"
		// fmt.Printf("--%s--\n",token)  // debug statement
		switch {
		case token == "q": // stop als invoer = "q"
			fmt.Println("Calculator stopped")
			return
		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc1.Push(i)
		case token == "+":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)+q.(int))
		case token == "-":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)-q.(int))

		case token == "*":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)*q.(int))

		case token == "/":
			q, _ := calc1.Pop()
			p, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)/q.(int))
		default:
			fmt.Println("No valid input")
		}
	}
}


------------

2018-10-16 自己更新


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type OutResult struct {
	symbol string
	result int
}

func main() {
	Machine()
}

func Machine() {
	var resultSlice int
	var v1 int //中间量
	outResultValue := new(OutResult)
	readValue, i := outResultValue.Read()
	fmt.Println("输出值", readValue)
	switch readValue[0] {
	case "+":
		fmt.Println("是+")
		for _, v := range readValue[1:] {
			v1 = transfer(v)
			resultSlice += v1
		}
	case "-":
		resultSlice = transfer(readValue[1:][0])
		for _, v := range readValue[2:] {
			v1 = transfer(v)
			resultSlice = resultSlice - v1
		}
	case "*":
		resultSlice = 1
		for _, v := range readValue[1:] {
			v1 = transfer(v)
			resultSlice *= v1
		}
	case "/":
		resultSlice = transfer(readValue[1:][0])
		for _, v := range readValue[2:] {
			v1 = transfer(v)
			resultSlice = resultSlice / v1
		}
	default:
		fmt.Println("发生错误")
	}
	outResultValue.result = resultSlice
	outResultValue.symbol = strconv.FormatInt(int64(i-1), 10)
	fmt.Println(outResultValue)
}
func (o *OutResult) String() string {
	result := strconv.FormatInt(int64(o.result), 10)
	return "显示结果" + result + "计算的次数" + o.symbol
}
func (o *OutResult) Read() ([]string, int) {
	i := 0
	fmt.Println("请输入正确的type,仅限 + - * /")
	re := make([]string, 0)
	reader := bufio.NewReader(os.Stdin) //创建读
	for {
		str, err := reader.ReadString('\n') //记录一次读取，每次读取的分隔符使用\n来确定
		str = str[0 : len(str)-1]           // remove "\r\n"
		if err != nil {
			fmt.Println("25:", err)
		}
		if str == "q" {
			fmt.Println("输入停止，正在计算数值，并且返回给标准输出")
			break
		}
		if i >= 1 {
			if transfer(str) > 999 {
				fmt.Println("请勿输入的数值大于999，请重新输入")
				continue
			}
		}
		re = append(re, str)
		i++
	}
	return re, i
}

func transfer(v string) int {
	value, _ := strconv.ParseInt(v, 10, 0)
	return int(value)
}
