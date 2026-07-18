# Модуль 10. Архитектура и паттерны 🔴

**Срок:** 4–6 недель  
**Цель модуля:** проектировать поддерживаемые системы. Без фанатизма, но с пониманием.

---

## Урок 10.1. SOLID по-гошному

**📚 Теория:**
- **S** (Single Responsibility): одна причина изменения
- - **O** (Open/Closed): расширяемость через интерфейсы
  - - **L** (Liskov): подтипы не ломают контракт
    - - **I** (Interface Segregation): маленькие интерфейсы
      - - **D** (Dependency Inversion): зависеть от абстракций
       
        - Без слепого переноса из Java. Go свободен от фабрик-фабрик-стратегий-посредников.
       
        - ---

        ## Урок 10.2. Clean Architecture

        **📚 Теория:** слои (entities → use cases → interface adapters → frameworks), правило зависимостей (внутрь), DTO на границах.

        **⚠️ Красный флаг:** Clean Architecture в простом CRUD на 200 строк.

        ---

        ## Урок 10.3. Hexagonal / Ports & Adapters

        **📚 Теория:** domain в центре, ports (интерфейсы) + adapters (HTTP/gRPC/CLI/DB) по краям, легкость замены технологий.

        **💻 Практика:** выдели domain в одном из своих сервисов, вынеси PostgreSQL в адаптер.

        ---

        ## Урок 10.4. DDD-минимум

        **📚 Теория:** bounded context, ubiquitous language, entities, value objects, агрегаты и aggregate root, domain events, anti-corruption layer.

        **✅ Чек:** приведи пример value object из своего кода.

        ---

        ## Урок 10.5. Event-driven и eventual consistency

        **📚 Теория:** событийная модель, idempotent consumers, projection, цена (debugging, eventual consistency, ordering).

        ---

        ## Урок 10.6. Идиоматичные паттерны Go

        **📚 Теория:**
        - **Functional options**: `WithTimeout(d time.Duration) Option`
        - - **Builder**: для сложных объектов
          - - **Decorator** через middleware
            - - **Strategy** через интерфейс
              - - **Observer** через каналы
               
                - **💻 Практика:** перепиши конфигурацию HTTP-клиента через functional options.
               
                - ---

                ## Урок 10.7. Антипаттерны

                **📚 Теория:**
                - **God package**: всё в пакете `utils`/`common`
                - - **Anemic domain**: структуры без поведения, вся логика в сервисах
                  - - **Преждевременная абстракция**: интерфейсы с одной реализацией «на всякий случай»
                    - - **Оверинжиниринг**: 7 слоёв для Hello World
                     
                      - ---

                      ## 🎯 Проект модуля

                      Перепиши маркетплейс из Модуля 8 с чистой архитектурой:
                      - выдели domain-слой с entities, value objects, aggregate roots
                      - - repository интерфейсы в domain, реализации в infrastructure
                        - - HTTP- и gRPC-транспорты как отдельные адаптеры
                          - - use cases посредине
                           
                            - **Критерии приёмки:** domain не знает о PostgreSQL/HTTP/gRPC, замена базы требует только нового адаптера, use cases покрыты тестами без реальной БД.
                           
                            - ---

                            ## 🏁 Чек-пойнт модуля

                            - [ ] Понимаешь SOLID, но не навязываешь везде
                            - [ ] - [ ] Отличаешь Clean от Hexagonal и понимаешь, когда что уместно
                            - [ ] - [ ] Проектируешь domain с entities и value objects
                            - [ ] - [ ] Используешь functional options
                            - [ ] - [ ] Избегаешь оверинжиниринга
                           
                            - [ ] **Предыдущий:** [Модуль 9 ←](09-devops.md) · **Следующий:** [Модуль 11. Performance и production →](11-performance.md)
                            - [ ] 
