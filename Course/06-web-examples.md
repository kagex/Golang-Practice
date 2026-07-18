# Web и API в Go — примеры кода от нуля до продвинутого

> Дополнение к [Модулю 6: Web и API](./06-web-api.md)

Запускаемые примеры по всему стеку: от `http.ListenAndServe` до JWT, WebSocket, gRPC и OpenAPI. Каждый блок — самодостаточный файл, который можно скопировать и запустить.

---

## Содержание

1. [Уровень 0 — Первый HTTP-сервер](#уровень-0--первый-http-сервер)
2. [Уровень 1 — Роутинг и middleware](#уровень-1--роутинг-и-middleware)
3. [Уровень 2 — REST API с chi](#уровень-2--rest-api-с-chi)
4. [Уровень 3 — Валидация и ошибки](#уровень-3--валидация-и-ошибки)
5. [Уровень 4 — JWT-аутентификация](#уровень-4--jwt-аутентификация)
6. [Уровень 5 — WebSocket](#уровень-5--websocket)
7. [Уровень 6 — gRPC](#уровень-6--grpc)
8. [Уровень 7 — Production-сервер](#уровень-7--production-сервер)

---

## Уровень 0 — Первый HTTP-сервер

### Пример 0.1 — Hello World

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Go!")
	})

	fmt.Println("Сервер запущен: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
```

> ⚠️ `http.ListenAndServe` без таймаутов — только для экспериментов. В продакшене используй явный `http.Server{}`.

---

### Пример 0.2 — Читаем параметры запроса

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		// Query параметр: /greet?name=Alice
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "незнакомец"
		}

		// Метод запроса
		fmt.Fprintf(w, "Привет, %s! Метод: %s\n", name, r.Method)
	})

	http.ListenAndServe(":8080", nil)
}
```

---

### Пример 0.3 — JSON ответ

```go
package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(Response{
		Message: "Всё хорошо",
		Time:    time.Now(),
	})
}

func main() {
	http.HandleFunc("/api/status", jsonHandler)
	http.ListenAndServe(":8080", nil)
}
```

---

## Уровень 1 — Роутинг и middleware

### Пример 1.1 — ServeMux Go 1.22+ с path params

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Go 1.22+: метод + path params в фигурных скобках
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id") // извлекаем path param
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"id": id, "name": "Alice"})
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, `{"status":"created"}`)
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.WriteHeader(http.StatusNoContent)
		_ = id
	})

	fmt.Println("Сервер: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
```

---

### Пример 1.2 — Middleware: цепочка хендлеров

```go
package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// Тип для middleware-функций
type Middleware func(http.Handler) http.Handler

// Chain применяет middleware в порядке слева направо
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

// Logger логирует метод, путь и время выполнения
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"duration", time.Since(start),
		)
	})
}

// Recover ловит panic и возвращает 500
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Error("panic recovered", "err", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// RequestID добавляет уникальный ID запроса в заголовок
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := fmt.Sprintf("%d", time.Now().UnixNano())
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello!")
	})
	mux.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("намеренная паника для теста!")
	})

	// Цепочка: Recover → Logger → RequestID → handler
	handler := Chain(mux, Recover, Logger, RequestID)

	http.ListenAndServe(":8080", handler)
}
```

---

## Уровень 2 — REST API с chi

### Пример 2.1 — CRUD API на chi

```go
// go get github.com/go-chi/chi/v5
package main

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// In-memory хранилище
type Store struct {
	mu    sync.RWMutex
	users map[string]User
}

func NewStore() *Store {
	return &Store{users: map[string]User{
		"1": {ID: "1", Name: "Alice", Age: 30},
		"2": {ID: "2", Name: "Bob", Age: 25},
	}}
}

func (s *Store) List() []User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]User, 0, len(s.users))
	for _, u := range s.users {
		list = append(list, u)
	}
	return list
}

func (s *Store) Get(id string) (User, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	u, ok := s.users[id]
	return u, ok
}

func (s *Store) Create(u User) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.users[u.ID] = u
}

func (s *Store) Delete(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.users[id]; !ok {
		return false
	}
	delete(s.users, id)
	return true
}

// writeJSON — хелпер для JSON-ответов
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func main() {
	store := NewStore()
	r := chi.NewRouter()

	// Встроенные middleware chi
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, store.List())
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var u User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				http.Error(w, "invalid JSON", http.StatusBadRequest)
				return
			}
			store.Create(u)
			writeJSON(w, http.StatusCreated, u)
		})

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				u, ok := store.Get(id)
				if !ok {
					http.Error(w, "not found", http.StatusNotFound)
					return
				}
				writeJSON(w, http.StatusOK, u)
			})

			r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
				id := chi.URLParam(r, "id")
				if !store.Delete(id) {
					http.Error(w, "not found", http.StatusNotFound)
					return
				}
				w.WriteHeader(http.StatusNoContent)
			})
		})
	})

	http.ListenAndServe(":8080", r)
}
```

---
## Уровень 3 — Валидация и ошибки

### Пример 3.1 — Структурированные ошибки API

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// APIError — стандартный формат ошибки
type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

var (
	ErrNotFound   = &APIError{Code: 404, Message: "resource not found"}
	ErrBadRequest = &APIError{Code: 400, Message: "bad request"}
	ErrInternal   = &APIError{Code: 500, Message: "internal server error"}
)

// errorHandler — централизованная обработка ошибок
func errorHandler(w http.ResponseWriter, err error) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(apiErr.Code)
		json.NewEncoder(w).Encode(apiErr)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(ErrInternal)
}

func getUser(id string) (map[string]any, error) {
	if id == "" {
		return nil, &APIError{Code: 400, Message: "id is required"}
	}
	if id != "42" {
		return nil, ErrNotFound
	}
	return map[string]any{"id": id, "name": "Alice"}, nil
}

func main() {
	http.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		user, err := getUser(id)
		if err != nil {
			errorHandler(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	})

	http.ListenAndServe(":8080", nil)
}
```

---

### Пример 3.2 — Валидация через go-playground/validator

```go
// go get github.com/go-playground/validator/v10
package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type RegisterRequest struct {
	Username string `json:"username" validate:"required,min=3,max=20,alphanum"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Age      int    `json:"age"      validate:"gte=18,lte=120"`
}

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func formatValidationErrors(err error) []ValidationError {
	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		return nil
	}

	result := make([]ValidationError, 0, len(errs))
	for _, e := range errs {
		result = append(result, ValidationError{
			Field:   e.Field(),
			Message: fmt.Sprintf("failed on '%s'", e.Tag()),
		})
	}
	return result
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(req); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]any{
			"errors": formatValidationErrors(err),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}
```

---

## Уровень 4 — JWT-аутентификация

### Пример 4.1 — Полный Auth: регистрация, логин, защищённый маршрут

```go
// go get github.com/golang-jwt/jwt/v5
// go get golang.org/x/crypto
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("super-secret-key-change-in-prod")

// User — модель пользователя
type User struct {
	ID       int
	Username string
	Password string // bcrypt hash
}

// Claims — payload JWT-токена
type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Простая in-memory "база"
var users = map[string]*User{}
var nextID = 1

// generateToken создаёт JWT с экспирацией 15 минут
func generateToken(user *User) (string, error) {
	claims := &Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// parseToken валидирует JWT и возвращает claims
func parseToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// authMiddleware проверяет Bearer-токен в заголовке Authorization
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := parseToken(tokenStr)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// Добавляем username в контекст через заголовок (упрощение)
		r.Header.Set("X-User", claims.Username)
		next.ServeHTTP(w, r)
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	if _, exists := users[body.Username]; exists {
		http.Error(w, "user already exists", http.StatusConflict)
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	user := &User{ID: nextID, Username: body.Username, Password: string(hash)}
	nextID++
	users[body.Username] = user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "registered"})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&body)

	user, ok := users[body.Username]
	if !ok {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := generateToken(user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"access_token": token})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("X-User")
	json.NewEncoder(w).Encode(map[string]string{
		"username": username,
		"message":  "это защищённый маршрут",
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/register", registerHandler)
	mux.HandleFunc("POST /auth/login", loginHandler)

	// Защищённый маршрут
	protected := http.NewServeMux()
	protected.HandleFunc("GET /profile", profileHandler)
	mux.Handle("/profile", authMiddleware(protected))

	fmt.Println("Auth API: http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
```

---
## Уровень 5 — WebSocket

### Пример 5.1 — Echo WebSocket (nhooyr.io/websocket)

```go
// go get nhooyr.io/websocket
package main

import (
	"context"
	"fmt"
	"net/http"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Принимаем WebSocket-соединение
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // только для dev!
	})
	if err != nil {
		return
	}
	defer conn.CloseNow()

	ctx := r.Context()
	for {
		var msg map[string]any
		if err := wsjson.Read(ctx, conn, &msg); err != nil {
			break
		}
		fmt.Printf("Получено: %v\n", msg)
		// Отправляем обратно
		if err := wsjson.Write(ctx, conn, msg); err != nil {
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", echoHandler)
	// HTML-клиент для теста
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, wsHTML)
	})
	fmt.Println("WebSocket echo: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

const wsHTML = `<!DOCTYPE html>
<html>
<body>
<input id="msg" placeholder="Введи сообщение">
<button onclick="send()">Отправить</button>
<pre id="log"></pre>
<script>
const ws = new WebSocket("ws://localhost:8080/ws");
ws.onmessage = e => document.getElementById("log").textContent += "Echo: " + e.data + "\n";
function send() {
  ws.send(JSON.stringify({text: document.getElementById("msg").value}));
}
</script>
</body></html>`
```

---

### Пример 5.2 — Чат: Hub + Broadcast

```go
// go get nhooyr.io/websocket
package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type Message struct {
	User string `json:"user"`
	Text string `json:"text"`
}

// Hub — центральный хаб чата
type Hub struct {
	mu      sync.RWMutex
	clients map[string]*websocket.Conn
}

func NewHub() *Hub {
	return &Hub{clients: make(map[string]*websocket.Conn)}
}

func (h *Hub) Register(id string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clients[id] = conn
}

func (h *Hub) Unregister(id string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	delete(h.clients, id)
}

func (h *Hub) Broadcast(ctx context.Context, msg Message) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	for id, conn := range h.clients {
		if err := wsjson.Write(ctx, conn, msg); err != nil {
			slog.Error("broadcast error", "client", id, "err", err)
		}
	}
}

var hub = NewHub()
var clientCounter int
var counterMu sync.Mutex

func chatHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return
	}

	counterMu.Lock()
	clientCounter++
	id := fmt.Sprintf("user-%d", clientCounter)
	counterMu.Unlock()

	hub.Register(id, conn)
	defer func() {
		hub.Unregister(id)
		conn.CloseNow()
		slog.Info("клиент отключился", "id", id)
	}()

	slog.Info("новый клиент", "id", id)
	hub.Broadcast(r.Context(), Message{User: "system", Text: id + " вошёл в чат"})

	for {
		var msg Message
		if err := wsjson.Read(r.Context(), conn, &msg); err != nil {
			break
		}
		msg.User = id
		hub.Broadcast(r.Context(), msg)
	}
}

func main() {
	http.HandleFunc("/ws/chat", chatHandler)
	fmt.Println("Чат: ws://localhost:8080/ws/chat")
	http.ListenAndServe(":8080", nil)
}
```

---

## Уровень 6 — gRPC

### Пример 6.1 — Unary gRPC сервис

Структура проекта:
```
grpc-example/
  proto/user.proto
  server/main.go
  client/main.go
```

**proto/user.proto:**
```protobuf
syntax = "proto3";
package user;
option go_package = "./proto";

message GetUserRequest { string id = 1; }
message GetUserResponse {
  string id   = 1;
  string name = 2;
  int32  age  = 3;
}
message ListUsersRequest  {}
message ListUsersResponse { repeated GetUserResponse users = 1; }

service UserService {
  rpc GetUser    (GetUserRequest)    returns (GetUserResponse);
  rpc ListUsers  (ListUsersRequest)  returns (stream GetUserResponse);
}
```

Генерация кода:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/user.proto
```

**server/main.go:**
```go
// go get google.golang.org/grpc
// go get google.golang.org/protobuf
package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

var fakeDB = map[string]*pb.GetUserResponse{
	"1": {Id: "1", Name: "Alice", Age: 30},
	"2": {Id: "2", Name: "Bob", Age: 25},
	"3": {Id: "3", Name: "Carol", Age: 28},
}

// GetUser — unary RPC
func (s *server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, ok := fakeDB[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "user %s not found", req.Id)
	}
	return user, nil
}

// ListUsers — server-side streaming RPC
func (s *server) ListUsers(_ *pb.ListUsersRequest, stream pb.UserService_ListUsersServer) error {
	for _, user := range fakeDB {
		if err := stream.Send(user); err != nil {
			return err
		}
		time.Sleep(100 * time.Millisecond) // имитируем задержку
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Interceptor для логирования
	logInterceptor := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		start := time.Now()
		resp, err := handler(ctx, req)
		fmt.Printf("[gRPC] %s — %v\n", info.FullMethod, time.Since(start))
		return resp, err
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(logInterceptor))
	pb.RegisterUserServiceServer(s, &server{})

	fmt.Println("gRPC сервер: :50051")
	s.Serve(lis)
}
```

**client/main.go:**
```go
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "grpc-example/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Unary: получить одного пользователя
	user, err := client.GetUser(ctx, &pb.GetUserRequest{Id: "1"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Пользователь: %s, возраст: %d\n", user.Name, user.Age)

	// Streaming: получить всех
	fmt.Println("\nВсе пользователи:")
	stream, err := client.ListUsers(ctx, &pb.ListUsersRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		u, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  - %s (id=%s, age=%d)\n", u.Name, u.Id, u.Age)
	}
}
```

---
## Уровень 7 — Production-сервер

### Пример 7.1 — HTTP-сервер с таймаутами, graceful shutdown и health-check

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// HealthResponse — ответ health-check эндпоинта
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
	Uptime  string `json:"uptime"`
}

var startTime = time.Now()

func buildRouter(version string) http.Handler {
	mux := http.NewServeMux()

	// Liveness probe — жив ли процесс
	mux.HandleFunc("GET /livez", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	// Readiness probe — готов ли принимать трафик
	mux.HandleFunc("GET /readyz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(HealthResponse{
			Status:  "ready",
			Version: version,
			Uptime:  time.Since(startTime).Round(time.Second).String(),
		})
	})

	// Основное API
	mux.HandleFunc("GET /api/v1/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.InfoContext(r.Context(), "hello request",
			"remote", r.RemoteAddr,
			"user-agent", r.UserAgent(),
		)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello from production Go!",
			"version": version,
		})
	})

	return mux
}

func main() {
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "dev"
	}

	// Настраиваем structured logging
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})))

	server := &http.Server{
		Addr:    ":8080",
		Handler: buildRouter(version),

		// Защита от медленных клиентов
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       120 * time.Second,

		// Кастомный error log через slog
		ErrorLog: slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
	}

	// Запускаем сервер в горутине
	go func() {
		slog.Info("server starting", "addr", server.Addr, "version", version)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("server error", "err", err)
			os.Exit(1)
		}
	}()

	// Ждём сигнала остановки (SIGINT/SIGTERM от Kubernetes или Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit

	slog.Info("shutdown signal received", "signal", sig.String())

	// Даём 30 секунд завершить инфлайт-запросы
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("shutdown error", "err", err)
		os.Exit(1)
	}

	slog.Info("server stopped gracefully",
		"uptime", time.Since(startTime).Round(time.Second),
	)
}
```

---

### Пример 7.2 — Middleware: HTTP-клиент с retry и timeout

```go
package main

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"time"
)

// RetryTransport — RoundTripper с экспоненциальным backoff
type RetryTransport struct {
	base       http.RoundTripper
	maxRetries int
}

func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var resp *http.Response
	var err error

	for attempt := 0; attempt <= t.maxRetries; attempt++ {
		if attempt > 0 {
			// Экспоненциальный backoff: 100ms, 200ms, 400ms...
			wait := time.Duration(math.Pow(2, float64(attempt-1))*100) * time.Millisecond
			slog.Info("retrying request", "attempt", attempt, "wait", wait)
			time.Sleep(wait)
		}

		resp, err = t.base.RoundTrip(req.Clone(req.Context()))
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		if resp != nil {
			resp.Body.Close()
		}
	}
	return resp, err
}

// NewHTTPClient создаёт клиент с таймаутами и retry
func NewHTTPClient(timeout time.Duration, maxRetries int) *http.Client {
	return &http.Client{
		Timeout: timeout,
		Transport: &RetryTransport{
			base:       http.DefaultTransport,
			maxRetries: maxRetries,
		},
	}
}

func fetchJSON(ctx context.Context, client *http.Client, url string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Статус: %d\n", resp.StatusCode)
	return nil
}

func main() {
	client := NewHTTPClient(10*time.Second, 3)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := fetchJSON(ctx, client, "https://httpbin.org/get"); err != nil {
		slog.Error("request failed", "err", err)
		return
	}
	fmt.Println("Запрос выполнен успешно")
}
```

---

## Шпаргалка: выбор инструмента

| Задача | Инструмент |
|---|---|
| Простой REST API (stdlib) | `http.NewServeMux` + `http.Server` |
| REST с группировкой маршрутов | `chi` |
| REST с максимальным количеством фич | `gin` или `echo` |
| Межсервисное общение | `gRPC + protobuf` |
| Реалтайм / чат | `WebSocket (nhooyr.io/websocket)` |
| Аутентификация | `JWT (golang-jwt/jwt/v5)` + `bcrypt` |
| Валидация входных данных | `go-playground/validator/v10` |
| Документация API | `swaggo/swag` или `oapi-codegen` |
| HTTP-клиент в продакшене | `http.Client` с таймаутами + retry |
| Таймаут всего запроса | `context.WithTimeout` |
| Graceful shutdown | `server.Shutdown(ctx)` + `signal.Notify` |
| Health checks | `/livez` (liveness) + `/readyz` (readiness) |

---

## Что изучить дальше

- [Модуль 6: Web и API — теория и уроки](./06-web-api.md)
- [Модуль 7: Testing — тестирование HTTP-сервисов](./07-testing.md)
- [Модуль 5: Databases — подключение PostgreSQL и Redis](./05-databases.md)
- [Модуль 3: Concurrency — горутины и каналы внутри HTTP-хендлеров](./03-concurrency-examples.md)

