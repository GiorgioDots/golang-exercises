package intermediate

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Log(msg string) error
}

type ConsoleLogger struct {
}

func (l ConsoleLogger) Log(msg string) error {
	_, err := fmt.Println(msg)
	return err
}

type FileLogger struct {
	filePath string
}

func (l *FileLogger) Log(msg string) error {
	f, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	formatted := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339), msg)
	_, err = f.WriteString(formatted)
	if err != nil {
		return err
	}
	return nil
}

func ExecuteCustomLogger() {
	consoleLogger := ConsoleLogger{}
	err := consoleLogger.Log("Hello from console")
	if err != nil {
		panic(err)
	}

	fileLogger := FileLogger{
		filePath: "./log.csv",
	}
	err = fileLogger.Log("Hello from file")
	if err != nil {
		consoleLogger.Log(err.Error())
	}
}
