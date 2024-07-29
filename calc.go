package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var strSplit, first, second []string
var a, b, res, res1, resfinish string
var sign, checkNum int
var err error

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите данные: ")
	str, _ := reader.ReadString('\n')
	strSplit = strings.Split(str, "")
	for i, v := range strSplit {
		if v == "+" || v == "-" || v == "*" || v == "/" {
			if strSplit[i-1] == " " && strSplit[i+1] == " " {
				sign = i
			}
		}
	}

	if sign == 0 {
		panic("Невозможная операция, доступны только следующие операции: +, -, *, /")
	}
	first = strSplit[:sign-1]
	second = strSplit[sign+2:]
	check()
	fmt.Println(calculate())
}

func check() (string, string) {
	if first[0] == `"` && first[len(first)-1] == `"` {
		first = first[1 : len(first)-1]
		a = strings.Join(first, "")
	} else {
		panic("Первый операнд должен быть строкой")
	}

	if second[0] == `"` && second[len(second)-2] == `"` {
		second = second[1 : len(second)-2]
		b = strings.Join(second, "")
	} else {
		second = second[:len(second)-1]
		b = strings.Join(second, "")
		checkNum, err = strconv.Atoi(b)
		if err != nil {
			panic("Второй операнд может быть либо строкой, либо целым числом")
		} else if checkNum < 1 || checkNum > 10 {
			panic("Второй операнд должен быть целым числом от 1 до 10")
		}
	}

	if len(a) < 1 || len(a) > 10 || len(b) < 1 || len(b) > 10 {
		panic("Операнды должны быть в диапазоне от 1 до 10 символов, не включая кавычки")
	}

	return a, b
}

func calculate() string {
	switch strSplit[sign] {
	case "+":
		res = a + b
	case "-":
		if strings.Contains(a, b) {
			res = strings.Replace(a, b, "", 1)
		} else {
			res = a
		}
	case "*":
		if checkNum != 0 {
			res = strings.Repeat(a, checkNum)
		} else {
			panic("При операции умножения второй аргумент может быть только числом")
		}
	case "/":
		if checkNum != 0 {
			checkNum = len(first) / checkNum
			res = strings.Join(first[:checkNum], "")
		} else {
			panic("При операции деления второй аргумент может быть только числом")
		}
	default:
		panic("Операция невозможна")
	}

	for index, _ := range res {
		if index > 39 {
			res1 = res[:40]
			resfinish = res1 + "..."
		} else {
			resfinish = res
		}
	}
	return resfinish
}
