# Пакет strings в Go

## Введение

Пакет `strings` предоставляет простые функции для работы со строками (тип `string`). Поскольку строки в Go **неизменяемы** (immutable), большинство функций возвращают новую строку, а не изменяют исходную.

```go
import "strings"
```

> 💡 **Важно:** Строки в Go — это последовательность байтов в кодировке UTF-8. Для работы с отдельными символами (рунами) используйте `rune` и пакет `unicode/utf8`.

---

## Поиск в строках

### Проверка наличия подстроки

```go
s := "Привет, мир!"

strings.Contains(s, "мир")        // true
strings.Contains(s, "hello")      // false

strings.ContainsAny(s, "abc")     // true (содержит 'а' или 'б' или 'в')
strings.ContainsAny(s, "xyz")     // false

strings.ContainsRune(s, 'П')      // true
strings.ContainsRune(s, 'z')      // false
```

### Проверка начала и конца строки

```go
s := "config.json"

strings.HasPrefix(s, "config")    // true
strings.HasPrefix(s, ".json")     // false

strings.HasSuffix(s, ".json")     // true
strings.HasSuffix(s, "config")    // false
```

### Поиск позиции подстроки

Возвращают индекс **первого байта** найденной подстроки или `-1`, если не найдено.

```go
s := "hello world"

strings.Index(s, "world")         // 6
strings.Index(s, "xyz")           // -1

strings.LastIndex(s, "l")         // 9 (последнее вхождение)

strings.IndexByte(s, 'o')         // 4
strings.IndexRune(s, 'w')         // 6

strings.IndexAny(s, "aeiou")      // 1 (первая гласная 'e')
strings.LastIndexAny(s, "aeiou")  // 7 ('o' в "world")
```

### ⚠️ Важно: индексы в байтах, а не в рунах

```go
s := "Привет"  // кириллица занимает 2 байта на символ

strings.Index(s, "вет")  // 6 (не 3!)
len(s)                   // 12 байт
len([]rune(s))           // 6 рун (символов)
```

---

## Сравнение строк

### `EqualFold` — сравнение без учёта регистра

```go
strings.EqualFold("Go", "go")       // true
strings.EqualFold("Привет", "привет") // true
strings.EqualFold("Hello", "World") // false
```

### `Compare` — лексикографическое сравнение

Возвращает:
- `-1` если `a < b`
- `0` если `a == b`
- `+1` если `a > b`

```go
strings.Compare("a", "b")  // -1
strings.Compare("b", "b")  // 0
strings.Compare("c", "b")  // 1
```

> 💡 **Совет:** В Go idiomatic использовать операторы `==`, `<`, `>` вместо `Compare`. Функция нужна в основном для сортировок.

---

## Преобразование регистра

```go
s := "Hello, Мир!"

strings.ToUpper(s)  // "HELLO, МИР!"
strings.ToLower(s)  // "hello, мир!"
strings.ToTitle(s)  // "HELLO, МИР!" (для большинства символов как ToUpper)

// Title case — каждое слово с заглавной (устарело, используйте golang.org/x/text/cases)
strings.Title("hello world")  // "Hello World"
```

### Специальные преобразования

```go
import "unicode"

// С учётом специфической локали
strings.ToUpperSpecial(unicode.TurkishCase, "istanbul")  // "İSTANBUL"
strings.ToLowerSpecial(unicode.TurkishCase, "İSTANBUL")  // "istanbul"
```

---

## Обрезка (Trim)

### `TrimSpace` — удаление пробелов по краям

```go
strings.TrimSpace("  hello  ")      // "hello"
strings.TrimSpace("\t\nhello\n\t")  // "hello"
```

### `Trim` — удаление указанных символов с обоих концов

```go
strings.Trim("!!!hello!!!", "!")       // "hello"
strings.Trim("  hello  ", " ")         // "hello"
strings.Trim("###hello###", "#")       // "hello"
```

### `TrimLeft` и `TrimRight`

```go
strings.TrimLeft("!!!hello!!!", "!")   // "hello!!!"
strings.TrimRight("!!!hello!!!", "!")  // "!!!hello"
```

### `TrimPrefix` и `TrimSuffix`

```go
strings.TrimPrefix("config.json", "config.")  // "json"
strings.TrimPrefix("config.json", "data.")    // "config.json" (не изменилась)

strings.TrimSuffix("config.json", ".json")    // "config"
strings.TrimSuffix("config.json", ".txt")     // "config.json"
```

---

## Разбиение и соединение

### `Split` — разбиение строки на срез

```go
s := "a,b,c,d"

strings.Split(s, ",")       // ["a", "b", "c", "d"]
strings.Split(s, "-")       // ["a,b,c,d"] (разделитель не найден)
strings.Split("", ",")      // [""]

// SplitN — ограничение количества частей
strings.SplitN("a,b,c,d", ",", 2)   // ["a", "b,c,d"]
strings.SplitN("a,b,c,d", ",", 3)   // ["a", "b", "c,d"]
strings.SplitN("a,b,c,d", ",", -1)  // ["a", "b", "c", "d"] (как Split)

// SplitAfter — сохраняет разделитель
strings.SplitAfter("a,b,c", ",")    // ["a,", "b,", "c"]
```

### `Cut` — разделение строки на две части

Элегантная и безопасная замена `strings.Index` и `strings.SplitN`, когда нужно разделить строку по первому вхождению разделителя.

```go
s := "key=value=123"

before, after, found := strings.Cut(s, "=")
// before: "key"
// after: "value=123"
// found: true

// Если разделитель не найден:
b, a, f := strings.Cut("hello", "=")
// b: "hello"
// a: ""
// f: false
```

### `Join` — соединение среза в строку

```go
parts := []string{"a", "b", "c", "d"}

strings.Join(parts, ",")    // "a,b,c,d"
strings.Join(parts, " - ")  // "a - b - c - d"
strings.Join(parts, "")     // "abcd"
```

### `Fields` — разбиение по пробельным символам

```go
s := "  hello   world   golang  "

strings.Fields(s)  // ["hello", "world", "golang"]
// Пустые строки между словами удаляются автоматически
```

### `FieldsFunc` — разбиение по кастомной функции

```go
f := func(c rune) bool {
    return c == ',' || c == ';'
}
strings.FieldsFunc("a,b;c,d", f)  // ["a", "b", "c", "d"]
```

---

## Замена

### `Replace` и `ReplaceAll`

```go
s := "hello world hello"

strings.Replace(s, "hello", "hi", 1)   // "hi world hello" (1 замена)
strings.Replace(s, "hello", "hi", 2)   // "hi world hi" (2 замены)
strings.Replace(s, "hello", "hi", -1)  // "hi world hi" (все замены)

strings.ReplaceAll(s, "hello", "hi")   // "hi world hi" (все замены)
```

### `Replacer` — множественная замена за один проход

Эффективнее, чем несколько `Replace`, когда нужно заменить много разных подстрок.

```go
replacer := strings.NewReplacer(
    "<", "&lt;",
    ">", "&gt;",
    "&", "&amp;",
)

html := replacer.Replace("<div>Hello & World</div>")
// "&lt;div&gt;Hello &amp; World&lt;/div&gt;"
```

> 💡 `Replacer` создаётся один раз и переиспользуется — он потокобезопасен.

---

## Другие операции

### `Repeat` — повторение строки

```go
strings.Repeat("Go", 3)     // "GoGoGo"
strings.Repeat("-", 10)     // "----------"
strings.Repeat("abc", 0)    // ""
```

> ⚠️ **Осторожно:** `Repeat("a", 1000000000)` вызовет `panic: strings: Repeat output length overflow` или `out of memory`.

### `Count` — подсчёт непересекающихся вхождений

```go
s := "hello world"

strings.Count(s, "l")      // 3
strings.Count(s, "lo")     // 1
strings.Count(s, "xyz")    // 0

// Пустая строка: возвращает len(s) + 1
strings.Count("hello", "")  // 6
```

### `Map` — преобразование каждой руны

```go
// Удаление всех гласных
f := func(r rune) rune {
    switch r {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        return -1  // -1 = удалить руну
    }
    return r
}
strings.Map(f, "Hello World")  // "Hll Wrld"

// Преобразование регистра
toUpper := func(r rune) rune {
    if r >= 'a' && r <= 'z' {
        return r - 32
    }
    return r
}
strings.Map(toUpper, "hello")  // "HELLO"
```

### `Clone` — клонирование строки

Предотвращает скрытые утечки памяти. В Go подстрока делит базовый массив с исходной строкой. Если вы извлекаете маленькую подстроку из гигантского текста, весь текст останется в памяти. `Clone` создаёт независимую копию.

```go
import "strings"

giantString := getHugeText() // гигантский текст на несколько мегабайт

// ❌ Плохо: giantString не удалится из памяти, так как small ссылается на её данные
small := giantString[:10] 

// ✅ Хорошо: создаётся новая независимая строка, giantString соберётся garbage collector'ом
smallSafe := strings.Clone(giantString[:10])
```

---

## `strings.Builder` — эффективное построение строк

При множественных конкатенациях используйте `strings.Builder` вместо `+=` — это **гораздо быстрее**, не создаёт промежуточные строки и более эффективно.

### Базовый пример

```go
var b strings.Builder

b.WriteString("Hello")
b.WriteString(", ")
b.WriteString("World")
b.WriteByte('!')
b.WriteRune('😀')

result := b.String()  // "Hello, World!😀"
```

### Сравнение производительности

```go
// ❌ Плохо — O(n²) из-за копирования строки на каждой итерации
s := ""
for i := 0; i < 10000; i++ {
    s += "a"
}

// ✅ Хорошо — O(n), аллокации минимальны
var b strings.Builder
b.Grow(10000)  // заранее резервируем память
for i := 0; i < 10000; i++ {
    b.WriteString("a")
}
s := b.String()
```

### Полный пример

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    var b strings.Builder
    
    // Предварительное резервирование памяти (опционально, но ускоряет)
    b.Grow(100)
    
    b.WriteString("User: ")
    b.WriteString("Ivan")
    b.WriteByte('\n')
    b.WriteString("Age: ")
    fmt.Fprintf(&b, "%d", 25)
    b.WriteByte('\n')
    
    fmt.Print(b.String())
}
```

### Методы `strings.Builder`

| Метод | Описание |
|-------|----------|
| `WriteString(s string)` | Записать строку |
| `WriteByte(c byte)` | Записать байт |
| `WriteRune(r rune)` | Записать руну |
| `Write(p []byte)` | Записать срез байт |
| `Grow(n int)` | Заранее зарезервировать `n` байт |
| `String() string` | Получить итоговую строку |
| `Len() int` | Текущая длина |
| `Reset()` | Сбросить билдер (переиспользование буфера) |

### ⚠️ Важно: Builder нельзя копировать

```go
var b strings.Builder
b.WriteString("hello")
b2 := b  // ❌ Ошибка: go vet выдаст предупреждение
// Builder содержит указатель на буфер, копирование ломает его
```

---

## `strings.Reader` — чтение из строки

Позволяет использовать строку как `io.Reader`, `io.ReaderAt`, `io.Seeker`, `io.WriterTo`, `io.ByteScanner` и `io.RuneScanner`.

```go
r := strings.NewReader("Hello, World!")

// Чтение как io.Reader
buf := make([]byte, 5)
n, _ := r.Read(buf)
fmt.Println(string(buf[:n]))  // "Hello"

// Позиционирование (Seek)
r.Seek(7, io.SeekStart)  // перейти на позицию 7
buf2 := make([]byte, 5)
n, _ = r.Read(buf2)
fmt.Println(string(buf2[:n]))  // "World"

// Размер
r.Len()   // оставшиеся байты
r.Size()  // общий размер
```

### Практический пример: парсинг как из файла

```go
// Часто используется для передачи строки в функции, ожидающие io.Reader
data := "key1=value1\nkey2=value2\n"
scanner := bufio.NewScanner(strings.NewReader(data))
for scanner.Scan() {
    fmt.Println(scanner.Text())
}
```

---

## `strings` vs `bytes`

| Критерий | `strings` | `bytes` |
|----------|-----------|---------|
| Тип данных | `string` (неизменяемая) | `[]byte` (изменяемая) |
| Когда использовать | Работа с готовым текстом | Манипуляции с изменяемыми данными |
| Производительность | Копирование при каждой операции | Изменение на месте |
| Связь | Легко конвертировать: `[]byte(s)` и `string(b)` | — |

### Когда что использовать

```go
// ✅ strings — когда данные не меняются
if strings.Contains(filename, ".json") {
    // ...
}

// ✅ bytes — когда нужно многократно изменять данные
buf := []byte("hello")
for i := range buf {
    if buf[i] == 'l' {
        buf[i] = 'L'  // изменяем на месте
    }
}
```

> ⚠️ **Важно:** Преобразование `string ↔ []byte` создаёт **копию данных** (кроме некоторых оптимизаций компилятора в Go 1.20+).

---

## Работа с UTF-8 и рунами

Строки в Go — это последовательность байт. Для корректной работы с Unicode используйте `rune` и пакет `unicode/utf8`.

### Пример: подсчёт символов, а не байт

```go
s := "Привет"

len(s)                    // 12 (байт)
len([]rune(s))            // 6 (символов/рун)
utf8.RuneCountInString(s) // 6 (символов, эффективнее)
```

### Итерация по рунам

```go
s := "Hello, 世界"

// ❌ Неправильно — итерируем по байтам
for i := 0; i < len(s); i++ {
    fmt.Printf("%c\n", s[i])  // сломается на кириллице/китайском
}

// ✅ Правильно — итерируем по рунам
for i, r := range s {
    fmt.Printf("позиция %d: %c\n", i, r)
}
```

### Пример: обращение строки с поддержкой UTF-8

```go
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

Reverse("Привет")  // "тевирП"
Reverse("Hello")   // "olleH"
```

---

## Таблица быстрого справочника

### Поиск

| Функция | Описание |
|---------|----------|
| `Contains(s, substr)` | Содержит ли `s` подстроку `substr` |
| `ContainsAny(s, chars)` | Содержит ли любой символ из `chars` |
| `ContainsRune(s, r)` | Содержит ли руну `r` |
| `HasPrefix(s, prefix)` | Начинается ли с `prefix` |
| `HasSuffix(s, suffix)` | Заканчивается ли на `suffix` |
| `Index(s, substr)` | Индекс первого вхождения (-1 если нет) |
| `LastIndex(s, substr)` | Индекс последнего вхождения |
| `IndexByte(s, c)` | Индекс байта `c` |
| `IndexRune(s, r)` | Индекс руны `r` |

### Разбиение и соединение

| Функция | Описание |
|---------|----------|
| `Split(s, sep)` | Разбить по разделителю |
| `SplitN(s, sep, n)` | Разбить на максимум `n` частей |
| `SplitAfter(s, sep)` | Разбить, сохраняя разделитель |
| `Join(elems, sep)` | Соединить срез строк |
| `Fields(s)` | Разбить по пробелам |

### Преобразование

| Функция | Описание |
|---------|----------|
| `ToUpper(s)` | В верхний регистр |
| `ToLower(s)` | В нижний регистр |
| `EqualFold(s, t)` | Сравнение без учёта регистра |
| `TrimSpace(s)` | Убрать пробелы по краям |
| `Trim(s, cutset)` | Убрать символы из `cutset` с краёв |
| `TrimPrefix(s, prefix)` | Убрать префикс |
| `TrimSuffix(s, suffix)` | Убрать суффикс |

### Другие

| Функция | Описание |
|---------|----------|
| `Replace(s, old, new, n)` | Заменить `n` вхождений |
| `ReplaceAll(s, old, new)` | Заменить все вхождения |
| `Repeat(s, count)` | Повторить `count` раз |
| `Count(s, substr)` | Подсчитать непересекающиеся вхождения |

---

## Частые ошибки

### 1. Изменение строки по индексу

```go
// ❌ Нельзя — строки неизменяемы
s := "hello"
s[0] = 'H'  // ошибка компиляции: cannot assign to s[0]

// ✅ Правильно — через []rune или []byte
runes := []rune(s)
runes[0] = 'H'
s = string(runes)  // "Hello"
```

### 2. Неэффективная конкатенация в цикле

```go
// ❌ Плохо — O(n²)
result := ""
for _, word := range words {
    result += word + " "
}

// ✅ Хорошо — O(n)
var b strings.Builder
for _, word := range words {
    b.WriteString(word)
    b.WriteByte(' ')
}
result := b.String()
```

### 3. Игнорирование UTF-8

```go
s := "Привет"

// ❌ Неправильно — работает с байтами
fmt.Println(s[0])  // 208 (первый байт 'П', не сам символ)

// ✅ Правильно — работает с рунами
runes := []rune(s)
fmt.Println(string(runes[0]))  // "П"
```

### 4. Забытая проверка результата Trim

```go
s := "config.json"

// ❌ Может быть неожиданным
result := strings.TrimPrefix(s, "data.")  // "config.json" (не изменилась!)

// ✅ Правильно — проверяем, был ли префикс
if strings.HasPrefix(s, "data.") {
    result := strings.TrimPrefix(s, "data.")
    // ...
}
```

### 5. Split с пустым результатом

```go
s := ""
parts := strings.Split(s, ",")  // [""], а не []!
fmt.Println(len(parts))         // 1, а не 0

// ✅ Проверяйте длину корректно
if len(parts) == 1 && parts[0] == "" {
    // строка была пустой
}
```

---

## Лучшие практики

### 1. Используйте `strings.Builder` для множественных конкатенаций

```go
var b strings.Builder
b.Grow(expectedSize)  // заранее резервируйте память
b.WriteString(...)
// ...
result := b.String()
```

### 2. Используйте `EqualFold` для сравнения без учёта регистра

```go
// ❌ Медленно и создаёт новые строки
if strings.ToLower(a) == strings.ToLower(b) { ... }

// ✅ Быстро и без аллокаций
if strings.EqualFold(a, b) { ... }
```

### 3. Используйте `strings.NewReader` для тестов

```go
// Вместо создания временного файла
reader := strings.NewReader("test data")
// Можно передать в любую функцию, ожидающую io.Reader
processData(reader)
```

### 4. Проверяйте HasPrefix перед TrimPrefix

```go
// ✅ Более явно и понятно
if strings.HasPrefix(s, "http://") {
    s = strings.TrimPrefix(s, "http://")
}
```

### 5. Помните про UTF-8 при работе с символами

```go
// Для подсчёта символов, а не байт
count := utf8.RuneCountInString(s)

// Для итерации по символам
for i, r := range s {
    // i — индекс байта, r — руна
}
```

### 6. Используйте Replacer для множественных замен

```go
// ✅ Один проход, быстрее
r := strings.NewReplacer("<", "&lt;", ">", "&gt;", "&", "&amp;")
result := r.Replace(input)
```

### 7. Используйте `Cut` для разбиения на две части (Go 1.18+)

```go
// ✅ Элегантное разделение
key, value, found := strings.Cut(line, "=")
```

### 8. Используйте `Clone` для предотвращения утечек памяти (Go 1.18+)

```go
// ✅ Безопасное извлечение подстроки
small := strings.Clone(hugeString[:100])
```

---

## Заключение

Пакет `strings` — один из самых используемых в Go. Ключевые моменты:

1. **Строки неизменяемы** — все функции возвращают новую строку
2. **Индексы в байтах** — используйте `[]rune` для работы с символами Unicode
3. **`strings.Builder`** — для эффективного построения строк
4. **`strings.Reader`** — для использования строки как `io.Reader`
5. **`EqualFold`** — для быстрого сравнения без учёта регистра
6. **`Replacer`** — для множественных замен за один проход
7. **`Cut`** — для элегантного разбиения на две части (Go 1.18+)
8. **`Clone`** — для предотвращения утечек памяти при подстроках (Go 1.18+)

Правильное использование `strings` делает код эффективным, читаемым и надёжным.

---

## Ссылки

- [Official Go Documentation - strings package](https://pkg.go.dev/strings)
- [Official Go Documentation - unicode/utf8 package](https://pkg.go.dev/unicode/utf8)
- [Effective Go - Strings](https://go.dev/doc/effective_go#strings)
- [Go Blog - Strings, bytes, runes and characters in Go](https://go.dev/blog/strings)
