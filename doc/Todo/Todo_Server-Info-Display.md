# Server Info Display

## Overview
Добавить вывод информации о сервере после успешного подключения. После команды connect или при быстром подключении через аргумент IP, клиент должен показывать:
- Product Name
- Version
- Manufacturer
- Server State

## Step-by-step
1. Изучить API gopcua для получения Server Status и Build Info
2. Создать функцию GetServerInfo() в пакете client для получения информации о сервере
3. Добавить структуру ServerInfo для хранения данных (ProductName, Version, Manufacturer, State)
4. Интегрировать вызов GetServerInfo() в функцию Connect() после успешного подключения
5. Реализовать форматированный вывод информации о сервере в консоль
6. Протестировать на реальном сервере

## Progress

### Шаг 1: Изучить API gopcua ✓
Изучена документация и примеры gopcua v0.8.0:
- Node IDs: ServerStatus (i=2256), BuildInfo (i=2260), ServerState (i=2259)
- Структуры: ua.BuildInfo, ua.ServerStatusDataType
- Метод чтения: client.Read() с ReadRequest

**Extension Objects:**
- ua.BuildInfo - Go-структура для OPC UA типа BuildInfo
- Extension Object - контейнер для передачи сложных типов данных
- При чтении сложных структур требуется регистрация типа: ua.RegisterExtensionObject()
- Регистрация говорит библиотеке, как декодировать бинарные данные в Go-структуру

**Стандартное адресное пространство OPC UA:**
Каждый OPC UA сервер, соответствующий спецификации, обязан иметь стандартные узлы:
- Server Object (i=2253) - корневой объект
  - ServerStatus (i=2256) - переменная типа ServerStatusDataType
    - BuildInfo (i=2260) - информация о сборке
      - ProductName (i=2261)
      - ManufacturerName (i=2262)
      - SoftwareVersion (i=2263)
      - ProductURI (i=2295)
      - BuildNumber (i=2264)
      - BuildDate (i=2265)
    - ServerState (i=2259) - текущее состояние
    - CurrentTime (i=2258) - текущее время
    - StartTime (i=2257) - время запуска
  - ServerCapabilities (i=2268)
  - NamespaceArray (i=2255)

Префикс "i=" означает namespace 0 (стандартный namespace OPC UA).
Эти Node ID одинаковы на всех серверах и работают универсально.

**Варианты чтения:**
- Вариант 1: Читать ServerStatus (i=2256) целиком - нужна регистрация extension object
- Вариант 2: Читать BuildInfo (i=2260) и State (i=2259) отдельно - нужна регистрация
- Вариант 3 (выбран): Читать простые поля напрямую - не требует регистрации
  * ProductName: i=2261
  * ManufacturerName: i=2262
  * SoftwareVersion: i=2263
  * ServerState: i=2259

## Summary
