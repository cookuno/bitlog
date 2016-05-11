package bitlog

import (
	"testing"
	"fmt"
	"github.com/cookuno/bitlog/writer"
)

func StdWriterTest(t *testing.T)  {
	log := New()
	textLineFormatter := NewTextLineFormatter()
	stdWriter := writer.NewStdWriter(false)
	log.AddOutput(LEVEL_DEBUG, textLineFormatter, stdWriter)
	log.Attr(AttrMap{"1":1, "x":false})
	fmt.Println("xxxxx")

}
