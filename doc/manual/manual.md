# Manual

## Getting started

### Quick connect

To connect to an OPC UA server, pass its IP address as argument:

    opcli <ip-address>

The client will connect to `opc.tcp://<ip-address>:4840` (default OPC UA port).

**Example:**

    opcli 10.10.10.95
    → Connecting to opc.tcp://10.10.10.95:4840...