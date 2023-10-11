Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


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
```

Ответ:
```
Вывод программы:
2
1

Объяснение:
В функции test() (x int) мы используем именованый результат
поэтому defer выполняется до возврата значения из функции

В функции anotherTest() int мы сначала возвращаем значение x
и только потом вызывается defer
```
