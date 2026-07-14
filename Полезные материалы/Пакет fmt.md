# fmt

Пакет `fmt` реализует форматированный ввод-вывод с функциями, аналогичными `printf` и `scanf` из C. Вербосы формата производны от C, но проще.

## Печать

Четыре семьи функций печати определяются их местом назначения вывода:
- `Print`, `Println`, `Printf` — вывод в `os.Stdout`
- `Sprint`, `Sprintln`, `Sprintf` — возвращают строку
- `Fprint`, `Fprintln`, `Fprintf` — запись в `io.Writer`
- `Append`, `Appendln`, `Appendf` — добавление в байтовый срез

### Функции печати без формата

#### `fmt.Print(a ...any) (n int, err error)`

Форматирует операнды в формате по умолчанию и выводит в стандартный вывод. Добавляет пробелы между операндами только когда оба не являются строками. Не добавляет новую строку в конце.

```go
fmt.Print("Hello", " world")  // Hello world
```

#### `fmt.Println(a ...any) (n int, err error)`

Аналогично `Print`, но всегда добавляет пробелы между операндами и добавляет новую строку в конце.

```go
fmt.Println("Hello", "world")  // Hello world\n
```

#### `fmt.Sprint(a ...any) string`

Аналогично `Print`, но возвращает строку.

```go
s := fmt.Sprint("Hello", " world")  // "Hello world"
```

#### `fmt.Sprintf(format string, a ...any) string`

Форматирует по спецификатору формата и возвращает строку.

```go
s := fmt.Sprintf("%s is %d years old", "Kim", 22)  // "Kim is 22 years old"
```

#### `fmt.Fprintf(w io.Writer, format string, a ...any) (n int, err error)`

Форматирует по спецификатору и записывает в `io.Writer`.

```go
fmt.Fprintf(os.Stdout, "Value: %d\n", 42)
```

#### `fmt.Fprint(w io.Writer, a ...any) (n int, err error)`

Форматирует в формате по умолчанию и записывает в `io.Writer`.

```go
fmt.Fprint(os.Stdout, "Hello", " world")
```

### Функции для работы с буферами (Append*)

#### `fmt.Append(b []byte, a ...any) []byte`

Форматирует операнды в формате по умолчанию, добавляет к буферу и возвращает обновленный срез.

```go
buf := []byte{}
buf = fmt.Append(buf, "Hello", " world")  // buf = []byte("Hello world")
```

#### `fmt.Appendf(b []byte, format string, a ...any) []byte`

Форматирует по спецификатору, добавляет к буферу и возвращает обновленный срез.

```go
buf := []byte{}
buf = fmt.Appendf(buf, "%d", 42)  // buf = []byte("42")
```

#### `fmt.Appendln(b []byte, a ...any) []byte`

Форматирует в формате по умолчанию, добавляет пробелы, новую строку и возвращает обновленный срез.

```go
buf := []byte{}
buf = fmt.Appendln(buf, "Hello", "world")  // buf = []byte("Hello world\n")
```

## Спецификаторы формата

### Общие

- `%v` — значение в формате по умолчанию
- `%+v` — добавляет имена полей для структур
- `%#v` — представление Go-синтаксиса значения
- `%T` — представление типа значения
- `%%` — буквенный знак процента

### Булевы

- `%t` — `true` или `false`

### Целые числа

- `%b` — основание 2 (двоичное)
- `%c` — символ, представленный Unicode-кодовой точкой
- `%d` — основание 10 (десятичное)
- `%o` — основание 8 (восмеричное)
- `%O` — основание 8 с префиксом `0o`
- `%q` — одинарно-кавычечный литерал символа с экранированием Go-синтаксисом
- `%x` — основание 16, строчные буквы для a-f
- `%X` — основание 16, заглавные буквы для A-F
- `%U` — формат Unicode: `U+1234`

### Числа с плавающей точкой и комплексные

- `%b` — научная нотация без десятичной точки, степень двойки (например `-123456p-78`)
- `%e` — научная нотация (например `-1.234456e+78`)
- `%E` — научная нотация (например `-1.234456E+78`)
- `%f` — десятичная точка без экспоненты (например `123.456`)
- `%F` — синоним `%f`
- `%g` — `%e` для больших экспонент, иначе `%f`
- `%G` — `%E` для больших экспонент, иначе `%F`
- `%x` — шестнадцатеричная нотация (например `-0x1.23abcp+20`)
- `%X` — заглавная шестнадцатеричная нотация (например `-0X1.23ABCP+20`)

### Строки и срезы байт

- `%s` — неинтерпретированные байты строки или среза
- `%q` — строка в двойных кавычках с экранированием Go-синтаксисом
- `%x` — шестнадцатеричная, строчные буквы, два символа на байт
- `%X` — шестнадцатеричная, заглавные буквы, два символа на байт

### Комплексные числа

- `%v`, `%+v`, `%g`, `%G` — формат комплексных чисел в скобках: `(real+imaginaryi)`
- `%f`, `%e`, `%E` — формат для каждой компоненты комплексного числа
- `%b`, `%x`, `%X` — не поддерживаются для комплексных чисел (ошибка во время форматирования)

### Указатели

- `%p` — шестнадцатеричная нотация с префиксом `0x`

## Флаги

- `+` — всегда печатать знак для числовых значений; ASCII-only для `%q`
- `-` — заполнение пробелами справа (левое выравнивание)
- `#` — альтернативный формат:
  - добавляет `0b` для двоичного (`%#b`)
  - добавляет `0` для восмеричного (`%#o`)
  - добавляет `0x` или `0X` для шестнадцатеричного (`%#x`, `%#X`)
  - для `%q` — использует backticks если возможно (если `strconv.CanBackquote` возвращает `true`)
  - всегда печатать десятичную точку для `%e`, `%E`, `%f`, `%F`, `%g`, `%G`
  - не удалять trailing zeros для `%g` и `%G`
- ` ` (space) — оставить пробел для опущенного знака в числах (`% d`)
- `0` — заполнение ведущими нулями вместо пробелов

## Ширина и точность

Ширина указывается необязательным десятичным числом перед вербосом. Точность указывается после ширины точкой и десятичным числом.

Примеры:
- `%9f` — ширина 9, точность по умолчанию
- `%.2f` — ширина по умолчанию, точность 2
- `%9.2f` — ширина 9, точность 2

```go
fmt.Printf("%6.2f", 12.345)  // "  12.35"
```

Для `%g` и `%G` точность задает максимальное количество значащих цифр.

Для комплексных чисел ширина и точность применяются независимо к каждой компоненте.

## Примеры использования

### Базовое форматирование

```go
package main

import (
	"fmt"
)

func main() {
	integer := 23
	fmt.Println(integer)          // 23
	fmt.Printf("%v\n", integer)   // 23
	fmt.Printf("%d\n", integer)   // 23
	fmt.Printf("%T\n", integer)   // int

	truth := true
	fmt.Printf("%v %t\n", truth, truth)  // true true

	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// 42 42 2a 52 101010

	pi := 3.14159
	fmt.Printf("%v %g %.2f %e\n", pi, pi, pi, pi)
	// 3.14159 3.14159 3.14 3.141593e+00

	smile := '😀'
	fmt.Printf("%v %d %c %q %U\n", smile, smile, smile, smile, smile)
	// 128512 128512 😀 '😀' U+1F600
}
```

### Строки

```go
s := `foo "bar"`
fmt.Printf("%v %s %q %#q\n", s, s, s, s)
// foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
```

### Структуры и карты

```go
person := struct {
	Name string
	Age  int
}{"Kim", 22}

fmt.Printf("%v\n", person)   // {Kim 22}
fmt.Printf("%+v\n", person)  // {Name:Kim Age:22}
fmt.Printf("%#v\n", person)  // struct { Name string; Age int }{Name:"Kim", Age:22}

isLegume := map[string]bool{
	"peanut":    true,
	"dachshund": false,
}
fmt.Printf("%v\n", isLegume)   // map[dachshund:false peanut:true]
fmt.Printf("%#v\n", isLegume)  // map[string]bool{"dachshund":false, "peanut":true}
```

### Указатели и срезы

```go
pointer := &person
fmt.Printf("%v %p\n", pointer, pointer)
// &{Kim 22} 0x...

greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
fmt.Printf("%v %q\n", greats, greats)
// [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]

kGreats := greats[:3]
fmt.Printf("%v %q\n", kGreats, kGreats)
// [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"]
```

### Byte срезы

```go
cmd := []byte("a⌘")
fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
// [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
```

## Сканирование

Функции сканирования:
- `Scan`, `Scanf`, `Scanln` — чтение из `os.Stdin`
- `Fscan`, `Fscanf`, `Fscanln` — чтение из `io.Reader`
- `Sscan`, `Sscanf`, `Sscanln` — чтение из строки

### Функции сканирования

#### `fmt.Scan(a ...any) (n int, err error)`

Сканирует текст из стандартного ввода, сохраняет значения в аргументы.

#### `fmt.Sscanf(str string, format string, a ...any) (n int, err error)`

Сканирует текст из строки по формату.

#### `fmt.Fscanf(r io.Reader, format string, a ...any) (n int, err error)`

Сканирует текст из `io.Reader` по формату.

### Примеры сканирования

```go
var i int
var s string
n, err := fmt.Sscanf("42 hello", "%d %s", &i, &s)
// i = 42, s = "hello", n = 2
```

## Интерфейсы форматирования

### `fmt.Stringer`

```go
type Stringer interface {
	String() string
}
```

Метод `String()` вызывается при форматировании операнда.

### `fmt.Formatter`

```go
type Formatter interface {
	Format(state State, verb rune)
}
```

Позволяет контролировать formatting verbs и flags.

### `fmt.GoStringer`

```go
type GoStringer interface {
	GoString() string
}
```

Метод `GoString()` вызывается при использовании `%#v`.

### `fmt.Scanner`

```go
type Scanner interface {
	Scan(state ScanState, verb rune) error
}
```

Метод `Scan()` используется при сканировании текста.

## Функция Errorf

### `fmt.Errorf(format string, a ...any) error`

Форматирует сообщение и возвращает значение типа `error`.

```go
err := fmt.Errorf("user %q (id %d) not found", "bueller", 17)
// user "bueller" (id 17) not found
```

## Примеры полных программ

### Программа печати

```go
package main

import (
	"fmt"
)

func main() {
	name, age := "Kim", 22
	
	fmt.Print(name, " is ", age, " years old.\n")
	fmt.Println(name, "is", age, "years old.")
	fmt.Printf("%s is %d years old.\n", name, age)
}
```

### Программа сканирования

```go
package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	
	fmt.Print("Enter name and age: ")
	n, err := fmt.Scan(&name, &age)
	
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	fmt.Printf("Scanned %d items: %s is %d\n", n, name, age)
}
```

### Кастомный форматировщик

```go
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func (p Person) Format(state fmt.State, verb rune) {
	switch verb {
	case 'v':
		if state.Flag('+') {
			fmt.Fprintf(state, "Name: %s, Age: %d", p.Name, p.Age)
			return
		}
		fallthrough
	case 's':
		fmt.Fprintf(state, "%s (%d)", p.Name, p.Age)
	case 'q':
		fmt.Fprintf(state, "%q", fmt.Sprintf("%s (%d)", p.Name, p.Age))
	}
}

func main() {
	p := Person{"Kim", 22}
	fmt.Printf("%v\n", p)   // Kim (22)
	fmt.Printf("%+v\n", p)  // Name: Kim, Age: 22
	fmt.Printf("%q\n", p)   // "Kim (22)"
}