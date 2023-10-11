Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
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
```

Ответ:
```
Вывод программы:
[3 2 3]

Объяснение:
Изначально i указывает на ту же область памяти, что и s, поэтому
успешно меняется 0-й элемент массива, на который указывает s.
Затем мы вызываем функцию append и поскольку мы превышаем cap исходного слайса,
выделяется память под новый слайс с большей ёмкостью, append возвращает уже его
все изменения происходят уже с ним, однако, на слайс s это уже никак не влияет. 
```