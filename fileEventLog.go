package watcher

import (
	"os"
)

type fileEventLog struct {
	file *os.File
}

func (f fileEventLog) Emit(event Event) error {
	_, err := f.file.Write(event.MarshalYAML())
	f.file.WriteString("\n")
	return err
}

func (f fileEventLog) Close() error {
	return nil
}

func NewFileEventLog(file *os.File) (EventLog, error) {
	o := &fileEventLog{}
	o.file = file
	return o, nil
}
