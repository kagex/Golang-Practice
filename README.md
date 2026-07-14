# Golang - Practice
Репозиторий содержит мою практику на языке Go.
Так же репозиторий содержит решения упражнений с платформы [Exercism](https://exercism.org/tracks/go).

## О репозитории

Здесь собраны мои решения задач из официального трека **Go** на Exercism и самостоятельная практика.  
Репозиторий используется для практики, закрепления материала и ведения портфолио решений.

## Структура репозитория

```
.
├── solutions/go/               # Основные решения упражнений Exercism в треке Go
│   ├── blackjack/
│       ├── 1 
│           ├── blackjack.go    # Принятое решение задачи blackjack
│       ├── blackjack_test.go   # Тесты для проверки решения
│       ├── blackjack.go        # Текущий код задачи (в основном будет дублировать решение)
│       ├── HELP.md             # Информация об взаимодействии с Exercism
│       ├── HINTS.md            # Подсказки к задаче
│       ├── README.md           # Описание задачи
│   │ 
│   ├── lasagna/
│   └── ... (и другие упражнения)
│
├── Practice/                  # Дополнительные самостоятельные практики
│
├── .github/workflows/         # GitHub Actions workflows
├── .gitignore
└── README.md
```

## Как запустить решения

### Запуск проверки отдельного упражнения с Exercism

```bash
cd solutions/go/hello-world
go test -v
```

### Запуск всех тестов (если используете Go workspaces)

```bash
go work init
go work use ./solutions/go/*
go test ./...
```

## Цели репозитория

- Практика синтаксиса и идиом Go
- Работа с основными концепциями: структуры, методы, интерфейсы, горутины и т.д.
- Подготовка к реальным проектам на Go
- Ведение истории обучения

## Полезные ресурсы

- [Exercism Go Track](https://exercism.org/tracks/go)
- [Go Official Documentation](https://go.dev/doc/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go by Example](https://gobyexample.com/)

---

⭐ Если репозиторий оказался полезным — можешь поставить звезду!

**Happy coding!** 🚀
