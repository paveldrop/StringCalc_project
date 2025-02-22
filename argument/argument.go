package argument

import (
	"os"
)

/*
Калькулятор умеет выполнять операции сложения строк,
вычитания строки из строки, умножения строки на число
и деления строки на число: "a" + "b", "a" - "b", "a" * b, "a" / b.
Данные передаются в одну строку (смотри пример ниже).
Решения, в которых каждая строка, число и арифметическая операция передаются
с новой строки, считаются неверными.

Значения строк, передаваемых в выражении, выделяются двойными кавычками.

Результатом сложения двух строк является строка, состоящая из переданных.

Результатом деления строки на число n является строка в n раз
короче исходной (смотри пример).

Результатом умножения строки на число n является строка,
в которой переданная строка повторяется ровно n раз.

Результатом вычитания строки из строки является строка,
в которой удалена переданная подстрока или сама исходная строка,
если в неё нет вхождения вычитаемой строки (смотри пример).

Калькулятор должен принимать на вход числа от 1 до 10 включительно,
не более, и строки длиной не более 10 символов.
Если строка, полученная в результате работы приложения,
длиннее 40 символов, то в выводе после 40 символа должны стоять три точки (...).

Калькулятор умеет работать только с целыми числами.

Первым аргументом выражения, подаваемым на вход, должна быть строка.
При вводе пользователем выражения вроде 3 + "hello", калькулятор
должен выдать панику и прекратить свою работу.

При вводе пользователем неподходящих чисел, строк или
неподдерживаемых операций (например, деление строки на строку)
приложение выдаёт панику и завершает свою работу.

При вводе пользователем выражения, не соответствующего одной из вышеописанных
арифметических операций, приложение выдаёт панику и завершает свою работу.
*/

const (
	countArgs = "Аргументов больше или меньше чем 3.\nПриложение завершит свою работу."
)

func GetArgs() (string, string, string) {
	arguments := os.Args[1:]
	if len(arguments) != 3 {
		panic(countArgs)
	}

	return arguments[0], arguments[1], arguments[2]
}
