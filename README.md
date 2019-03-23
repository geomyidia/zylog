# zylog

*A simple wrapper for customized logrus usage*

## Usage

Here's an example of setting up global logging for use by your app in your
app's `logging` package, based on configuration pulled in by Viper (from either
a config file or ENV variables):


```go
package logging

import (
	cfg "github.com/spf13/viper"
	log "github.com/zylisp/zylog/zylog"
)

func init() {
	log.SetupLogging(&log.ZyLogOptions{
		Colored:      cfg.GetBool("logging.colored"),
		Level:        cfg.GetString("logging.level"),
		Output:       cfg.GetString("logging.output"),
		ReportCaller: cfg.GetBool("logging.report-caller"),
	})
}
```


## License

© 2019, ZYLISP Project

Apache License, Version 2.0
