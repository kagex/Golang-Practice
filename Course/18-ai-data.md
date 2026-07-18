# 🤖 Модуль 18. AI/ML & Data Engineering на Go

> **Срок:** 3–4 недели  
> **Уровень:** 🔴 Senior  
> **Цель модуля:** строить на Go сервисы, которые интегрируются с LLM, vector DB и стримами данных — там, где Python обычно прототипирует, а Go держит прод.

---

## Урок 18.1. LLM-интеграции из Go

- 🎯 **Цель:** работать с OpenAI/Anthropic/локальными моделями (Ollama, llama.cpp) через HTTP/streaming, грамотно строить prompt-пайплайны.
- 📚 **Теория:** OpenAI/Anthropic API docs, статьи про prompt engineering, Eugene Yan blog про LLM в продукте.
- 💻 **Практика:** Go-клиент с streaming SSE для chat completions, поддержкой tool calling и retry. Свой минимальный LangChain-like граф (chain of prompts).
- ✅ **Чек:** что такое context window и почему он критичен. Зачем streaming.
- ⚠️ **Красный флаг:** SDK без таймаутов и retry для LLM — медленные модели легко вешают сервис.

## Урок 18.2. Embeddings и vector databases

- 🎯 **Цель:** понимать embedding-модели, метрики расстояния, индексы (HNSW, IVF), работать с pgvector/Qdrant/Weaviate/Milvus.
- 📚 **Теория:** статьи про HNSW, документация pgvector и Qdrant.
- 💻 **Практика:** построить семантический поиск по корпусу markdown-файлов: embedding + pgvector + Go API. Сравнить cosine и L2.
- ✅ **Чек:** что значит recall@k и почему он важнее latency. Когда HNSW проигрывает brute force.
- ⚠️ **Красный флаг:** vector DB без тестов на качество поиска.

## Урок 18.3. RAG и продакшен-паттерны

- 🎯 **Цель:** retrieval-augmented generation: chunking, reranking, hybrid search (BM25 + vectors), guardrails.
- 📚 **Теория:** статьи Pinecone/Anthropic про RAG, Contextual Retrieval, BM25.
- 💻 **Практика:** RAG-сервис на Go: загрузка документов → chunking → embeddings → hybrid search → LLM ответ с citations. Eval-набор и metric (faithfulness, context recall).
- ✅ **Чек:** зачем reranker. Что такое hallucination и как меряется faithfulness.
- ⚠️ **Красный флаг:** RAG без eval — невозможно объяснить, стало лучше или хуже.

## Урок 18.4. Agents, tool calling, MCP

- 🎯 **Цель:** строить агенты с tool calling, понимать MCP (Model Context Protocol), безопасно давать инструменты.
- 📚 **Теория:** Anthropic "Building Effective Agents", MCP spec.
- 💻 **Практика:** Go-агент с tools (search, fetch, calc) и MCP-сервером для своего домена. Sandbox для tool execution.
- ✅ **Чек:** когда нужен агент, а когда хватит pipeline. Какие риски tool calling в проде.
- ⚠️ **Красный флаг:** агент с file system tool без sandbox.

## Урок 18.5. Stream processing на Go

- 🎯 **Цель:** строить пайплайны на Kafka/NATS/Redpanda, понимать watermarks, windowing, exactly-once.
- 📚 **Теория:** Tyler Akidau "Streaming Systems", статьи Confluent про EOS, Benthos/Redpanda Connect docs.
- 💻 **Практика:** stream-пайплайн на Go (или Benthos): Kafka → enrichment через HTTP → tumbling windows → ClickHouse. Поддержка late events.
- ✅ **Чек:** разница event time и processing time. Зачем watermarks.
- ⚠️ **Красный флаг:** «processing time достаточно» для аналитики — данные приходят с лагом всегда.

## Урок 18.6. OLAP и аналитические БД из Go

- 🎯 **Цель:** ClickHouse, DuckDB, Apache Arrow, Parquet, columnar storage.
- 📚 **Теория:** ClickHouse docs, Arrow Go docs, статьи про columnar formats.
- 💻 **Практика:** Go-сервис, который пишет события в ClickHouse через native protocol, читает через Arrow Flight. Локальная аналитика на DuckDB по Parquet.
- ✅ **Чек:** почему columnar выигрывает на аналитике. Когда Parquet удобнее ClickHouse.
- ⚠️ **Красный флаг:** хранить аналитику в Postgres «потому что уже есть».

## Урок 18.7. ML inference на Go

- 🎯 **Цель:** запускать ONNX-модели из Go (onnxruntime, gorgonia), gRPC-обвязка вокруг моделей, batching, autoscaling.
- 📚 **Теория:** ONNX runtime docs, Triton Inference Server, статьи про batch inference.
- 💻 **Практика:** Go-сервис, который грузит ONNX-модель (например, классификатор изображений) и отдаёт через gRPC с dynamic batching.
- ✅ **Чек:** зачем dynamic batching. Когда выносить инференс из Go-сервиса в Triton.
- ⚠️ **Красный флаг:** загружать модель на каждый запрос.

---

## 🎯 Проект модуля

**Production-grade RAG-сервис на Go.**

Критерии приёмки:
- Пайплайн: документы → chunking → embeddings → vector DB (pgvector) → hybrid search → LLM с citations.
- Streaming SSE-ответы, tool calling для запросов к внутренним API.
- Eval-набор и метрики (faithfulness, context recall, answer relevance).
- Guardrails: PII-фильтр, rate limit per user, prompt injection detection.
- Observability: latency p99, token usage, cost per request, RAG quality dashboard.
- A/B тест двух embedding-моделей с метрикой recall@10.

---

## 🏁 Чек-пойнт модуля

- [ ] Делал RAG end-to-end с eval'ом.
- [ ] Понимаю vector search и indices глубже, чем «вызвал API».
- [ ] Считаю стоимость LLM-запросов и оптимизировал её.
- [ ] Запускал ONNX-модель из Go.
- [ ] Строил streaming-пайплайн с windowing и watermarks.

---

[← Модуль 17. Cloud Native & K8s Operators](./17-cloud-native.md) | [📚 К содержанию курса](../README.md)
