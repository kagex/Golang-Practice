# Тестирование в Go

## Введение

Тестирование — важная часть разработки программного обеспечения. Go встроенными средствами поддерживает модульное тестирование через пакет `testing` и команду `go test`.

## Основы тестирования

### Структура тестового файла

- Файлы тестов имеют суффикс `_test.go`
- Пакет теста совпадает с пакетом тестируемого кода
- Импортируется пакет `testing`
- Используются функции `TestXxx(*testing.T)`

```go
package main

import "testing"

func TestAdd(t *testing.T) {
    result := add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %d", result)
    }
}
```

### Именование тестов

- Функции тестов начинаются с `Test` и следующей заглавной буквы
- Используется PascalCase: `TestFunctionName`, `TestTypeName`
- Можно добавлять детализацию через `t.Run()`

## Основные методы `testing.T`

| Метод | Описание |
|-------|----------|
| `Error()` | Записывает ошибку, тест продолжается |
| `Errorf()` | Форматированная ошибка |
| `Fail()` | Помечает тест как проваленный |
| `FailNow()` | Немедленно останавливает тест |
| `Fatal()` | Ошибка + `FailNow()` |
| `Fatalf()` | Форматированная `Fatal()` |
| `Log()` | Записывает сообщение в лог |
| `Logf()` | Форматированное сообщение |
| `Skip()` | Пропускает тест |
| `Skipf()` | Форматированный `Skip()` |
| `Helper()` | Помечает функцию как вспомогательную |

## Виды тестов

### Unit-тесты

Тестируют отдельные функции и методы.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -2, -3, -5},
        {"mixed", -2, 3, 1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("add(%d, %d): expected %d, got %d", tt.a, tt.b, tt.expected, result)
            }
        })
    }
}
```

### Интеграционные тесты

Тестируют взаимодействие компонентов.

```go
func TestDatabaseConnection(t *testing.T) {
    db, err := connectToDB()
    if err != nil {
        t.Fatalf("failed to connect: %v", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        t.Errorf("database ping failed: %v", err)
    }
}
```

### Функциональные тесты (black-box)

Тестируют поведение программы как единого целого.

```go
func runProgram(input string) (string, error) {
    cmd := exec.Command("go", "run", "main.go")
    cmd.Stdin = bytes.NewBufferString(input)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        return "", err
    }

    output := stdout.String()
    if stderr.Len() > 0 {
        return "", errors.New(stderr.String())
    }

    return strings.TrimSpace(output), nil
}

func TestProgram(t *testing.T) {
    output, err := runProgram("6\n")
    if err != nil {
        t.Fatalf("Error running program: %v", err)
    }

    expected := "Сейчас 6ч. - утро"
    if !strings.Contains(output, expected) {
        t.Errorf("Expected to contain %q, got %q", expected, output)
    }
}
```

## Параметризованные тесты

### Subtests (Go 1.7+)

```go
func TestTimeOfDay(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {"morning", "6\n", "утро"},
        {"day", "12\n", "день"},
        {"evening", "18\n", "вечер"},
        {"night", "23\n", "ночь"},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            output, err := runProgram(tt.input)
            if err != nil {
                t.Fatalf("Error: %v", err)
            }
            if !strings.Contains(output, tt.expected) {
                t.Errorf("Expected %q in %q", tt.expected, output)
            }
        })
    }
}
```

## Тесты производительности (benchmarks)

Используются для измерения скорости выполнения кода.

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(1, 2)
    }
}
```

Запуск: `go test -bench=.`
- `-bench=.`
- `-benchmem` — показать использование памяти

## Фазз-тестирование (fuzzing)

Начиная с Go 1.18, фаззинг встроен прямо в стандартную библиотеку через `testing.F`. Это мощный инструмент для поиска скрытых багов и неучтённых граничных случаев, при котором язык сам генерирует случайные входные данные.

### Базовый пример

```go
func FuzzAdd(f *testing.F) {
    f.Add(1, 2)
    f.Add(-5, 10)

    f.Fuzz(func(t *testing.T, a, b int) {
        result := add(a, b)
        if result != a+b {
            t.Errorf("add(%d, %d) = %d, expected %d", a, b, result, a+b)
        }
    })
}
```

### Фаззинг с бинарными данными

```go
func FuzzParse(f *testing.F) {
    f.Add([]byte("hello"))
    f.Add([]byte("world"))

    f.Fuzz(func(t *testing.T, data []byte) {
        result := parse(data)
        if result == nil {
            t.Fatal("parse returned nil for valid input")
        }
    })
}
```

### Важные правила фаззинга
- Имя функции должно начинаться с `Fuzz`, а не `Test`
- Параметр должен быть `*testing.F`, а не `*testing.T`
- Используйте `f.Add()` для добавления начальных примеров (seed corpus)
- Фаззер автоматически расширяет corpus новыми входами
- Тест работает бесконечно, пока не найдёт панику или ошибку
- Запуск: `go test -fuzz=FuzzAdd -fuzztime=10s`

## Библиотека Testify

Хотя философия Go поощряет использование только стандартных инструментов, на практике подавляющее большинство проектов использует пакет `github.com/stretchr/testify`. Он избавляет от необходимости постоянно писать громоздкие проверки и делает тестовый код намного чище.

### Установка

```bash
go get github.com/stretchr/testify
```

### Основные пакеты

| Пакет | Описание |
|-------|----------|
| `assert` | Проверки, которые не останавливают тест |
| `require` | Проверки, которые останавливают тест при ошибке |
| `mock` | Создание моков для интерфейсов |
| `suite` | Организация групп тестов |

### Примеры использования

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
    result := add(2, 3)

    // assert — не останавливает тест
    assert.Equal(t, 5, result, "2 + 3 should equal 5")
    assert.Greater(t, result, 0, "result should be positive")

    // require — останавливает тест при ошибке
    user, err := getUser(1)
    require.NoError(t, err, "should not fail to get user")
    assert.Equal(t, "John", user.Name)
}
```

### Распространённые методы `assert`

```go
assert.Equal(t, expected, actual)           // Равенство
assert.NotEqual(t, a, b)                     // Неравенство
assert.True(t, condition)                    // Булево
assert.Nil(t, value)                         // Нулевое значение
assert.NotNil(t, value)                      // Ненулевое значение
assert.Contains(t, slice, item)             // Содержит элемент
assert.NoError(t, err)                       // Нет ошибки
assert.Panics(t, func() { ... })            // Паникует
assert.WithinDuration(t, time1, time2, d)   // Время с допуском
```

## Управление жизненным циклом тестов

### `TestMain`

Позволяет выполнить код до/после всех тестов.

```go
func TestMain(m *testing.M) {
    // подготовка
    setupDatabase()

    code := m.Run()

    // cleanup
    teardownDatabase()

    os.Exit(code)
}
```

### `t.Parallel()`

Позволяют выполнять тесты параллельно.

```go
func TestA(t *testing.T) {
    t.Parallel()
    // тест A
}

func TestB(t *testing.T) {
    t.Parallel()
    // тест B
}
```

### `t.Helper()`

Помечает функцию как вспомогательную — при ошибке стек будет указывать на место вызова, а не на строку внутри функции.

```go
func assertEqual(t *testing.T, expected, actual int) {
    t.Helper()
    if expected != actual {
        t.Errorf("expected %d, got %d", expected, actual)
    }
}
```

### `t.TempDir()`

Создаёт временную директорию, которая автоматически удаляется после завершения теста.

```go
func TestFileProcessing(t *testing.T) {
    dir := t.TempDir()
    path := filepath.Join(dir, "test.txt")
    os.WriteFile(path, []byte("hello"), 0644)
    // не нужно defer os.RemoveAll(dir)
}
```

### `t.Setenv()`

Устанавливает переменную окружения, которая автоматически восстанавливается после теста.

```go
func TestWithEnv(t *testing.T) {
    t.Setenv("DATABASE_URL", "postgres://test")
    // после теста старое значение будет восстановлено
}
```

### `testing.Short()`

Позволяет пропускать долгие тесты при запуске `go test -short`.

```go
func TestHeavyComputation(t *testing.T) {
    if testing.Short() {
        t.Skip("пропуск в коротком режиме")
    }
    // долгие вычисления...
}
```

### `t.Cleanup()`

Регистрирует функцию, которая будет вызвана при завершении теста.

```go
func TestFile(t *testing.T) {
    tmpFile, err := os.CreateTemp("", "test")
    if err != nil {
        t.Fatal(err)
    }

    t.Cleanup(func() {
        tmpFile.Close()
        os.Remove(tmpFile.Name())
    })

    // тест
}
```

## Покрытие кода тестами

```bash
go test -cover
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Вспомогательные функции

### `testing.TB` — интерфейс

Общий интерфейс для `*testing.T` и `*testing.B`.

## Лучшие практики

1. **Тесты должны быть независимыми** — порядок выполнения не должен влиять на результат
2. **Используйте `t.Run()` для группировки** — улучшает читаемость отчётов
3. **Не дублируйте код** — выносите общую логику в вспомогательные функции
4. **Проверяйте ошибки** — используйте `t.Fatalf()` для критических ошибок
5. **Тестируйте граничные случаи** — пограничные значения, пустые входы, ошибочные данные
6. **Используйте таблицы тестов** — для параметризованных тестов
7. **Используйте `t.Helper()`** — для точных сообщений об ошибках
8. **Используйте `t.TempDir()`** — для временных файлов и папок
9. **Используйте `t.Setenv()`** — для переменных окружения

## Частые ошибки

1. **Игнорирование ошибок** в тестах
2. **Зависимость от порядка выполнения**
3. **Отсутствие тестов на ошибочные входы**
4. **Тесты с "магическими числами" без пояснений**
5. **Слишком сложные тесты** — тестируют не одну функцию, а многое сразу

## Структура тестового файла для time-of-day

```go
package main

import (
    "bytes"
    "errors"
    "os/exec"
    "strings"
    "testing"
)

func runProgram(input string) (string, error) {
    cmd := exec.Command("go", "run", "main.go")
    cmd.Stdin = bytes.NewBufferString(input)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        return "", err
    }

    output := stdout.String()
    if stderr.Len() > 0 {
        return "", errors.New(stderr.String())
    }

    return strings.TrimSpace(output), nil
}

type testCase struct {
    name     string
    input    string
    expected string
}

func TestTimeOfDay(t *testing.T) {
    tests := []testCase{
        {"утро 6ч", "6\n", "Сейчас 6ч. - утро."},
        {"утро 7ч", "7\n", "Сейчас 7ч. - утро."},
        {"утро 11ч", "11\n", "Сейчас 11ч. - утро."},
        {"день 12ч", "12\n", "Сейчас 12ч. - день."},
        {"день 13ч", "13\n", "Сейчас 13ч. - день."},
        {"день 17ч", "17\n", "Сейчас 17ч. - день."},
        {"вечер 18ч", "18\n", "Сейчас 18ч. - вечер."},
        {"вечер 21ч", "21\n", "Сейчас 21ч. - вечер."},
        {"вечер 22ч", "22\n", "Сейчас 22ч. - вечер."},
        {"ночь 23ч", "23\n", "Сейчас 23ч. - ночь."},
        {"ночь 0ч", "0\n", "Сейчас 0ч. - ночь."},
        {"ночь 5ч", "5\n", "Сейчас 5ч. - ночь."},
        {"отрицательное число", "-5\n", "Неверно задано время"},
        {"больше 24", "25\n", "Неверно задано время"},
        {"текст", "abc\n", "Неверно задано время"},
        {"пустая строка", "\n", "Неверно задано время"},
        {"граничное 5ч", "5\n", "Сейчас 5ч. - ночь."},
        {"граничное 11ч", "11\n", "Сейчас 11ч. - утро."},
        {"граничное 17ч", "17\n", "Сейчас 17ч. - день."},
        {"граничное 22ч", "22\n", "Сейчас 22ч. - вечер."},
        {"граничное 24ч", "24\n", "Сейчас 24ч. - ночь."},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            output, err := runProgram(tt.input)
            if err != nil {
                t.Fatalf("Error running program: %v", err)
            }
            if !strings.Contains(output, tt.expected) {
                t.Errorf("Expected to contain %q, got %q", tt.expected, output)
            }
        })
    }
}
```

## Оптимизация функциональных тестов

В примере выше `runProgram` использует `exec.Command("go", "run", "main.go")`. Для обучающих примеров это нормально, но в коммерческой разработке такой подход неэффективен — Go заново компилирует программу на каждый тестовый кейс.

### Проблема

- Каждый вызов `runProgram` компилирует `main.go` заново
- При 20+ тестовых кейсах — 20+ компиляций
- Это замедляет тесты в разы

### Решение: однократная компиляция

```go
var testBinary string

func TestMain(m *testing.M) {
    // Компилируем бинарник один раз
    tmpFile, err := os.CreateTemp("", "test-binary-*.exe")
    if err != nil {
        log.Fatal(err)
    }
    defer os.Remove(tmpFile.Name())

    buildCmd := exec.Command("go", "build", "-o", tmpFile.Name(), "main.go")
    if output, err := buildCmd.CombinedOutput(); err != nil {
        log.Fatalf("Build failed: %v, output: %s", err, output)
    }
    testBinary = tmpFile.Name()

    code := m.Run()

    // Очистка
    if testBinary != "" {
        os.Remove(testBinary)
    }

    os.Exit(code)
}

func runTestProgram(input string) (string, error) {
    cmd := exec.Command(testBinary)
    cmd.Stdin = bytes.NewBufferString(input)
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        return "", err
    }

    output := stdout.String()
    if stderr.Len() > 0 {
        return "", errors.New(stderr.String())
    }

    return strings.TrimSpace(output), nil
}
```

Теперь компиляция происходит **один раз** в `TestMain`, а все тесты используют готовый бинарник.

## Заключение

Go предоставляет мощные встроенные средства для тестирования:
- `testing` — встроенная библиотека с полным функционалом
- Фаззинг (`testing.F`) — автоматический поиск багов через случайные данные
- Таблицы тестов и subtests — читаемая параметризация
- `t.Cleanup()` и `TestMain` — управление ресурсами

Для коммерческих проектов часто добавляют `testify` для более чистого кода проверок, а функциональные тесты оптимизируют через однократную компиляцию бинарника.

Правильное тестирование делает код надёжным, поддерживаемым и готовым к масштабированию.
