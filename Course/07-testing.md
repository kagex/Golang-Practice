# Модуль 7. Тестирование 🟢

**Срок:** 2–3 недели  
**Цель модуля:** писать надёжные тесты — unit, integration, benchmark, fuzz. Без этого нельзя выкатываться в production.

---

## Урок 7.1. Базовый testing

**📚 Теория:** `func TestXxx(t *testing.T)`, `t.Run` (subtests), table-driven tests, `t.Parallel`, `t.Helper`, `t.Cleanup`, golden files.

**💻 Практика:** покрой table-driven тестами функцию парсинга IP-адресов.

---

## Урок 7.2. testify

**📚 Теория:** `assert` (продолжает тест) vs `require` (останавливает), `suite.Suite` для setup/teardown.

**💻 Практика:** перепиши тесты из 7.1 с testify.

---

## Урок 7.3. Моки

**📚 Теория:** `gomock` (официальный), `mockery` (популярный, работает с testify), ручные стабы для простых интерфейсов.

**💻 Практика:** замокай `TaskStore` в тестах HTTP-хендлеров.

---

## Урок 7.4. Интеграционные тесты

**📚 Теория:** `testcontainers-go` для Postgres/Redis в Docker, build tags (`//go:build integration`), отдельная команда запуска.

**💻 Практика:** тест на `UserRepo`, который поднимает настоящий Postgres в контейнере.

---

## Урок 7.5. HTTP-тесты

**📚 Теория:** `httptest.NewServer`, `httptest.NewRecorder`, тест хендлера без поднятия реального сервера.

**💻 Практика:** тесты на все эндпоинты мини-соцсети.

---

## Урок 7.6. Бенчмарки и pprof

**📚 Теория:** `func BenchmarkXxx(b *testing.B)`, `b.ResetTimer`, `b.ReportAllocs`, `go test -bench=. -benchmem`, `benchstat` для сравнения версий.

**💻 Практика:** бенчмарк 3 реализаций конкатенации строк: `+`, `strings.Builder`, `bytes.Buffer`.

---

## Урок 7.7. Fuzzing (Go 1.18+)

**📚 Теория:** `func FuzzXxx(f *testing.F)`, `f.Add` (seed corpus), `f.Fuzz`, `go test -fuzz=Fuzz`, хранение падающих входов в `testdata/fuzz/`.

**💻 Практика:** fuzz-тест для парсера URL.

---

## Урок 7.8. Покрытие кода

**📚 Теория:** `go test -cover`, `go test -coverprofile=c.out`, HTML-отчёт `go tool cover -html=c.out`, line vs statement vs branch coverage.

**✅ Чек:** 70%+ покрытие на своём проекте.

---

## 🎯 Практика модуля

Покрой проект из Модуля 5 (Библиотека книг):
- unit-тесты 70%+
- - минимум 5 интеграционных тестов
  - - 2–3 бенчмарка для критичных путей
    - - 1 fuzz-тест
     
      - ---

      ## 🏁 Чек-пойнт модуля

      - [ ] Пишешь table-driven тесты
      - [ ] - [ ] Используешь моки без фанатизма
      - [ ] - [ ] Пишешь интеграционные тесты с testcontainers
      - [ ] - [ ] Понимаешь benchstat и интерпретируешь результаты
      - [ ] - [ ] Добился 70%+ покрытия
     
      - [ ] **Предыдущий:** [Модуль 6 ←](06-web-api.md) · **Следующий:** [Модуль 8. Микросервисы и сети →](08-microservices.md)
      - [ ] 
