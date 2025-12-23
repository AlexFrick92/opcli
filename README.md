# opcli

Interactive console client for OPC UA servers, similar to psql for PostgreSQL.

## Features

- **Interactive shell** - connect to OPC UA server and execute commands
- **Browse** - explore nodes tree structure
- **Read/Write** - read and write node values
- **Methods** - call server methods
- **Subscriptions** - subscribe to node value changes

## Usage

```bash
opcli connect opc.tcp://localhost:4840
opcli> browse
opcli> read ns=2;i=1000
opcli> write ns=2;i=1000 42
opcli> subscribe ns=2;i=1000
opcli> exit
```