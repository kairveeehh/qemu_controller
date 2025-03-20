

This project demonstrates how to programmatically control a QEMU virtual machine using Go and QEMU Machine Protocol (QMP). The Go application launches a QEMU VM, establishes a QMP connection, sends control commands (e.g., start, query status, shutdown), and handles responses.

## what it does 

- Launch a QEMU virtual machine
- Connect to QEMU using QMP
- Query VM status
- Gracefully shutdown the VM

## Local Setup

### Clone the Repository

```bash
git clone https://github.com/kairveeehh/qemu-go-control.git
cd qemu-go-control
go mod init github.com/kairveeehh/qemu-go-control
go mod tidy
```

### Download a Minimal VM Image

Download Alpine Linux minimal image for testing:

```bash
wget https://dl-cdn.alpinelinux.org/alpine/v3.19/releases/x86_64/alpine-virt-3.19.1-x86_64.qcow2 -O alpine.qcow2
```

### Project Structure

```
qemu-go-control/
├── main.go
├── alpine.qcow2
├── go.mod
└── go.sum
```

### Build and Run

```bash
go build -o qemu-controller
./qemu-controller
```

## Output screenshot 
![alt text](<Screenshot from 2025-03-19 23-33-45.png>)

## Further Use Case Extension

This mock implementation can be extended to:

- Dynamically create VM instances
- Manage multiple VMs concurrently
- Integrate into a Lima plugin or gRPC service
- Monitor VM resource usage in real-time
