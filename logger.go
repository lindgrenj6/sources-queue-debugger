package main

import "log"

type Logger struct{}

func (l Logger) Debugf(msg string, args ...interface{}) {
	log.Printf(msg+"\n", args)
}

func (l Logger) Errorf(msg string, args ...interface{}) {
	log.Printf(msg+"\n", args)
}
