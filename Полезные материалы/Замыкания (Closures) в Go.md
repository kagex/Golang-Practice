# Замыкания (Closures) в Go

## Введение

Замыкание (closure) — это функция, которая **захватывает переменные** из внешней области видимости и сохраняет доступ к ним даже после того, как внешняя функция завершила выполнение.

В Go замыкания реализуются через **анонимные функции** (function literals), которые могут ссылаться на переменные, объявленные за их пределами.

```go
func outer() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

> 💡 **Ключевая особенность Go:** Замыкания захватывают переменные **по ссылке**, а не по значению. Все обращения к переменной из замыкания работают с одной и той же ячейкой памяти.

---

## Что такое замыкание

Замыкание состоит из двух частей:
1. **Функция** (обычно анонимная)
2. **Окружение** (environment) — набор переменных из внешней области видимости, которые функция "запомнила"

```go
func makeAdder(x int) func(int) int {
    return func(y int) int {
        return x + y  // x захвачена из makeAdder
    }
}

add5 := makeAdder(5)   // замыкание с x = 5
add10 := makeAdder(10) // замыкание с x = 10

fmt.Println(add5(3))   // 8  (5 + 3)
fmt.Println(add10(3))  // 13 (10 + 3)
```

Каждый вызов `makeAdder` создаёт **новое замыкание** со своим собственным окружением. Переменная `x` в `add5` и `add10` — это **разные** ячейки памяти.

---

## Как работает захват переменных

### Захват по ссылке

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c()) // 1
fmt.Println(c()) // 2
fmt.Println(c()) // 3
```

Переменная `count` **продолжает существовать** после завершения `counter()`, потому что на неё ссылается замыкание. Сборщик мусора не освободит её, пока существует замыкание.

### Переменная "живёт", пока живёт замыкание

```go
func createClosure() func() {
    data := make([]int, 1000000) // большой массив
    return func() {
        fmt.Println(len(data))
    }
}

closure := createClosure()
// `data` не будет освобождён, пока существует closure
```

> ⚠️ Это может приводить к **утечкам памяти**, если замыкание хранится дольше, чем нужно.

---

## Как устроено замыкание под капотом

Компилятор Go не создаёт "магических" объектов. Замыкание — это **структура**, содержащая:
1. Указатель на функцию
2. Указатели на захваченные переменные (или их копии)

### Эквивалентное представление

То, что пишет программист:
```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

Примерно то, что генерирует компилятор:
```go
// Структура окружения
type counterEnv struct {
    count *int  // указатель на захваченную переменную
}

// Функция замыкания получает указатель на окружение
func counterInvoke(env *counterEnv) int {
    *env.count++
    return *env.count
}

func counter() func() int {
    count := 0           // размещается в куче (escape to heap)
    env := &counterEnv{count: &count}
    
    // Возвращаем "связку" функции и окружения
    return func() int {
        return counterInvoke(env)
    }
}
```

### Ключевые следствия

1. **Захваченные переменные "убегают" в кучу** — они должны пережить внешнюю функцию
2. **Каждое замыкание имеет свою структуру окружения** — поэтому `add5` и `add10` независимы
3. **Обращение к переменной = разыменование указателя** — есть небольшой overhead
4. **Несколько замыканий могут делить одно окружение** — если захватывают одни и те же переменные

```go
func shared() (func(), func()) {
    x := 0
    inc := func() { x++ }
    get := func() int { return x }
    return inc, get  // оба замыкания ссылаются на ОДНУ переменную x
}

inc, get := shared()
inc()
inc()
fmt.Println(get()) // 2 — оба замыкания работают с одной x
```

---

## Базовые примеры

### 1. Счётчик (stateful function)

```go
func NewCounter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

counter1 := NewCounter()
counter2 := NewCounter()

fmt.Println(counter1()) // 1
fmt.Println(counter1()) // 2
fmt.Println(counter2()) // 1 (отдельный счётчик!)
fmt.Println(counter1()) // 3
```

### 2. Инкремент/декремент с общей переменной

Можно вернуть **несколько замыканий**, разделяющих общую переменную:

```go
func NewCounter() (inc func(), dec func(), val func() int) {
    count := 0
    inc = func() { count++ }
    dec = func() { count-- }
    val = func() int { return count }
    return
}

inc, dec, val := NewCounter()
inc(); inc(); inc(); dec()
fmt.Println(val()) // 2
```

### 3. Приватные переменные (инкапсуляция)

Замыкания эмулируют приватные поля, как в ООП:

```go
func NewWallet(initial int) (
    deposit func(int),
    withdraw func(int) error,
    balance func() int,
) {
    amount := initial // приватная переменная
    
    deposit = func(x int) {
        if x > 0 { amount += x }
    }
    
    withdraw = func(x int) error {
        if x <= 0 { return errors.New("сумма должна быть положительной") }
        if x > amount { return errors.New("недостаточно средств") }
        amount -= x
        return nil
    }
    
    balance = func() int { return amount }
    return
}

dep, wd, bal := NewWallet(100)
dep(50)
wd(200)                   // error: недостаточно средств
fmt.Println(bal())        // 150
// amount недоступна извне
```

---

## Практические паттерны

### 1. Генераторы последовательностей

#### Генератор ID

```go
func NewIDGenerator(prefix string) func() string {
    id := 0
    return func() string {
        id++
        return fmt.Sprintf("%s-%05d", prefix, id)
    }
}

userID := NewIDGenerator("USER")
fmt.Println(userID())  // USER-00001
fmt.Println(userID())  // USER-00002
```

#### Генератор Фибоначчи

```go
func Fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a+b
        return a
    }
}

fib := Fibonacci()
for i := 0; i < 10; i++ {
    fmt.Print(fib(), " ") // 1 1 2 3 5 8 13 21 34 55
}
```

#### Итератор

```go
func Range(start, end int) func() (int, bool) {
    current := start
    return func() (int, bool) {
        if current >= end { return 0, false }
        v := current
        current++
        return v, true
    }
}

iter := Range(1, 5)
for {
    v, ok := iter()
    if !ok { break }
    fmt.Println(v) // 1, 2, 3, 4
}
```

### 2. Middleware (декораторы)

Самый частый паттерн в веб-разработке:

```go
type Handler func(string) string

func WithLogging(next Handler) Handler {
    return func(input string) string {
        log.Printf("→ %q", input)
        result := next(input)
        log.Printf("← %q", result)
        return result
    }
}

func WithTiming(next Handler) Handler {
    return func(input string) string {
        start := time.Now()
        result := next(input)
        log.Printf("time: %v", time.Since(start))
        return result
    }
}

process := func(s string) string { return strings.ToUpper(s) }
handler := WithLogging(WithTiming(process))
fmt.Println(handler("hello"))
```

#### HTTP Rate Limiter (состояние в замыкании)

```go
func RateLimiter(rpm int, next http.HandlerFunc) http.HandlerFunc {
    tokens := rpm                      // состояние
    mu := sync.Mutex{}                 // состояние
    lastRefill := time.Now()           // состояние
    
    return func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        // логика пополнения токенов...
        if tokens <= 0 {
            mu.Unlock()
            http.Error(w, "rate limit", 429)
            return
        }
        tokens--
        mu.Unlock()
        next(w, r)
    }
}
```

Замыкание захватывает `tokens`, `mu` и `lastRefill` — они сохраняются между HTTP-запросами.

### 3. Memoization (кэширование)

```go
func Memoize[T comparable, R any](f func(T) R) func(T) R {
    cache := make(map[T]R)
    mu := sync.RWMutex{}
    
    return func(key T) R {
        mu.RLock()
        if v, ok := cache[key]; ok {
            mu.RUnlock()
            return v
        }
        mu.RUnlock()
        
        mu.Lock()
        defer mu.Unlock()
        if v, ok := cache[key]; ok { return v } // двойная проверка
        
        result := f(key)
        cache[key] = result
        return result
    }
}

slowSquare := func(n int) int {
    time.Sleep(100 * time.Millisecond)
    return n * n
}
fastSquare := Memoize(slowSquare)

fmt.Println(fastSquare(5)) // 25 (100ms)
fmt.Println(fastSquare(5)) // 25 (мгновенно!)
```

### 4. Фабрики объектов

```go
type Logger struct{ log func(string) }

func NewLogger(prefix string) *Logger {
    return &Logger{
        log: func(msg string) {
            fmt.Printf("[%s] %s: %s\n",
                time.Now().Format("15:04:05"), prefix, msg)
        },
    }
}

dbLogger := NewLogger("DB")
apiLogger := NewLogger("API")
dbLogger.log("connected") // [14:30:00] DB: connected
```

### 5. Частичное применение (Currying)

```go
func MultiplyBy(a int) func(int) int {
    return func(b int) int { return a * b }
}

double := MultiplyBy(2)
triple := MultiplyBy(3)
fmt.Println(double(5))  // 10
fmt.Println(triple(5))  // 15
```

### 6. Отложенные вычисления (Lazy Evaluation)

```go
type Lazy[T any] struct {
    once  sync.Once
    value T
    init  func() T
}

func NewLazy[T any](init func() T) *Lazy[T] {
    return &Lazy[T]{init: init}
}

func (l *Lazy[T]) Get() T {
    l.once.Do(func() { l.value = l.init() })
    return l.value
}

config := NewLazy(func() map[string]string {
    fmt.Println("Загрузка конфига...")
    return loadConfig()
})

fmt.Println(config.Get()) // Загрузка... (только при первом вызове)
fmt.Println(config.Get()) // уже загружен
```

### 7. Обработчики событий

```go
type EventEmitter struct {
    handlers map[string][]func(string)
}

func (e *EventEmitter) On(event string, handler func(string)) {
    e.handlers[event] = append(e.handlers[event], handler)
}

func (e *EventEmitter) Emit(event, data string) {
    for _, h := range e.handlers[event] { h(data) }
}

emitter := &EventEmitter{handlers: make(map[string][]func(string))}

callCount := 0
emitter.On("click", func(data string) {
    callCount++
    fmt.Printf("Клик #%d: %s\n", callCount, data)
})

emitter.Emit("click", "button1") // Клик #1: button1
```

### 8. Functional Options

Идиоматичный Go-паттерн для конфигурации:

```go
type Server struct {
    host    string
    port    int
    timeout time.Duration
}

type Option func(*Server)

func WithPort(p int) Option           { return func(s *Server) { s.port = p } }
func WithTimeout(d time.Duration) Option { return func(s *Server) { s.timeout = d } }

func NewServer(host string, opts ...Option) *Server {
    s := &Server{host: host, port: 8080, timeout: 30 * time.Second}
    for _, opt := range opts { opt(s) }
    return s
}

server := NewServer("localhost", WithPort(9000), WithTimeout(60*time.Second))
```

Каждая функция `With*` возвращает замыкание, захватывающее аргумент.

### 9. Тестирование (Mocking)

```go
// В продакшен-коде
var timeNow = time.Now

func IsWeekend() bool {
    d := timeNow().Weekday()
    return d == time.Saturday || d == time.Sunday
}

// В _test.go
func TestIsWeekend(t *testing.T) {
    original := timeNow
    defer func() { timeNow = original }()
    
    timeNow = func() time.Time {
        return time.Date(2025, 7, 20, 0, 0, 0, 0, time.UTC) // воскресенье
    }
    
    if !IsWeekend() { t.Error("ожидался выходной") }
}
```

### 10. Method Values

Присваивание метода переменной создаёт неявное замыкание:

```go
type Processor struct{ prefix string }
func (p *Processor) Do(s string) { fmt.Println(p.prefix, s) }

p := &Processor{prefix: "LOG:"}
handler := p.Do          // неявное замыкание: захватывает p
handler("test")          // LOG: test
```

---

## Подводные камни

### 1. ⚠️ Захват переменных цикла (до Go 1.22)

**До Go 1.22** все итерации использовали **одну и ту же** переменную:

```go
// ❌ До Go 1.22
var handlers []func()
for i := 0; i < 3; i++ {
    handlers = append(handlers, func() { fmt.Println(i) })
}
for _, h := range handlers { h() } // Вывод: 3 3 3
```

**Решение (до Go 1.22):**
```go
for i := 0; i < 3; i++ {
    i := i // локальная копия
    handlers = append(handlers, func() { fmt.Println(i) })
}
```

**Go 1.22+:** проблема решена на уровне языка — каждая итерация создаёт новую переменную.

### 2. ⚠️ Захват указателей из `range`

Даже в Go 1.22+ захват указателей требует осторожности:

```go
func processUsers(users []*User) []func() {
    var handlers []func()
    for _, u := range users {
        handlers = append(handlers, func() {
            fmt.Println(u.Name) // все увидят последнего!
        })
    }
    return handlers
}

// Решение: явная копия
for _, u := range users {
    u := u  // копия указателя
    handlers = append(handlers, func() { fmt.Println(u.Name) })
}
```

> **Правило:** при захвате указателей всегда создавайте явную копию.

### 3. ⚠️ Утечки памяти

#### Сценарий 1: захват большого слайса

```go
// ❌ Плохо
func createHandler() func() {
    huge := make([]byte, 100*1024*1024) // 100 MB
    huge[0] = 42
    return func() { fmt.Println(huge[0]) } // весь слайс в памяти!
}

// ✅ Хорошо: копируем нужное значение
func createHandler() func() {
    huge := make([]byte, 100*1024*1024)
    first := huge[0]
    return func() { fmt.Println(first) }
}
```

#### Сценарий 2: незакрытые ресурсы

```go
// ❌ Плохо: resp.Body удерживается замыканием
func handler(resp *http.Response) func() {
    return func() { fmt.Println(resp.StatusCode) }
}

// ✅ Хорошо
func handler(resp *http.Response) func() {
    code := resp.StatusCode
    resp.Body.Close()
    return func() { fmt.Println(code) }
}
```

### 4. ⚠️ Race Conditions

Несколько горутин, использующих одно замыкание с общей переменной:

```go
// ❌ Race condition
counter := func() {
    count := 0
    return func() int {
        count++
        return count
    }
}()
for i := 0; i < 1000; i++ { go counter() }

// ✅ Потокобезопасно
func SafeCounter() func() int {
    count := 0
    mu := sync.Mutex{}
    return func() int {
        mu.Lock()
        defer mu.Unlock()
        count++
        return count
    }
}

// Или с atomic
func AtomicCounter() func() int {
    var count atomic.Int64
    return func() int { return int(count.Add(1)) }
}
```

### 5. ⚠️ Затенение переменных (Shadowing)

```go
// ❌ := создаёт НОВУЮ переменную err
func example() func() error {
    var err error
    return func() error {
        result, err := someOperation() // новая err!
        return err
    }
    // внешняя err так и останется nil
}

// ✅ используем = для изменения внешней
func example() func() error {
    var err error
    return func() error {
        var result string
        result, err = someOperation() // изменяем внешнюю
        return err
    }
}
```

---

## Замыкания и горутины

Замыкания часто используются для передачи данных в горутины:

```go
// Явная передача (работает во всех версиях)
for _, item := range items {
    go func(it string) {
        process(it)
    }(item)
}

// Захват через замыкание (безопасно в Go 1.22+)
for _, item := range items {
    go func() {
        process(item)
    }()
}
```

> 💡 Явная передача через аргументы понятнее и безопаснее во всех версиях Go.

---

## Производительность и Escape Analysis

### Как компилятор решает, где размещать переменные

```go
// Случай 1: переменная убегает в кучу
func f() func() int {
    x := 5
    return func() int { return x }
    // x escapes to heap — должна пережить f()
}

// Случай 2: переменная остаётся на стеке
func g() int {
    x := 5
    return x  // x stays on stack
}
```

Проверка через `go build -gcflags="-m"`:

```
./main.go:3:2: moved to heap: x       // в функции f
./main.go:9:2: g x does not escape    // в функции g
```

### Накладные расходы

Замыкания создают:
1. **Аллокацию структуры окружения** (в куче)
2. **Косвенные вызовы** (через указатель на функцию)
3. **Разыменование указателей** при доступе к захваченным переменным

В высоконагруженном коде предпочитайте явные аргументы:

```go
// ❌ Плохо: аллокация при каждом Handle
func (s *Server) Handle(req *Request) {
    go func() { process(req) }()
}

// ✅ Лучше: worker pool
func (s *Server) Handle(req *Request) {
    s.workerPool <- req
}
```

### Бенчмарк

```go
func BenchmarkClosure(b *testing.B) {
    for i := 0; i < b.N; i++ {
        adder := makeAdder(i)
        _ = adder(10)
    }
}

func BenchmarkDirect(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = directAdd(i, 10)
    }
}
```

Для большинства задач overhead несущественен, но в hot path стоит измерять.

---

## Замыкания vs Структуры с методами

```go
// Через замыкание
func NewCounter() func() int {
    count := 0
    return func() int { count++; return count }
}

// Через структуру
type Counter struct{ count int }
func (c *Counter) Next() int { c.count++; return c.count }
```

| Критерий | Замыкание | Структура |
|----------|-----------|-----------|
| Простота | ✅ Проще | ❌ Больше кода |
| Производительность | ⚠️ Аллокации | ✅ Эффективнее |
| Расширяемость | ❌ Сложно | ✅ Легко |
| Сериализация | ❌ Нельзя | ✅ Можно |
| Отладка | ❌ Анонимные функции | ✅ Видимые методы |
| Интерфейсы | ❌ Один функциональный | ✅ Несколько |
| Потокобезопасность | ⚠️ Нужен mutex внутри | ✅ Можно добавить |

**Правило:** замыкания — для простых случаев (1–3 переменные). Для сложной логики — структуры с методами.

---

## Сравнение с другими языками

### JavaScript
```javascript
function counter() {
    let count = 0;
    return function() { return ++count; };
}
```
Аналогично Go, проблема с циклами была до ES6 (let/const).

### Python
```python
def counter():
    count = 0
    def inner():
        nonlocal count  # нужно явно для изменения
        count += 1
        return count
    return inner
```
В Python нужно `nonlocal` для изменения внешней переменной. В Go — нет.

### Rust
```rust
fn counter() -> impl FnMut() -> i32 {
    let mut count = 0;
    move || { count += 1; count }  // move = захват по значению
}
```
Rust даёт явный контроль: захват по ссылке или по значению. В Go всегда по ссылке.

---

## Лучшие практики

### 1. Документируйте, что захватывает замыкание

```go
// NewHandler создаёт HTTP-обработчик.
// Замыкание захватывает db и config — они должны жить дольше handler.
func NewHandler(db *sql.DB, cfg *Config) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) { ... }
}
```

### 2. Минимизируйте количество захватываемых переменных

```go
// ❌ huge останется в памяти
func process() func() {
    huge := loadHugeData()
    small := extractSmall(huge)
    return func() { fmt.Println(small) }
}

// ✅ захватываем только нужное
func process() func() {
    huge := loadHugeData()
    small := extractSmall(huge)
    huge = nil
    return func() { fmt.Println(small) }
}
```

### 3. Защищайте общие переменные мьютексами

```go
func SharedCounter() func() int {
    count := 0
    mu := sync.Mutex{}
    return func() int {
        mu.Lock()
        defer mu.Unlock()
        count++
        return count
    }
}
```

### 4. Защищайте горутины от panic

```go
go func() {
    defer func() {
        if r := recover(); r != nil {
            log.Printf("goroutine panic: %v", r)
        }
    }()
    // работа
}()
```

### 5. Передавайте данные в горутины явно

```go
for _, item := range items {
    go func(it string) { process(it) }(item)
}
```

---

## FAQ

### Замыкание копирует переменную?
**Нет.** В Go замыкания всегда захватывают переменные **по ссылке**. Если нужно захватить значение — создайте явную копию перед замыканием:
```go
v := getValue()
vCopy := v  // копия
f := func() { use(vCopy) }
```

### Каждый вызов создаёт новое окружение?
**Да.** Каждый вызов функции, возвращающей замыкание, создаёт новую структуру окружения. Поэтому `counter1` и `counter2` независимы.

### Все ли замыкания вызывают Escape Analysis?
**Нет.** Если переменная не переживает функцию (например, передаётся в `sort.Slice` и используется только там), она может остаться на стеке. Escape Analysis срабатывает только когда переменная действительно "убегает".

### Замыкания медленные?
**Обычно нет.** Overhead — одна аллокация в кучу и косвенный вызов. Для большинства задач это несущественно. Проблемы возникают только в hot path с миллионами вызовов в секунду.

### Можно ли сериализовать замыкание?
**Нет.** Замыкание содержит указатели на функции и память, которые нельзя сериализовать. Если нужно сохранять состояние — используйте структуру.

### Что будет, если изменить захваченную переменную извне?
**Замыкание увидит новое значение**, потому что захват по ссылке:
```go
x := 10
f := func() { fmt.Println(x) }
x = 20
f() // выведет 20
```

### Почему `defer` в цикле ведёт себя неожиданно?
Замыкание в `defer` захватывает переменные по ссылке. К моменту выполнения `defer` значение может измениться:
```go
for i := 0; i < 3; i++ {
    defer func() { fmt.Println(i) }() // выведет 3 три раза
}
// Решение: передать как аргумент
for i := 0; i < 3; i++ {
    defer func(v int) { fmt.Println(v) }(i) // 2, 1, 0
}
```

---

## Заключение

Замыкания в Go — мощный, но требующий дисциплины инструмент. Они позволяют писать элегантный код для:

- создания stateful функций и генераторов
- реализации middleware и декораторов
- функциональных опций
- тестирования (моки и spies)
- отложенных и кэшированных вычислений

**Под капотом** замыкание — это структура с указателями на захваченные переменные + функция. Компилятор размещает захваченные переменные в куче (escape to heap).

**Главное правило:**
> Используйте замыкания, когда нужно захватить **1–3 переменные** и логика кода простая.  
> При усложнении — переходите на структуры с методами.

---

## Ссылки

- [Go Specification - Function literals](https://go.dev/ref/spec#Function_literals)
- [Go Memory Model](https://go.dev/ref/mem)
- [Effective Go - Functions](https://go.dev/doc/effective_go#functions)
- [Go Blog - Functions, closures, generics](https://go.dev/blog/intro-generics)
- [Go 1.22 Release Notes - Loop variable scoping](https://go.dev/doc/go1.22#language)
- [Go Escape Analysis](https://tip.golang.org/doc/gc-guide#Escape_analysis)
