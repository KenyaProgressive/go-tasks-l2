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