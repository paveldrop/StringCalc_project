package calculate

import "fmt"

func OperationCalc(memb1, sign string, memb2 interface{}) string {
	switch sign {
	case "+":
		fmt.Println("operation add")
		return operationAdd(memb1, memb2)
	case "-":
		fmt.Println("operation remove")

	case "/":
		fmt.Println("operation division")
		return operationDiv(memb1, memb2)
	case "*":
		fmt.Println("operation multiplication")
		return operationMultipl(memb1, memb2)
	default:
		return ""
	}
	return ""
}

func operationAdd(memb1 string, memb2 interface{}) string {
	memb1 += memb2.(string)
	return memb1
}

// func operationRemove(memb1 string, memb2 interface{}) string {

// }

func operationMultipl(memb1 string, memb2 interface{}) string {
	num := memb2.(int)
	temp := memb1
	for i := 0; i < num; i++ {
		memb1 += temp
	}
	return memb1
}

func operationDiv(memb1 string, memb2 interface{}) string {
	n := len(memb1)
	num := memb2.(int)
	return memb1[:n/num]
}
