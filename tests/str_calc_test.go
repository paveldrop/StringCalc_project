package tests

import (
	"fmt"
	calc "str_calc/calculate"
	valid "str_calc/validator"
	"testing"
)

func TestCase1(t *testing.T) {
	inputStr := "\"String\""
	inputSign := "/"
	inputStr2 := "2"
	expected := "Str"
	var1, sign, var2, err := valid.Validator(inputStr, inputSign, inputStr2)
	if err != nil {
		t.Errorf("Error on validator")
	}
	result := calc.OperationCalc(var1, sign, var2)
	fmt.Println(result)
	if result != expected {
		t.Errorf("Error in calc")
	}
}

func TestCase2(t *testing.T) {
	inputStr := "\"String\""
	inputSign := "+"
	inputStr2 := "\"str\""
	expected := "Stringstr"
	var1, sign, var2, err := valid.Validator(inputStr, inputSign, inputStr2)
	if err != nil {
		t.Errorf("Error on validator")
	}
	result := calc.OperationCalc(var1, sign, var2)
	fmt.Println(result)
	if result != expected {
		t.Errorf("Error in calc")
	}
}

func TestCase3(t *testing.T) {
	inputStr := "\"Hi World!\""
	inputSign := "-"
	inputStr2 := "\"World!\""
	expected := "Hi "
	var1, sign, var2, err := valid.Validator(inputStr, inputSign, inputStr2)
	if err != nil {
		t.Errorf("Error on validator")
	}
	result := calc.OperationCalc(var1, sign, var2)
	fmt.Println(result)
	if result != expected {
		t.Errorf("Error in calc")
	}
}

func TestCase4(t *testing.T) {
	inputStr := "\"Bye-bye!\""
	inputSign := "-"
	inputStr2 := "\"World!\""
	expected := "Bye-bye!"
	var1, sign, var2, err := valid.Validator(inputStr, inputSign, inputStr2)
	if err != nil {
		t.Errorf("Error on validator")
	}
	result := calc.OperationCalc(var1, sign, var2)
	fmt.Println(result)
	if result != expected {
		t.Errorf("Error in calc")
	}
}

func TestCase5(t *testing.T) {
	inputStr := "\"Golang\""
	inputSign := "*"
	inputStr2 := "5"
	expected := "GolangGolangGolangGolangGolang"
	var1, sign, var2, err := valid.Validator(inputStr, inputSign, inputStr2)
	if err != nil {
		t.Errorf("Error on validator")
	}
	result := calc.OperationCalc(var1, sign, var2)
	fmt.Println(result)
	if result != expected {
		t.Errorf("Error in calc\n%s\n%s", expected, result)
	}
}

func TestCase6(t *testing.T) {
	inputStr := "\"Example!!!\""
	inputSign := "/"
	inputStr2 := "3"
	expected := "Exa"
	var1, sign, var2, err := valid.Validator(inputStr, inputSign, inputStr2)
	if err != nil {
		t.Errorf("Error on validator")
	}
	result := calc.OperationCalc(var1, sign, var2)
	fmt.Println(result)
	if result != expected {
		t.Errorf("Error in calc\n%s\n%s", expected, result)
	}
}
