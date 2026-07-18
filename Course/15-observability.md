# 🔭 Модуль 15. Observability & SRE

> **Срок:** 3–4 недели  
> **Уровень:** 🔴 Senior  
> **Цель модуля:** перестать «дебажить логами» — выстраивать observability как продукт, мыслить в терминах SLO/SLI/error budget, готовить системы к chaos.

---

## Урок 15.1. Three pillars и почему этого мало

- 🎯 **Цель:** понимать различия и связь metrics/logs/traces, знать о continuous profiling и events.
- 📚 **Теория:** Cindy Sridharan "Distributed Systems Observability", OpenTelemetry docs.
- 💻 **Практика:** обвешать Go-сервис тремя сигналами, прокинуть trace_id из логов в Jaeger и обратно.
- ✅ **Чек:** почему metrics ≠ observability. Зачем exemplars.
- ⚠️ **Красный флаг:** «у нас Grafana, значит observability есть».

## Урок 15.2. OpenTelemetry в Go глубоко

- 🎯 **Цель:** SDK, exporters, propagators, baggage, sampling, OTel Collector pipeline.
- 📚 **Теория:** opentelemetry.io/docs/languages/go, спека OTLP.
- 💻 **Практика:** настроить OTel SDK с tail-based sampling, прокинуть W3C trace context через HTTP+Kafka+gRPC, поднять Collector с processors.
- ✅ **Чек:** разница head vs tail sampling. Что такое span links и когда они нужны.
- ⚠️ **Красный флаг:** 100% sampling в проде «чтобы ничего не потерять».

## Урок 15.3. Prometheus и метрики уровня production

- 🎯 **Цель:** RED/USE, histogram vs summary, cardinality, recording rules, federation, remote write.
- 📚 **Теория:** Brendan Gregg USE method, Tom Wilkie RED method, Prometheus docs про high cardinality.
- 💻 **Практика:** запилить RED-дашборд для сервиса. Сравнить native histograms и classic. Поймать cardinality explosion и починить.
- ✅ **Чек:** когда histogram_quantile врёт. Чем плох label с user_id.
- ⚠️ **Красный флаг:** метрика с label, содержащий request_id или email.

## Урок 15.4. Логи: структурированность, sampling, корреляция

- 🎯 **Цель:** `log/slog`, JSON-логи, level guidelines, log sampling под нагрузкой, корреляция с trace.
- 📚 **Теория:** Go blog про slog, статьи про log levels (Dave Cheney).
- 💻 **Практика:** перевести проект на slog с context-aware handler, который добавляет trace_id/span_id. Внедрить sampler для debug-логов.
- ✅ **Чек:** почему `info` логов должно быть мало. Зачем sampling на error-логах в throttling-сценариях.
- ⚠️ **Красный флаг:** `log.Println` в hot path без уровней.

## Урок 15.5. Continuous profiling и runtime-визибилити

- 🎯 **Цель:** Pyroscope/Parca, pprof в проде, runtime/metrics, eBPF observability.
- 📚 **Теория:** Parca docs, Polar Signals блог, статьи про eBPF (Brendan Gregg).
- 💻 **Практика:** включить continuous profiling в Go-сервисе, найти регрессию по CPU между двумя релизами. Подключить runtime/metrics к Prometheus.
- ✅ **Чек:** что покажет block-профиль, чего не покажут metrics. Как читать flamegraph.
- ⚠️ **Красный флаг:** «у нас CPU 80%, наверное, аллокации» — без pprof.

## Урок 15.6. SLO, SLI, error budget, alerting

- 🎯 **Цель:** определять SLI по journey пользователя, считать SLO и error budget, делать burn-rate alerts.
- 📚 **Теория:** Google SRE Book (главы про SLO), SRE Workbook, статьи Alex Hidalgo.
- 💻 **Практика:** сформулировать SLO для сервиса (availability, latency, freshness), настроить multi-window burn-rate alerts в Prometheus.
- ✅ **Чек:** почему alert на «CPU > 80%» — antipattern. Что значит fast/slow burn.
- ⚠️ **Красный флаг:** alert на каждую метрику без связки с user impact.

## Урок 15.7. Incident response и chaos engineering

- 🎯 **Цель:** runbooks, postmortems blameless, chaos-эксперименты, game days.
- 📚 **Теория:** Google SRE Book (incident management), книга "Chaos Engineering" (Rosenthal).
- 💻 **Практика:** провести chaos-эксперимент (network delay, kill pod, disk fill) на staging. Написать blameless postmortem по реальному инциденту.
- ✅ **Чек:** структура хорошего postmortem. Принципы chaos engineering Netflix.
- ⚠️ **Красный флаг:** «виноват Вася» в postmortem.

---

## 🎯 Проект модуля

**Observability stack для микросервиса.**

Критерии приёмки:
- Сервис с RED-метриками, slog-логами, OTel traces.
- OTel Collector с tail-based sampling и фильтрами.
- 3 SLO (availability, latency, freshness) с burn-rate alerts.
- Runbook на каждый alert.
- Chaos-эксперимент: network partition + отчёт.
- Дашборд: USE для инфры, RED для сервиса, golden signals.

---

## 🏁 Чек-пойнт модуля

- [ ] Настроил OTel end-to-end сам.
- [ ] Знаю, как ловить high cardinality в Prometheus.
- [ ] Считаю SLO и error budget руками.
- [ ] Писал blameless postmortem.
- [ ] Делал chaos-эксперимент с гипотезой и выводами.

---

[← Модуль 14. Distributed Systems](./14-distributed-systems.md) | [Модуль 16. Security в Go →](./16-security.md)
