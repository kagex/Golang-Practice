# 🧬 Модуль 13. Go Runtime & Internals

> **Срок:** 4–6 недель  
> **Уровень:** 🔴 Senior  
> **Цель модуля:** уметь объяснить, как Go работает изнутри, тюнить GC и scheduler под прод, читать runtime-код и дампы, осознанно использовать unsafe.

---

## Урок 13.1. Memory model и happens-before

- 🎯 **Цель:** понимать, какие гарантии даёт Go Memory Model, какие операции синхронизируются.
- 📚 **Теория:** [Go Memory Model](https://go.dev/ref/mem), статьи Russ Cox про memory models, lock-free programming.
- 💻 **Практика:** написать примеры data race, поймать их через `-race`, объяснить каждый. Реализовать spinlock и сравнить с `sync.Mutex`.
- ✅ **Чек:** объясни, почему `sync.Once` корректен. Что такое sequenced before / synchronizes with.
- ⚠️ **Красный флаг:** «atomic быстрее, потому что без блокировок» — без понимания memory ordering это карго-культ.

## Урок 13.2. Scheduler (G-M-P) глубоко

- 🎯 **Цель:** понимать work-stealing, sysmon, preemption, netpoller, LockOSThread.
- 📚 **Теория:** статья Dmitry Vyukov, доклады Kavya Joshi, исходники `runtime/proc.go`.
- 💻 **Практика:** трассировать программу через `go tool trace`, найти STW-паузы, syscalls, блокировки P. Эксперимент: что делает `runtime.LockOSThread` и когда он нужен.
- ✅ **Чек:** объясни, как горутина переходит из running → runnable → waiting. Что такое handoff P.
- ⚠️ **Красный флаг:** «GOMAXPROCS = число горутин» — путаница M, P и G.

## Урок 13.3. Garbage Collector: tricolor, write barrier, pacer

- 🎯 **Цель:** понимать concurrent mark-sweep, write barrier, GC pacer, GOGC и GOMEMLIMIT.
- 📚 **Теория:** Go GC guide (go.dev/doc/gc-guide), доклады Rick Hudson и Austin Clements.
- 💻 **Практика:** снять `GODEBUG=gctrace=1`, разобрать строки. Сравнить latency приложения при GOGC=50/100/300 и с GOMEMLIMIT.
- ✅ **Чек:** что делает write barrier и почему он нужен. Когда полезен `runtime.GC()` (почти никогда).
- ⚠️ **Красный флаг:** «уменьшил GOGC — стало быстрее» без замеров p99 latency и CPU.

## Урок 13.4. Escape analysis и аллокации

- 🎯 **Цель:** видеть, где компилятор аллоцирует на куче, и осознанно уводить на стек.
- 📚 **Теория:** `go build -gcflags='-m -m'`, статьи про escape analysis от Damian Gryski.
- 💻 **Практика:** взять hot path, прогнать `-gcflags='-m'`, убрать аллокации (sync.Pool, pre-allocated buffers, избегание interface boxing).
- ✅ **Чек:** почему возврат указателя из функции часто escape'ит. Когда interface вызывает аллокацию.
- ⚠️ **Красный флаг:** микрооптимизации без бенчмарков и pprof.

## Урок 13.5. unsafe, cgo, assembly

- 🎯 **Цель:** уметь применять `unsafe.Pointer`, понимать стоимость и риски cgo, читать Go-assembler (Plan 9).
- 📚 **Теория:** [unsafe](https://pkg.go.dev/unsafe), [cgo wiki](https://github.com/golang/go/wiki/cgo), Plan 9 assembler guide.
- 💻 **Практика:** zero-copy конверсия `[]byte ↔ string` через unsafe. Минимальная cgo-обёртка над C-библиотекой. Прочитать ассемблер `runtime.memmove`.
- ✅ **Чек:** четыре легальных паттерна unsafe.Pointer из доки. Стоимость cgo-вызова (~150–200 нс) и почему.
- ⚠️ **Красный флаг:** unsafe ради «производительности» без замеров. cgo в hot path.

## Урок 13.6. Reflection и generics под капотом

- 🎯 **Цель:** понимать стоимость reflect, как реализованы generics (GC shape stenciling).
- 📚 **Теория:** Russ Cox "The Laws of Reflection", доклад про generics implementation.
- 💻 **Практика:** бенчмарк reflect vs code-gen vs type switch. Написать дженерик-функцию, посмотреть на сгенерированный код через `objdump`.
- ✅ **Чек:** почему `reflect.Value.Interface()` аллоцирует. Что такое dictionary в дженериках.
- ⚠️ **Красный флаг:** reflect в hot path вместо генерации кода.

---

## 🎯 Проект модуля

**High-performance in-memory cache с low GC pressure.**

Критерии приёмки:
- Concurrent get/set/delete с шардированием.
- p99 < 1 ms под нагрузкой 100k RPS.
- GC pause < 1 ms (через GOMEMLIMIT и pre-allocation).
- Отчёт: pprof heap/cpu/block/mutex до и после оптимизаций.
- Сравнение с `sync.Map` и `freecache` по бенчмаркам.

---

## 🏁 Чек-пойнт модуля

- [ ] Объясняю G-M-P и work-stealing без подглядываний.
- [ ] Читаю `gctrace` и `schedtrace`.
- [ ] Тюнил GOGC/GOMEMLIMIT под реальную нагрузку.
- [ ] Использовал `unsafe` минимум один раз осознанно и с тестами.
- [ ] Знаю, когда reflect недопустим, и есть альтернативы.

---

[← Модуль 12. Подготовка к собеседованиям](./12-interview.md) | [Модуль 14. Distributed Systems →](./14-distributed-systems.md)
