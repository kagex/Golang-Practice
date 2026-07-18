# Модуль 8. Микросервисы и сети 🟠

**Срок:** 6–8 недель  
**Цель модуля:** научиться разрабатывать распределённые системы — брокеры, паттерны устойчивости, tracing.

---

## Урок 8.1. Когда микросервисы нужны, когда нет

**📚 Теория:** monolith-first принцип, Conway's Law, цена распределённости (latency, eventual consistency, операционная сложность), когда пора разбивать.

**⚠️ Красный флаг:** микросервисы в проекте на 5 человек, где хватило бы монолита.

---

## Урок 8.2. Брокеры сообщений

**📚 Теория:**
- **Kafka:** партиции, consumer groups, гарантии (at-least-once / at-most-once / exactly-once), `confluent-kafka-go` или `segmentio/kafka-go`
- - **RabbitMQ:** exchange, queue, routing key, `amqp091-go`
  - - **NATS:** subjects, JetStream, `nats.go`
   
    - **💻 Практика:** producer + consumer на Kafka с учётом гарантий.
   
    - ---

    ## Урок 8.3. Saga и Outbox

    **📚 Теория:** распределённые транзакции без 2PC, choreography vs orchestration saga, Outbox pattern для надёжной публикации событий, idempotency keys.

    **💻 Практика:** Outbox в PostgreSQL + relay в Kafka.

    ---

    ## Урок 8.4. CQRS и Event Sourcing

    **📚 Теория:** разделение команд и запросов, проекции (read models), повторное воссоздание состояния из событий, когда это оправдано.

    **⚠️ Красный флаг:** Event Sourcing везде, где видел команду.

    ---

    ## Урок 8.5. Устойчивость

    **📚 Теория:** Circuit Breaker (`sony/gobreaker`), Retry с экспоненциальным backoff и jitter, bulkhead, timeout, idempotency, deadlines.

    **💻 Практика:** обёртка HTTP-клиента с retry + circuit breaker.

    ---

    ## Урок 8.6. Service Discovery и API Gateway

    **📚 Теория:** Consul/etcd, DNS-based discovery, Kubernetes Services, API Gateway (Kong, Traefik, Ambassador), edge вс internal traffic.

    ---

    ## Урок 8.7. gRPC interceptors

    **📚 Теория:** unary и stream interceptors, цепочки, логирование, auth, retry, recovery, metrics, tracing (через `otelgrpc`).

    **💻 Практика:** напиши auth interceptor, проверяющий JWT в metadata.

    ---

    ## Урок 8.8. Distributed Tracing

    **📚 Теория:** OpenTelemetry (стандарт), span, trace context, propagation через HTTP/gRPC headers, exporters (OTLP), Jaeger/Tempo для визуализации.

    **💻 Практика:** инструментируй два своих сервиса, посмотри trace в Jaeger.

    ---

    ## 🎯 Проект модуля: Мини-маркетплейс

    4 сервиса:
    - **users** — регистрация, логин, профиль
    - - **catalog** — товары
      - - **orders** — заказы (saga: резерв товара → оплата → отгрузка)
        - - **notifier** — уведомления (читает Kafka)
         
          - Коммуникация: gRPC между сервисами + Kafka для событий. Каждый в своём контейнере.
         
          - **Критерии приёмки:** Outbox реализован, distributed tracing проходит через все сервисы, Circuit Breaker на вызовах внешних систем.
         
          - ---

          ## 🏁 Чек-пойнт модуля

          - [ ] Работаешь хотя бы с одним брокером на практике
          - [ ] - [ ] Понимаешь разницу между at-least-once и exactly-once
          - [ ] - [ ] Реализовываешь Saga и Outbox
          - [ ] - [ ] Пишешь Circuit Breaker и retry с jitter
          - [ ] - [ ] Настраиваешь distributed tracing
         
          - [ ] **Предыдущий:** [Модуль 7 ←](07-testing.md) · **Следующий:** [Модуль 9. DevOps минимум →](09-devops.md)
          - [ ] 
