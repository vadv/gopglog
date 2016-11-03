package parser

import (
	"fmt"
	"regexp"
)

// level regexp
var (
	LevelRegexp       = regexp.MustCompile(`(LOG|DETAIL)\:`)
	UnusedLevelRegexp = regexp.MustCompile(`(WARNING|ERROR|FATAL|PANIC|HINT|STATEMENT|CONTEXT|LOCATION|DEBUG\d)\:`)
)

// timestamp regexp
var (
	TimeRegexpError       = fmt.Errorf("Unknown timestamp format")
	TimeTimestampRegexp   = regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})(?: [A-Z\+\-\d]{3,6})?`)
	TimeTimestampLayout   = "2006-01-02 15:04:05 MST"
	TimeEpochRegexp       = regexp.MustCompile(`(\d{10})\.(\d{3})`)
	TimeTimestampMsLayout = "2006-01-02 15:04:05.000 MST"
	TimeTimestampMsRegexp = regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2})\.\d+(?: [A-Z\+\-\d]{3,6})?`)
)

// workerId regexp
var (
	WorkerIdRegexpError = fmt.Errorf("Unknown worker id (pid) format")
	WorkerIdRegexp      = regexp.MustCompile(`\[(\d+)\]\:`)
)

// DBName regexp
var (
	DbNameRegexpError = fmt.Errorf("Unknown db name format")
	DbNameRegexp      = regexp.MustCompile(`(db\=|database\=)([0-9a-zA-Z\_\[\]\-\.]*)`)
)

// UserName regexp
var (
	UserNameRegexpError = fmt.Errorf("Unknown user name format")
	UserNameRegexp      = regexp.MustCompile(`(user\=|username\=)([0-9a-zA-Z\_\[\]\-\.]*)`)
)

// Type and param regexp
var (
	LogTypeSqlRegexp    = regexp.MustCompile(`LOG\: .* (\w+)( <unnamed>| \w+)?\: (.*)$`)
	DetailTypeSqlRegexp = regexp.MustCompile(`DETAIL\: .* (\w+)( <unnamed>| \w+)?\: (.*)$`)
	SplitParamRegexp    = regexp.MustCompile(`\$(\d+) = '(.*?)'`)
	TypeSqlRegexpError  = fmt.Errorf("Can't parse type of sql")
	ParamRegexpError    = fmt.Errorf("Parameters regexp error")
)
