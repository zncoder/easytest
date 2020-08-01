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

func (tt T) Nil(err error) {
	if err == nil {
		return
	}
	tt.Helper()
	tt.fatal(fmt.Sprintf("UNEXPECTED ERR:%v", err))
}

func (tt T) True(ok bool) {
	if ok {
		return
	}
	tt.Helper()
	tt.fatal("FALSE")
}

func (tt T) DeepEqual(want, got interface{}) {
	if reflect.DeepEqual(want, got) {
		return
	}
	tt.Helper()
	tt.fatal(fmt.Sprintf("NOT EQUAL: want:%v got:%v", want, got))
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
