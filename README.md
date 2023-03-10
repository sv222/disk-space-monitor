# Disk Space Monitor

A CLI tool written in Golang for monitoring disk usage on a server and sending alerts if usage exceeds a specified threshold.

## Features

- Monitor disk usage on a specified path (defaults to `/`)
- Send alert via Telegram if disk usage exceeds the specified threshold
- Specify threshold and interval for checking disk usage via command line options
- Run as a daemon, checking disk usage at specified intervals

## Usage

Usage:

```shell
disk-space-monitor [options]
```

Options:

```shell
-i int
```

Interval in seconds for checking disk usage (default 60)

```shell
-p string
```

Path to monitor disk usage (default "/")

```shell
-t int
```

Threshold for disk usage percentage (default 90)

## Example of Usage

```shell
disk-space-monitor -t 80 -i 300
```
This will run the disk space monitor with a threshold of 80% and an interval of 300 seconds (5 minutes).

## Installation

1. Clone the repository:

```shell
git clone https://github.com/<username>/disk-space-monitor.git
```
2. Build the binary:

```shell
go build -o disk-space-monitor main.go
```

3. Install the binary and make it executable:

```shell
sudo mv disk-space-monitor /usr/local/bin \
sudo chmod +x /usr/local/bin/disk-space-monitor
```

## Here is an example of the config.json file:

```json
{
    "Token": "your-bot-token",
    "ChatID": "your-chat-id"
}
```

## Requirements

- Golang 1.13 or higher
- Access to a Telegram bot API key

## License

This project is licensed under the MIT License.


