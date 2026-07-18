# Конкурентность в Go — примеры кода от нуля до продвинутого

> Дополнение к [Модулю 3: Конкурентность](./03-concurrency.md)

Этот файл содержит полный спектр примеров — от самого простого запуска горутины до продвинутых паттернов, которые используются в production. Каждый пример можно скопировать и запустить локально.

---

## Содержание

1. [Уровень 0 — Первые шаги](#уровень-0--первые-шаги)
2. [Уровень 1 — Каналы](#уровень-1--каналы)
3. [Уровень 2 — sync-примитивы](#уровень-2--sync-примитивы)
4. [Уровень 3 — context.Context](#уровень-3--contextcontext)
5. [Уровень 4 — Паттерны конкурентности](#уровень-4--паттерны-конкурентности)
6. [Уровень 5 — Продвинутые техники](#уровень-5--продвинутые-техники)
7. [Уровень 6 — Production-паттерны](#уровень-6--production-паттерны)

---

## Уровень 0 — Первые шаги

### Пример 0.1 — Запуск горутины

Самый простой способ запустить код параллельно:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	// Запускаем горутину — лёгкий поток Go-рантайма
	go func() {
		fmt.Println("Привет из горутины!")
	}()

	// Без этого main завершится раньше, чем горутина успеет напечатать
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Привет из main!")
}
```

> ⚠️ `time.Sleep` — плохая синхронизация. Дальше увидим правильный способ.

---

### Пример 0.2 — Несколько горутин

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // сообщаем WaitGroup, что появилась ещё одна горутина
		go func(n int) {
			defer wg.Done() // сигнализируем о завершении
			fmt.Printf("Горутина %d работает\n", n)
		}(i)
	}

	wg.Wait() // ждём пока все горутины вызовут Done()
	fmt.Println("Все горутины завершились")
}
```

---

## Уровень 1 — Каналы

Каналы — основной способ общения между горутинами. «Не общайся через общую память, делись памятью через общение».

### Пример 1.1 — Небуферизованный канал

Отправитель блокируется до тех пор, пока получатель не примет значение:

```go
package main

import "fmt"

func sum(numbers []int, result chan<- int) {
	total := 0
	for _, n := range numbers {
		total += n
	}
	result <- total // отправка блокирует до момента получения
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ch := make(chan int) // небуферизованный канал

	// Считаем сумму первой и второй половины параллельно
	go sum(numbers[:5], ch)
	go sum(numbers[5:], ch)

	// Получаем два результата
	a, b := <-ch, <-ch
	fmt.Println("Сумма:", a+b) // 55
}
```

---

### Пример 1.2 — Буферизованный канал

Отправитель не блокируется, пока в буфере есть место:

```go
package main

import "fmt"

func main() {
	// Канал с буфером на 3 элемента
	ch := make(chan string, 3)

	ch <- "первый"  // не блокирует
	ch <- "второй"  // не блокирует
	ch <- "третий"  // не блокирует
	// ch <- "четвёртый" // заблокировало бы — буфер полон

	close(ch) // закрываем, чтобы range завершился

	for msg := range ch {
		fmt.Println(msg)
	}
}
```

---

### Пример 1.3 — select: мультиплексирование каналов

```go
package main

import (
	"fmt"
	"time"
)

func fastWorker(ch chan<- string) {
	time.Sleep(100 * time.Millisecond)
	ch <- "быстрый результат"
}

func slowWorker(ch chan<- string) {
	time.Sleep(500 * time.Millisecond)
	ch <- "медленный результат"
}

func main() {
	fast := make(chan string, 1)
	slow := make(chan string, 1)

	go fastWorker(fast)
	go slowWorker(slow)

	// Ждём любого из двух или таймаут
	select {
	case result := <-fast:
		fmt.Println("Первым ответил:", result)
	case result := <-slow:
		fmt.Println("Первым ответил:", result)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Таймаут!")
	}
}
```

---

### Пример 1.4 — Генератор (producer) и потребитель (consumer)

```go
package main

import "fmt"

// generate возвращает канал, в который пишет числа в отдельной горутине
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

// square читает из in и пишет квадраты в новый канал
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			out <- n * n
		}
	}()
	return out
}

func main() {
	// Строим конвейер: generate → square → print
	gen := generate(2, 3, 4, 5)
	sq := square(gen)

	for result := range sq {
		fmt.Println(result) // 4, 9, 16, 25
	}
}
```

---

## Уровень 2 — sync-примитивы

### Пример 2.1 — Мьютекс: защита общего состояния

```go
package main

import (
	"fmt"
	"sync"
)

// SafeCounter — потокобезопасный счётчик
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	counter := &SafeCounter{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Итог:", counter.Value()) // всегда 1000
}
```

---

### Пример 2.2 — RWMutex: много читателей, один писатель

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Cache — кэш с блокировкой чтения/записи
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock() // эксклюзивная блокировка для записи
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // разделяемая блокировка для чтения
	defer c.mu.RUnlock()
	v, ok := c.data[key]
	return v, ok
}

func main() {
	cache := NewCache()
	cache.Set("user:1", "Alice")

	var wg sync.WaitGroup

	// Запускаем 10 читателей параллельно
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(10 * time.Millisecond)
			if v, ok := cache.Get("user:1"); ok {
				fmt.Printf("Читатель %d прочитал: %s\n", id, v)
			}
		}(i)
	}

	wg.Wait()
}
```

---

### Пример 2.3 — sync.Once: инициализация один раз

```go
package main

import (
	"fmt"
	"sync"
)

type Database struct {
	connection string
}

var (
	db   *Database
	once sync.Once
)

func GetDB() *Database {
	once.Do(func() {
		fmt.Println("Инициализируем соединение с БД...")
		db = &Database{connection: "postgres://localhost:5432/mydb"}
	})
	return db
}

func main() {
	var wg sync.WaitGroup

	// 10 горутин пытаются получить соединение
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			d := GetDB()
			fmt.Printf("Горутина %d использует: %s\n", id, d.connection)
		}(i)
	}

	wg.Wait()
	// "Инициализируем..." напечатается ровно один раз
}
```

---

### Пример 2.4 — sync/atomic: без блокировок

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter atomic.Int64 // Go 1.19+

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println("Atomic счётчик:", counter.Load()) // всегда 1000
}
```

> **Когда atomic, когда mutex?** Atomic — для простых чисел и указателей. Mutex — когда нужно защитить составную структуру или выполнить несколько операций атомарно.

---

## Уровень 3 — context.Context

### Пример 3.1 — Отмена горутины через context

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d остановлен: %v\n", id, ctx.Err())
			return
		default:
			fmt.Printf("Воркер %d работает...\n", id)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(600 * time.Millisecond)
	fmt.Println("Отменяем все горутины...")
	cancel() // сигнализируем всем горутинам об остановке

	time.Sleep(100 * time.Millisecond) // даём время завершиться
}
```

---

### Пример 3.2 — Таймаут через context

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context, url string) (string, error) {
	// Имитируем медленный запрос
	select {
	case <-time.After(2 * time.Second):
		return "данные от " + url, nil
	case <-ctx.Done():
		return "", fmt.Errorf("запрос отменён: %w", ctx.Err())
	}
}

func main() {
	// Даём операции не более 500ms
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	result, err := fetchData(ctx, "https://example.com")
	if err != nil {
		fmt.Println("Ошибка:", err) // context deadline exceeded
		return
	}
	fmt.Println("Результат:", result)
}
```

---

### Пример 3.3 — Передача значений через context (request-scoped)

```go
package main

import (
	"context"
	"fmt"
)

// Используем typedKey, чтобы избежать коллизий ключей
type contextKey string

const requestIDKey contextKey = "requestID"

func middleware(ctx context.Context) context.Context {
	// Добавляем request ID в контекст (только request-scoped данные!)
	return context.WithValue(ctx, requestIDKey, "req-abc-123")
}

func handler(ctx context.Context) {
	// Извлекаем только нужный тип
	if reqID, ok := ctx.Value(requestIDKey).(string); ok {
		fmt.Println("Обрабатываем запрос:", reqID)
	}
}

func main() {
	ctx := context.Background()
	ctx = middleware(ctx)
	handler(ctx)
}
```

> ⚠️ `context.WithValue` только для request-scoped данных (request ID, trace ID, auth token). Никогда — для бизнес-логики.

---

## Уровень 4 — Паттерны конкурентности

### Пример 4.1 — Worker Pool (пул воркеров)

Классический паттерн: фиксированное число воркеров обрабатывают задачи из очереди:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Имитируем работу
		time.Sleep(10 * time.Millisecond)
		result := Result{
			JobID:  job.ID,
			Output: fmt.Sprintf("воркер-%d обработал [%s]", id, job.Data),
		}
		results <- result
	}
}

func main() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Запускаем воркеры
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Отправляем задачи
	for i := 1; i <= numJobs; i++ {
		jobs <- Job{ID: i, Data: fmt.Sprintf("задача-%d", i)}
	}
	close(jobs) // сигнал воркерам: новых задач не будет

	// Закрываем results когда все воркеры завершились
	go func() {
		wg.Wait()
		close(results)
	}()

	// Собираем результаты
	for result := range results {
		fmt.Printf("Задача %d: %s\n", result.JobID, result.Output)
	}
}
```

---

### Пример 4.2 — Fan-Out / Fan-In

Fan-Out: одна горутина раздаёт работу многим. Fan-In: собираем результаты от многих в один канал:

```go
package main

import (
	"fmt"
	"sync"
)

// fanOut распределяет задачи по N горутинам
func fanOut(input <-chan int, n int) []<-chan int {
	outputs := make([]<-chan int, n)
	for i := 0; i < n; i++ {
		out := make(chan int)
		outputs[i] = out
		go func(ch chan<- int) {
			defer close(ch)
			for v := range input {
				ch <- v * v // каждая горутина возводит в квадрат
			}
		}(out)
	}
	return outputs
}

// fanIn сливает несколько каналов в один
func fanIn(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	output := func(ch <-chan int) {
		defer wg.Done()
		for v := range ch {
			merged <- v
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	input := make(chan int)
	go func() {
		defer close(input)
		for i := 1; i <= 8; i++ {
			input <- i
		}
	}()

	// Раздаём 3 воркерам и собираем обратно
	workers := fanOut(input, 3)
	results := fanIn(workers...)

	for v := range results {
		fmt.Println(v)
	}
}
```

---

### Пример 4.3 — Pipeline (конвейер)

Каждый этап получает данные из предыдущего и передаёт дальше:

```go
package main

import (
	"fmt"
	"strings"
)

func generate(words ...string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for _, w := range words {
			out <- w
		}
	}()
	return out
}

func toUpper(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- strings.ToUpper(s)
		}
	}()
	return out
}

func addExclamation(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for s := range in {
			out <- s + "!"
		}
	}()
	return out
}

func main() {
	// Конвейер: generate → toUpper → addExclamation
	words := generate("hello", "world", "golang")
	upper := toUpper(words)
	excited := addExclamation(upper)

	for result := range excited {
		fmt.Println(result) // HELLO!, WORLD!, GOLANG!
	}
}
```

---

### Пример 4.4 — Semaphore через буферизованный канал

Ограничиваем количество параллельных операций:

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const maxConcurrent = 3 // не более 3 одновременных операций
	const totalJobs = 10

	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup

	for i := 1; i <= totalJobs; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{} // захватываем семафор (блокирует если занято)
			defer func() { <-sem }() // освобождаем при выходе

			fmt.Printf("Задача %d начата (активных: %d)\n", id, len(sem))
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Задача %d завершена\n", id)
		}(i)
	}

	wg.Wait()
}
```

---

## Уровень 5 — Продвинутые техники

### Пример 5.1 — errgroup: управление группой горутин с ошибками

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func checkURL(ctx context.Context, url string) error {
	client := &http.Client{Timeout: 2 * time.Second}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("создание запроса для %s: %w", url, err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("запрос к %s: %w", url, err)
	}
	defer resp.Body.Close()
	fmt.Printf("%s — статус: %d\n", url, resp.StatusCode)
	return nil
}

func main() {
	urls := []string{
		"https://go.dev",
		"https://pkg.go.dev",
		"https://golang.org",
	}

	// errgroup автоматически отменяет контекст при первой ошибке
	g, ctx := errgroup.WithContext(context.Background())

	for _, url := range urls {
		url := url // захватываем переменную (Go < 1.22)
		g.Go(func() error {
			return checkURL(ctx, url)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Все URL доступны!")
}
```

---

### Пример 5.2 — singleflight: защита от thundering herd

Если несколько горутин запрашивают одно и то же — выполняем запрос только один раз:

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

type CacheWithSingleflight struct {
	sf    singleflight.Group
	mu    sync.RWMutex
	cache map[string]string
}

func (c *CacheWithSingleflight) Get(key string) (string, error) {
	c.mu.RLock()
	if v, ok := c.cache[key]; ok {
		c.mu.RUnlock()
		return v, nil
	}
	c.mu.RUnlock()

	// singleflight: только одна горутина сходит в БД
	v, err, shared := c.sf.Do(key, func() (interface{}, error) {
		fmt.Printf("Идём в БД за ключом '%s'...\n", key)
		time.Sleep(100 * time.Millisecond)
		return "значение_" + key, nil
	})

	if err != nil {
		return "", err
	}

	fmt.Printf("Результат для '%s' был общим: %v\n", key, shared)

	c.mu.Lock()
	if c.cache == nil {
		c.cache = make(map[string]string)
	}
	c.cache[key] = v.(string)
	c.mu.Unlock()

	return v.(string), nil
}

func main() {
	cache := &CacheWithSingleflight{}
	var wg sync.WaitGroup

	// 10 горутин одновременно запрашивают один ключ
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			v, _ := cache.Get("user:1")
			fmt.Printf("Горутина %d получила: %s\n", id, v)
		}(i)
	}

	wg.Wait()
	// "Идём в БД..." напечатается только один раз
}
```

---

### Пример 5.3 — Rate Limiter на каналах

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		requests <- i
	}
	close(requests)

	// Тикер — сигнализирует каждые 200ms (5 req/s)
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for req := range requests {
		<-ticker.C
		fmt.Printf("Обрабатываем запрос %d в %s\n", req, time.Now().Format("15:04:05.000"))
	}
}
```

---

## Уровень 6 — Production-паттерны

### Пример 6.1 — Graceful Shutdown HTTP-сервера

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond) // имитация работы
		fmt.Fprintln(w, "OK")
	})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Запускаем сервер в горутине
	go func() {
		fmt.Println("Сервер запущен на :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Ошибка сервера: %v\n", err)
			os.Exit(1)
		}
	}()

	// Ждём сигнала остановки
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Получен сигнал остановки, завершаем gracefully...")

	// Даём 30 секунд на завершение активных запросов
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Принудительная остановка: %v\n", err)
	}

	fmt.Println("Сервер остановлен")
}
```

---

### Пример 6.2 — Полный Worker Pool с контекстом, graceful shutdown и метриками

```go
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

type Task struct {
	ID      int
	Payload string
}

type WorkerPool struct {
	numWorkers int
	tasks      chan Task
	processed  atomic.Int64
	wg         sync.WaitGroup
}

func NewWorkerPool(workers int, queueSize int) *WorkerPool {
	return &WorkerPool{
		numWorkers: workers,
		tasks:      make(chan Task, queueSize),
	}
}

func (p *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < p.numWorkers; i++ {
		p.wg.Add(1)
		go p.worker(ctx, i)
	}
}

func (p *WorkerPool) worker(ctx context.Context, id int) {
	defer p.wg.Done()
	fmt.Printf("[воркер %d] запущен\n", id)

	for {
		select {
		case task, ok := <-p.tasks:
			if !ok {
				fmt.Printf("[воркер %d] канал задач закрыт, завершаю\n", id)
				return
			}
			p.processTask(ctx, id, task)
		case <-ctx.Done():
			fmt.Printf("[воркер %d] получен сигнал отмены: %v\n", id, ctx.Err())
			return
		}
	}
}

func (p *WorkerPool) processTask(ctx context.Context, workerID int, task Task) {
	select {
	case <-ctx.Done():
		fmt.Printf("[воркер %d] пропускаем задачу %d — контекст отменён\n", workerID, task.ID)
		return
	default:
	}

	time.Sleep(50 * time.Millisecond)
	p.processed.Add(1)
	fmt.Printf("[воркер %d] задача %d выполнена: %s\n", workerID, task.ID, task.Payload)
}

func (p *WorkerPool) Submit(task Task) bool {
	select {
	case p.tasks <- task:
		return true
	default:
		return false // очередь заполнена
	}
}

func (p *WorkerPool) Shutdown() {
	close(p.tasks)
	p.wg.Wait()
	fmt.Printf("\nВсего обработано задач: %d\n", p.processed.Load())
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool := NewWorkerPool(5, 100)
	pool.Start(ctx)

	for i := 1; i <= 20; i++ {
		pool.Submit(Task{
			ID:      i,
			Payload: fmt.Sprintf("обработать данные #%d", i),
		})
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Для примера — останавливаем через 500ms
	go func() {
		time.Sleep(500 * time.Millisecond)
		quit <- syscall.SIGTERM
	}()

	<-quit
	fmt.Println("\nОстанавливаем пул воркеров...")
	cancel()
	pool.Shutdown()
}
```

---

### Пример 6.3 — Обнаружение Data Race

Запустите с флагом `go run -race main.go`:

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// ПЛОХО: гонка данных — запустите с -race чтобы увидеть
func badCounter() int {
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // DATA RACE: несинхронизированный доступ
		}()
	}

	wg.Wait()
	return counter
}

// ХОРОШО: через atomic
func goodCounterAtomic() int64 {
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}

	wg.Wait()
	return counter.Load()
}

// ХОРОШО: через канал
func goodCounterChannel() int {
	ch := make(chan struct{}, 1000)
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- struct{}{}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	total := 0
	for range ch {
		total++
	}
	return total
}

func main() {
	fmt.Println("Atomic:", goodCounterAtomic())   // всегда 1000
	fmt.Println("Channel:", goodCounterChannel()) // всегда 1000
	// fmt.Println("Bad:", badCounter()) // раскомментируй и запусти с -race
}
```

---

## Шпаргалка: когда что использовать

| Ситуация | Инструмент |
|---|---|
| Запустить код параллельно | `go func() + sync.WaitGroup` |
| Передать результат между горутинами | канал |
| Несколько горутин, один получатель | fan-in |
| Один источник, несколько получателей | fan-out / worker pool |
| Цепочка преобразований | pipeline |
| Защита общего состояния | `sync.Mutex` или `sync/atomic` |
| Инициализация один раз | `sync.Once` |
| Много читателей, мало писателей | `sync.RWMutex` |
| Ограничить параллелизм | семафор через `chan struct{}` |
| Отмена по сигналу | `context.WithCancel` |
| Таймаут операции | `context.WithTimeout` |
| Группа горутин с обработкой ошибок | `errgroup` |
| Дедупликация одинаковых запросов | `singleflight` |
| Rate limiting | `time.Ticker` + канал |
| Корректная остановка сервиса | `signal.Notify` + `Shutdown` |

---

## Что изучить дальше

- [Модуль 3: Конкурентность — теория и уроки](./03-concurrency.md)
- [Модуль 11: Performance — pprof, escape analysis, оптимизации](./11-performance.md)
- Запустите все примеры с `go run -race` — убедитесь, что гонок нет
- Изучите исходники `sync` и `context` в стандартной библиотеке Go

