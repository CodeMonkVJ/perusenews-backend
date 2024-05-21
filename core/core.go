package core 

import (
    "fmt"
    "log"
    "os"
    "os/exec"
)

func init() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <URL>")
        os.Exit(1)
    }

    url := os.Args[1]

    // Check if Python script exists
    pythonScript := "./fetch_and_prettify.py"
    if _, err := os.Stat(pythonScript); os.IsNotExist(err) {
        log.Fatalf("Python script not found: %s", pythonScript)
    }

    // Execute the Python script with the URL as an argument
    cmd := exec.Command("python3", pythonScript, url)

    // Capture the output of the Python script
    output, err := cmd.Output()
    if err != nil {
        log.Fatalf("Failed to execute Python script: %s", err)
    }

    err2 := os.WriteFile("./script_output.txt", output, 0644)
    if err2 != nil {
        log.Fatalf("Failed: %s", err2)
    }
    fmt.Println("done")
}
