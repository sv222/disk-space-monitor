package main

import (
    "encoding/json"
    "fmt"
	"log"
    "net/http"
    "os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/jessevdk/go-flags"
)

type Configuration struct {
    Token string
    ChatID string
}

func readConfig() Configuration {
    file, _ := os.Open("config.json")
    defer file.Close()
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        log.Fatalf("Error reading configuration: %s", err)
    }
    return configuration
}

var token, chatID string

type options struct {
	Interval int    `short:"i" long:"interval" default:"60" description:"Interval in seconds for checking disk usage"`
	Path     string `short:"p" long:"path" default:"/" description:"Path to monitor disk usage"`
	Threshold int    `short:"t" long:"threshold" default:"90" description:"Threshold for disk usage percentage"`
}

func main() {

    config := readConfig()

    token = config.Token
    chatID = config.ChatID

	var opts options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("Error parsing options: %s", err)
	}

	for {
		output, err := exec.Command("df", "-h", opts.Path).Output()
		if err != nil {
			log.Fatalf("Error running command 'df': %s", err)
		}

		lines := strings.Split(string(output), "\n")
		if len(lines) < 2 {
			log.Fatalf("Unexpected output from 'df'")
		}

		fields := strings.Fields(lines[1])
		if len(fields) < 5 {
			log.Fatalf("Unexpected output from 'df'")
		}

		usage, err := strconv.Atoi(strings.TrimRight(fields[4], "%"))
		if err != nil {
			log.Fatalf("Error parsing disk usage: %s", err)
		}

		if usage >= opts.Threshold {
			fmt.Printf("Disk usage exceeded threshold of %d%%\n", opts.Threshold)
			// Send alert via Telegram
			sendTelegramAlert(opts.Threshold, usage)
		}

		time.Sleep(time.Duration(opts.Interval) * time.Second)
	}
}

func sendTelegramAlert(threshold int, usage int) {
    msg := fmt.Sprintf("Disk usage on server has exceeded the threshold of %d%% (current usage: %d%%)", threshold, usage)

    apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s", token, chatID, msg)
    _, err := http.Get(apiURL)
    if err != nil {
        log.Printf("Error sending Telegram message: %s", err)
    }
}
