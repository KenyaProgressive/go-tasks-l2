# go-tasks-l2
Задачи уровня L2 -- WB.Техношкола | Tasks from L2 -- WB.Technoschool

# Задание 1

Что выведет программа?

Объясните вывод программы.

package main

import "fmt" 

func main() {
  a := [5]int{76, 77, 78, 79, 80}
  var b []int = a[1:4]
  fmt.Println(b)
}

**Запуск решения**: `go run 1/main.go` | `make task1`

# Задание 2


Что выведет программа?

Объяснить порядок выполнения defer функций и итоговый вывод.

package main

import "fmt"

func test() (x int) {
  defer func() {
    x++
  }()
  x = 1
  return
}

func anotherTest() int {
  var x int
  defer func() {
    x++
  }()
  x = 1
  return x
}

func main() {
  fmt.Println(test())
  fmt.Println(anotherTest())
}

**Запуск решения**: `go run 2/main.go` | `make task2`

# Задание 3

Что выведет программа?

Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

package main

import (
  "fmt"
  "os"
)

func Foo() error {
  var err *os.PathError = nil
  return err
}

func main() {
  err := Foo()
  fmt.Println(err)
  fmt.Println(err == nil)
}

**Запуск решения**: `go run 3/main.go` | `make task3`


# Задание 4

Что выведет программа?

Объяснить вывод программы.

func main() {
  ch := make(chan int)
  go func() {
    for i := 0; i < 10; i++ {
    ch <- i
  }
}()
  for n := range ch {
    println(n)
  }
}


**Запуск решения**: `go run 4/main.go` | `make task4`

# Задание 5

Что выведет программа?

Объяснить вывод программы.

package main

type customError struct {
  msg string
}

func (e *customError) Error() string {
  return e.msg
}

func test() *customError {
  // ... do something
  return nil
}

func main() {
  var err error
  err = test()
  if err != nil {
    println("error")
    return
  }
  println("ok")
}

**Запуск решения**: `go run 5/main.go` | `make task5`

# Задание 6

Что выведет программа?

Объяснить поведение срезов при передаче их в функцию.

package main

import (
  "fmt"
)

func main() {
  var s = []string{"1", "2", "3"}
  modifySlice(s)
  fmt.Println(s)
}

func modifySlice(i []string) {
  i[0] = "3"
  i = append(i, "4")
  i[1] = "5"
  i = append(i, "6")
}

**Запуск решения**: `go run 6/main.go` | `make task6`

# Задание 7

Что выведет программа?

Объяснить работу конвейера с использованием select.

package main

import (
  "fmt"
  "math/rand"
  "time"
)

func asChan(vs ...int) <-chan int {
  c := make(chan int)
  go func() {
    for _, v := range vs {
      c <- v
      time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    }
  close(c)
}()
  return c
}

func merge(a, b <-chan int) <-chan int {
  c := make(chan int)
  go func() {
    for {
      select {
        case v, ok := <-a:
          if ok {
            c <- v
          } else {
            a = nil
          }
        case v, ok := <-b:
          if ok {
            c <- v
          } else {
            b = nil
          }
        }
        if a == nil && b == nil {
          close(c)
          return
        }
     }
   }()
  return c
}

  func main() {
    rand.Seed(time.Now().Unix())
    a := asChan(1, 3, 5, 7)
    b := asChan(2, 4, 6, 8)
    c := merge(a, b)
    for v := range c {
    fmt.Print(v)
  }
}

**Запуск решения**: `go run 7/main.go` | `make task7`

# Задание 8

Создать программу, печатающую точное текущее время с использованием NTP-сервера.

Реализовать проект как модуль Go.

Использовать библиотеку ntp для получения времени.

Программа должна выводить текущее время, полученное через NTP (Network Time Protocol).

Необходимо обрабатывать ошибки библиотеки: в случае ошибки вывести её текст в STDERR и вернуть ненулевой код выхода.

Код должен проходить проверки (vet и golint), т.е. быть написан идиоматически корректно.

**Запуск решения**: `cd task8-ntp && go run . && cd .. ` | `make task8`
**Проверка линтерами**: `make task8-pretty-check`

# Задание 9

Написать функцию Go, осуществляющую примитивную распаковку строки, содержащей повторяющиеся символы/руны.

Примеры работы функции:

Вход: "a4bc2d5e"
Выход: "aaaabccddddde"

Вход: "abcd"
Выход: "abcd" (нет цифр — ничего не меняется)

Вход: "45"
Выход: "" (некорректная строка, т.к. в строке только цифры — функция должна вернуть ошибку)

Вход: ""
Выход: "" (пустая строка -> пустая строка)

Дополнительное задание
Поддерживать escape-последовательности вида \:

Вход: "qwe\4\5"
Выход: "qwe45" (4 и 5 не трактуются как числа, т.к. экранированы)

Вход: "qwe\45"
Выход: "qwe44444" (\4 экранирует 4, поэтому распаковывается только 5)

Требования к реализации
Функция должна корректно обрабатывать ошибочные случаи (возвращать ошибку, например, через error), и проходить unit-тесты.

Код должен быть статически анализируем (vet, golint).


**Запуск решения**: `cd task9 && go run . && cd .. ` | `make task9`
**Проверка линтерами**: `make task9-pretty-check`