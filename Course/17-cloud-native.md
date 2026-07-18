# ☸️ Модуль 17. Cloud Native & Kubernetes Operators

> **Срок:** 4–6 недель  
> **Уровень:** 🔴 Senior  
> **Цель модуля:** уметь не просто деплоить в Kubernetes, а расширять его — писать контроллеры и операторы, понимать controller-runtime, проектировать CRD как API.

---

## Урок 17.1. Kubernetes изнутри для Go-разработчика

- 🎯 **Цель:** API-машина (api-server, etcd, controllers), reconcile loop, declarative API.
- 📚 **Теория:** [Kubernetes the Hard Way](https://github.com/kelseyhightower/kubernetes-the-hard-way), книга "Programming Kubernetes" (Hausenblas & Schimanski).
- 💻 **Практика:** поднять kind/k3d, развернуть Go-сервис с liveness/readiness/startup probes, HPA по custom metric.
- ✅ **Чек:** что такое level-triggered vs edge-triggered. Почему k8s «declarative».
- ⚠️ **Красный флаг:** «kubectl apply и забыл» вместо понимания reconcile.

## Урок 17.2. client-go и informers

- 🎯 **Цель:** работать с k8s API через client-go, понимать informers, listers, work queues.
- 📚 **Теория:** sample-controller, документация client-go.
- 💻 **Практика:** написать контроллер на client-go, который реагирует на создание ConfigMap с лейблом и пишет ивент. Использовать SharedIndexInformer и rate-limited workqueue.
- ✅ **Чек:** зачем resync period. Что такое DeltaFIFO.
- ⚠️ **Красный флаг:** опрос API в цикле вместо informer.

## Урок 17.3. CRD как API: проектирование

- 🎯 **Цель:** проектировать CRD по Kubernetes API conventions, делать subresources, conversion webhooks, валидацию через CEL.
- 📚 **Теория:** Kubernetes API conventions, статьи про CRD versioning, CEL validation.
- 💻 **Практика:** спроектировать CRD `Database` с status subresource, printer columns, CEL-валидацией, conversion v1alpha1 → v1.
- ✅ **Чек:** зачем status subresource. Когда нужны conversion webhooks.
- ⚠️ **Красный флаг:** свалка полей без spec/status разделения.

## Урок 17.4. Kubebuilder и controller-runtime

- 🎯 **Цель:** писать операторы на kubebuilder/operator-sdk, понимать менеджер, reconciler, webhooks.
- 📚 **Теория:** kubebuilder book, controller-runtime docs.
- 💻 **Практика:** оператор для управления приложением (Deployment+Service+Ingress+ConfigMap из одного CR). Финализаторы для аккуратного удаления.
- ✅ **Чек:** разница predicates и event filters. Зачем owner references.
- ⚠️ **Красный флаг:** reconcile, который не идемпотентен.

## Урок 17.5. Admission webhooks, policy и Gatekeeper

- 🎯 **Цель:** validating/mutating webhooks, OPA/Gatekeeper, Kyverno, политики на проде.
- 📚 **Теория:** Kubernetes admission docs, Gatekeeper и Kyverno tutorials.
- 💻 **Практика:** написать validating webhook на controller-runtime: запрещать Pod без resources.requests. Эквивалентная политика в Kyverno.
- ✅ **Чек:** разница validating и mutating. Почему политики через CRD удобнее кода.
- ⚠️ **Красный флаг:** mutating webhook, который меняет всё подряд.

## Урок 17.6. Helm, Kustomize, GitOps

- 🎯 **Цель:** упаковка приложений, шаблонизация, GitOps через ArgoCD/Flux.
- 📚 **Теория:** Helm docs, Kustomize docs, ArgoCD getting started.
- 💻 **Практика:** Helm chart для своего оператора. Развернуть кластер через ArgoCD app-of-apps.
- ✅ **Чек:** когда Helm уместнее Kustomize и наоборот. Что такое ApplicationSet.
- ⚠️ **Красный флаг:** kubectl apply из CI вместо GitOps.

## Урок 17.7. Multi-cluster, edge, serverless на k8s

- 🎯 **Цель:** Karmada/Cluster API, KubeEdge, Knative, обзор service mesh для multi-cluster.
- 📚 **Теория:** Cluster API docs, Knative docs, статьи CNCF про edge.
- 💻 **Практика:** развернуть Knative Service для Go-функции с scale-to-zero. Cluster API на двух kind-кластерах.
- ✅ **Чек:** что даёт scale-to-zero и какой ценой (cold start). Зачем Cluster API.
- ⚠️ **Красный флаг:** «multi-cluster, потому что круто» — без бизнес-кейса.

---

## 🎯 Проект модуля

**Оператор приложения уровня production.**

Критерии приёмки:
- CRD `AppService` со spec/status, printer columns, CEL-валидацией.
- Reconciler идемпотентный, с финализатором и owner references.
- Validating webhook: проверка имени/ресурсов.
- Метрики Prometheus (controller-runtime), structured logs (logr).
- Helm chart + ArgoCD Application для оператора.
- Тесты: envtest + ginkgo, e2e через kind в CI.

---

## 🏁 Чек-пойнт модуля

- [ ] Написал минимум один оператор сам, не из туториала.
- [ ] Понимаю informer cache и work queue.
- [ ] Проектирую CRD по conventions.
- [ ] Знаю разницу validating/mutating и когда писать webhook.
- [ ] Деплою через GitOps, а не kubectl apply из терминала.

---

[← Модуль 16. Security в Go](./16-security.md) | [Модуль 18. AI/ML & Data Engineering на Go →](./18-ai-data.md)
