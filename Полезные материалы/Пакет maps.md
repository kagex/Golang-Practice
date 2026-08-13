# Пакет `maps` в Go

Пакет `maps` — это стандартный пакет Go, который содержит generic-функции для работы с картами (`map`). Он не заменяет встроенный тип `map`, а предоставляет удобные вспомогательные функции: клонирование, копирование, сравнение, очистку, удаление по условию, а начиная с Go 1.23 — итераторные функции.

Официальная документация пакета: `https://pkg.go.dev/maps`.

---

## 1. Что такое пакет `maps`

В Go есть встроенный тип:

```go
map[string]int
```

А есть стандартный пакет `maps`, который помогает работать с такими картами.

Импорт:

```go
import "maps"
```

Пакет `maps` появился в стандартной библиотеке начиная с **Go 1.21**.

---

## 2. Версии Go и доступность функций

| Версия Go | Что доступно |
| --- | --- |
| Go 1.21 | Базовые функции пакета `maps`: `Clear`, `Clone`, `Copy`, `DeleteFunc`, `Equal`, `EqualFunc` |
| Go 1.23 | Итераторные функции: `All`, `Keys`, `Values`, `Collect`, `Insert` |

Локально проверить доступные функции можно так:

```bash
go doc maps
```

---

## 3. Зачем нужен пакет `maps`

До пакета `maps` многие операции приходилось писать вручную.

Например, копирование карты:

```go
dst := make(map[string]int, len(src))

for k, v := range src {
 dst[k] = v
}
```

С пакетом `maps`:

```go
dst := maps.Clone(src)
```

Сравнение двух карт:

```go
same := maps.Equal(m1, m2)
```

Удаление элементов по условию:

```go
maps.DeleteFunc(m, func(k string, v int) bool {
 return v < 0
})
```

Основные преимущества пакета:

1. **Читаемость** — операция видна из имени функции.
2. **Типобезопасность** — используются generics.
3. **Удобные функции для частых операций**.
4. **Поддержка именованных типов карт** благодаря ограничению `~map[K]V`.
5. **Интеграция с итераторами** в Go 1.23+.

---

## 4. Важное уточнение: пакет `maps` не делает карты потокобезопасными

Пакет `maps` — это набор функций для обычных карт. Он не добавляет синхронизацию.

Если с картой работают несколько горутин, нужно использовать:

- `sync.Mutex`;
- `sync.RWMutex`;
- или `sync.Map` в подходящих сценариях.

Пример небезопасного кода:

```go
m := make(map[string]int)

go func() {
 m["a"] = 1
}()

go func() {
 maps.DeleteFunc(m, func(k string, v int) bool {
  return v == 0
 })
}()
```

Так делать нельзя без синхронизации.

# 5. Основные функции пакета `maps`

---

## 5.1. `maps.Clear`

Функция `maps.Clear` удаляет все элементы из карты.

```go
import "maps"

m := map[string]int{
 "a": 1,
 "b": 2,
}

maps.Clear(m)

fmt.Println(len(m)) // 0
```

Также в Go начиная с 1.21 есть встроенный `clear`:

```go
clear(m)
```

Оба варианта очищают карту.

Важно: очистка карты не обязательно освобождает память так же, как создание новой карты. Если нужно начать с полностью новой карты, обычно делают:

```go
m = make(map[string]int)
```

---

## 5.2. `maps.Clone`

Функция `maps.Clone` создаёт копию карты.

```go
original := map[string]int{
 "a": 1,
 "b": 2,
}

clone := maps.Clone(original)

clone["a"] = 100

fmt.Println(original["a"]) // 1
fmt.Println(clone["a"])    // 100
```

Это **поверхностная копия**.

То есть копируется сама карта: ключи и значения присваиваются в новую карту. Но если значения являются указателями, срезами, картами, каналами или интерфейсами, содержащими ссылочные данные, то копируются сами значения этих ссылок, а не объекты, на которые они указывают.

Пример:

```go
type User struct {
 Name string
}

original := map[string]*User{
 "alice": {Name: "Alice"},
}

clone := maps.Clone(original)

clone["alice"].Name = "Changed"

fmt.Println(original["alice"].Name) // Changed
```

Карта скопирована, но указатель на `User` остался тем же.

---

### Поведение `maps.Clone` с `nil`

Если передана `nil`-карта, `maps.Clone` возвращает `nil`.

```go
var m map[string]int

clone := maps.Clone(m)

fmt.Println(clone == nil) // true
```

---

## 5.3. `maps.Copy`

Функция `maps.Copy` копирует все пары ключ-значение из одной карты в другую.

```go
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

Особенности:

- `dst` изменяется;
- существующие ключи перезаписываются;
- `dst` не очищается перед копированием.

Если нужно, чтобы `dst` стала точной копией `src`, можно сначала очистить `dst`:

```go
clear(dst)
maps.Copy(dst, src)
```

Или использовать `maps.Clone`:

```go
dst := maps.Clone(src)
```

---

## 5.4. `maps.DeleteFunc`

Функция `maps.DeleteFunc` удаляет из карты элементы, для которых функция-предикат возвращает `true`.

```go
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

Предикат получает ключ и значение:

```go
func(k KeyType, v ValueType) bool
```

Если вернуть `true`, элемент будет удалён.

Ручной аналог:

```go
for k, v := range m {
 if v < 0 {
  delete(m, k)
 }
}
```

---

## 5.5. `maps.Equal`

Функция `maps.Equal` проверяет, содержат ли две карты одинаковые пары ключ-значение.

```go
m1 := map[string]int{
 "a": 1,
 "b": 2,
}

m2 := map[string]int{
 "a": 1,
 "b": 2,
}

fmt.Println(maps.Equal(m1, m2)) // true
```

Важно:

- порядок элементов в `map` не учитывается;
- длина карт должна совпадать;
- значения сравниваются через `==`.

Поэтому `maps.Equal` требует, чтобы значения были сравнимыми.

---

### `maps.Equal` и `nil` / пустые карты

`nil`-карта и пустая карта считаются равными, потому что обе не содержат элементов.

```go
var m1 map[string]int
m2 := make(map[string]int)

fmt.Println(maps.Equal(m1, m2)) // true
```

---

### Ограничение с float и NaN

Так как сравнение идёт через `==`, значения `NaN` не считаются равными сами себе.

```go
import "math"

m1 := map[string]float64{"x": math.NaN()}
m2 := map[string]float64{"x": math.NaN()}

fmt.Println(maps.Equal(m1, m2)) // false
```

Если нужна особая логика сравнения, используйте `maps.EqualFunc`.

---

## 5.6. `maps.EqualFunc`

Функция `maps.EqualFunc` сравнивает две карты с помощью пользовательской функции сравнения значений.

```go
a := map[string]int{
 "x": 1,
}

b := map[string]int64{
 "x": 1,
}

equal := maps.EqualFunc(a, b, func(v1 int, v2 int64) bool {
 return int64(v1) == v2
})

fmt.Println(equal) // true
```

Это полезно, когда:

- типы значений разные, но их можно сравнить;
- нужна нестандартная логика равенства;
- значения несравнимы через `==`, например срезы.

Пример со срезами:

```go
import "slices"

m1 := map[string][]int{
 "a": {1, 2},
}

m2 := map[string][]int{
 "a": {1, 2},
}

equal := maps.EqualFunc(m1, m2, func(v1, v2 []int) bool {
 return slices.Equal(v1, v2)
})

fmt.Println(equal) // true
```

---

# 6. Итераторные функции, Go 1.23+

Начиная с Go 1.23, пакет `maps` поддерживает итераторы из пакета `iter`.

Основные функции:

| Функция | Назначение |
| --- | --- |
| `maps.All` | итерация по парам ключ-значение |
| `maps.Keys` | итерация по ключам |
| `maps.Values` | итерация по значениям |
| `maps.Collect` | собрать пары ключ-значение в новую карту |
| `maps.Insert` | получить функцию для вставки пар в существующую карту |

Эти функции особенно полезны вместе с пакетом `slices`.

---

## 6.1. `maps.All`

`maps.All` возвращает итератор по парам ключ-значение.

```go
m := map[string]int{
 "a": 1,
 "b": 2,
}

for k, v := range maps.All(m) {
 fmt.Println(k, v)
}
```

Фактически это функциональный аналог обычного:

```go
for k, v := range m {
 fmt.Println(k, v)
}
```

Но `maps.All` удобно использовать в итераторных цепочках.

Порядок обхода не гарантируется, как и у обычного `range` по карте.

---

## 6.2. `maps.Keys`

`maps.Keys` возвращает итератор по ключам карты.

```go
m := map[string]int{
 "a": 1,
 "b": 2,
}

for k := range maps.Keys(m) {
 fmt.Println(k)
}
```

Собрать ключи в срез можно через `slices.Collect`:

```go
import (
 "maps"
 "slices"
)

keys := slices.Collect(maps.Keys(m))

fmt.Println(keys)
```

---

## 6.3. `maps.Values`

`maps.Values` возвращает итератор по значениям карты.

```go
m := map[string]int{
 "a": 1,
 "b": 2,
}

for v := range maps.Values(m) {
 fmt.Println(v)
}
```

Собрать значения в срез:

```go
values := slices.Collect(maps.Values(m))

fmt.Println(values)
```

---

## 6.4. Отсортированные ключи

Ключи карты сами по себе не имеют порядка. Если нужен детерминированный порядок, ключи нужно отсортировать.

В Go 1.23+ это можно сделать так:

```go
import (
 "fmt"
 "maps"
 "slices"
)

m := map[string]int{
 "banana": 2,
 "apple": 1,
 "cherry": 3,
}

for _, k := range slices.Sorted(maps.Keys(m)) {
 fmt.Println(k, m[k])
}
```

Ожидаемый порядок ключей:

```text
apple 1
banana 2
cherry 3
```

Для старых версий Go обычно используют ручной сбор и сортировку ключей:

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

---

## 6.5. `maps.Collect`

`maps.Collect` собирает пары ключ-значение из итератора в новую карту.

```go
src := map[string]int{
 "a": 1,
 "b": 2,
}

dst := maps.Collect(maps.All(src))

fmt.Println(dst) // map[a:1 b:2]
```

Если в итераторе встречаются повторяющиеся ключи, обычно побеждает более поздняя пара, потому что происходит обычная запись в карту.

---

## 6.6. `maps.Insert`

`maps.Insert` возвращает функцию, которая вставляет пары ключ-значение в существующую карту.

Эта функция полезна при работе с итераторами.

```go
src := map[string]int{
 "a": 1,
 "b": 2,
}

dst := make(map[string]int)

insert := maps.Insert(dst)

for k, v := range maps.All(src) {
 insert(k, v)
}

fmt.Println(dst) // map[a:1 b:2]
```

Также возвращаемую функцию можно использовать как приёмник итератора:

```go
maps.All(src)(maps.Insert(dst))
```

Звучит непривычно, но это просто вызов итератора с передачей ему функции-приёмника.

---

# 7. Сводная таблица функций

| Функция | Назначение | Изменяет исходную карту? | Версия |
| --- | --- | ---: | ---: |
| `maps.Clear` | удаляет все элементы | да | Go 1.21+ |
| `maps.Clone` | создаёт поверхностную копию | нет | Go 1.21+ |
| `maps.Copy` | копирует элементы из одной карты в другую | изменяет `dst` | Go 1.21+ |
| `maps.DeleteFunc` | удаляет элементы по условию | да | Go 1.21+ |
| `maps.Equal` | сравнивает две карты | нет | Go 1.21+ |
| `maps.EqualFunc` | сравнивает две карты с пользовательской функцией | нет | Go 1.21+ |
| `maps.All` | итератор по парам ключ-значение | нет | Go 1.23+ |
| `maps.Keys` | итератор по ключам | нет | Go 1.23+ |
| `maps.Values` | итератор по значениям | нет | Go 1.23+ |
| `maps.Collect` | собирает итератор пар в новую карту | создаёт новую карту | Go 1.23+ |
| `maps.Insert` | возвращает функцию для вставки пар в карту | изменяет переданную карту | Go 1.23+ |

---

# 8. Сигнатуры основных функций

Ниже приведены типовые сигнатуры функций пакета `maps`. Они полезны, чтобы понимать ограничения типов.

## Базовые функции

```go
func Clear[M ~map[K]V, K comparable, V any](m M)
```

```go
func Clone[M ~map[K]V, K comparable, V any](m M) M
```

```go
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)
```

```go
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool)
```

```go
func Equal[M1 ~map[K]V, M2 ~map[K]V, K comparable, V comparable](m1 M1, m2 M2) bool
```

```go
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](
 m1 M1,
 m2 M2,
 eq func(V1, V2) bool,
) bool
```

---

## Итераторные функции, Go 1.23+

```go
func All[M ~map[K]V, K comparable, V any](m M) iter.Seq2[K, V]
```

```go
func Keys[M ~map[K]V, K comparable, V any](m M) iter.Seq[K]
```

```go
func Values[M ~map[K]V, K comparable, V any](m M) iter.Seq[V]
```

```go
func Collect[K comparable, V any](seq iter.Seq2[K, V]) map[K]V
```

```go
func Insert[M ~map[K]V, K comparable, V any](m M) func(K, V) bool
```

Здесь:

- `K comparable` — ключи карты должны быть сравнимыми;
- `V any` — значения могут быть любого типа;
- `~map[K]V` — поддерживаются в том числе именованные типы, чей базовый тип является картой;
- `iter.Seq` и `iter.Seq2` — типы итераторов из пакета `iter`.

---

# 9. Именованные типы карт

Пакет `maps` хорошо работает с именованными типами карт.

Например:

```go
type Scores map[string]int

original := Scores{
 "alice": 10,
 "bob":   8,
}

clone := maps.Clone(original)

fmt.Printf("%T\n", clone) // Scores
```

Это возможно благодаря generic-ограничению `~map[K]V`.

---

# 10. Поверхностное копирование

`maps.Clone` и `maps.Copy` не делают глубокую копию.

Если значения являются ссылочными типами, обе карты будут ссылаться на одни и те же данные.

Пример:

```go
type Item struct {
 Name string
}

original := map[string]*Item{
 "a": {Name: "original"},
}

clone := maps.Clone(original)

clone["a"].Name = "changed"

fmt.Println(original["a"].Name) // changed
```

Если нужна глубокая копия, её нужно реализовывать вручную.

Например:

```go
func deepClone(m map[string]*Item) map[string]*Item {
 result := make(map[string]*Item, len(m))

 for k, v := range m {
  copied := *v
  result[k] = &copied
 }

 return result
}
```

---

# 11. Сравнение карт

В Go карты нельзя сравнивать оператором `==`, кроме сравнения с `nil`.

Нельзя:

```go
m1 := map[string]int{"a": 1}
m2 := map[string]int{"a": 1}

// m1 == m2 // ошибка компиляции
```

Можно:

```go
var m map[string]int
fmt.Println(m == nil) // true
```

Пакет `maps` решает проблему сравнения содержимого:

```go
fmt.Println(maps.Equal(m1, m2))
```

Для нестандартных значений:

```go
fmt.Println(maps.EqualFunc(m1, m2, func(a, b []int) bool {
 return slices.Equal(a, b)
}))
```

---

# 12. Практический пример

Допустим, есть исходная конфигурация. Нужно:

1. Скопировать её.
2. Удалить отрицательные значения.
3. Сравнить с оригиналом.
4. Вывести ключи в отсортированном порядке.

```go
package main

import (
 "fmt"
 "maps"
 "slices"
)

func main() {
 original := map[string]int{
  "timeout":  30,
  "retries":  -1,
  "max_conn": 10,
  "delay":    0,
 }

 working := maps.Clone(original)

 maps.DeleteFunc(working, func(k string, v int) bool {
  return v < 0
 })

 fmt.Println("equal:", maps.Equal(original, working))

 for _, k := range slices.Sorted(maps.Keys(working)) {
  fmt.Printf("%s = %d\n", k, working[k])
 }
}
```

Пример вывода:

```text
equal: false
delay = 0
max_conn = 10
timeout = 30
```

Здесь `slices.Sorted(maps.Keys(working))` требует Go 1.23+.

---

# 13. Типичные ошибки

## 1. Ожидать, что `maps.Clone` делает глубокую копию

Нет. Это поверхностная копия.

```go
clone := maps.Clone(original)
```

Копируются ключи и значения как значения. Если значения являются указателями или срезами, копируются ссылки.

---

## 2. Думать, что `maps.Copy` очищает принимающую карту

`maps.Copy` только добавляет и перезаписывает пары.

```go
maps.Copy(dst, src)
```

Если нужно полное замещение, используйте:

```go
clear(dst)
maps.Copy(dst, src)
```

или:

```go
dst = maps.Clone(src)
```

---

## 3. Использовать пакет `maps` для конкурентного доступа

Пакет `maps` не синхронизирует доступ.

Если несколько горутин читают и пишут карту, нужна защита:

```go
mu.Lock()
m[key] = value
mu.Unlock()
```

---

## 4. Ожидать порядок от `maps.Keys`, `maps.Values`, `maps.All`

Порядок обхода карты не гарантируется.

Если нужен порядок, сортируйте ключи:

```go
for _, k := range slices.Sorted(maps.Keys(m)) {
 fmt.Println(k, m[k])
}
```

---

## 5. Использовать `maps.Equal` для значений с `NaN`

`NaN != NaN`, поэтому:

```go
maps.Equal(m1, m2)
```

может вернуть `false`, даже если вы ожидаете логического равенства.

Для таких случаев используйте `maps.EqualFunc`.

---

# 14. Что пакет `maps` не делает

Пакет `maps`:

- не делает карты потокобезопасными;
- не создаёт глубокие копии автоматически;
- не добавляет порядок хранения ключей;
- не заменяет встроенный тип `map`;
- не предоставляет тип `set`;
- не решает автоматически проблемы памяти после удаления большого числа элементов.

---

# 15. Когда использовать пакет `maps`

Используйте пакет `maps`, когда нужно:

- скопировать карту;
- скопировать одну карту в другую;
- сравнить две карты;
- удалить элементы по условию;
- очистить карту;
- получить итератор по ключам или значениям;
- собрать карту из итератора пар ключ-значение;
- отсортировать ключи через `maps.Keys` и `slices.Sorted` в Go 1.23+.

---

# Источники

1. **Package `maps` documentation**  
   <https://pkg.go.dev/maps>

2. **Go 1.21 Release Notes**  
   <https://go.dev/doc/go1.21>  
   Здесь описано появление пакета `maps` и связанных функций.

3. **Go 1.23 Release Notes**  
   <https://go.dev/doc/go1.23>  
   Здесь описаны итераторы и функции, связанные с пакетом `iter`.

4. **Package `iter` documentation**  
   <https://pkg.go.dev/iter>

5. **Package `slices` documentation**  
   <https://pkg.go.dev/slices>  
   Используется вместе с `maps` для сбора и сортировки ключей/значений.

6. **The Go Programming Language Specification — Map types**  
   <https://go.dev/ref/spec#Map_types>
