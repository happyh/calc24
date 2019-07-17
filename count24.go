package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(a, b float64, fun byte) float64 {
	r := 0.0
	switch fun {
	case '+':
		r = float64(a + b)
	case '-':
		r = float64(a - b)
	case '*':
		r = float64(a * b)
	case '/':
		r = float64(a) / float64(b)
	}
	return r
}

func dsf2(value []float64, str []string) bool {
	for i, v1 := range value {
		for j, v2 := range value {
			if i != j {
				for _, f := range fun {
					v := calc(v1, v2, f)
					if len(value) == 2 && v == 24.0 {
						fmt.Printf("(%s%s%s)\n", str[i], string(f), str[j])
						return true
					} else if len(value) > 2 {
						begin := i
						end := j
						if i > j {
							begin = j
							end = i
						}

						s := make([]float64, 0)
						s = append(s, value[0:begin]...)
						s = append(s, value[begin+1:end]...)
						s = append(s, value[end+1:]...)
						s = append(s, v)

						s2 := make([]string, 0)
						s2 = append(s2, str[0:begin]...)
						s2 = append(s2, str[begin+1:end]...)
						s2 = append(s2, str[end+1:]...)
						funstr := fmt.Sprintf("(%s%s%s)", str[i], string(f), str[j])
						s2 = append(s2, funstr)

						if dsf2(s, s2) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}

var fun [4]byte

func main() {
	fun[0] = '+'
	fun[1] = '-'
	fun[2] = '*'
	fun[3] = '/'

	var str string

	for str == "" {
		fmt.Print("请输入数字(q for quit):")
		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()

		str = string(data)

		if str == "q" {
			break
		}

		sInt := make([]float64, 0)
		sStr := make([]string, 0)

		checked := true
		ints := strings.Split(str, " ")
		for _, v := range ints {
			if v == "" {
				continue
			}
			num, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("输入不正确：", v, ",err:", err.Error())
				checked = false
				break
			}
			sInt = append(sInt, float64(num))
			sStr = append(sStr, v)
		}
		if checked {
			if !dsf2(sInt, sStr) {
				fmt.Println("不能生成24！")
			}
		}
		str = ""
	}
}
