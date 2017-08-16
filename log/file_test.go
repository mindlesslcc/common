package logs

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	log := newFileWriter()
	if log == nil {
		t.Fatal("new loggerwriter fail")
	}
	if err := log.Init("log"); err != nil {
		t.Fatal("write init error!")
	}
	if err := log.WriteMsg(DEBUG, "debug"); err != nil {
		t.Fatal("write debug error")
	}
	if err := log.WriteMsg(WARN, "warn"); err != nil {
		t.Fatal("write warn error")
	}
	if err := log.WriteMsg(ERROR, "error"); err != nil {
		t.Fatal("write error error")
	}
	if err := log.WriteMsg(INFO, "info"); err != nil {
		t.Fatal("write info error")
	}
	log.Flush()
	log.Destroy()

	f, err := os.Open("log")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	defer os.Remove("log")
}
