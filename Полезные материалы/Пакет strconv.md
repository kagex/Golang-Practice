# strconv

Пакет `strconv` предоставляет функции для преобразования строк в базовые типы данных и обратно.

## Числовые преобразования

### `strconv.Atoi(s string) (int, error)`

Преобразует строку в целое число типа `int`. Эквивалентно `ParseInt(s, 10, 0)`.

```go
i, err := strconv.Atoi("42")
// i = 42, err = nil
```

### `strconv.ParseBool(str string) (bool, error)`

Парсит строку в boolean. Допустимые значения: `1`, `t`, `T`, `TRUE`, `true`, `0`, `f`, `F`, `FALSE`, `false`.

```go
b, err := strconv.ParseBool("true")
// b = true, err = nil
```

### `strconv.ParseFloat(s string, bitSize int) (float64, error)`

Парсит строку в число с плавающей точкой. `bitSize` может быть 32 или 64.

```go
f, err := strconv.ParseFloat("3.1415", 64)
// f = 3.1415, err = nil
```

### `strconv.ParseInt(s string, base int, bitSize int) (int64, error)`

Парсит строку в целое число с указанным основанием системы счисления и размером бита.

```go
n, err := strconv.ParseInt("1010", 2, 64)
// n = 10, err = nil

n, err := strconv.ParseInt("FF", 16, 64)
// n = 255, err = nil
```

### `strconv.ParseUint(s string, base int, bitSize int) (uint64, error)`

Парсит строку в беззнаковое целое число с указанным основанием и размером бита.

```go
n, err := strconv.ParseUint("1010", 2, 64)
// n = 10, err = nil

u, _ := strconv.ParseUint("42", 10, 64)
// u = 42
```

### `strconv.FormatBool(b bool) string`

Преобразует boolean в строку ("true" или "false").

```go
s := strconv.FormatBool(true)
// s = "true"
```

### `strconv.FormatFloat(f float64, fmt byte, prec, bitSize int) string`

Преобразует `float64` в строку с указанным форматом и точностью.

Параметр `fmt`:
- `'b'` — экспоненциальная форма без дробной части: `-123p-45`
- `'e'` — экспоненциальная форма: `-1.234456e+78`
- `'E'` — экспоненциальная форма: `-1.234456E+78`
- `'f'` — без экспоненты: `123.456`
- `'g'` — как `'e'`, но без нулей в конце
- `'G'` — как `'E'`, но без нулей в конце
- `'x'` — шестнадцатеричное представление с p экспонентой (например `0x1.9p+2`)
- `'X'` — шестнадцатеричное представление с P экспонентой (например `0X1.9P+2`)

Параметр `prec`:
- Для `'e'`, `'E'`, `'f'`: количество знаков после точки
- Для `'g'`, `'G'`: максимальное количество значащих цифр
- Для `'b'`: игнорируется
- `-1`: использует минимально необходимое количество цифр

```go
s := strconv.FormatFloat(3.14, 'f', 2, 64)
// s = "3.14"

s := strconv.FormatFloat(123.456, 'g', 4, 64)
// s = "123.5"
```

### `strconv.FormatInt(i int64, base int) string`

Преобразует `int64` в строку с указанным основанием системы счисления (от 2 до 36).

```go
s := strconv.FormatInt(10, 2)
// s = "1010"

s := strconv.FormatInt(255, 16)
// s = "ff"
```

### `strconv.Itoa(i int) string`

Преобразует `int` в строку. Эквивалентно `FormatInt(int64(i), 10)`.

```go
s := strconv.Itoa(42)
// s = "42"
```

### `strconv.ParseComplex(s string, bitSize int) (complex128, error)`

Парсит строку в комплексное число. Добавлено в go1.15.

```go
c, err := strconv.ParseComplex("1+2i", 128)
// c = (1+2i), err = nil
```

### `strconv.FormatComplex(c complex128, fmt byte, prec, bitSize int) string`

Преобразует комплексное число в строку. Добавлено в go1.15.

```go
s := strconv.FormatComplex(1+2i, 'E', -1, 128)
// s = "(1+2i)" или "(1E+00+2E+00i)" в зависимости от формата
```

## Функции для работы с буферами (Append*)

### `strconv.AppendBool(dst []byte, b bool) []byte`

Добавляет "true" или "false" к буферу.

```go
buf := []byte{}
buf = strconv.AppendBool(buf, true)
// buf = []byte("true")
```

### `strconv.AppendFloat(dst []byte, f float64, fmt byte, prec, bitSize int) []byte`

Добавляет отформатированное число с плавающей точкой к буферу.

```go
buf := []byte{}
buf = strconv.AppendFloat(buf, 3.14, 'f', 2, 64)
// buf = []byte("3.14")
```

### `strconv.AppendInt(dst []byte, i int64, base int) []byte`

Добавляет отформатированное целое число к буферу.

```go
buf := []byte{}
buf = strconv.AppendInt(buf, -42, 16)
// buf = []byte("-2a")
```

### `strconv.AppendUint(dst []byte, i uint64, base int) []byte`

Добавляет отформатированное беззнаковое целое число к буферу.

```go
buf := []byte{}
buf = strconv.AppendUint(buf, 42, 16)
// buf = []byte("2a")
```

### `strconv.AppendQuote(dst []byte, s string) []byte`

Добавляет к буферу двойную кавычку Go string literal.

```go
buf := []byte{}
buf = strconv.AppendQuote(buf, `"Fran & Freddie's Diner"`)
// buf = []byte("\"Fran & Freddie's Diner\"")
```

### `strconv.AppendQuoteRune(dst []byte, r rune) []byte`

Добавляет к буферу одинарную кавычку Go character literal.

```go
buf := []byte{}
buf = strconv.AppendQuoteRune(buf, '☺')
// buf = []byte("'☺'")
```

### `strconv.AppendQuoteRuneToASCII(dst []byte, r rune) []byte`

Добавляет к буферу ASCII-safe character literal.

```go
buf := []byte{}
buf = strconv.AppendQuoteRuneToASCII(buf, '☺')
// buf = []byte("'\\u263a'")
```

### `strconv.AppendQuoteRuneToGraphic(dst []byte, r rune) []byte`

Добавляет к буферу graphic character literal. Добавлено в go1.6.

### `strconv.AppendQuoteToASCII(dst []byte, s string) []byte`

Добавляет к буферу ASCII-safe string literal.

```go
buf := []byte{}
buf = strconv.AppendQuoteToASCII(buf, `"Fran & Freddie's Diner"`)
// buf = []byte("\"Fran & Freddie's Diner\"")
```

### `strconv.AppendQuoteToGraphic(dst []byte, s string) []byte`

Добавляет к буферу graphic string literal. Добавлено в go1.6.

## Строковые преобразования (Quote/Unquote)

### `strconv.Quote(s string) string`

Возвращает двойную кавычку Go string literal.

```go
s := strconv.Quote("Hello, 世界")
// s = `"Hello, 世界"`
```

### `strconv.QuoteToASCII(s string) string`

Возвращает ASCII-safe string literal.

```go
s := strconv.QuoteToASCII("Hello, 世界")
// s = `"Hello, \u4e16\u754c"`
```

### `strconv.QuoteToGraphic(s string) string`

Возвращает graphic string literal. Добавлено в go1.6.

### `strconv.QuoteRune(r rune) string`

Возвращает одинарную кавычку Go character literal.

```go
s := strconv.QuoteRune('☺')
// s = `'☺'`
```

### `strconv.QuoteRuneToASCII(r rune) string`

Возвращает ASCII-safe character literal.

```go
s := strconv.QuoteRuneToASCII('☺')
// s = `'\u263a'`
```

### `strconv.QuoteRuneToGraphic(r rune) string`

Возвращает graphic character literal. Добавлено в go1.6.

### `strconv.Unquote(s string) (string, error)`

Интерпретирует Go string literal и возвращает значение.

```go
s, err := strconv.Unquote(`"Hello"`)
// s = "Hello", err = nil
```

### `strconv.UnquoteChar(s string, quote byte) (value rune, multibyte bool, tail string, err error)`

Декодирует первый символ в escaped string.

```go
r, multibyte, tail, err := strconv.UnquoteChar(`\n`, '"')
// r = '\n', multibyte = false, tail = "", err = nil
```

### `strconv.QuotedPrefix(s string) (string, error)`

Возвращает quoted string в начале s. Добавлено в go1.17.

```go
q, err := strconv.QuotedPrefix(`"double-quoted" with text`)
// q = `"double-quoted"`, err = nil
```

### `strconv.CanBackquote(s string) bool`

Проверяет, может ли строка быть представлена в виде backquoted string.

```go
b := strconv.CanBackquote("Fran & Freddie's Diner ☺")
// b = true

b = strconv.CanBackquote("`can't backquote this`")
// b = false
```

## Функции проверки

### `strconv.IsPrint(r rune) bool`

Проверяет, является ли символ печатаемым в Go.

```go
b := strconv.IsPrint('☺')
// b = true
```

### `strconv.IsGraphic(r rune) bool`

Проверяет, является ли символ графическим (включает пробел). Добавлено в go1.6.

```go
b := strconv.IsGraphic('☘')
// b = true
```

## Типы ошибок

### `strconv.NumError`

Тип ошибки для ошибок преобразования чисел.

```go
type NumError struct {
	Func string // имя функции, вызвавшей ошибку
	Num  string // входная строка
	Err  error  // описание ошибки
}
```

### `strconv.ErrRange`

Значение вне допустимого диапазона.

### `strconv.ErrSyntax`

Неправильный синтаксис.

## Константы

### `strconv.IntSize`

Размер `int` и `uint` в битах.

## Примеры использования

### Преобразование строки в число

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", n)
}
```

### Преобразование числа в строку

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(42)
	fmt.Println("Result:", s)
}
```

### Парсинг чисел в разных системах счисления

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	n1, _ := strconv.ParseInt("1010", 2, 64)
	fmt.Println("Binary:", n1)

	n2, _ := strconv.ParseInt("FF", 16, 64)
	fmt.Println("Hex:", n2)

	n3, _ := strconv.ParseInt("777", 8, 64)
	fmt.Println("Octal:", n3)

	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println("Uint:", u)
}
```

### Форматирование чисел

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := strconv.FormatInt(10, 2)
	fmt.Println("Binary:", s1)

	s2 := strconv.FormatFloat(3.14159, 'f', 2, 64)
	fmt.Println("Float:", s2)

	s3 := strconv.FormatFloat(123456.789, 'e', 3, 64)
	fmt.Println("Exponential:", s3)

	s4 := strconv.FormatFloat(1.5, 'x', -1, 64)
	fmt.Println("Hex float:", s4)
}
```

### Работа со строками в кавычках

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := strconv.Quote("hello\nworld")
	fmt.Println("Quoted:", s1)

	s2, _ := strconv.Unquote(`"hello\nworld"`)
	fmt.Println("Unquoted:", s2)

	s3 := strconv.QuoteToASCII("hello\nworld")
	fmt.Println("ASCII-safe:", s3)

	complexNum, _ := strconv.ParseComplex("1+2i", 128)
	fmt.Println("Complex:", complexNum)

	s4 := strconv.FormatComplex(1+2i, 'E', -1, 128)
	fmt.Println("Format complex:", s4)
}
```