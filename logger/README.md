# Context aware logger


```go
// helper function configures the logging.
func setupLogging(c config.Config) {
	logger.Default = logger.Logrus(
		logrus.NewEntry(
			logrus.StandardLogger(),
		),
	)
	if c.Logs.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if c.Logs.Trace {
		logrus.SetLevel(logrus.TraceLevel)
	}
	if c.Logs.Pretty == false {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}
```