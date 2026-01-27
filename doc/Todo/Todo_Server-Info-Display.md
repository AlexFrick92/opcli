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

## Summary
