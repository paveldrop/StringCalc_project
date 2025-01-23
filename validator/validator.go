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
	errorTwoMemb      = "Ошибка при вводе второго члена выражения, не является допуститмым числом или строкой в кавычках.\nПриложение завершит свою работу."
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

func validateNum(num string) (int, error) {
	result, err := strconv.Atoi(num)
	if err != nil {
		return 0, fmt.Errorf(errorInt)
	}
	if result < 1 || result > 10 {
		return 0, fmt.Errorf(errorLsGtInt)
	}
	return result, nil
}

func validateStrOrNum(memb2 string) (interface{}, error) {
	var result interface{}
	fmt.Println(memb2)
	if memb2[0] == '"' && memb2[len(memb2)-1] == '"' {
		result, err := validateStr(memb2)
		if err == nil {
			return result, nil
		}
	} else {
		result, err := validateNum(memb2)
		if err == nil {
			return result, nil
		}
	}
	return result, fmt.Errorf(errorTwoMemb)
}

func validateCalculation(memberOne, sign string, memberTwo interface{}) error {
	switch memberTwo.(type) {
	case string:
		if sign == "/" || sign == "*" {
			return fmt.Errorf(templateErrString, memberOne, sign, memberTwo)
		}
	case int:
		if sign == "+" || sign == "-" {
			return fmt.Errorf(templateErrString, memberOne, sign, memberTwo)
		}
	default:
		return fmt.Errorf(unknownType)
	}
	return nil
}
