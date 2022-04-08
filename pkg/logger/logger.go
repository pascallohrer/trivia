package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileLogger struct {
	filename string
}

func NewFileLogger(path ...string) *FileLogger {
	filename := filepath.FromSlash("./log/trivia.log")
	if len(path) > 0 {
		filename = path[0]
	}
	os.MkdirAll(filepath.Dir(filename), os.ModeDir)
	return &FileLogger{
		filename: filename,
	}
}

func (l *FileLogger) Print(v ...interface{}) {
	f, _ := os.OpenFile(l.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	f.WriteString(time.Now().String() + ":\n")
	f.WriteString(fmt.Sprint(v...) + "\n\n")
	f.Close()
}

func (l *FileLogger) Printf(format string, v ...interface{}) {
	f, _ := os.OpenFile(l.filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	f.WriteString(time.Now().String() + ":\n")
	f.WriteString(fmt.Sprintf(format, v...) + "\n\n")
	f.Close()
}

func (l *FileLogger) Fatal(v ...interface{}) {
	l.Print(v...)
	os.Exit(1)
}

func (l *FileLogger) Fatalf(format string, v ...interface{}) {
	l.Printf(format, v...)
	os.Exit(1)
}
