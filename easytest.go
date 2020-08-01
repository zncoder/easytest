package easytest

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

type T struct {
	*testing.T
}

func New(t *testing.T) (tt T) {
	return T{T: t}
}

func (tt T) Nil(err error, fmtArgs ...interface{}) {
	if err == nil {
		return
	}
	tt.Helper()
	tt.fatal(fmt.Sprintf("UNEXPECTED ERR:%v %s", err, formatArgs(fmtArgs)))
}

func formatArgs(fmtArgs []interface{}) string {
	if len(fmtArgs) == 0 {
		return ""
	}
	s, ok := fmtArgs[0].(string)
	if !ok {
		log.Fatalf("first arg of fmtArgs must be the format string: %v", fmtArgs)
	}
	return fmt.Sprintf(s, fmtArgs[1:]...)
}

func (tt T) True(ok bool, fmtArgs ...interface{}) {
	if ok {
		return
	}
	tt.Helper()
	tt.fatal("FALSE " + formatArgs(fmtArgs))
}

func (tt T) DeepEqual(a, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}
	tt.Helper()
	tt.fatal(fmt.Sprintf("NOT EQUAL: a:%v b:%v", a, b))
}

func (tt T) Logf(format string, args ...interface{}) {
	tt.Helper()
	if testing.Verbose() {
		log.Output(2, fmt.Sprintf(format, args...))
	} else {
		tt.T.Logf(format, args...)
	}
}

func (tt T) fatal(s string) {
	tt.Helper()
	// log failure inline
	if testing.Verbose() {
		log.Output(3, s)
	}
	tt.T.Fatal(s)
}

func (tt T) Run(name string, f func(T)) bool {
	return tt.T.Run(name, func(t *testing.T) {
		f(New(t))
	})
}

// NewDir creates a tempdir.
func (tt T) NewDir() string {
	dir, err := ioutil.TempDir("", "testdir")
	tt.Nil(err)
	return dir
}

func (tt T) RemoveDir(dir string) {
	err := os.RemoveAll(dir)
	tt.Nil(err)
}
