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