# Указатели (Pointers)

## Введение

Указатель — это переменная, которая хранит адрес памяти другой переменной. Указатели позволяют косвенно обращаться к данным, передавать данные по ссылке и создавать динамические структуры данных.

### Назначение указателей
- Эффективная передача больших структур в функции (без копирования)
- Изменение значений внутри функций
- Создание связанных структур данных (списки, деревья)
- Работа с низкоуровневыми операциями памяти

## Объявление и инициализация

### Операторы
- `&` — оператор взятия адреса (address operator)
- `*` — оператор разыменования (dereference operator)
- `*T` — тип указателя на `T`
- `nil` — нулевой указатель (zero value для указателей)

### Примеры

```go
var x int = 42
p := &x        // p хранит адрес x, тип p: *int
v := *p        // v получает значение по адресу p (42)

fmt.Println(x)  // 42
fmt.Println(*p) // 42
```

```go
var num int = 10
var ptr *int = &num

fmt.Println(num)    // 10
fmt.Println(ptr)    // 0xc0000100a0 (адрес в памяти)
fmt.Println(*ptr)   // 10 (значение по адресу)
```

### Использование `new`

```go
p := new(int)  // выделяет память для int, возвращает *int
*p = 42        // записывает значение 42 по адресу
fmt.Println(*p) // 42
```

## Основные операции

### Взятие адреса (`&`)

```go
x := 100
p := &x  // p: *int, хранит адрес x
```

### Разыменование (`*`)

```go
p := &x
fmt.Println(*p)  // читает значение по адресу
*p = 200         // записывает новое значение
```

### Проверка на `nil`

```go
var p *int
if p == nil {
    fmt.Println("Указатель не инициализирован")
}
```

## Указатели в функциях

### Передача по ссылке

```go
func increment(n *int) {
    *n = *n + 1
}

func main() {
    x := 5
    increment(&x)
    fmt.Println(x)  // 6
}
```

### Изменение структур

```go
type Person struct {
    Name string
    Age  int
}

func setAge(p *Person, age int) {
    p.Age = age
}

func main() {
    person := Person{Name: "Alice"}
    setAge(&person, 30)
    fmt.Println(person.Age)  // 30
}
```

## Указатели на структуры

### Доступ к полям

```go
type Point struct {
    X, Y int
}

p := &Point{X: 1, Y: 2}
fmt.Println(p.X)     // 1 (p.X — синтаксический сахар, Go автоматически разыменовывает указатель)
fmt.Println((*p).Y)  // 2 (эквивалентно p.X, явное разыменование)
```

## Указатели и срезы

### Срезы уже содержат указатель

```go
s := []int{1, 2, 3}
// Заголовок среза содержит указатель на массив данных
// При передаче в функцию копируется только заголовок среза
```

### Изменение элементов среза

```go
func double(nums []int) {
    for i := range nums {
        nums[i] *= 2
    }
}
// Элементы изменяются, так как заголовок среза содержит указатель на данные
```



## Указатели и массивы

### Массивы и указатели

```go
var arr [3]int = [3]int{1, 2, 3}
p := &arr  // p: *[3]int

fmt.Println((*p)[0])  // 1 (явное разыменование)
fmt.Println(p[1])     // 2 (синтаксический сахар: p[1] работает благодаря автоматическому разыменованию указателя на массив)
fmt.Println((*p)[2])  // 3
```

## Безопасные указатели (`unsafe.Pointer`)

### Пакет `unsafe`

```go
import "unsafe"

// Пример: преобразование указателей
var x int = 42
p := unsafe.Pointer(&x)
// p: unsafe.Pointer
```

### Ограничения

- `unsafe.Pointer` обходит проверку типов
- Использование может привести к неопределенному поведению
- Рекомендуется использовать только при крайней необходимости

## Сравнение указателей

```go
func main() {
    x, y := 10, 20
    px, py := &x, &y
    
    fmt.Println(px == py)  // false (разные адреса)
    fmt.Println(px != py)  // true
    
    var p *int
    fmt.Println(p == nil)  // true
}
```

## Двойные указатели

```go
x := 42
p := &x      // p: *int
pp := &p     // pp: **int

fmt.Println(x)   // 42
fmt.Println(*p)  // 42 (значение по адресу p)
fmt.Println(**pp) // 42 (значение по адресу, который хранит pp)
```

## Практические примеры

### Обмен значений

```go
func swap(a, b *int) {
    *a, *b = *b, *a
}

func main() {
    x, y := 1, 2
    swap(&x, &y)
    fmt.Println(x, y)  // 2 1
}
```

### Построение связного списка

```go
type Node struct {
    Value int
    Next  *Node
}

func main() {
    head := &Node{Value: 1}
    head.Next = &Node{Value: 2}
    head.Next.Next = &Node{Value: 3}
}
```

### Определение типа через указатель

```go
func getType(v any) {
    switch v.(type) {
    case *int:
        fmt.Println("Указатель на int")
    case *string:
        fmt.Println("Указатель на string")
    }
}

// Альтернатива: reflect
import "reflect"
x := 42
ptr := &x
t := reflect.TypeOf(ptr)  // t: *int
```

## Лучшие практики

1. **Используйте указатели** для больших структур при передаче в функции
2. **Используйте указатели** когда нужно изменить значение внутри функции
3. **Избегайте указателей** для простых типов, если изменение не требуется
4. **Проверяйте на `nil`** перед разыменованием
5. **Не используйте `unsafe`** без крайней необходимости
6. **Помните** что срезы и карты уже содержат указатели внутри
7. **Используйте `p.X`** вместо `(*p).X` для структур (синтаксический сахар Go автоматически разыменовывает указатели при доступе к полям)

## Частые ошибки

### Разыменование `nil`

```go
var p *int
fmt.Println(*p)  // panic: nil pointer dereference
```

### Взятие адреса константы или результата выражения

```go
const x = 42
p := &x  // ошибка: cannot take the address of x
```

## Ссылки

- [Official Go Documentation - Pointer types](https://go.dev/ref/spec#Pointer_types)
- [Effective Go - Pointers vs. Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go Specification - Address operators](https://go.dev/ref/spec#Address_operators)
