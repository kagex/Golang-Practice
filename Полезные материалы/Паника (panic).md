# Panic и Recover в Go

## Введение

В Go есть два механизма обработки исключительных ситуаций:
- **Ошибки (`error`)** — для ожидаемых проблем (валидация, I/O, сеть)
- **Panic/Recover** — для фатальных, непредвиденных ошибок

> 💡 **Философия Go:** Panic — это **не** замена исключениям из Java/C#/Python. Это механизм для ситуаций, когда программа **не может продолжать работу корректно**.

---

## Что такое Panic

`panic` — это встроенная функция, которая:
1. Немедленно прекращает выполнение текущей функции
2. Начинает **раскрутку стека** (stack unwinding)
3. Выполняет все отложенные вызовы (`defer`) в обратном порядке
4. Если panic не перехвачена через `recover` — программа завершается с ошибкой

```go
func main() {
    fmt.Println("Start")
    panic("что-то пошло не так")
    fmt.Println("End") // ❌ Никогда не выполнится
}
// Вывод:
// Start
// panic: что-то пошло не так
// goroutine 1 [running]: ...
```

### Сигнатура

```go
func panic(v any)
```

Аргумент `v` может быть любого типа (`any`), но по соглашению обычно передаётся строка или `error`.

---

## Когда использовать Panic

### ✅ Допустимые случаи

| Ситуация | Пример |
|----------|--------|
| Неинвариантное состояние | Нарушены внутренние предположения кода |
| Ошибка программиста | Неверные аргументы в библиотеке |
| Ошибки инициализации | Невозможно загрузить критический конфиг при старте |
| Баги, которые нельзя обработать | Индекс вне диапазона в алгоритме, который должен быть корректным |
| Прототипирование / тесты | Быстрая проверка гипотез |

```go
// ✅ Ошибка программиста — неверный аргумент
func SetTimeout(d time.Duration) {
    if d < 0 {
        panic(fmt.Sprintf("timeout must be non-negative, got %v", d))
    }
    // ...
}

// ✅ Нарушение инварианта
func (s *Stack) Pop() int {
    if s.len == 0 {
        panic("Pop called on empty stack") // баг в логике вызывающего кода
    }
    s.len--
    return s.data[s.len]
}

// ✅ Критическая ошибка инициализации
var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        panic(fmt.Sprintf("failed to connect to database: %v", err))
    }
}
```

### ❌ НЕ используйте Panic для

| Ситуация | Что делать вместо этого |
|----------|------------------------|
| Ошибки ввода-вывода | `return err` |
| Ошибки валидации пользовательских данных | `return ErrValidation` |
| Ошибки сети | `return err` + retry |
| Отсутствие файла | `return os.ErrNotExist` |
| Любая ожидаемая ошибка | Возвращайте `error` |

```go
// ❌ Плохо — это ожидаемая ошибка
func ReadConfig(path string) Config {
    data, err := os.ReadFile(path)
    if err != nil {
        panic(err) // ❌ Файл может отсутствовать — это нормально
    }
    // ...
}

// ✅ Хорошо
func ReadConfig(path string) (Config, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return Config{}, fmt.Errorf("read config %q: %w", path, err)
    }
    // ...
}
```

---

## Механизм раскрутки стека

Когда вызывается `panic`, Go начинает **раскрутку стека**:

1. Текущая функция немедленно останавливается
2. Все `defer`-вызовы в текущей функции выполняются в порядке **LIFO**
3. Управление возвращается к вызывающей функции
4. Процесс повторяется для каждого уровня стека
5. Если ни один `recover` не перехватил panic — программа завершается

```go
func a() {
    defer fmt.Println("defer in a") // ✅ Выполнится даже при panic
    b()
    fmt.Println("after b in a")     // ❌ Не выполнится
}

func b() {
    defer fmt.Println("defer in b") // ✅ Выполнится даже при panic
    c()
    fmt.Println("after c in b")     // ❌ Не выполнится
}

func c() {
    defer fmt.Println("defer in c") // ✅ Выполнится даже при panic
    panic("boom!")
}

func main() {
    defer fmt.Println("defer in main") // ✅ Выполнится даже при panic
    a()
    fmt.Println("after a in main")     // ❌ Не выполнится
}

// Вывод:
// defer in c       ← LIFO внутри c()
// defer in b       ← LIFO внутри b()
// defer in a       ← LIFO внутри a()
// defer in main    ← LIFO внутри main()
// panic: boom!
```

> ⚠️ **Важно:** `defer` выполняется **даже при panic**. Это ключевой механизм для очистки ресурсов.

---

## Recover — перехват Panic

`recover` — это встроенная функция, которая позволяет перехватить panic и продолжить выполнение программы.

### Сигнатура

```go
func recover() any
```

Возвращает значение, переданное в `panic()`, или `nil`, если:
- Panic не было
- `recover` вызван **вне** `defer` (в этом случае **всегда** возвращает `nil`)

### ⚠️ Критические правила

1. **`recover` работает ТОЛЬКО внутри `defer`**
2. Перехватывает panic только из **текущей горутины**
3. После успешного recover выполнение продолжается **после defer**, а не после panic

```go
func safeCall() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("Перехвачено: %v\n", r)
        }
    }()
    
    panic("ошибка!")
    fmt.Println("После panic") // ❌ Не выполнится
}

func main() {
    safeCall()
    fmt.Println("Программа продолжает работу") // ✅ Выполнится
}

// Вывод:
// Перехвачено: ошибка!
// Программа продолжает работу
```

### ❌ Типичные ошибки с recover

```go
// ❌ Ошибка 1: recover вне defer — ВСЕГДА возвращает nil
func bad1() {
    r := recover() // r всегда nil здесь, даже если была panic
    panic("test")
}

// ❌ Ошибка 2: recover в обычной функции, вызванной из defer
func handler() {
    r := recover() // НЕ работает! Это не прямой вызов из defer
    fmt.Println(r)
}

func bad2() {
    defer handler() // handler вызывается из defer, но recover внутри handler не работает
    panic("test")
}

// ✅ Правильно: recover напрямую в анонимной функции defer
func good() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("OK:", r)
        }
    }()
    panic("test")
}
```

> 💡 **Почему так?** Спецификация Go требует, чтобы `recover` был вызван **напрямую** из deferred-функции. Вызов через промежуточную функцию не считается прямым.

---

## Паттерны использования Recover

### 1. Защита HTTP-хендлеров

Самый распространённый паттерн в веб-серверах:

```go
func recoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("PANIC recovered: %v\nStack: %s", r, debug.Stack())
                http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            }
        }()
        next.ServeHTTP(w, r)
    })
}
```

### 2. Преобразование panic в error

```go
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic in divide: %v", r)
        }
    }()
    
    if b == 0 {
        panic("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("Ошибка:", err) // panic in divide: division by zero
    } else {
        fmt.Println("Результат:", result)
    }
}
```

### 3. Защита горутин

**Panic в горутине убивает ВСЮ программу**, если не перехвачен:

```go
// ❌ Опасно — паника убьёт всю программу
go func() {
    doSomethingRisky()
}()

// ✅ Безопасно — перехватываем panic
go func() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("Goroutine panic: %v", r)
        }
    }()
    doSomethingRisky()
}()
```

### 4. Тестирование кода, который должен паниковать

```go
func TestPanicOnInvalidInput(t *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            t.Error("Expected panic, but didn't get one")
        }
    }()
    
    SetTimeout(-1 * time.Second) // Должна паниковать
}
```

Или с `testify`:
```go
func TestPanicOnInvalidInput(t *testing.T) {
    assert.Panics(t, func() {
        SetTimeout(-1 * time.Second)
    })
}
```

---

## Defer и Panic: взаимодействие

### Defer выполняется при panic

Это гарантирует очистку ресурсов даже в аварийных ситуациях:

```go
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close() // ✅ Закроется даже при panic
    
    data := make([]byte, 1000000)
    n, err := file.Read(data)
    if n > maxAllowed {
        panic("file too large") // file.Close() всё равно вызовется
    }
    // ...
    return nil
}
```

### Модификация возвращаемых значений при panic

```go
func safeOperation() (err error) {
    defer func() {
        if r := recover(); r != nil {
            // Перезаписываем возвращаемую ошибку
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    
    panic("something broke")
    return nil // Не выполнится, но err будет установлен в defer
}
```

### Порядок выполнения нескольких defer при panic

```go
func example() {
    defer fmt.Println("first defer")
    defer fmt.Println("second defer")
    defer fmt.Println("third defer")
    panic("stop!")
}

// Вывод (LIFO):
// third defer
// second defer
// first defer
// panic: stop!
```

---

## Аргумент Panic

### Что передавать в panic

```go
// ✅ Строка — самый частый вариант
panic("unexpected state")

// ✅ error — когда есть контекст
panic(fmt.Errorf("invalid config key %q: %w", key, err))

// ✅ Кастомный тип — для структурированной информации
type PanicInfo struct {
    Component string
    Reason    string
    Context   map[string]any
}

panic(PanicInfo{
    Component: "parser",
    Reason:    "unexpected token",
    Context:   map[string]any{"line": 42, "token": "}"},
})

// ⚠️ panic(nil) — см. раздел ниже про Go 1.21+
```

### ⚠️ `panic(nil)` и Go 1.21+

До Go 1.21 `panic(nil)` создавал проблему: `recover()` возвращал `nil`, и было невозможно отличить перехваченную пустую панику от ситуации, когда паники не было вообще.

**Начиная с Go 1.21**, `recover()` при перехвате `panic(nil)` возвращает специальный тип `*runtime.PanicNilError`:

```go
// Go 1.21+
defer func() {
    r := recover()
    if r != nil {
        // r будет *runtime.PanicNilError, а не nil
        fmt.Printf("Type: %T, Value: %v\n", r, r)
        // Type: *runtime.PanicNilError, Value: panic called with nil argument
    }
}()
panic(nil)
```

> 💡 **Рекомендация:** Даже с учётом этого исправления, избегайте `panic(nil)`. Всегда передавайте осмысленное сообщение.

### Получение значения из recover

```go
defer func() {
    r := recover()
    switch v := r.(type) {
    case string:
        log.Printf("String panic: %s", v)
    case error:
        log.Printf("Error panic: %v", v)
    case PanicInfo:
        log.Printf("Structured panic: %+v", v)
    case *runtime.PanicNilError:
        log.Printf("Panic with nil argument (Go 1.21+)")
    default:
        log.Printf("Unknown panic type %T: %v", v, v)
    }
}()
```

---

## Panic в горутинах

### Критическое правило

> ⚠️ **Panic в одной горутине НЕ может быть перехвачена другой горутиной.** Каждая горутина имеет свой стек, и recover работает только в пределах текущего стека.

```go
// ❌ Этот recover НЕ перехватит panic из горутины
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r) // Не сработает!
        }
    }()
    
    go func() {
        panic("goroutine panic")
    }()
    
    time.Sleep(time.Second)
}
// Программа ЗАВЕРШИТСЯ с ошибкой
```

### Правильная защита горутин

Каждая горутина должна сама защищать себя:

```go
func safeGoroutine(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Goroutine panic recovered: %v", r)
            }
        }()
        fn()
    }()
}

// Использование
safeGoroutine(func() {
    doSomethingRisky()
})
```

### Библиотеки для защиты горутин

В продакшене часто используют обёртки:

```go
// Пример из стандартной практики
func GoSafe(fn func()) {
    go func() {
        defer func() {
            if r := recover(); r != nil {
                buf := make([]byte, 64<<10)
                n := runtime.Stack(buf, false)
                log.Printf("goroutine panic: %v\n%s", r, buf[:n])
            }
        }()
        fn()
    }()
}
```

---

## Runtime Panics

Go runtime автоматически вызывает panic в следующих случаях:

| Ситуация | Сообщение |
|----------|-----------|
| Индекс вне диапазона | `index out of range [X] with length Y` |
| Разыменование nil-указателя | `nil pointer dereference` |
| Деление на ноль | `integer divide by zero` |
| Запись в nil-map | `assignment to entry in nil map` |
| Закрытие закрытого канала | `close of closed channel` |
| Отправка в закрытый канал | `send on closed channel` |
| Недостаточно памяти | `out of memory` |
| Переполнение стека | `stack overflow` |
| `panic(nil)` (Go 1.21+) | `*runtime.PanicNilError` |

```go
// Все эти операции вызывают runtime panic:
var s []int
_ = s[10]           // index out of range

var p *int
_ = *p              // nil pointer dereference

_ = 1 / 0           // integer divide by zero

var m map[string]int
m["key"] = 1        // assignment to entry in nil map

ch := make(chan int)
close(ch)
close(ch)           // close of closed channel
```

> 💡 Эти panic можно перехватить через recover, но **обычно не стоит** — они указывают на баги в коде.

---

## Debug: получение стектрейса

### Через debug.Stack()

```go
import "runtime/debug"

defer func() {
    if r := recover(); r != nil {
        log.Printf("Panic: %v\nStack trace:\n%s", r, debug.Stack())
    }
}()
```

### Через runtime.Callers

```go
import "runtime"

func getStackTrace() string {
    pcs := make([]uintptr, 32)
    n := runtime.Callers(2, pcs) // пропускаем getStackTrace и recover
    frames := runtime.CallersFrames(pcs[:n])
    
    var sb strings.Builder
    for {
        frame, more := frames.Next()
        fmt.Fprintf(&sb, "%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
        if !more {
            break
        }
    }
    return sb.String()
}
```

### Переменная окружения GOTRACEBACK

```bash
# Показать полный стектрейс при crash
GOTRACEBACK=all ./myapp

# Значения:
# 0      — только информация о panic
# 1      — стектрейс пользовательского кода (по умолчанию)
# all    — стектрейс включая runtime
# system — как all, но с системными горутинами
# crash  — как system + core dump (для отладки через gdb/lldb)
```

> 💡 **Совет:** В продакшене используйте `GOTRACEBACK=all` для сбора полных стектрейсов при падении сервиса. Это значительно упрощает диагностику.

---

## Таблица сравнения: Error vs Panic

| Критерий | Error | Panic |
|----------|-------|-------|
| Назначение | Ожидаемые проблемы | Неожиданные/фатальные ошибки |
| Обработка | Явная проверка `if err != nil` | Автоматическая раскрутка стека |
| Производительность | Минимальные накладные расходы | Дорогая операция (раскрутка стека) |
| Восстановление | Продолжение работы | Только через recover в defer |
| Видимость в API | Часть сигнатуры функции | Скрытое поведение |
| Когда использовать | Всегда, когда возможно | Только когда невозможно вернуть error |
| Пример | Файл не найден | Инвариант нарушен |

---

## Лучшие практики

### 1. Предпочитайте errors над panic

```go
// ✅ Хорошо
func ParseInt(s string) (int, error) {
    // ...
}

// ❌ Плохо
func MustParseInt(s string) int {
    // panic при ошибке
}
```

### 2. Используйте panic только для невозможных состояний

```go
// ✅ Допустимо — это баг, если сюда попали
switch action {
case "create", "update", "delete":
    // обработка
default:
    panic(fmt.Sprintf("unknown action %q: this is a bug", action))
}
```

### 3. Всегда защищайте горутины

```go
go func() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("panic: %v", r)
        }
    }()
    // работа
}()
```

### 4. Логируйте стектрейс при recover

```go
defer func() {
    if r := recover(); r != nil {
        log.Printf("PANIC: %v\n%s", r, debug.Stack())
    }
}()
```

### 5. Документируйте функции, которые могут паниковать

```go
// Divide divides a by b.
// Panics if b is zero.
func Divide(a, b int) int {
    if b == 0 {
        panic("divide by zero")
    }
    return a / b
}
```

### 6. Не используйте panic для управления потоком

```go
// ❌ Ужасно — panic как goto
func findItem(items []Item, id int) Item {
    for _, item := range items {
        if item.ID == id {
            panic(item) // ❌ Использовать panic для возврата значения
        }
    }
    panic(nil)
}

// ✅ Хорошо
func findItem(items []Item, id int) (Item, bool) {
    for _, item := range items {
        if item.ID == id {
            return item, true
        }
    }
    return Item{}, false
}
```

### 7. Используйте GOTRACEBACK в продакшене

```bash
# Для полного стектрейса при падении
GOTRACEBACK=all ./myapp
```

---

## Частые ошибки

### 1. Recover вне defer

```go
// ❌ Никогда не сработает — recover() всегда вернёт nil
func bad() {
    r := recover() // всегда nil
    panic("test")
}
```

### 2. Попытка перехватить panic другой горутины

```go
// ❌ Не сработает
defer recover()
go func() { panic("other goroutine") }()
```

### 3. Panic для ожидаемых ошибок

```go
// ❌ Плохо
func GetUser(id int) *User {
    user, err := db.Find(id)
    if err != nil {
        panic(err) // ❌ Ошибка БД — ожидаемая ситуация
    }
    return user
}
```

### 4. Забытый стектрейс

```go
// ❌ Плохо — потеряна информация о месте ошибки
defer func() {
    if r := recover(); r != nil {
        log.Println(r) // Нет стектрейса!
    }
}()

// ✅ Хорошо
defer func() {
    if r := recover(); r != nil {
        log.Printf("%v\n%s", r, debug.Stack())
    }
}()
```

### 5. Panic с nil

```go
// ⚠️ Допустимо в Go 1.21+ (возвращает *runtime.PanicNilError),
// но всё равно сложно отлаживать
panic(nil)

// ✅ Лучше передать осмысленное сообщение
panic("unexpected nil value in field X")
```

---

## Заключение

Panic и Recover — мощные, но опасные инструменты:

1. **Panic** — для фатальных ошибок и невозможных состояний
2. **Recover** — только внутри `defer`, только в текущей горутине
3. **Defer** — выполняется даже при panic, гарантирует очистку
4. **Горутины** — каждая должна сама защищать себя от panic
5. **Errors** — предпочтительнее panic в 99% случаев
6. **Go 1.21+** — `panic(nil)` теперь возвращает `*runtime.PanicNilError`

> 🎯 **Золотое правило:** **Panic — для фатальных случаев, Error — для всего остального.** Если ошибку можно вернуть через `return err` — верните её. Используйте panic только когда продолжение работы **невозможно или бессмысленно**.

---

## Ссылки

- [Official Go Specification - Handling panics](https://go.dev/ref/spec#Handling_panics)
- [Effective Go - Panic and Recover](https://go.dev/doc/effective_go#panic)
- [Go Blog - Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
- [Package runtime/debug - Stack](https://pkg.go.dev/runtime/debug#Stack)
- [Package builtin - panic](https://pkg.go.dev/builtin#panic)
- [Package builtin - recover](https://pkg.go.dev/builtin#recover)
- [Go 1.21 Release Notes - PanicNilError](https://go.dev/doc/go1.21#runtime)
