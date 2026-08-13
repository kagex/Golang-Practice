# Map в Go

## 1. Что такое `map`

`map` в Go — это встроенный ссылочный тип, который представляет собой ассоциативный массив, то есть отображение ключей на значения. Внутри это хеш-таблица.

Общая идея:

```text
map[KeyType]ValueType
```

Пример:

```go
map[string]int
map[int]string
map[string][]string
map[struct{}]bool
```

Каждому ключу соответствует значение. Ключи в `map` уникальны.

---

## 2. Объявление и создание map

### Nil map

```go
var m map[string]int
```

Такой `map` имеет нулевое значение `nil`.

```go
fmt.Println(m == nil) // true
fmt.Println(len(m))   // 0
```

Из nil-карты можно читать, она ведёт себя как пустая:

```go
v := m["missing"]
fmt.Println(v) // 0
```

Но писать в nil-карту нельзя:

```go
m["a"] = 1 // panic: assignment to entry in nil map
```

---

### Пустая карта через `make`

```go
m := make(map[string]int)
```

Это уже nil-карта, в неё можно писать:

```go
m["a"] = 1
```

---

### Пустая карта через литерал

```go
m := map[string]int{}
```

Эквивалентно `make(map[string]int)`.

---

### Карта с начальными значениями

```go
ages := map[string]int{
 "Alice": 25,
 "Bob":   30,
}
```

---

## 3. Nil map и пустая map

| Операция | nil map | пустая map |
| --- | ---: | ---: |
| чтение | безопасно | безопасно |
| `len` | `0` | `0` |
| `range` | 0 итераций | 0 итераций |
| `delete` | безопасно, no-op | безопасно, no-op |
| запись | panic | работает |

Пример:

```go
var nilMap map[string]int
emptyMap := make(map[string]int)

fmt.Println(len(nilMap))  // 0
fmt.Println(len(emptyMap)) // 0

// nilMap["x"] = 1 // panic
emptyMap["x"] = 1  // ok
```

Если функция может получать nil-карту и должна в неё писать, обычно карту нужно инициализировать заранее или создавать внутри.

---

## 4. Чтение значений

Обычное чтение возвращает значение или нулевое значение типа значения:

```go
m := map[string]int{
 "a": 1,
}

fmt.Println(m["a"])       // 1
fmt.Println(m["missing"]) // 0
```

Если значением является, например, `bool`, то отсутствие ключа даст `false`:

```go
m := map[string]bool{
 "ok": false,
}

fmt.Println(m["ok"])      // false
fmt.Println(m["missing"]) // false
```

Поэтому для проверки наличия ключа используется форма с двумя возвращаемыми значениями.

---

## 5. Проверка наличия ключа: "идиома `ok`"

```go
v, ok := m["key"]
```

- `v` — значение;
- `ok` — `true`, если ключ существует.

Пример:

```go
ages := map[string]int{
 "Alice": 25,
}

age, ok := ages["Alice"]
fmt.Println(age, ok) // 25 true

age, ok = ages["Bob"]
fmt.Println(age, ok) // 0 false
```

Важно: если ключ существует, но значение равно нулевому значению типа, `ok` всё равно будет `true`.

```go
m := map[string]int{
 "zero": 0,
}

v, ok := m["zero"]
fmt.Println(v, ok) // 0 true
```

---

## 6. Запись и обновление значений

```go
m := make(map[string]int)

m["a"] = 1
m["a"] = 2

fmt.Println(m["a"]) // 2
```

Если ключ уже существует, значение перезаписывается.

---

## 7. Удаление элементов: `delete`

Встроенная функция `delete` удаляет элемент по ключу:

```go
m := map[string]int{
 "a": 1,
 "b": 2,
}

delete(m, "a")

fmt.Println(m) // map[b:2]
```

Если ключа нет, `delete` ничего не делает.

Если карта nil, `delete` тоже безопасна:

```go
var m map[string]int
delete(m, "any") // no-op
```

`delete` ничего не возвращает. Если нужно узнать, существовал ли ключ до удаления, используйте:

```go
v, existed := m["a"]
if existed {
 delete(m, "a")
}
```

---

## 8. Длина карты

```go
m := map[string]int{
 "a": 1,
 "b": 2,
}

fmt.Println(len(m)) // 2
```

`cap` для карт не существует.

---

## 9. Итерация по map

### Полная итерация

```go
m := map[string]int{
 "a": 1,
 "b": 2,
 "c": 3,
}

for key, value := range m {
 fmt.Println(key, value)
}
```

### Только ключи

```go
for key := range m {
 fmt.Println(key)
}
```

### Только значения

```go
for _, value := range m {
 fmt.Println(value)
}
```

---

## 10. Порядок итерации не гарантируется

Порядок обхода `map` в Go не детерминирован. Он может отличаться от запуска к запуску и от итерации к итерации.

Нельзя полагаться на такой код:

```go
for k, v := range m {
 fmt.Println(k, v)
}
```

если вам важен порядок.

---

## 11. Как получить детерминированный порядок

Нужно собрать ключи, отсортировать их, затем обойти карту по ключам.

### Классический способ

```go
import "sort"

keys := make([]string, 0, len(m))
for k := range m {
 keys = append(keys, k)
}

sort.Strings(keys)

for _, k := range keys {
 fmt.Println(k, m[k])
}
```

### С использованием пакетов `maps` и `slices` в Go 1.23+

```go
import (
 "maps"
 "slices"
)

for _, k := range slices.Sorted(maps.Keys(m)) {
 fmt.Println(k, m[k])
}
```

---

## 12. Изменение карты во время итерации

Во время `range` можно удалять элементы из той же карты:

```go
for k, v := range m {
 if v < 0 {
  delete(m, k)
 }
}
```

Это безопасно в рамках одной горутины.

Если во время итерации добавлять новые элементы:

```go
for k, v := range m {
 m[k+"!"] = v
}
```

новые элементы могут быть обойдены, а могут и не быть. Поведение не гарантируется.

---

## 13. Какие типы можно использовать как ключи

Ключи в `map` должны быть сравнимыми, то есть поддерживать операции `==` и `!=`.

Можно использовать:

```go
map[string]int
map[int]string
map[bool]string
map[uint64]float64
map[complex128]string
map[uintptr]int
map[chan int]string
map[*MyStruct]string
map[[3]int]string
map[struct{ X int }]string
```

Нельзя использовать:

```go
map[[]int]string
map[map[string]int]string
map[func()]string
```

потому что срезы, карты и функции не сравнимы.

---

## 14. Ключи-структуры

Структуру можно использовать как ключ, если все её поля сравнимы.

```go
type Point struct {
 X int
 Y int
}

m := map[Point]string{}

m[Point{X: 1, Y: 2}] = "point 1-2"
```

Если в структуре есть срез, карта или функция, такая структура не будет сравнимой:

```go
type BadKey struct {
 ID   int
 Tags []string
}

// map[BadKey]string{} // ошибка: ключ несравним
```

---

## 15. Ключи-интерфейсы

Интерфейсный тип может быть ключом:

```go
m := map[interface{}]string{}
```

Но если динамическим значением окажется несравнимый тип, возможна паника при сравнении.

Например:

```go
var k interface{} = []int{1, 2, 3}

m := map[interface{}]string{}
m[k] = "value" // может привести к панике
```

Поэтому использовать `interface{}` как ключ нужно осторожно.

---

## 16. Ключи с плавающей точкой и NaN

Формально float-типы сравнимы, но `NaN != NaN`.

Из-за этого ключи вида `math.NaN()` ведут себя неочевидно. На практике лучше избегать использования float-ключей, особенно если может встречаться `NaN`.

---

## 17. Нулевые значения

Если ключ не найден, возвращается нулевое значение типа значения.

Примеры нулевых значений:

| Тип значения | Нулевое значение |
| --- | --- |
| `int` | `0` |
| `string` | `""` |
| `bool` | `false` |
| `pointer` | `nil` |
| `slice` | `nil` |
| `map` | `nil` |
| `interface` | `nil` |
| `struct` | нулевой struct |

Пример:

```go
m := map[string][]int{}

v := m["missing"]
fmt.Println(v == nil) // true
```

---

## 18. Карта значений-указателей

В `map` можно хранить указатели:

```go
type User struct {
 Name string
}

m := map[string]*User{
 "alice": {Name: "Alice"},
}

m["alice"].Name = "Changed"
```

Но нужно помнить, что это поверхностное хранение: карта хранит указатели, а не копии объектов.

---

## 19. Карта из срезов

```go
groups := make(map[string][]string)

groups["fruits"] = append(groups["fruits"], "apple")
groups["fruits"] = append(groups["fruits"], "banana")

fmt.Println(groups["fruits"]) // [apple banana]
```

Это частый паттерн группировки.

---

## 20. Карта из карт

```go
config := map[string]map[string]string{
 "db": {
  "host": "localhost",
  "port": "5432",
 },
}

fmt.Println(config["db"]["host"]) // localhost
```

---

## 21. `map` как множество, set

В Go нет встроенного типа `set`, но его можно реализовать через `map`.

Часто используют `struct{}` как значение, потому что он занимает ноль байт:

```go
set := make(map[string]struct{})

set["done"] = struct{}{}

if _, ok := set["done"]; ok {
 fmt.Println("done exists")
}
```

Удаление:

```go
delete(set, "done")
```

---

## 22. `map` как счётчик

```go
words := []string{"go", "rust", "go", "python", "go"}

counts := make(map[string]int)

for _, word := range words {
 counts[word]++
}

fmt.Println(counts) // map[go:3 python:1 rust:1]
```

Это работает, потому что при обращении к несуществующему ключу возвращается `0`.

---

## 23. `map` как кэш

```go
cache := make(map[string]int)

func get(key string) int {
 if v, ok := cache[key]; ok {
  return v
 }

 v := compute(key)
 cache[key] = v
 return v
}
```

В реальном коде нужно учитывать конкурентный доступ, инвалидацию и ограничения памяти.

---

## 24. Передача `map` в функцию

Значение `map` — это ссылка на внутреннюю хеш-таблицу. Если передать карту в функцию и изменить содержимое, изменения будут видны снаружи.

```go
func increment(m map[string]int) {
 m["a"]++
}

func main() {
 m := map[string]int{"a": 1}
 increment(m)
 fmt.Println(m["a"]) // 2
}
```

Но если внутри функции присвоить переменной новую карту, внешняя переменная не изменится:

```go
func replace(m map[string]int) {
 m = make(map[string]int)
 m["new"] = 1
}

func main() {
 m := map[string]int{"old": 1}
 replace(m)
 fmt.Println(m) // map[old:1]
}
```

---

## 25. Копирование `map`

Присваивание карты копирует ссылку на ту же внутреннюю структуру, а не элементы:

```go
m1 := map[string]int{"a": 1}
m2 := m1

m2["a"] = 2

fmt.Println(m1["a"]) // 2
fmt.Println(m2["a"]) // 2
```

Если нужна копия карты, её нужно создавать вручную или использовать пакет `maps`.

### Ручное копирование

```go
func cloneMap(m map[string]int) map[string]int {
 result := make(map[string]int, len(m))
 for k, v := range m {
  result[k] = v
 }
 return result
}
```

### Копирование через `maps.Clone`

```go
import "maps"

m1 := map[string]int{"a": 1}
m2 := maps.Clone(m1)

m2["a"] = 2

fmt.Println(m1["a"]) // 1
fmt.Println(m2["a"]) // 2
```

Важно: `maps.Clone` делает поверхностную копию. Если значения являются указателями, срезами, картами или каналами, копируются сами значения этих ссылок, а не глубокие объекты.

---

## 26. Сравнение карт

Карты нельзя сравнивать между собой оператором `==`, кроме сравнения с `nil`.

Можно:

```go
var m map[string]int

fmt.Println(m == nil) // true
```

Нельзя:

```go
m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}

// fmt.Println(m1 == m2) // ошибка компиляции
```

Для сравнения содержимого можно использовать пакет `maps`:

```go
import "maps"

m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}

fmt.Println(maps.Equal(m1, m2)) // true
```

Для несравнимых значений используется `maps.EqualFunc`:

```go
equal := maps.EqualFunc(m1, m2, func(v1, v2 []string) bool {
 return slices.Equal(v1, v2)
})
```

---

## 27. Очистка карты

### Встроенный `clear`

Начиная с Go 1.21 доступен встроенный `clear`:

```go
m := map[string]int{
 "a": 1,
 "b": 2,
}

clear(m)

fmt.Println(len(m)) // 0
```

### `maps.Clear`

Также есть функция в пакете `maps`:

```go
import "maps"

maps.Clear(m)
```

Если нужно гарантированно освободить память и начать с новой карты, обычно создают новую:

```go
m = make(map[string]int)
```

---

## 28. Удаление элементов по условию

В пакете `maps` есть `DeleteFunc`.

```go
import "maps"

m := map[string]int{
 "a": 1,
 "b": -2,
 "c": 3,
}

maps.DeleteFunc(m, func(k string, v int) bool {
 return v < 0
})

fmt.Println(m) // map[a:1 c:3]
```

То же самое вручную:

```go
for k, v := range m {
 if v < 0 {
  delete(m, k)
 }
}
```

---

## 29. Копирование из одной карты в другую

Функция `maps.Copy` копирует элементы из одной карты в другую.

```go
import "maps"

dst := map[string]int{
 "b": 10,
}

src := map[string]int{
 "a": 1,
 "b": 2,
}

maps.Copy(dst, src)

fmt.Println(dst) // map[a:1 b:2]
```

`maps.Copy` не очищает `dst`, а только добавляет или перезаписывает ключи.

---

## 30. Пакет `maps`

Стандартный пакет `maps` появился в Go 1.21 и содержит функции для работы с картами.

Импорт:

```go
import "maps"
```

Основные функции:

| Функция | Назначение |
| --- | --- |
| `maps.Clone` | поверхностная копия карты |
| `maps.Copy` | копирование элементов из одной карты в другую |
| `maps.Equal` | сравнение карт |
| `maps.EqualFunc` | сравнение карт с пользовательской функцией |
| `maps.DeleteFunc` | удаление элементов по условию |
| `maps.Clear` | очистка карты |
| `maps.Keys` | итерация по ключам в Go 1.23+ |
| `maps.Values` | итерация по значениям в Go 1.23+ |
| `maps.All` | итерация по парам ключ-значение в Go 1.23+ |

Важно: итераторные функции зависят от версии Go. В Go 1.23+ они работают с пакетом `iter` и поддерживают `range`.

---

## 31. Итераторы для карт в Go 1.23+

Начиная с Go 1.23, в пакете `maps` есть функции, возвращающие итераторы.

### `maps.All`

```go
import "maps"

m := map[string]int{
 "a": 1,
 "b": 2,
}

for k, v := range maps.All(m) {
 fmt.Println(k, v)
}
```

### `maps.Keys`

```go
for k := range maps.Keys(m) {
 fmt.Println(k)
}
```

### `maps.Values`

```go
for v := range maps.Values(m) {
 fmt.Println(v)
}
```

### Отсортированные ключи

```go
import (
 "maps"
 "slices"
)

for _, k := range slices.Sorted(maps.Keys(m)) {
 fmt.Println(k, m[k])
}
```

---

## 32. Производительность

В среднем операции чтения, записи и удаления в `map` работают за `O(1)`.

Но:

- в худшем случае операции могут деградировать;
- порядок итерации не гарантирован;
- карта не является структурой с фиксированным размером;
- при большом количестве элементов важна правильная оценка начального размера.

Если заранее известно количество элементов, можно дать подсказку:

```go
m := make(map[string]int, 1000)
```

Это не гарантирует ёмкость, но помогает уменьшить количество расширений.

---

## 33. `map` и конкурентность

Обычная `map` в Go не является потокобезопасной.

Нельзя одновременно читать и писать из разных горутин без синхронизации:

```go
m := make(map[string]int)

go func() {
 m["a"] = 1
}()

go func() {
 _ = m["a"]
}()
```

Это может привести к гонке данных и падению программы.

---

## 34. Защита с помощью мьютекса

```go
type SafeCounter struct {
 mu     sync.Mutex
 counts map[string]int
}

func NewSafeCounter() *SafeCounter {
 return &SafeCounter{
  counts: make(map[string]int),
 }
}

func (c *SafeCounter) Inc(name string) {
 c.mu.Lock()
 defer c.mu.Unlock()

 c.counts[name]++
}

func (c *SafeCounter) Get(name string) int {
 c.mu.Lock()
 defer c.mu.Unlock()

 return c.counts[name]
}
```

Для чтения можно использовать `sync.RWMutex`, если чтение сильно преобладает.

---

## 35. `sync.Map`

В пакете `sync` есть `sync.Map`.

```go
import "sync"

var m sync.Map

m.Store("a", 1)

v, ok := m.Load("a")
if ok {
 fmt.Println(v) // 1
}

m.Delete("a")
```

`sync.Map` подходит не для всех сценариев. Обычно её рассматривают, когда:

- есть преимущественно чтение уже записанных ключей;
- разные горутины работают с непересекающимися наборами ключей;
- нужен lock-free доступ в специфических паттернах.

Для большинства обычных случаев `map + sync.Mutex` или `map + sync.RWMutex` проще и предсказуемее.

---

## 36. Частые ошибки

### 1. Запись в nil map

```go
var m map[string]int
// m["a"] = 1 // panic
```

Нужно:

```go
m = make(map[string]int)
```

---

### 2. Ожидание порядка итерации

```go
for k, v := range m {
 // порядок не гарантируется
}
```

Если нужен порядок — сортируйте ключи.

---

### 3. Сравнение карт через `==`

```go
m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}

// m1 == m2 // нельзя
```

Используйте `maps.Equal` или ручное сравнение.

---

### 4. Конкурентный доступ без синхронизации

Обычная `map` не потокобезопасна.

---

### 5. Ожидание глубокой копии

```go
m2 := maps.Clone(m1)
```

`Clone` копирует карту, но не делает deep copy значений.

---

### 6. Использование несравнимых ключей

Нельзя:

```go
map[[]int]string
map[map[string]int]string
map[func()]string
```

---

### 7. Попытка взять адрес элемента карты

```go
m := map[string]int{"a": 1}

// p := &m["a"] // нельзя
```

Элементы карты не являются адресуемыми.

Если нужен указатель, храните указатель как значение:

```go
m := map[string]*int{}
```

---

### 8. Удаление большого числа элементов и ожидание автоматического уменьшения памяти

Удаление элементов не обязательно возвращает память операционной системе. Если карта стала ненужной или нужно освободить память, лучше создать новую:

```go
m = make(map[string]int)
```

---

## 37. Практический пример

Подсчёт слов и вывод в отсортированном порядке:

```go
package main

import (
 "fmt"
 "sort"
)

func main() {
 text := "go map go slice map go"

 counts := make(map[string]int)

 words := []string{"go", "map", "go", "slice", "map", "go"}

 for _, word := range words {
  counts[word]++
 }

 keys := make([]string, 0, len(counts))
 for k := range counts {
  keys = append(keys, k)
 }

 sort.Strings(keys)

 for _, k := range keys {
  fmt.Printf("%s: %d\n", k, counts[k])
 }

 _ = text
}
```

Вывод:

```text
go: 3
map: 2
slice: 1
```

---

## 38. Когда использовать `map`

Используйте `map`, когда нужно:

- быстро находить значение по ключу;
- хранить уникальные ключи;
- группировать данные;
- считать частоты;
- реализовывать set;
- кэшировать результаты;
- строить индексы.

Не используйте `map`, когда:

- важен порядок сам по себе;
- нужен последовательный доступ как к массиву;
- ключи несравнимы;
- нужна конкурентность без дополнительной синхронизации;
- нужна глубокая копия автоматически.

---

## Источники

1. **The Go Programming Language Specification — Map types**  
   <https://go.dev/ref/spec#Map_types>

2. **Effective Go — Maps**  
   <https://go.dev/doc/effective_go#maps>

3. **Go Blog — Go maps in action**  
   <https://go.dev/blog/maps>

4. **Package `maps` documentation**  
   <https://pkg.go.dev/maps>

5. **Go 1.21 Release Notes**  
   <https://go.dev/doc/go1.21>  
   Здесь описаны встроенный `clear` и появление пакета `maps`.

6. **Go 1.23 Release Notes**  
   <https://go.dev/doc/go1.23>  
   Здесь описаны итераторы и связанные с ними функции.

7. **Package `sync` — `Map`**  
   <https://pkg.go.dev/sync#Map>
