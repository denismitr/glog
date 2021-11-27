# Convenient logger interface and wrapper around std logger

### Interface
```go
type Logger interface {
	Error(err error)
	Debugf(format string, args ...interface{})
	Noticef(format string, args ...interface{})
}
```

### Usage with std logger 
Ideal to use with **Kubernates** where you usually write logs to stdout and stderr.
method `Error` will write to stderr and methods `Debugf` and `Noticef` will write
to stdout. `Debugf` writes only in `glog.Dev` environment and `Error` and `Noticef` both
in `glog.Dev` and `glog.Prod`.

```go
lg := glog.NewStdoutLogger(glog.Dev, "my-service")

lg.Error(errors.New("some error"))
lg.Debugf("Something to debug %s", "foo")
lg.Noticef("Something to debug %s", "foo")
```

### Null logger
```go
lg := glog.NullLogger{}

// nothing will be written
lg.Error(errors.New("some error"))
lg.Debugf("Something to debug %s", "foo")
lg.Noticef("Something to debug %s", "foo")
```