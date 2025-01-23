package main

import (
	"fmt"
	arg "str_calc/argument"
	valid "str_calc/validator"
)

func main() {
	arg1, arg2, arg3 := arg.GetArgs()
	memb1, sign, memb2, err := valid.Validator(arg1, arg2, arg3)
	if err != nil {
		panic(err)
	}
	fmt.Println(memb1, sign, memb2)
	fmt.Println("Приложение завершено корректно")
}
