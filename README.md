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