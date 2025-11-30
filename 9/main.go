package main

import (
	"fmt"
	"unicode"
)

func main() {

	a := "a4bc2d5e"

	/// Проверка работоспособности функции

	// b := "abcd"
	// c := "45"
	// d := ""
	// e := "ab\\4\\5"
	// f := "qwe\\45"

	// answA := "aaaabccddddde"
	// answB := "abcd"
	// answCD := ""

	// fmt.Println(a, b, c, d, e, f, answA, answCD, answB)

	ea := extractString(a)
	if ea == "" {
		fmt.Println("Распаккованная строка: '' ")
	} else {
		fmt.Println("Распаккованная строка:", ea)
	}

}

func extractString(s string) string {
	sType := getStringType(s)
	switch sType {
	case "emptyType":
		fmt.Println("Была введена пустая строка. Распакованная строка тоже является пустой.")
		return ""
	case "OnlyLetterType":
		fmt.Println("Строка из букв. Распакованная строка идентчина этой.")
		return s
	case "onlyNumsInvalidType":
		fmt.Println("Ошибка. Введённая строка состоит только из цифр.")
		return ""
	default:
		result := doNormalExtraction(s)
		return result
	}
}

func doNormalExtraction(s string) string {
	result := ""
	lengthString := len(s)
	symbols := []rune(s)
	isEscape := false
	var lastChar rune

	for i, symb := range s {
		if symb == '\\' {
			isEscape = true
			continue
		}

		if unicode.IsDigit(symb) { // цифра
			if isEscape { // если цифра экранирована , добалвяем как обычный
				result += string(symb)
				lastChar = symb
				isEscape = false
			} else {
				if (i != 0) && (unicode.IsLetter(lastChar)) { // число без \
					result += multipleSymbol(lastChar, runeToInt(symb))
				} else if unicode.IsDigit(lastChar) { // если прилетела цифра, значит она была экранирована, значит добалвляем её нужное количество раз (не считая добавленную цифру)
					result += multipleSymbol(lastChar, runeToInt(symb)-1)
				}
			}
		} else if unicode.IsLetter(symb) {
			if isEscape { // Экранированную букву добавляем как обычно
				result += string(symb)
				lastChar = symb
				isEscape = false
			} else if (i != lengthString-1) && (unicode.IsDigit(symbols[i+1])) { // Случаи по типу a5 -- сохраняем букву и размножим когда встретим цифру
				lastChar = symb
				continue
			} else if i == lengthString-1 { // Последняя буква
				result += string(symb)
			} else if (i != lengthString-1) && (!unicode.IsDigit(symbols[i+1])) { // Если за буквой идет буква, а не цифра
				result += string(symb)
				lastChar = symb
			}
		}
	}

	return result
}

func multipleSymbol(symb rune, count int) string { // Размножаем буквы и цифры
	res := ""
	for i := 0; i < count; i++ {
		res += string(symb)
	}
	return res
}

func getStringType(s string) string { // Определяем тип строки
	if len(s) == 0 { // пустая
		return "emptyType"
	}

	numCounter := 0
	letterCounter := 0

	for _, symb := range s {
		if unicode.IsDigit(symb) {
			numCounter++
		} else if unicode.IsLetter(symb) {
			letterCounter++
		}
	}

	if numCounter == len(s) { // чисто из цифр
		return "onlyNumsInvalidType"
	} else if letterCounter == len(s) { // чисто из букв
		return "OnlyLetterType"
	} else { // общего типа
		return "normalType"
	}
}

func runeToInt(r rune) int { // Переводим цифру-символ в числовой тип через код в таблице символов
	return int(r - '0')
}
