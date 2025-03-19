package main

import (
    "bufio"
    "encoding/json"
    "fmt"
    "log"
    "net"
    "os"
    "os/exec"
    "time"
)

// QMP Command structure
type QMPCommand struct {
    Execute   string                 `json:"execute"`
    Arguments map[string]interface{} `json:"arguments,omitempty"`
}

func main() {
    // Start QEMU with QMP enabled
    qemuCmd := exec.Command("qemu-system-x86_64",
        "-m", "512M",
        "-smp", "2",
        "-drive", "file=alpine.qcow2,format=qcow2",
        "-qmp", "unix:/tmp/qmp-sock,server,nowait",
        "-nographic",
    )

    err := qemuCmd.Start()
    if err != nil {
        log.Fatalf("Failed to start QEMU: %v", err)
    }
    log.Println("QEMU started with PID", qemuCmd.Process.Pid)

    // Wait for QEMU QMP socket to be ready
    time.Sleep(3 * time.Second)

    // Connect to QMP
    conn, err := net.Dial("unix", "/tmp/qmp-sock")
    if err != nil {
        log.Fatalf("Failed to connect to QMP: %v", err)
    }
    defer conn.Close()

    reader := bufio.NewReader(conn)
    // Read QMP greeting
    greeting, _ := reader.ReadString('\n')
    fmt.Println("QMP Greeting:", greeting)

    // Enable QMP capabilities
    enableQMP := QMPCommand{Execute: "qmp_capabilities"}
    sendQMP(conn, enableQMP, reader)

    // Get VM status
    queryStatus := QMPCommand{Execute: "query-status"}
    sendQMP(conn, queryStatus, reader)

    // Wait a bit (simulate work)
    time.Sleep(5 * time.Second)

    // Shutdown VM gracefully
    shutdownCmd := QMPCommand{Execute: "system_powerdown"}
    sendQMP(conn, shutdownCmd, reader)

    log.Println("Sent shutdown command, waiting for VM to exit...")

    qemuCmd.Wait()
    log.Println("QEMU VM exited")
}

func sendQMP(conn net.Conn, cmd QMPCommand, reader *bufio.Reader) {
    data, _ := json.Marshal(cmd)
    conn.Write(append(data, '\n'))
    // Read QMP response
    res, _ := reader.ReadString('\n')
    fmt.Println("QMP Response:", res)
}
