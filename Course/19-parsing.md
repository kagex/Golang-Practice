# 19. Парсинг и веб-скрапинг — продвинутый модуль (с нуля до профи)

> Модуль для тех, кто хочет уметь вытаскивать данные откуда угодно: HTML, JSON, XML, CSV, PDF, бинарные форматы, гигабайтные логи, динамические SPA на JS. На Go это делается быстро, идиоматично и без боли — если знать правильные библиотеки и подводные камни.

**Длительность:** 3–4 недели
**Префикс по грейдам:** Junior проходит первые 2 недели, Middle — все 4, Senior использует модуль как референс по производительности и обходу антибот-защит.

---

## 🎯 Чему вы научитесь

- Парсить HTML/XML через `golang.org/x/net/html`, `goquery`, `colly`.
- Работать с динамическим JS-контентом через `chromedp` и `playwright-go`.
- Стримить и разбирать JSON/NDJSON объёмом в десятки гигабайт без OOM.
- Писать собственные лексеры и парсеры (рекурсивный спуск, Pratt-парсер, `participle`, `goyacc`).
- Парсить бинарные форматы (`encoding/binary`, MessagePack, CBOR).
- Делать промышленные скрейперы с rate-limit, retries, прокси-ротацией, кэшем и распределённой очередью задач.
- Обходить базовые антибот-защиты этично (User-Agent, cookies, TLS-fingerprint через `utls`).
- Тестировать парсеры через fuzzing и golden-файлы.

---

## 🗺 Дорожная карта модуля

| Неделя | Тема | Что закрываем |
|--------|------|---------------|
| 1 | Базовый парсинг текста и форматов | strings, regexp, bufio, encoding/* |
| 2 | HTML и веб-скрапинг | net/http, goquery, colly |
| 3 | JS-рендеринг + продвинутый скрапинг | chromedp, playwright-go, прокси, rate-limit |
| 4 | Свои парсеры + бинарные форматы + проект | participle, goyacc, binary, fuzzing |

---

## 📚 Неделя 1 — фундамент: текст, regexp, потоки

### Урок 1.1 — `strings`, `strconv`, `unicode`

**Теория:** `strings.Builder` (почему не `+=`), `strings.Cut` (1.18+), нормализация Unicode через `golang.org/x/text/unicode/norm`, NFC vs NFD.
**Анти-паттерны:** склейка строк через `+` в цикле → O(n²); сравнение Unicode без нормализации (`café` ≠ `café`).
**Практика:** написать функцию `Tokenize(s string) []string`, которая режет произвольный текст на слова, корректно работая с UTF-8, эмодзи и составными графемами (`uniseg`).

### Урок 1.2 — `regexp` без боли

**Теория:** RE2-семантика (нет lookahead/backreferences — и слава богу, нет ReDoS), `MustCompile` на пакетном уровне, named groups, `FindAllStringSubmatchIndex`, бенчмарк compiled vs inline.
**Практика:** парсер логов nginx combined-формата через одну `regexp`, выдаёт `[]LogEntry`. Сравнить скорость с ручным разбором через `strings.Cut` — обычно ручной в 3–5 раз быстрее, и это важный урок.

### Урок 1.3 — потоковое чтение через `bufio`

**Теория:** `bufio.Scanner` и его ловушка `MaxScanTokenSize` (64KB по умолчанию — длинная строка лога → `bufio.Scanner: token too long`). Решение: `scanner.Buffer(make([]byte, 1024*1024), 64*1024*1024)` или `bufio.Reader.ReadString('\n')`.
**Практика:** обработать 10 ГБ access.log, посчитать топ-100 IP по числу запросов. RAM не должна превышать 100 МБ.

### Урок 1.4 — `encoding/csv`, `encoding/json`, `encoding/xml`

**Теория:** потоковый `json.Decoder` vs `json.Unmarshal` (вся в память), `Decoder.Token()` для огромных массивов, `xml.Decoder.Token()` для SAX-стиля, `csv.Reader.ReuseRecord = true` (экономит аллокации).
**Бенчмарк:** распарсить 1 ГБ JSON через `Unmarshal` (упадёт по памяти) и через `Decoder` стримом (отработает за минуты на 50 МБ RAM).

**Альтернативы стандартной либе:**

- [`goccy/go-json`](https://github.com/goccy/go-json) — drop-in замена, в 2–3 раза быстрее.
- [`bytedance/sonic`](https://github.com/bytedance/sonic) — JIT-ускоренный JSON, x86_64-only.
- [`valyala/fastjson`](https://github.com/valyala/fastjson) — zero-allocation парсер, если структура известна.
- [`tidwall/gjson`](https://github.com/tidwall/gjson) — JSONPath-стиль `gjson.Get(data, "users.#(age>30).name")`.

### Мини-проект недели 1

Утилита `logtop` — читает stdin или файл с access-логами (nginx/apache/JSON), агрегирует топ-N по любому полю (IP, URL, UA, status), отдаёт результат как JSON или красивую таблицу. Покрытие тестами ≥ 85%, бенчмарк против стандартного `awk`.

---

## 🌐 Неделя 2 — HTML и веб-скрапинг

### Урок 2.1 — HTTP-клиент в продакшен-режиме

**Теория:** свой `http.Client` с таймаутами, `Transport` с настройкой пула соединений, `http.CookieJar`, `gzip`/`brotli` decompression (`andybalholm/brotli`), `context` на каждый запрос, retries с jitter (`avast/retry-go`, `hashicorp/go-retryablehttp`).
**Анти-паттерн:** `http.Get` в проде. Нет таймаута → один зависший хост валит сервис.

### Урок 2.2 — `golang.org/x/net/html` (низкий уровень)

**Теория:** токенайзер vs дерево, итерация по узлам, восстановление после битого HTML.
**Практика:** распарсить страницу Hacker News, вытащить заголовки + ссылки + score. Без сторонних библиотек — только stdlib + x/net/html.

### Урок 2.3 — `PuerkitoBio/goquery` (jQuery-стиль)

**Теория:** CSS-селекторы, `.Find().Each()`, `.Attr()`, `.Text()`, нормализация URL через `net/url.ResolveReference`.
**Подводные камни:** `Text()` склеивает всё подряд без пробелов; для аккуратного текста — обходите DOM сами.
**Практика:** скрейпер Habr-главной → JSON со статьями, авторами, тегами, рейтингом.

### Урок 2.4 — `gocolly/colly` (фреймворк для скрейпинга)

**Теория:** `Collector`, `OnHTML`, `OnRequest`, `OnResponse`, `OnError`, асинхронность, `LimitRule` (rate-limit на хост), кэш через `colly/extensions`, distributed-режим через Redis.
**Практика:** обойти многостраничный каталог (Project Gutenberg или ваш собственный учебный сайт), уважая `robots.txt`, с rate-limit 1 RPS на хост, сохраняя в SQLite.
**Этика:** всегда проверяйте `robots.txt` (`temoto/robotstxt`), уважайте `Crawl-delay`, ставьте честный `User-Agent` с контактом. Скрейпинг публичных данных — ок, обход капчи и авторизации без разрешения — нет.

### Мини-проект недели 2

`newscrawler` — собирает заголовки + ссылки + краткое описание с 5 новостных сайтов, складывает в SQLite, отдаёт через REST API с фильтрами по дате/источнику. Уважает robots.txt, имеет rate-limit, retries с exponential backoff и метрики Prometheus.

---

## 🤖 Неделя 3 — JS-рендеринг и продвинутый скрапинг

### Урок 3.1 — `chromedp` (headless Chrome через CDP)

**Теория:** Chrome DevTools Protocol, контексты, таргеты, ожидание элементов (`chromedp.WaitVisible`), выполнение JS (`chromedp.Evaluate`), скриншоты, перехват сети (`network.Enable`).
**Подводные камни:** chromedp жрёт RAM (200–500 МБ на инстанс), нужен пул через `chromedp.NewExecAllocator` и таймауты на КАЖДЫЙ шаг.
**Практика:** залогиниться на тестовый сайт (`the-internet.herokuapp.com`), снять скриншот авторизованной зоны, вытащить данные из таблицы, отрендеренной через React.

### Урок 3.2 — `playwright-community/playwright-go`

**Теория:** Playwright vs Puppeteer/CDP, auto-waiting, локаторы (`page.Locator(...)`), сценарии с несколькими браузерами, перехват и mock-ование сетевых запросов.
**Когда что:** chromedp — легче и быстрее для простых задач; Playwright — мощнее для сложных SPA и кросс-браузерности.
**Практика:** scenario-тест на сайт-SPA с infinite scroll.

### Урок 3.3 — TLS-fingerprint и обход базовых антибот-систем (этично!)

**Теория:** JA3/JA4 fingerprints, почему стандартный `crypto/tls` палится антибот-системами, `refraction-networking/utls` для имитации Chrome/Firefox-хендшейка. HTTP/2 fingerprint, header order matters.
**Важно:** мы говорим о парсинге публичных данных, разрешённых ToS или ваших собственных сервисов. Обход капч, защит Cloudflare, авторизаций без разрешения — вне рамок этого курса и часто незаконен.
**Практика:** написать клиент, который отдаёт корректный JA3-fingerprint Chrome и проходит тестовый сайт `tls.peet.ws`.

### Урок 3.4 — Прокси, ротация, очереди задач

**Теория:** HTTP/SOCKS5 прокси, ротация через `Transport.Proxy = func(*Request) (*url.URL, error)`, sticky-сессии через cookie jar, распределение через Redis/NATS, идемпотентность задач.
**Практика:** превратить скрейпер из недели 2 в распределённый — N воркеров читают задачи из NATS JetStream, дедуп по URL через Redis, результаты складывают в PostgreSQL. Graceful shutdown по SIGTERM.

### Мини-проект недели 3

`spa-scraper` — универсальный скрейпер SPA на chromedp с пулом браузерных контекстов, прокси-ротацией, очередью задач в NATS, retries, экспортом результатов в JSON/Parquet и метриками в Prometheus.

---

## 🧬 Неделя 4 — свои парсеры, бинарные форматы, проект

### Урок 4.1 — `text/scanner` и ручной лексер

**Теория:** конечные автоматы, `bufio.Scanner` с кастомным `SplitFunc`, идиома lexer-as-state-function (Rob Pike — «Lexical Scanning in Go»).
**Практика:** написать лексер простого языка выражений (`1 + 2 * (3 - foo)`).

### Урок 4.2 — Парсер рекурсивным спуском и Pratt-парсер

**Теория:** грамматика в EBNF, приоритеты операторов, ассоциативность, обработка ошибок с восстановлением (panic mode), AST + visitor pattern.
**Практика:** дописать парсер из 4.1 до полноценного калькулятора с переменными и функциями. Получится мини-DSL.

### Урок 4.3 — `alecthomas/participle` — парсеры через struct-теги

**Теория:** PEG-стиль, теги `@@`, `@Ident`, lookahead, lexer-customization. Когда participle сильно проще ручного, а когда — нет.
**Практика:** парсер своего конфиг-формата (что-то между TOML и HCL) на 50 строк participle.

### Урок 4.4 — `goyacc` (LALR(1)) — для серьёзных грамматик

**Теория:** YACC-нотация, конфликты shift/reduce, как Go-команда сама использует goyacc в `go/parser` (раньше).
**Практика:** мини-SQL парсер (SELECT с WHERE, ORDER BY, LIMIT) на goyacc, выдающий AST.
**Когда брать:** ручной парсер — для маленьких DSL; participle — для средних; goyacc — для языков и сложных грамматик.

### Урок 4.5 — Бинарные форматы

**Теория:** `encoding/binary` (BigEndian/LittleEndian), выравнивание, structs с `binary.Size`, парсинг ELF/PE/PNG/ZIP заголовков, MessagePack (`vmihailenco/msgpack`), CBOR (`fxamacker/cbor`), Protobuf без `.proto` через `protoreflect`.
**Практика:** парсер PNG-файла — читает chunks (IHDR, IDAT, IEND), валидирует CRC32, выдаёт метаданные (размер, bit depth, colour type).

### Урок 4.6 — Тестирование парсеров: fuzzing + golden files

**Теория:** `testing.F.Fuzz`, корпус из реальных данных в `testdata/fuzz`, минимизация падений, golden-файлы (`-update` флаг), property-based testing (`leanovate/gopter`).
**Практика:** запустить fuzz на парсере из 4.2 на 10 минут — обычно находится 1–2 паники на edge-cases. Чините, добавляете golden-тесты, отправляете в CI.

### Урок 4.7 — Производительность парсинга

**Теория:** zero-copy через `[]byte` вместо `string`, `unsafe` string↔bytes конверсия (когда оправдано), пулы буферов (`sync.Pool`), SIMD-парсинг JSON (sonic), профилирование pprof, allocs/op в бенчмарках.
**Практика:** оптимизировать парсер логов из недели 1 — цель: уменьшить allocs/op в 5+ раз, MB/s — увеличить в 2+ раза. Замер через `benchstat` до/после.

---

## 🏁 Финальный капстоун модуля

**`gocrawl` — production-grade скрейпинг-платформа на Go.**

Требования:

1. **Источники:** статический HTML (goquery), SPA с JS (chromedp), JSON-API, RSS/Atom (gofeed). Драйверы плагинами.
2. **Очередь задач:** NATS JetStream или Kafka с at-least-once, дедуп URL через Redis (TTL).
3. **Хранилище:** PostgreSQL для метаданных + S3/MinIO для сырых HTML-снапшотов.
4. **Соблюдение приличий:** `robots.txt`, `Crawl-delay`, rate-limit на хост (token bucket), вежливый User-Agent с email-контактом.
5. **Устойчивость:** retries с exponential backoff + jitter, circuit breaker per host, dead letter queue.
6. **Observability:** structured logs (slog), метрики Prometheus (requests, errors, latency, queue depth), OpenTelemetry traces до уровня одного URL.
7. **DX:** CLI (`gocrawl add <url>`, `gocrawl stats`), HTTP API, конфиг YAML, Docker Compose для локального запуска.
8. **Тесты:** unit + integration через testcontainers (Postgres, Redis, NATS), fuzz на парсерах, ≥ 80% coverage, e2e против локального httpbin/wiremock.
9. **CI:** GitHub Actions с lint (golangci-lint), test, race, fuzz (5 минут на PR), docker build, govulncheck.

Что вы покажете на собеседовании:

- Понимание HTTP, HTML, JS-рендеринга, бинарных форматов.
- Идиоматичный Go: интерфейсы для драйверов, context везде, graceful shutdown.
- Распределёнка: очереди, идемпотентность, dedup.
- Production-mindset: observability, метрики, безопасность, этичность.

---

## 📚 Бесплатные ресурсы по модулю

### Документация и спецификации

- [Effective Go — IO patterns](https://go.dev/doc/effective_go)
- [pkg.go.dev — encoding/json](https://pkg.go.dev/encoding/json), [encoding/xml](https://pkg.go.dev/encoding/xml), [encoding/binary](https://pkg.go.dev/encoding/binary)
- [golang.org/x/net/html docs](https://pkg.go.dev/golang.org/x/net/html)
- [HTML5 Parsing spec — WHATWG](https://html.spec.whatwg.org/multipage/parsing.html)
- [JSON RFC 8259](https://www.rfc-editor.org/rfc/rfc8259)

### Статьи и доклады

- Rob Pike — *Lexical Scanning in Go* (видео + слайды) — основа любого ручного парсера на Go.
- Dave Cheney — *High Performance JSON in Go*.
- Eli Bendersky — *A recursive descent parser in Go*.
- Three Dots Labs — *Writing scrapers that don't break*.

### Библиотеки (всё OSS)

- **HTML/CSS:** `PuerkitoBio/goquery`, `andybalholm/cascadia`, `gocolly/colly`.
- **JS-рендеринг:** `chromedp/chromedp`, `playwright-community/playwright-go`, `go-rod/rod`.
- **JSON:** `goccy/go-json`, `bytedance/sonic`, `valyala/fastjson`, `tidwall/gjson`.
- **XML:** stdlib + `antchfx/xmlquery` (XPath).
- **CSV:** stdlib + `gocarina/gocsv` (struct-теги).
- **Feeds:** `mmcdole/gofeed` (RSS/Atom/JSON Feed).
- **Парсер-генераторы:** `alecthomas/participle`, `goyacc`, `mna/pigeon` (PEG).
- **Бинарные:** `vmihailenco/msgpack`, `fxamacker/cbor`, `google.golang.org/protobuf/reflect/protoreflect`.
- **HTTP-fingerprint:** `refraction-networking/utls`.
- **robots.txt:** `temoto/robotstxt`.
- **Retries:** `avast/retry-go`, `hashicorp/go-retryablehttp`.

### Практика

- [Exercism Go track — parsers](https://exercism.org/tracks/go)
- [Advent of Code](https://adventofcode.com/) — половина задач сводится к парсингу.
- [Crafting Interpreters](https://craftinginterpreters.com/) (Bob Nystrom) — не на Go, но идеи переносятся идеально.

---

## ⚠️ Этика и закон

Парсинг — мощный инструмент, и им легко навредить. Правила хорошего тона:

1. **Читайте `robots.txt` и ToS сайта.** Если запрещено — не парсите.
2. **Уважайте rate-limit.** 1 RPS на хост — почти всегда безопасно. Не валите чужой сервер.
3. **Не парсите персональные данные** без законного основания (GDPR, локальное законодательство).
4. **Не обходите технические защиты** (капчи, авторизации) без явного разрешения владельца — это может быть нарушением CFAA/закона о неправомерном доступе.
5. **Идентифицируйте себя.** В User-Agent — название бота + email/URL контакта. Так с вами свяжутся прежде, чем забанить.
6. **Кэшируйте.** Если данные не изменились — не дёргайте сайт повторно. Conditional GET (`If-Modified-Since`, `ETag`) — ваш друг.

---

## 🔥 Чек-лист «я прошёл модуль»

- [ ] Понимаю разницу между `Unmarshal` и `Decoder` и когда что брать.
- [ ] Знаю, почему `bufio.Scanner` падает на длинных строках и как чинить.
- [ ] Могу написать парсер логов и оптимизировать его через pprof.
- [ ] Могу скрейпить статический и динамический контент, уважая `robots.txt`.
- [ ] Понимаю, как делать retries, rate-limit, circuit breaker и dedup.
- [ ] Написал хотя бы один лексер и парсер вручную (без библиотек).
- [ ] Использовал participle или goyacc для DSL.
- [ ] Парсил бинарный формат через `encoding/binary`.
- [ ] Прогнал fuzz и починил найденные паники.
- [ ] Капстоун `gocrawl` лежит на GitHub с README, тестами, CI и Docker Compose.

После этого вы — реально сильный парсинг/скрейпинг-инженер на Go, которого берут на позиции Data Engineering, Web Scraping, и серьёзные backend-роли, где нужен ETL.
