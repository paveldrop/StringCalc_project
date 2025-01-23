package validator

import (
	"fmt"
	"strconv"
)

const (
	errorStr          = "Первый член выражения не является строкой или не обозначено кавычками.\nПриложение завершит свою работу."
	lenStr            = "Строка более 10 символов.\nПриложение завершит свою работу."
	errorSign         = "Знак не определен.\nПриложение завершит свою работу."
	errorInt          = "Член выражения не является целочисленным или строкой.\nПриложение завершит свою работу."
	errorLsGtInt      = "Число больше 10 или меньше 1.\nПриложение завершит свою работу."
	unknownType       = "Неизвестный тип результата для второго члена выражения\n"
	templateErrString = "Невозможна калькуляция типов в выражении: %T %v %T\n"
)

type MemberTwo interface{}

func Validator(var1, var2, var3 string) (string, string, interface{}, error) {
	memberOne, err := validateStr(var1)
	if err != nil {
		return "", "", "", err
	}
	sign, err := validateSign(var2)
	if err != nil {
		return "", "", "", err
	}
	MemberTwo, err := validateStrOrNum(var3)
	if err != nil {
		return "", "", "", err
	}
	err = validateCalculation(memberOne, sign, MemberTwo)
	if err != nil {
		return "", "", "", err
	}
	return memberOne, sign, MemberTwo, nil
}

func validateStr(var1 string) (string, error) {
	n := len(var1)
	// fmt.Println(n)
	if var1[0] != '"' || var1[n-1] != '"' {
		return "", fmt.Errorf(errorStr)
	} else if n-2 > 10 {
		return "", fmt.Errorf(lenStr)
	} else {
		return var1[1 : n-1], nil
	}

}

func validateSign(varSign string) (string, error) {
	sign := ""
	switch varSign {
	case "+":
		sign = "+"
	case "-":
		sign = "-"
	case "/":
		sign = "/"
	case "*":
		sign = "*"
	default:
		return "", fmt.Errorf(errorSign)
	}
	return sign, nil
}

func validateNum(num string) error {
	num1, err := strconv.Atoi(num)
	if err != nil {
		return fmt.Errorf(errorInt)
	}
	if num1 < 1 || num1 > 10 {
		return fmt.Errorf(errorLsGtInt)
	}
	return nil
}

func validateStrOrNum(input string) (interface{}, error) {
	var result interface{}
	err := validateNum(input)
	if err == nil {
		result, _ = strconv.Atoi(input)
	} else {
		strResult, err := validateStr(input)
		if err == nil {
			result = strResult
		} else {
			return nil, err
		}
	}

	return result, nil
}

func validateCalculation(memberOne, sign string, memberTwo interface{}) error {
	strError := fmt.Sprintf(templateErrString, memberOne, sign, memberTwo)
	switch memberTwo.(type) {
	case string:
		if sign == "/" || sign == "*" {
			return fmt.Errorf(strError)
		}
	case int:
		if sign == "+" || sign == "-" {
			return fmt.Errorf(strError)
		}
	default:
		return fmt.Errorf(unknownType)
	}
	return nil
}
