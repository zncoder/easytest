package easytest

import (
	"fmt"
	"log"
	"reflect"
	"testing"
)

type T struct {
	*testing.T
}

func New(t *testing.T) (tt T) {
	return T{T: t}
}

func (tt T) Nil(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}
	tt.Helper()
	tt.fatal(fmt.Sprintf("UNEXPECTED ERR:%v: %s", err, fmt.Sprintf(format, args...)))
}

func (tt T) True(ok bool, format string, args ...interface{}) {
	if ok {
		return
	}
	tt.Helper()
	tt.fatal(fmt.Sprintf("FALSE: %s", fmt.Sprintf(format, args...)))
}

func (tt T) DeepEqual(want, got interface{}, format string, args ...interface{}) {
	if reflect.DeepEqual(want, got) {
		return
	}
	tt.Helper()
	tt.fatal(fmt.Sprintf("NOT EQUAL: want:%v got:%v: %s", want, got, fmt.Sprintf(format, args...)))
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
