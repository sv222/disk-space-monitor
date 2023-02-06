package main

import (
    "bytes"
    "flag"
    "fmt"
    "net/http"
    "os"
    "syscall"
)

const threshold = 90
const telegramAPI = "https://api.telegram.org/bot<BOT_TOKEN>/sendMessage"
const chatID = "<CHAT_ID>"

func checkDiskUsage(path string) (int, error) {
    fs := syscall.Statfs_t{}
    err := syscall.Statfs(path, &fs)
    if err != nil {
        return 0, err
    }
    used := 100 * (fs.Blocks - fs.Bfree) / fs.Blocks
    return used, nil
}

func sendTelegramAlert(message string) error {
    body := []byte(fmt.Sprintf("chat_id=%s&text=%s", chatID, message))
    resp, err := http.Post(telegramAPI, "application/x-www-form-urlencoded", bytes.NewBuffer(body))
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("Unexpected status code: %d", resp.StatusCode)
    }
    return nil
}

func main() {
    var showHelp bool
    flag.BoolVar(&showHelp, "h", false, "show help")
    flag.BoolVar(&showHelp, "help", false, "show help")
    flag.Parse()

    if showHelp {
        fmt.Println("Usage: disk-space-monitor [OPTION]...")
        fmt.Println("Monitor disk usage and send an alert via Telegram if disk usage exceeds a specified threshold.")
        fmt.Println("")
        fmt.Println("Options:")
        fmt.Println("  -h, --help     show help")
        fmt.Println("")
        os.Exit(0)
    }

    path := "/"
    if len(os.Args) > 1 {
        path = os.Args[1]
    }
    used, err := checkDiskUsage(path)
    if err != nil {
        fmt.Fprintln(os.Stderr, "Error:", err)
        os.Exit(1)
    }
    if used > threshold {
        message := fmt.Sprintf("Disk usage on %s exceeds threshold of %d%%\n", path, threshold)
        err := sendTelegramAlert(message)
        if err != nil {
            fmt.Fprintln(os.Stderr, "Error:", err)
            os.Exit(1)
        }
        fmt.Println(message)
        os.Exit(1)
    }
    fmt.Printf("Disk usage on %s is below threshold of %d%%\n", path, threshold)
    os.Exit(0)
}
