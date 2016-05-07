package bitlog

import "fmt"

type StdWriter struct  {
}

func (self StdWriter) Write(p []byte) (n int, err error)  {
	return fmt.Println(string(p))
}
