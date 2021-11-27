package glog

type NullLogger struct {}

func (n NullLogger) Error(error) {}
func (n NullLogger) Debugf(string, ...interface{}) {}
func (n NullLogger) Noticef(string, ...interface{}) {}