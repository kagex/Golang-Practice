# Пакет `slices` в Go

## 1. Что такое пакет `slices`

Пакет `slices` — это стандартный пакет Go, который содержит generic-функции для типичных операций над срезами: поиск, сортировка, сравнение, вставка, удаление, замена, копирование, компактификация и другие.

Пакет появился в стандартной библиотеке начиная с **Go 1.21**.

Импорт:

```go
import "slices"
```

Устанавливать через `go get` его не нужно: пакет входит в стандартную библиотеку Go.

---

## 2. Зачем нужен пакет `slices`

До появления пакета многие операции приходилось писать вручную или использовать пакет `sort`.

Например, проверка наличия элемента:

```go
found := false
for _, v := range nums {
    if v == target {
        found = true
        break
    }
}
```

С пакетом `slices`:

```go
found := slices.Contains(nums, target)
```

Сортировка:

```go
slices.Sort(nums)
```

Поиск в отсортированном срезе:

```go
i, ok := slices.BinarySearch(nums, target)
```

Основные преимущества:

1. **Читаемость** — операция видна из имени функции.
2. **Типобезопасность** — используются generics.
3. **Стандартность** — не нужно писать велосипеды.
4. **Производительность** — функции оптимизированы и не используют рефлексию там, где она не нужна.
5. **Удобные варианты с функциями обратного вызова** — `SortFunc`, `ContainsFunc`, `IndexFunc`, `EqualFunc` и другие.

---

## 3. Важные версии Go

Базовый пакет `slices` доступен с **Go 1.21**.

Некоторые функции были добавлены позже:

| Версия Go | Что добавлено |
| --- | --- |
| Go 1.21 | Базовый пакет `slices` |
| Go 1.22 | В том числе `Concat`, `DeleteFunc` |
| Go 1.23 | Итераторные функции: `All`, `Values`, `Backward`, `Collect`, `AppendSeq`, `Sorted`, `SortedFunc`, `SortedStableFunc` |

Если вы используете Go 1.21, части примеров с `DeleteFunc`, `Concat` и итераторами могут быть недоступны.

---

## 4. Основные понятия

Срез в Go состоит из трёх компонентов:

```go
slice := []int{1, 2, 3}
```

У него есть:

- длина: `len(slice)`;
- ёмкость: `cap(slice)`;
- указатель на базовый массив.

Пакет `slices` работает именно со срезами, а не с массивами. Многие функции изменяют срез in-place, многие возвращают новый срез или изменённый срез.

Важное правило: если функция возвращает срез, результат обычно нужно присваивать обратно:

```go
s = slices.Delete(s, 1, 3)
s = slices.Insert(s, 0, 42)
s = slices.Compact(s)
```

А вот функции вроде `Sort` и `Reverse` ничего не возвращают, потому что изменяют срез на месте:

```go
slices.Sort(nums)
slices.Reverse(nums)
```

---

## 5. Поиск и проверка наличия элемента

### `Contains`

Функция `slices.Contains` проверяет, содержится ли значение в срезе.

```go
nums := []int{1, 2, 3, 4}

fmt.Println(slices.Contains(nums, 3)) // true
fmt.Println(slices.Contains(nums, 9)) // false
```

Элементы должны быть сравнимы оператором `==`, поэтому тип элементов должен быть `comparable`.

---

### `ContainsFunc`

Если нужно искать по условию, используется `slices.ContainsFunc`.

```go
nums := []int{1, 3, 5, 8, 9}

hasEven := slices.ContainsFunc(nums, func(n int) bool {
    return n%2 == 0
})

fmt.Println(hasEven) // true
```

Функция возвращает `true`, если хотя бы один элемент удовлетворяет условию.

---

### `Index`

Функция `slices.Index` возвращает индекс первого вхождения значения или `-1`, если значение не найдено.

```go
nums := []int{10, 20, 30, 20}

fmt.Println(slices.Index(nums, 20)) // 1
fmt.Println(slices.Index(nums, 99)) // -1
```

---

### `IndexFunc`

Поиск индекса по условию:

```go
nums := []int{1, 3, 4, 7}

i := slices.IndexFunc(nums, func(n int) bool {
    return n%2 == 0
})

fmt.Println(i) // 2
```

Если подходящий элемент не найден, возвращается `-1`.

---

## 6. Бинарный поиск

### `BinarySearch`

Функция `slices.BinarySearch` ищет значение в **отсортированном** срезе.

```go
nums := []int{1, 2, 2, 3, 5}

i, found := slices.BinarySearch(nums, 2)
fmt.Println(i, found) // 1 true
```

Возвращаются два значения:

1. индекс;
2. найдено ли значение.

Если значение не найдено, возвращается позиция, в которую его можно вставить, чтобы сохранить сортировку.

```go
nums := []int{1, 2, 2, 3, 5}

i, found := slices.BinarySearch(nums, 4)
fmt.Println(i, found) // 4 false
```

Значение `4` можно вставить перед элементом с индексом `4`, то есть перед `5`.

Важно: срез должен быть отсортирован. Если он не отсортирован, результат не гарантирован.

---

### `BinarySearchFunc`

`BinarySearchFunc` позволяет искать по произвольному правилу сравнения.

```go
package main

import (
 "cmp"
 "fmt"
 "slices"
)

type User struct {
 Name string
 Age  int
}

func main() {
 users := []User{
  {"Alice", 25},
  {"Bob", 30},
  {"Carol", 35},
 }

 i, found := slices.BinarySearchFunc(users, "Bob", func(u User, name string) int {
  return cmp.Compare(u.Name, name)
 })

 fmt.Println(i, found) // 1 true
}
```

Здесь целевое значение имеет тип `string`, а элементы среза — тип `User`. Это удобно, когда не хочется создавать полноценный объект только для поиска.

---

## 7. Сортировка

### `Sort`

Функция `slices.Sort` сортирует срез по возрастанию.

```go
nums := []int{3, 1, 4, 1, 5}

slices.Sort(nums)

fmt.Println(nums) // [1 1 3 4 5]
```

`Sort` работает с типами, которые удовлетворяют ограничению `cmp.Ordered`: числа, строки и другие упорядоченные типы.

---

### Сортировка копией

`slices.Sort` изменяет исходный срез. Если нужно получить отсортированную копию, сначала сделайте клон:

```go
original := []int{3, 1, 2}

sorted := slices.Clone(original)
slices.Sort(sorted)

fmt.Println(original) // [3 1 2]
fmt.Println(sorted)   // [1 2 3]
```

---

### `SortFunc`

Если нужен собственный порядок сортировки, используется `slices.SortFunc`.

Функция сравнения должна возвращать:

- отрицательное число, если первый элемент меньше второго;
- `0`, если элементы равны;
- положительное число, если первый элемент больше второго.

Пример:

```go
package main

import (
 "cmp"
 "fmt"
 "slices"
)

type User struct {
 Name string
 Age  int
}

func main() {
 users := []User{
  {"Bob", 30},
  {"Alice", 25},
  {"Alice", 22},
 }

 slices.SortFunc(users, func(a, b User) int {
  if n := cmp.Compare(a.Name, b.Name); n != 0 {
   return n
  }
  return cmp.Compare(a.Age, b.Age)
 })

 fmt.Println(users)
 // [{Alice 22} {Alice 25} {Bob 30}]
}
```

Здесь сортировка сначала идёт по имени, а при равных именах — по возрасту.

---

### `SortStableFunc`

`SortStableFunc` — это стабильная сортировка. Она сохраняет относительный порядок элементов, которые функция сравнения считает равными.

```go
type Task struct {
 Priority int
 Name     string
}

tasks := []Task{
 {2, "deploy"},
 {1, "build"},
 {2, "test"},
}

slices.SortStableFunc(tasks, func(a, b Task) int {
 return cmp.Compare(a.Priority, b.Priority)
})
```

Если два элемента имеют одинаковый приоритет, их исходный порядок будет сохранён.

Обычная `SortFunc` стабильность не гарантирует.

---

### `IsSorted`

Проверяет, отсортирован ли срез.

```go
nums := []int{1, 2, 3, 4}

fmt.Println(slices.IsSorted(nums)) // true
```

---

### `IsSortedFunc`

Проверяет сортировку по собственной функции сравнения:

```go
type User struct {
 Name string
 Age  int
}

users := []User{
 {"Alice", 25},
 {"Bob", 30},
}

sorted := slices.IsSortedFunc(users, func(a, b User) int {
 return cmp.Compare(a.Name, b.Name)
})

fmt.Println(sorted) // true
```

---

## 8. Минимум и максимум

### `Min`

Возвращает минимальный элемент.

```go
nums := []int{3, 1, 4}

fmt.Println(slices.Min(nums)) // 1
```

---

### `Max`

Возвращает максимальный элемент.

```go
nums := []int{3, 1, 4}

fmt.Println(slices.Max(nums)) // 4
```

---

### Пустой срез

Если срез пуст, `Min` и `Max` вызывают panic.

```go
var nums []int

// slices.Min(nums) // panic
// slices.Max(nums) // panic
```

Поэтому перед вызовом проверяйте длину:

```go
if len(nums) > 0 {
 min := slices.Min(nums)
 max := slices.Max(nums)
 fmt.Println(min, max)
}
```

---

### `MinFunc` и `MaxFunc`

Если нужно сравнивать сложные структуры, используйте варианты с `Func`.

```go
type User struct {
 Name string
 Age  int
}

users := []User{
 {"Alice", 25},
 {"Bob", 30},
 {"Carol", 22},
}

oldest := slices.MaxFunc(users, func(a, b User) int {
 return cmp.Compare(a.Age, b.Age)
})

fmt.Println(oldest.Name) // Bob
```

---

## 9. Сравнение срезов

### `Equal`

Функция `slices.Equal` проверяет, одинаковы ли два среза по длине и элементам.

```go
a := []int{1, 2, 3}
b := []int{1, 2, 3}
c := []int{1, 2}

fmt.Println(slices.Equal(a, b)) // true
fmt.Println(slices.Equal(a, c)) // false
```

Ёмкость не учитывается.

---

### `EqualFunc`

Сравнение с собственной функцией равенства:

```go
type User struct {
 Name string
 Age  int
}

a := []User{{"Alice", 25}}
b := []User{{"Alice", 25}}

equal := slices.EqualFunc(a, b, func(x, y User) bool {
 return x.Name == y.Name && x.Age == y.Age
})

fmt.Println(equal) // true
```

Это полезно, когда тип элементов нельзя напрямую сравнивать через `==`, или когда нужна частичная equality-логика.

---

### `Compare`

Функция `slices.Compare` лексикографически сравнивает два среза.

Возвращаемое значение:

- отрицательное число, если первый срез меньше второго;
- `0`, если срезы равны;
- положительное число, если первый срез больше второго.

```go
fmt.Println(slices.Compare([]int{1, 2}, []int{1, 2}))    // 0
fmt.Println(slices.Compare([]int{1, 2}, []int{1, 2, 3})) // -1
fmt.Println(slices.Compare([]int{1, 3}, []int{1, 2}))    // 1
```

Если один срез является префиксом другого, более короткий считается меньшим.

---

### `CompareFunc`

То же самое, но с пользовательской функцией сравнения.

```go
type User struct {
 Name string
 Age  int
}

a := []User{{"Alice", 25}}
b := []User{{"Alice", 30}}

res := slices.CompareFunc(a, b, func(x, y User) int {
 return cmp.Compare(x.Age, y.Age)
})

fmt.Println(res) // -1
```

---

## 10. Клонирование и работа с capacity

### `Clone`

Функция `slices.Clone` создаёт копию среза.

```go
orig := []int{1, 2, 3}

copySlice := slices.Clone(orig)
copySlice[0] = 99

fmt.Println(orig)      // [1 2 3]
fmt.Println(copySlice) // [99 2 3]
```

Важно: это **поверхностная копия**.

Если элементы среза сами являются указателями, картами, каналами или функциями, копируются сами значения этих указателей/ссылок, а не данные, на которые они указывают.

Пример:

```go
type Item struct {
 Name string
}

orig := []*Item{{Name: "a"}}
cloned := slices.Clone(orig)

cloned[0].Name = "changed"

fmt.Println(orig[0].Name) // changed
```

То есть сам срез скопирован, но элементы-указатели указывают на те же объекты.

---

### Поведение `Clone` с `nil`

Если передан `nil`-срез, `Clone` возвращает `nil`.

```go
var s []int
c := slices.Clone(s)

fmt.Println(c == nil) // true
```

Если передан пустой, но не nil-срез, результат будет пустым не-nil-срезом.

---

### `Clip`

Функция `slices.Clip` уменьшает ёмкость среза до его длины.

Фактически это эквивалент:

```go
s = s[:len(s):len(s)]
```

Пример:

```go
s := make([]int, 0, 10)
s = append(s, 1, 2, 3)

fmt.Println(len(s), cap(s)) // 3 10

s = slices.Clip(s)

fmt.Println(len(s), cap(s)) // 3 3
```

`Clip` не копирует данные и не освобождает память сам по себе. Он лишь убирает запас ёмкости. После этого следующий `append` будет вынужден аллоцировать новый базовый массив.

---

## 11. Изменение среза

### `Insert`

Функция `slices.Insert` вставляет элементы перед указанным индексом.

```go
s := []int{1, 2, 5}

s = slices.Insert(s, 2, 3, 4)

fmt.Println(s) // [1 2 3 4 5]
```

Индекс может быть от `0` до `len(s)` включительно.

Если нужно добавить элементы в конец, обычно проще использовать встроенный `append`:

```go
s = append(s, 6, 7)
```

---

### `Delete`

Функция `slices.Delete` удаляет элементы из диапазона `[i, j)`.

```go
s := []int{1, 2, 3, 4, 5}

s = slices.Delete(s, 1, 3)

fmt.Println(s) // [1 4 5]
```

Здесь удаляются элементы с индексами `1` и `2`, то есть `2` и `3`.

Важно присваивать результат обратно:

```go
s = slices.Delete(s, i, j)
```

---

### `DeleteFunc`

`DeleteFunc` удаляет элементы, для которых функция возвращает `true`.

Доступна начиная с Go 1.22.

```go
nums := []int{1, 2, 3, 4, 5}

nums = slices.DeleteFunc(nums, func(n int) bool {
 return n%2 == 0
})

fmt.Println(nums) // [1 3 5]
```

---

### `Replace`

Функция `slices.Replace` заменяет диапазон `[i, j)` новыми элементами.

```go
s := []int{1, 2, 3, 4, 5}

s = slices.Replace(s, 1, 3, 9, 8, 7)

fmt.Println(s) // [1 9 8 7 3 4 5]
```

Диапазон `[1, 3)` — это элементы `2` и `3`. Они заменяются на `9, 8, 7`.

Если количество новых элементов меньше количества удаляемых, срез уменьшается:

```go
s := []int{1, 2, 3, 4, 5}

s = slices.Replace(s, 1, 4, 9)

fmt.Println(s) // [1 9 5]
```

Если больше — срез увеличивается и может быть выделен новый базовый массив.

---

## 12. Компактификация: удаление соседних дубликатов

### `Compact`

Функция `slices.Compact` удаляет **последовательные повторяющиеся** элементы.

```go
s := []int{1, 1, 2, 2, 3, 1}

s = slices.Compact(s)

fmt.Println(s) // [1 2 3 1]
```

Обратите внимание: последняя единица осталась, потому что она не была соседним дубликатом с предыдущей единицей.

Если нужно удалить все дубликаты, срез можно сначала отсортировать:

```go
s := []int{1, 3, 1, 2, 3, 2}

slices.Sort(s)
s = slices.Compact(s)

fmt.Println(s) // [1 2 3]
```

---

### `CompactFunc`

Удаляет соседние элементы, которые считаются равными по пользовательской функции.

```go
type User struct {
 Name string
 Age  int
}

users := []User{
 {"Alice", 25},
 {"Alice", 30},
 {"Bob", 20},
}

uniqueByName := slices.CompactFunc(users, func(a, b User) bool {
 return a.Name == b.Name
})

fmt.Println(uniqueByName)
// [{Alice 25} {Bob 20}]
```

`CompactFunc` тоже работает только с соседними элементами.

---

## 13. Разворот среза

Функция `slices.Reverse` разворачивает срез in-place.

```go
s := []int{1, 2, 3}

slices.Reverse(s)

fmt.Println(s) // [3 2 1]
```

Она ничего не возвращает.

---

## 14. Объединение срезов

### `Concat`

Функция `slices.Concat` объединяет несколько срезов в один новый срез.

Доступна начиная с Go 1.22.

```go
a := []int{1, 2}
b := []int{3}
c := []int{4, 5}

all := slices.Concat(a, b, c)

fmt.Println(all) // [1 2 3 4 5]
```

Это удобнее и часто чище, чем вручную делать:

```go
all := append(append(append([]int{}, a...), b...), c...)
```

---

## 15. Итераторы, появившиеся в Go 1.23

Начиная с Go 1.23, пакет `slices` поддерживает итераторы из пакета `iter`.

Основные функции:

| Функция | Назначение |
| --- | --- |
| `All` | итерация по парам индекс-значение |
| `Values` | итерация только по значениям |
| `Backward` | итерация с конца |
| `Collect` | собрать итератор в срез |
| `AppendSeq` | добавить значения из итератора в существующий срез |
| `Sorted` | собрать итератор и отсортировать результат |
| `SortedFunc` | собрать итератор и отсортировать с функцией сравнения |
| `SortedStableFunc` | стабильная сортировка значений из итератора |

---

### `All`

Итерация по индексу и значению:

```go
nums := []string{"a", "b", "c"}

for i, v := range slices.All(nums) {
 fmt.Println(i, v)
}
```

Вывод:

```text
0 a
1 b
2 c
```

---

### `Values`

Итерация только по значениям:

```go
nums := []int{10, 20, 30}

for v := range slices.Values(nums) {
 fmt.Println(v)
}
```

---

### `Backward`

Итерация с конца:

```go
nums := []int{1, 2, 3}

for i, v := range slices.Backward(nums) {
 fmt.Println(i, v)
}
```

Вывод:

```text
2 3
1 2
0 1
```

---

### `Collect`

Собирает значения из итератора в срез.

```go
seq := slices.Values([]int{3, 1, 2})

collected := slices.Collect(seq)

fmt.Println(collected) // [3 1 2]
```

---

### `AppendSeq`

Добавляет значения из итератора к существующему срезу.

```go
base := []int{1, 2}

seq := slices.Values([]int{3, 4})

result := slices.AppendSeq(base, seq)

fmt.Println(result) // [1 2 3 4]
```

---

### `Sorted`

Собирает значения из итератора и возвращает отсортированный срез.

```go
seq := slices.Values([]int{3, 1, 2})

sorted := slices.Sorted(seq)

fmt.Println(sorted) // [1 2 3]
```

---

### `SortedFunc`

Сортировка значений из итератора с собственной функцией сравнения:

```go
users := []User{
 {"Bob", 30},
 {"Alice", 25},
}

sorted := slices.SortedFunc(slices.Values(users), func(a, b User) int {
 return cmp.Compare(a.Name, b.Name)
})

fmt.Println(sorted)
```

---

## 16. Сводная таблица функций

Ниже приведён основной набор функций пакета `slices`.

### Поиск

| Функция | Назначение |
| --- | --- |
| `Contains` | проверяет наличие значения |
| `ContainsFunc` | проверяет наличие значения по условию |
| `Index` | возвращает индекс значения |
| `IndexFunc` | возвращает индекс первого элемента, удовлетворяющего условию |
| `BinarySearch` | бинарный поиск в отсортированном срезе |
| `BinarySearchFunc` | бинарный поиск с пользовательской функцией сравнения |

---

### Сортировка и экстремумы

| Функция | Назначение |
| --- | --- |
| `Sort` | сортирует срез по возрастанию |
| `SortFunc` | сортирует с пользовательской функцией сравнения |
| `SortStableFunc` | стабильная сортировка |
| `IsSorted` | проверяет, отсортирован ли срез |
| `IsSortedFunc` | проверяет сортировку по пользовательской функции |
| `Min` | возвращает минимальный элемент |
| `Max` | возвращает максимальный элемент |
| `MinFunc` | возвращает минимальный элемент по функции сравнения |
| `MaxFunc` | возвращает максимальный элемент по функции сравнения |

---

### Сравнение

| Функция | Назначение |
| --- | --- |
| `Equal` | проверяет равенство срезов |
| `EqualFunc` | проверяет равенство по пользовательской функции |
| `Compare` | лексикографически сравнивает срезы |
| `CompareFunc` | сравнивает срезы с пользовательской функцией |

---

### Копирование и ёмкость

| Функция | Назначение |
| --- | --- |
| `Clone` | создаёт копию среза |
| `Clip` | уменьшает ёмкость до длины |

---

### Изменение

| Функция | Назначение |
| --- | --- |
| `Insert` | вставляет элементы |
| `Delete` | удаляет диапазон элементов |
| `DeleteFunc` | удаляет элементы по условию |
| `Replace` | заменяет диапазон элементов |
| `Compact` | удаляет соседние дубликаты |
| `CompactFunc` | удаляет соседние дубликаты по функции равенства |
| `Reverse` | разворачивает срез |
| `Concat` | объединяет срезы |

---

### Итераторы, Go 1.23+

| Функция | Назначение |
| --- | --- |
| `All` | итератор по индексу и значению |
| `Values` | итератор по значениям |
| `Backward` | обратный итератор |
| `Collect` | собирает итератор в срез |
| `AppendSeq` | добавляет элементы из итератора к срезу |
| `Sorted` | собирает итератор и сортирует результат |
| `SortedFunc` | собирает итератор и сортирует с функцией сравнения |
| `SortedStableFunc` | стабильная сортировка значений из итератора |

---

## 17. Частые ошибки

### 1. Забыть присвоить результат

Неправильно:

```go
slices.Delete(s, 1, 3)
```

Правильно:

```go
s = slices.Delete(s, 1, 3)
```

Это касается функций, которые возвращают изменённый срез: `Delete`, `Insert`, `Replace`, `Compact`, `Clone`, `Concat`, `DeleteFunc` и других.

---

### 2. Использовать `BinarySearch` на неотсортированном срезе

```go
nums := []int{3, 1, 2}

i, found := slices.BinarySearch(nums, 2)
```

Так делать нельзя. Сначала нужно отсортировать:

```go
slices.Sort(nums)
i, found := slices.BinarySearch(nums, 2)
```

---

### 3. Ожидать, что `Compact` удалит все дубликаты

`Compact` удаляет только подряд идущие дубликаты.

```go
s := []int{1, 2, 1}

s = slices.Compact(s)

fmt.Println(s) // [1 2 1]
```

Если нужно удалить все дубликаты, сначала отсортируйте срез:

```go
slices.Sort(s)
s = slices.Compact(s)
```

---

### 4. Вызывать `Min` или `Max` на пустом срезе

```go
var nums []int

// slices.Min(nums) // panic
// slices.Max(nums) // panic
```

Нужна проверка:

```go
if len(nums) > 0 {
 min := slices.Min(nums)
}
```

---

### 5. Думать, что `Clone` делает глубокую копию

`Clone` копирует сам срез, но не делает глубокое копирование элементов.

Если элементы являются указателями, оба среза будут ссылаться на одни и те же объекты.

---

### 6. Ожидать стабильности от `SortFunc`

Если важен порядок равных элементов, используйте:

```go
slices.SortStableFunc(...)
```

---

## 18. Производительность и сложность

Ниже — приблизительная сложность типичных операций.

| Операция | Сложность |
| --- | --- |
| `Contains`, `Index` | O(n) |
| `BinarySearch` | O(log n), но требует отсортированный срез |
| `Sort`, `SortFunc`, `SortStableFunc` | O(n log n) |
| `Reverse` | O(n) |
| `Clone` | O(n) |
| `Concat` | O(суммарная длина срезов) |
| `Insert`, `Delete`, `Replace` | O(n), так как элементы могут сдвигаться |
| `Compact` | O(n) |
| `Clip` | O(1) |

---

## 19. Практический пример

Допустим, у нас есть список чисел. Нужно:

1. удалить дубликаты;
2. отсортировать;
3. найти минимальное и максимальное значение;
4. проверить, есть ли число 42.

```go
package main

import (
 "fmt"
 "slices"
)

func main() {
 nums := []int{5, 3, 8, 3, 1, 5, 8, 42}

 // Убираем дубликаты.
 slices.Sort(nums)
 nums = slices.Compact(nums)

 fmt.Println("После удаления дубликатов:", nums)

 if len(nums) == 0 {
  fmt.Println("срез пуст")
  return
 }

 min := slices.Min(nums)
 max := slices.Max(nums)

 fmt.Println("min:", min)
 fmt.Println("max:", max)

 fmt.Println("contains 42:", slices.Contains(nums, 42))

 i, found := slices.BinarySearch(nums, 42)
 fmt.Println("binary search:", i, found)
}
```

Пример вывода:

```text
После удаления дубликатов: [1 3 5 8 42]
min: 1
max: 42
contains 42: true
binary search: 4 true
```

---

## 20. Когда что использовать

| Задача | Что использовать |
| --- | --- |
| Проверить наличие элемента | `Contains` |
| Проверить наличие по условию | `ContainsFunc` |
| Найти индекс | `Index` |
| Найти индекс по условию | `IndexFunc` |
| Быстрый поиск в отсортированном срезе | `BinarySearch` |
| Отсортировать числа или строки | `Sort` |
| Отсортировать структуры | `SortFunc` |
| Стабильная сортировка | `SortStableFunc` |
| Найти минимум | `Min` |
| Найти максимум | `Max` |
| Сравнить срезы | `Equal`, `Compare` |
| Скопировать срез | `Clone` |
| Вставить элементы | `Insert` |
| Удалить элементы | `Delete`, `DeleteFunc` |
| Заменить элементы | `Replace` |
| Удалить соседние дубликаты | `Compact` |
| Объединить несколько срезов | `Concat` |
| Итерация в стиле range | `All`, `Values`, `Backward` |

---

## Источники

1. **Официальная документация пакета `slices`**  
   <https://pkg.go.dev/slices>  

2. **Go 1.21 Release Notes**  
   <https://go.dev/doc/go1.21>  

3. **Go 1.22 Release Notes**  
   <https://go.dev/doc/go1.22>  

4. **Go 1.23 Release Notes**  
   <https://go.dev/doc/go1.23>  

5. **Документация пакета `cmp`**  
   <https://pkg.go.dev/cmp>  

6. **Документация пакета `iter`**  
   <https://pkg.go.dev/iter>  

7. **Исходный код пакета `slices` в репозитории Go**  
   <https://github.com/golang/go/tree/master/src/slices>  
