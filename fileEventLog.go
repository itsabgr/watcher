package watcher

import (
	"fmt"
	"os"
)

type fileEventLog struct {
	file *os.File
}

func (f fileEventLog) Emit(event Event) error {
	var str string
	switch event.(type) {
	case EthereumFinEvent:
		str = ethereumTxToString(event.(EthereumFinEvent).Tx())
	case EthereumTxEvent:
		str = ethereumTxToString(event.(EthereumTxEvent).Tx())
	default:
		return fmt.Errorf("unsupported event %s:%s", event.Net(), event.Kind())
	}
	_, err := f.file.WriteString(str + "\n")
	return err
}

func (f fileEventLog) Close() error {
	return f.Close()
}

func NewFileEventLog(file *os.File) (EventLog, error) {
	o := &fileEventLog{}
	o.file = file
	return o, nil
}
