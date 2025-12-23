# Go Code Style Guide

## Именование файлов и папок

- **Файлы**: lowercase, snake_case для многословных имён
  - `main.go`
  - `opc_client.go`
  - `command_handler.go`

- **Папки/Пакеты**: lowercase, одно слово (избегать snake_case)
  - `cmd`, `client`, `handlers`, `formatter`
  - Если необходимо составное имя: слитно, например `opcuaclient`

## Именование в коде

### Переменные и функции

- **Экспортируемые** (публичные): PascalCase
  ```go
  func ConnectToServer() {}
  var DefaultTimeout int
  ```

- **Неэкспортируемые** (приватные): camelCase
  ```go
  func parseCommand() {}
  var connectionTimeout int
  ```

- **Константы**: PascalCase или SCREAMING_SNAKE_CASE для группы констант
  ```go
  const MaxRetries = 3
  const (
      STATUS_CONNECTED = "connected"
      STATUS_DISCONNECTED = "disconnected"
  )
  ```

### Структуры и интерфейсы

- **Имена типов**: PascalCase (существительные)
  ```go
  type OPCClient struct {}
  type CommandHandler struct {}
  ```

- **Интерфейсы**: обычно с суффиксом `-er`
  ```go
  type Reader interface {}
  type Writer interface {}
  type CommandExecutor interface {}
  ```

### Методы

- **Receiver**: короткая аббревиатура (1-2 буквы), lowercase
  ```go
  func (c *OPCClient) Connect() {}
  func (ch *CommandHandler) Execute() {}
  ```

## Общие принципы

1. **Краткость**: Go предпочитает короткие, ясные имена
   - Локальные переменные: `i`, `err`, `buf`
   - Чем уже область видимости, тем короче имя

2. **Избегать stuttering**: не дублировать имя пакета
   - ✅ `client.New()` вместо ❌ `client.NewClient()`
   - ✅ `handler.Execute()` вместо ❌ `handler.ExecuteHandler()`

3. **Getters без префикса Get**:
   - ✅ `client.Status()` вместо ❌ `client.GetStatus()`
   - Setters с префиксом Set: `client.SetStatus()`

4. **Акронимы пишутся одинаково**:
   - ✅ `OPCUA`, `HTTPServer`, `urlParser`
   - ❌ `OpcUa`, `HttpServer`

5. **Имена пакетов**:
   - Одно слово, lowercase
   - Должны быть существительными, описывающими содержимое
   - Избегать `util`, `common`, `base`
