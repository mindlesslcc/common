package logs

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
	"sync"
	"time"
)

const (
	DEFAULT_LOG_LEVEL   = DEBUG
	DEFAULT_CHANNEL_LEN = 20
)

type fileLogWriter struct {
	msgChan    chan *logMsg
	signalChan chan string
	wg         sync.WaitGroup
	logMsgPool *sync.Pool

	mu sync.Mutex
	// The opened file
	filename   string `json:"filename"`
	fileWriter *os.File

	level int `json:"level"`
}

type logMsg struct {
	level int
	msg   string
	when  time.Time
}

func newFileWriter() Logger {
	w := &fileLogWriter{
		fileWriter: nil,
		filename:   "log",
		level:      DEFAULT_LOG_LEVEL,
	}
	w.msgChan = make(chan *logMsg, DEFAULT_CHANNEL_LEN)
	w.signalChan = make(chan string)
	w.logMsgPool = &sync.Pool{
		New: func() interface{} {
			return &logMsg{}
		},
	}
	return w
}

func (w *fileLogWriter) Init(filename string) error {
	w.filename = filename
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		w.fileWriter, err = os.Create(filename)
		if err != nil {
			return err
		}
	}

	w.fileWriter, err = os.Open(filename)
	if err != nil {
		return err
	}

	w.wg.Add(1)
	go w.startLogger()
	return err
}

func (w *fileLogWriter) startLogger() error {
	exit := false
	for {
		select {
		case bm := <-w.msgChan:
			w.fileWriter.WriteString(fmt.Sprintf("%s:%d:%s\n", bm.when.String(), bm.level, bm.msg))
			w.logMsgPool.Put(bm)
		case sg := <-w.signalChan:
			w.Flush()
			if sg == "close" {
				w.Destroy()
				exit = true
			}
			w.wg.Done()
		}
		if exit {
			break
		}
	}
	return nil
}

func (w *fileLogWriter) WriteMsg(logLevel int, msg string) error {
	when := time.Now()
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg = "[" + filename + ":" + strconv.FormatInt(int64(line), 10) + "]" + msg
	lm := w.logMsgPool.Get().(*logMsg)
	lm.level = logLevel
	lm.msg = msg
	lm.when = when
	w.msgChan <- lm
	return nil
}

func (w *fileLogWriter) Destroy() {
	w.signalChan <- "close"
	w.wg.Wait()
	close(w.msgChan)
	close(w.signalChan)
	w.fileWriter.Close()
}

func (w *fileLogWriter) Flush() {
	for {
		if len(w.msgChan) > 0 {
			bm := <-w.msgChan
			w.fileWriter.WriteString(fmt.Sprintf("%s:%d:%s\n", bm.when.String(), bm.level, bm.msg))
			w.logMsgPool.Put(bm)
			continue
		}
		break
	}
	w.fileWriter.Sync()
}

func init() {
	Register("file", newFileWriter)
}
