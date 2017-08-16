package logs

import (
	"testing"
)

func TestFile1(t *testing.T) {
	log := newFileWriter()
	log.Init("log")
	log.WriteMsg(DEBUG, "debug")
	log.WriteMsg(WARN, "warn")
	log.WriteMsg(ERROR, "error")
	log.WriteMsg(INFO, "info")
	/*	f, err := os.Open("test.log")
		if err != nil {
			t.Fatal(err)
		}
		b := bufio.NewReader(f)
		lineNum := 0
		for {
			line, _, err := b.ReadLine()
			if err != nil {
				break
			}
			if len(line) > 0 {
				lineNum++
			}
		}
		var expected = LevelDebug + 1
		if lineNum != expected {
			t.Fatal(lineNum, "not "+strconv.Itoa(expected)+" lines")
		}
		os.Remove("test.log")*/
}
