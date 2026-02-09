# opcli

Interactive console client for OPC UA servers, similar to psql for PostgreSQL or ping command.

## Features

- **Interactive shell** - connect to OPC UA server and execute commands
- **Quick connect** - pass IP address as argument to connect automatically
- **Single connection** - supports one active connection at a time

## Usage

Quick connect with IP address:
```bash
opcli 10.10.10.95
```

Connect with full endpoint:
```bash
opcli connect opc.tcp://localhost:4840
opcli> disconnect
opcli> exit
```

## Limitations

- Only one active connection is supported at a time
- Connecting to a new server automatically disconnects from the previous one



- this is madness