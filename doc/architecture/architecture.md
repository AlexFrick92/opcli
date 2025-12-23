# Architecture

## Stack

- **Language**: Go
- **OPC UA**: gopcua library

## Components

1. **CLI Parser** - parse and route commands
2. **OPC UA Client** - connection and session management
3. **Command Handlers** - execute operations (browse, read, write, methods, subscriptions)
4. **Output Formatter** - display results in console

## Command Flow

```
User Input -> Parser -> Handler -> OPC UA Client -> Server
                                         ↓
User Output <- Formatter <- Response <-
```