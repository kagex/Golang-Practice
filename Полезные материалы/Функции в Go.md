# Функции в Go

## Введение

Функции — это основной способ организации кода в Go. Они позволяют группировать логику, переиспользовать код и создавать абстракции. В Go функции являются **first-class citizens** — их можно передавать как значения, возвращать из других функций и присваивать переменным.

> 💡 **Когда использовать функции вместо методов:**
> - **Функции** — когда логика не привязана к конкретному типу
> - **Методы** — когда логика зависит от состояния типа

```go
func имя(параметры) (возвращаемые значения) {
    // тело функции
    return значения
}
```

> 💡 **Ключевая особенность Go:** Функции могут возвращать **несколько значений**, что активно используется для возврата результата и ошибки одновременно.

---

## Содержание

- [Введение](#введение)
- [Базовый синтаксис](#базовый-синтаксис)
- [Параметры и возвращаемые значения](#параметры-и-возвращаемые-значения)
- [Именованные возвращаемые значения](#именованные-возвращаемые-значения)
- [Variadic функции](#variadic-функции)
- [Анонимные функции](#анонимные-функции)
- [Замыкания](#замыкания-closures)
- [Generics](#generics-обобщения-go-118)
- [Функции как значения](#функции-как-значения)
- [Рекурсия](#рекурсия)
- [Defer](#defer-отложенные-вызовы)
- [Panic и Recover](#panic-и-recover)
- [Методы](#методы)
- [Init функции](#init-функции)
- [Таблица быстрого справочника](#таблица-быстрого-справочника)
- [Лучшие практики](#лучшие-практики)
- [Частые ошибки](#частые-ошибки)
- [Заключение](#заключение)
- [Ссылки](#ссылки)

---

## Базовый синтаксис

### Простая функция

```go
package main

import (
    "fmt"
)

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println(result) // 8
}
```

### Функция без параметров и возвращаемых значений

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello, World!")
}

func main() {
    greet() // Hello, World!
}
```

### Функция с несколькими возвращаемыми значениями

```go
package main

import "fmt"

func minMax(nums ...int) (int, int) {
    if len(nums) == 0 {
        return 0, 0
    }
    min, max := nums[0], nums[0]
    for _, n := range nums {
        if n < min {
            min = n
        }
        if n > max {
            max = n
        }
    }
    return min, max
}

func main() {
    min, max := minMax(5, 2, 8, 1, 9)
    fmt.Printf("Min: %d, Max: %d\n", min, max) // Min: 1, Max: 9
}
```

---

## Параметры и возвращаемые значения

### Несколько параметров одного типа

В Go можно указатьывать тип один раз для нескольких последовательных параметров:

```go
func multiply(a, b, c int) int {
    return a * b * c
}

// Эквивалентно:
func multiplyVerbose(a int, b int, c int) int {
    return a * b * c
}
```

### Несколько возвращаемых значений

Одна из ключевых особенностей Go — возможность возвращать несколько значений:

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Результат:", result) // 5
}
```

### Игнорирование возвращаемых значений

Используйте `_` для игнорирования ненужных значений:

```go
result, _ := divide(10, 2) // игнорируем ошибку
_, err := divide(10, 0)    // игнорируем результат
```

---

## Именованные возвращаемые значения

Можно давать имена возвращаемым значениям — они становятся переменными внутри функции:

```go
package main

import (
    "fmt"
    "log"
)

func calculate(x, y int) (sum, product int) {
    sum = x + y
    product = x * y
    return // "голый" return — возвращает sum и product
}

func main() {
    s, p := calculate(3, 4)
    fmt.Println("Sum:", s, "Product:", p) // Sum: 7 Product: 12
}
```

### Классический паттерн с defer и именованными возвратами

```go
package main

import (
    "fmt"
    "log"
)

func trace(name string) func() {
    log.Printf("ENTER %s", name)
    return func() {
        log.Printf("EXIT %s", name)
    }
}

func foo() (result int) {
    defer trace("foo")()
    result = 42
    return
}
```

---

### ⚠️ Когда использовать именованные возвраты

✅ **Хорошо для:**
- Документации (имена появляются в godoc)
- Коротких функций
- Defer с модификацией возвращаемых значений

❌ **Плохо для:**
- Длинных функций (путаница с переменными)
- Когда имена не добавляют ясности

```go
// ✅ Хорошо — короткий и понятный
func location(city string) (lat, lng float64) {
    // ...
    return
}

// ❌ Плохо — слишком много именованных переменных
func process(data []byte) (result []byte, err error, count int, modified bool) {
    // ...
}
```

---

## Variadic функции (переменное число аргументов)

Функции могут принимать переменное количество аргументов одного типа:

```go
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2))          // 3
    fmt.Println(sum(1, 2, 3, 4, 5)) // 15
    fmt.Println(sum())              // 0
}
```

### Передача среза как variadic аргументов

Используйте `...` для распаковки среза:

```go
numbers := []int{1, 2, 3, 4, 5}
result := sum(numbers...) // распаковываем срез
fmt.Println(result)       // 15
```

### Комбинация обычных и variadic параметров

Variadic параметр **всегда последний**:

```go
func greet(greeting string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s, %s!\n", greeting, name)
    }
}

func main() {
    greet("Hello", "Alice", "Bob", "Charlie")
    // Hello, Alice!
    // Hello, Bob!
    // Hello, Charlie!
}
```

### Пример: fmt.Println

`fmt.Println` — классический пример variadic функции:

```go
package main

import (
    "fmt"
    "os"
)

func Println(a ...any) (n int, err error) {
    return Fprintln(os.Stdout, a...)
}

func main() {
    fmt.Println("Hello", 42, true, 3.14)
}
```

---

## Анонимные функции

Функции без имени, которые можно объявлять и вызывать на месте:

```go
package main

import "fmt"

func main() {
    func() {
        fmt.Println("Hello from anonymous function!")
    }()
    
    result := func(a, b int) int {
        return a + b
    }(5, 3)
    fmt.Println(result) // 8
}
```

### Присваивание анонимной функции переменной

```go
package main

import "fmt"

func main() {
    multiply := func(a, b int) int {
        return a * b
    }

    result := multiply(4, 5)
    fmt.Println(result) // 20
}
```

---

## Замыкания (Closures)

Замыкание — это анонимная функция, которая захватывает переменные из внешней области видимости:

```go
package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c1 := counter()
    fmt.Println(c1()) // 1
    fmt.Println(c1()) // 2
    fmt.Println(c1()) // 3
    
    c2 := counter()   // новый счётчик
    fmt.Println(c2()) // 1
}
```

### Практический пример: генератор последовательности

```go
package main

import "fmt"

func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

func main() {
    fib := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Print(fib(), " ")
    }
    // 1 1 2 3 5 8 13 21 34 55
}
```

### ⚠️ Замыкания и циклы (до Go 1.22)

До Go 1.22 все замыкания в цикле использовали **одну и ту же переменную**:

```go
package main

import "fmt"

func main() {
    // ❌ Проблема до Go 1.22
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Println(i) // все выведут 3
        }()
    }
    
    // ✅ Решение: создать локальную копию
    for i := 0; i < 3; i++ {
        j := i // локальная копия для каждой итерации
        go func() {
            fmt.Println(j)
        }()
    }
}
```

## Generics (Обобщения) — Go 1.18+

Появились в Go 1.18 — позволяют писать обобщённый код, который работает с любыми типами.

### Базовый синтаксис

```go
// Минимум из двух значений (только упорядоченные типы)
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

func main() {
    fmt.Println(Min[int](5, 3))       // 5, 3 → int
    fmt.Println(Min(float64(2.5), 1.8)) // 2.5, 1.8 → float64 (инференс)
}
```

### Type Constraints (ограничения типов)

Go использует интерфейсы для ограничения допустимых типов:

```go
package main

import (
    "constraints"
    "fmt"
)

// Типы, которые можно сравнивать (comparable)
func Equal[T comparable](a, b T) bool {
    return a == b
}

// Типы, которые можно упорядочить (Ordered)
func Max[T constraints.Ordered](nums ...T) T {
    if len(nums) == 0 {
        panic("empty slice")
    }
    max := nums[0]
    for _, n := range nums {
        if n > max {
            max = n
        }
    }
    return max
}

func main() {
    fmt.Println(Equal(5, 5))     // true
    fmt.Println(Max(5, 2, 8, 1)) // 8
}
```

### Общие паттерны использования

#### 1. Map для срезов

```go
package main

import "fmt"

func Map[T, R any](slice []T, f func(T) R) []R {
    result := make([]R, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4}
    squared := Map(nums, func(x int) int { return x * x })
    fmt.Println(squared) // [1 4 9 16]
}
```

#### 2. Filter для срезов

```go
package main

import "fmt"

func Filter[T any](slice []T, pred func(T) bool) []T {
    result := make([]T, 0, len(slice))
    for _, v := range slice {
        if pred(v) {
            result = append(result, v)
        }
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5, 6}
    evens := Filter(nums, func(x int) bool { return x%2 == 0 })
    fmt.Println(evens) // [2 4 6]
}
```

#### 3. Reduce для срезов

```go
package main

import "fmt"

func Reduce[T, R any](slice []T, initial R, f func(R, T) R) R {
    result := initial
    for _, v := range slice {
        result = f(result, v)
    }
    return result
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    sum := Reduce(nums, 0, func(acc, x int) int { return acc + x })
    fmt.Println(sum) // 15
}
```

### Когда использовать generics, а когда interface{}

### Когда использовать generics, а когда interface{}

| Сценарий | generics | interface{} |
|----------|----------|-------------|
| Операции с типами (сравнение, арифметика) | ✅ | ❌ |
| Хранение данных разных типов | ❌ | ✅ |
| Обработка срезов/карты | ✅ | ✅ |
| Функции высшего порядка | ✅ | ✅ |
| Максимальная производительность | ✅ | ⚠️ (type assertion) |

### Ограничения generics в Go

- Нет runtime reflection для типов
- Нельзя использовать `==` или `<` без constraints.Ordered/comparable
- Нет generics для методов (только для функций и типов)
- Можно объявлять переменные типа `T` через `var x T`, но нельзя использовать `T` в type assertion (`x.(T)` не работает)

### Пример: общий Map с constraints

```go
package main

import "fmt"

func MapKeys[T any, K comparable, V any](slice []T, keyFn func(T) K, valFn func(T) V) map[K]V {
    result := make(map[K]V)
    for _, v := range slice {
        result[keyFn(v)] = valFn(v)
    }
    return result
}

type Person struct {
    Name string
    Age  int
}

func main() {
    people := []Person{
        {Name: "Alice", Age: 25},
        {Name: "Bob", Age: 30},
    }
    
    nameToAge := MapKeys(people,
        func(p Person) string { return p.Name },
        func(p Person) int { return p.Age },
    )
    fmt.Println(nameToAge) // map[Alice:25 Bob:30]
}
```

---

## Функции как значения

Функции можно присваивать переменным, передавать как аргументы и возвращать:

### Тип функции

```go
package main

import "fmt"

// Объявление типа функции
type Operation func(int, int) int

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func calculate(op Operation, a, b int) int {
    return op(a, b)
}

func main() {
    fmt.Println(calculate(add, 10, 5))      // 15
    fmt.Println(calculate(subtract, 10, 5)) // 5
}
```

### Функция как аргумент

```go
package main

import (
    "fmt"
    "strings"
)

func apply(f func(string) string, s string) string {
    return f(s)
}

func toUpper(s string) string {
    return strings.ToUpper(s)
}

func main() {
    result := apply(toUpper, "hello")
    fmt.Println(result) // HELLO
}
```

### Функция как возвращаемое значение

```go
package main

import "fmt"

func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := multiplier(2)
    triple := multiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```

---

## Recursion (Рекурсия)

Функция может вызывать саму себя:

```go
func factorial(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println(factorial(5)) // 120
}
```

### Пример: числа Фибоначчи

```go
package main

import "fmt"

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    for i := 0; i < 10; i++ {
        fmt.Print(fibonacci(i), " ")
    }
    // 0 1 1 2 3 5 8 13 21 34
}
```

### ⚠️ Оптимизация: мемоизация

Наивная рекурсия Фибоначчи неэффективна (O(2^n)). Используйте мемоизацию:

```go
package main

import "fmt"

func fibonacciMemo(n int, memo map[int]int) int {
    if val, ok := memo[n]; ok {
        return val
    }
    if n <= 1 {
        return n
    }
    result := fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
    memo[n] = result
    return result
}

func main() {
    memo := make(map[int]int)
    fmt.Println(fibonacciMemo(50, memo)) // быстро!
}
```

---

## Defer (Отложенные вызовы)

`defer` откладывает выполнение функции до возврата из окружающей функции. Вызовы выполняются в порядке **LIFO** (последний добавленный — первый выполнится):

```go
package main

import "fmt"

func main() {
    fmt.Println("Start")
    
    defer fmt.Println("Deferred 1")
    defer fmt.Println("Deferred 2")
    defer fmt.Println("Deferred 3")
    
    fmt.Println("End")
}

// Вывод:
// Start
// End
// Deferred 3
// Deferred 2
// Deferred 1
```

### Основные применения defer

#### 1. Закрытие ресурсов

```go
package main

import (
    "fmt"
    "io"
    "os"
)

func readFile(path string) ([]byte, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close() // закроется при выходе из функции
    
    data, err := io.ReadAll(file)
    return data, err
}

func main() {
    data, err := readFile("example.txt")
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Данные прочитаны:", len(data), "байт")
}
```

#### 2. Разблокировка мьютексов

```go
package main

import (
    "sync"
)

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock() // разблокируется при выходе
    c.value++
}
```

---

## Panic и Recover

### Panic

`panic` прерывает нормальное выполнение функции и начинает раскрутку стека:

```go
package main

import (
    "fmt"
    "strconv"
)

func mustParse(s string) int {
    n, err := strconv.Atoi(s)
    if err != nil {
        panic(fmt.Sprintf("не удалось преобразовать %q: %v", s, err))
    }
    return n
}

func main() {
    fmt.Println(mustParse("42"))     // 42
    fmt.Println(mustParse("abc"))    // panic!
}
```

### ⚠️ Когда использовать panic

✅ **Подходит:**
- Невозможные состояния (инварианты нарушены)
- Ошибки программиста
- Ошибки инициализации при старте

❌ **НЕ подходит:**
- Ошибки ввода-вывода
- Ошибки валидации
- Ожидаемые ошибки

### Recover

`recover` перехватывает panic и позволяет продолжить выполнение:

```go
package main

import (
    "fmt"
)

func safeCall() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered: %v", r)
        }
    }()
    
    panic("что-то пошло не так")
    return
}

func main() {
    if err := safeCall(); err != nil {
        fmt.Println("Ошибка:", err)
    }
    fmt.Println("Программа продолжает работу")
}
```

### 📊 Когда использовать panic/recover

| Ситуация | Действие |
|----------|----------|
| Ошибка ввода-вывода | `return err` |
| Невозможное состояние | `panic` |
| Горутина в продакшене | `defer recover()` |
| Ошибка валидации | `return ErrValidation` |



```go
package main

import (
    "fmt"
)

// ❌ Не сработает
func badRecover() {
    if r := recover(); r != nil { // вне defer — не работает
        fmt.Println(r)
    }
    panic("test")
}

// ✅ Правильно
func goodRecover() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
    }()
    panic("test")
}

func main() {
    goodRecover() // паника будет перехвачена
    fmt.Println("Программа продолжает работу")
}
```

---

## Методы

Методы — это функции с особым параметром **получателем** (receiver):

### Embedded типы (композиция)

Go не имеет наследования, но поддерживает композицию через встраивание структур:

```go
package main

import "fmt"

type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return fmt.Sprintf("%s издаёт звук", a.Name)
}

type Dog struct {
    Animal  // embedded — методы Animal "наследуются"
    Breed string
}

func main() {
    d := Dog{Animal: Animal{Name: "Барон"}, Breed: "Двортерьер"}
    fmt.Println(d.Speak()) // Барон издаёт звук
    fmt.Println(d.Breed)   // Двортерьер
}
```

**Важно:**
- Встраивание — это композиция, а не наследование
- Embedded поля доступны напрямую: `dog.Name` вместо `dog.Animal.Name`
- Методы embedded типов "всплывают" на верхний уровень

```go
package main

import (
    "fmt"
    "math"
)

type Person struct {
    Name string
    Age  int
}

// Метод с получателем-значением
func (p Person) Greet() string {
    return fmt.Sprintf("Привет, я %s", p.Name)
}

// Метод с получателем-указателем
func (p *Person) Birthday() {
    p.Age++ // может изменять объект
}

func main() {
    p := Person{Name: "Alice", Age: 25}
    fmt.Println(p.Greet()) // Привет, я Alice
    p.Birthday()
    fmt.Println(p.Age)     // 26
}
```

### Получатель-значение vs Получатель-указатель

| Критерий | Значение `(p Person)` | Указатель `(p *Person)` |
|----------|----------------------|------------------------|
| Изменение объекта | ❌ Нет (работает с копией) | ✅ Да |
| Производительность | Копирование (дорого для больших структур) | Передача указателя (дешево) |
| Nil-safety | Не может быть nil | Может быть nil |
| Когда использовать | Неизменяемые типы, маленькие структуры | Большие структуры, когда нужно изменение |

```go
package main

import (
    "database/sql"
    "fmt"
    "math"
)

// ✅ Маленькая структура — можно значение
type Point struct {
    X, Y int
}

func (p Point) Distance() float64 {
    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// ✅ Большая структура или изменение — указатель
type Database struct {
    Connection *sql.DB
    Config     Config
    Cache      map[string]any
    // ... много полей
}

func (db *Database) Close() error {
    return db.Connection.Close()
}

func main() {
    p := Point{3, 4}
    fmt.Println(p.Distance()) // 5
}
```

---

## Init функции

Специальные функции `init()` вызываются автоматически при инициализации пакета:

```go
package main

import (
    "fmt"
)

var globalVar = initializeGlobal()

func initializeGlobal() int {
    fmt.Println("Инициализация глобальной переменной")
    return 42
}

func init() {
    fmt.Println("init() вызвана")
}

func main() {
    fmt.Println("main() вызвана")
}

// Порядок выполнения:
// 1. Инициализация глобальной переменной
// 2. init()
// 3. main()
```

```go
package config

import (
    "log"
)

var Config map[string]string

func init() {
    Config = make(map[string]string)
    // загрузка конфигурации
    log.Println("Конфигурация загружена")
}

func init() {
    // валидация конфигурации
    if Config == nil {
        log.Fatal("Конфигурация не инициализирована")
    }
}
```

### Особенности init

- Можно объявлять **несколько** `init()` в одном файле
- Если в пакете несколько файлов с `init()`, они выполняются в **алфавитном порядке имён файлов**
- Внутри одного файла `init()` выполняются в порядке объявления
- Выполняются **до** `main()`
- Выполняются **один раз** при загрузке пакета
- Не принимают параметров и не возвращают значений

---

## Таблица быстрого справочника

| Концепция | Синтаксис | Пример |
|-----------|-----------|--------|
| Простая функция | `func name(params) return` | `func add(a, b int) int` |
| Множественный возврат | `func name() (type1, type2)` | `func divide() (float64, error)` |
| Именованный возврат | `func name() (x, y int)` | `return` (голый return) |
| Variadic | `func name(args ...type)` | `func sum(nums ...int)` |
| Анонимная функция | `func(params) return { }` | `func() { fmt.Println() }()` |
| Замыкание | Захват внешних переменных | `count++` внутри возвращаемой функции |
| Defer | `defer func()` | `defer file.Close()` |
| Panic | `panic(value)` | `panic("error")` |
| Recover | `recover()` в defer | `if r := recover(); r != nil` |
| Метод | `func (r Receiver) name()` | `func (p Person) Greet()` |

---

## Лучшие практики

### 1. Используйте множественные возвраты для ошибок

```go
package main

import (
    "fmt"
    "os"
)

// ✅ Хорошо
func readFile(path string) ([]byte, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("чтение файла: %w", err)
    }
    return data, nil
}

// ❌ Плохо
func readFile(path string) []byte {
    data, err := os.ReadFile(path)
    if err != nil {
        panic(err) // не используйте panic для ожидаемых ошибок
    }
    return data
}
```

### 2. Всегда закрывайте ресурсы через defer

```go
package main

import (
    "fmt"
    "os"
)

// ✅ Хорошо
func processData() error {
    file, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    defer file.Close()
    
    // работа с файлом
    return nil
}

// ❌ Плохо — можно забыть закрыть
func processData() error {
    file, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    
    // работа с файлом
    file.Close() // может не выполниться при ошибке выше
    return nil
}
```

---

### 3. Используйте именованные возвраты с defer

```go
package main

import "fmt"

func double() (result int) {
    result = 5
    defer func() {
        result *= 2
    }()
    return
}

func main() {
    fmt.Println(double()) // 10
}
```

### 4. Минимизируйте количество параметров

```go
package main

import (
    "time"
)

// ✅ Хорошо — используйте структуру для многих параметров
type Config struct {
    Host     string
    Port     int
    Timeout  time.Duration
    Retries  int
}

func connect(cfg Config) error {
    // ...
    return nil
}

// ❌ Плохо — слишком много параметров
func connect(host string, port int, timeout time.Duration, retries int) error {
    // ...
    return nil
}
```

### 5. Избегайте глубокой рекурсии

```go
// ✅ Хорошо — итеративное решение
func factorial(n int) int {
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result
}

// ⚠️ Осторожно — рекурсия может вызвать stack overflow
func factorialRecursive(n int) int {
    if n <= 1 {
        return 1
    }
    return n * factorialRecursive(n-1)
}
```

### 6. Используйте recover для защиты от panic в горутинах

```go
package main

import (
    "log"
)

func main() {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Goroutine panic: %v", r)
            }
        }()
        // код, который может паниковать
    }()
}
```

---

## Частые ошибки

### 1. Забытый defer для закрытия ресурсов

```go
package main

import (
    "fmt"
    "os"
)

// ❌ Плохо — ресурс не закрыт
func bad() error {
    file, _ := os.Open("file.txt")
    // забыли defer file.Close()
    return nil
}

// ✅ Хорошо
func good() error {
    file, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer file.Close()
    return nil
}

func main() {
    data, _ := good()
    fmt.Println("Прочитано:", len(data))
}
```

### 2. Использование panic для обычных ошибок

```go
package main

import "errors"

// ❌ Плохо — не используйте panic для ожидаемых ошибок
func divide(a, b float64) float64 {
    if b == 0 {
        panic("деление на ноль")
    }
    return a / b
}

// ✅ Хорошо — возвращайте ошибку
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}
```

### 3. Неправильное использование именованных возвратов

```go
package main

import "fmt"

// ❌ Плохо — путаница с переменными
func confusing() (result int) {
    result := 10 // создаёт новую переменную, не модифицирует возвращаемую
    return       // вернёт 0, а не 10
}

// ✅ Хорошо
func clear() (result int) {
    result = 10  // модифицирует возвращаемую переменную
    return
}

func main() {
    fmt.Println(confusing()) // 0
    fmt.Println(clear())     // 10
}
```

### 4. Замыкания в циклах (до Go 1.22)

```go
package main

import "fmt"

func main() {
    // ❌ Проблема до Go 1.22
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Println(i) // все выведут 3
        }()
    }
    
    // ✅ Решение
    for i := 0; i < 3; i++ {
        j := i // локальная копия
        go func() {
            fmt.Println(j)
        }()
    }
}
```

---

---

## Заключение

Функции в Go — мощный инструмент с рядом уникальных особенностей:

1. **Множественные возвраты** — стандартный паттерн для ошибок
2. **First-class functions** — функции как значения
3. **Замыкания** — захват переменных из внешней области
4. **Defer** — гарантированная очистка ресурсов
5. **Panic/Recover** — обработка фатальных ошибок
6. **Методы** — функции, привязанные к типам
7. **Generics** — обобщённый код для любых типов (Go 1.18+)

Правильное использование функций делает код читаемым, поддерживаемым и идиоматичным.

---

## Ссылки

- [Official Go Specification - Function types](https://go.dev/ref/spec#Function_types)
- [Official Go Specification - Function declarations](https://go.dev/ref/spec#Function_declarations)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Effective Go - Defer, Panic, and Recover](https://go.dev/doc/effective_go#defer)
- [Go Blog - Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Tour of Go - Functions](https://go.dev/tour/moretypes/25)
- [Go Blog - Generics](https://go.dev/blog/intro-generics)
- [Go Generics Documentation](https://go.dev/doc/tutorial/generics)