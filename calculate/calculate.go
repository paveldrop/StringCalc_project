package calculate

import (
	"fmt"
	"strings"
)

func OperationCalc(memb1, sign string, memb2 interface{}) string {
	switch sign {
	case "+":
		fmt.Println("operation add")
		return operationAdd(memb1, memb2)
	case "-":
		fmt.Println("operation remove")
		return operationRemove(memb1, memb2)
	case "/":
		fmt.Println("operation division")
		return operationDiv(memb1, memb2)
	case "*":
		fmt.Println("operation multiplication")
		return operationMultipl(memb1, memb2)
	default:
		return ""
	}
}

func operationAdd(memb1 string, memb2 interface{}) string {
	memb1 += memb2.(string)
	return memb1
}

func operationRemove(memb1 string, memb2 interface{}) string {
	subStr := memb2.(string)
	result := strings.Replace(memb1, subStr, "", 1)
	return result
}

func operationMultipl(memb1 string, memb2 interface{}) string {
	num := memb2.(int)
	temp := memb1
	result := ""
	for i := 0; i < num; i++ {
		result += temp
	}
	return result
}

func operationDiv(memb1 string, memb2 interface{}) string {
	n := len(memb1)
	num := memb2.(int)
	return memb1[:n/num]
}
