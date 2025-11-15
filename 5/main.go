package main

type customError struct { // Структура customError имеет одно поле -- msg
	msg string
}

func (e *customError) Error() string { // метод для возврата сообщения об ошибке
	return e.msg
}

func test() *customError { // конструктор customError
	// ... do something
	return nil
}

func main() {
	var err error
	err = test()    // Передаем значение nil в переменную err
	if err != nil { // err не равно nil потому что err теперь имеет тип *customError, даже несмотря на nil (только когда переменная интерфейсного типа не имеет типа, реализующего интерфейс, и значения, она равна nil)
		println("error")
		return
	}
	println("ok")
}
