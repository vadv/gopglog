package parser

import (
	"fmt"
	"strconv"
	"time"
)

type Line struct {
	multiLinePart bool // if true this Line is a part of another message, usable field is sqlPart only
	isPrepare     bool // if true this statement if prepare
	beginTime     time.Time
	workerId      string
	dbName        string
	userName      string
	level         string // level: LOG|DETAIL ...
	sqlType       string // statement|prepare|parse|bind|execute|parameters
	stmtName      string
	parameters    map[int]string
	sqlPart       []byte
}

func ParseLine(source []byte) (*Line, error) {

	result, messageHeader := &Line{}, source
	if len(messageHeader) > 150 {
		messageHeader = messageHeader[:150]
	}
	levels := LevelRegexp.FindSubmatch(messageHeader)

	if len(levels) == 2 {

		// set level
		result.level = string(levels[1])

		// parse time
		if times := TimeTimestampRegexp.FindSubmatch(messageHeader); len(times) > 0 {
			// like "2006-01-02 15:04:05 MST"
			if resultTime, err := time.Parse(TimeTimestampLayout, string(times[0])); err == nil {
				result.beginTime = resultTime
			} else {
				return nil, err
			}
		} else {
			if times := TimeTimestampMsRegexp.FindSubmatch(messageHeader); len(times) > 0 {
				// like "2006-01-02 15:04:05.000 MST"
				if resultTime, err := time.Parse(TimeTimestampMsLayout, string(times[0])); err == nil {
					result.beginTime = resultTime
				} else {
					return nil, err
				}
			} else {
				// like "9876543210.000"
				if times := TimeEpochRegexp.FindSubmatch(messageHeader); len(times) == 3 {
					sec, _ := strconv.ParseInt(string(times[1]), 0, 64)
					nsec, _ := strconv.ParseInt(string(times[2]), 0, 64)
					result.beginTime = time.Unix(sec, nsec)
				} else {
					return nil, TimeRegexpError
				}
			}
		}

		// parse pid
		if data := WorkerIdRegexp.FindSubmatch(messageHeader); len(data) == 2 {
			result.workerId = string(data[1])
		} else {
			return nil, WorkerIdRegexpError
		}

		// parse db
		if data := DbNameRegexp.FindSubmatch(messageHeader); len(data) == 3 {
			result.dbName = string(data[2])
		} else {
			return nil, DbNameRegexpError
		}

		// parse username
		if data := UserNameRegexp.FindSubmatch(messageHeader); len(data) == 3 {
			result.userName = string(data[2])
		} else {
			return nil, UserNameRegexpError
		}

		// parse type
		if data := LogTypeSqlRegexp.FindSubmatch(source); len(data) == 4 {
			result.sqlType = string(data[1])
			result.stmtName = string(data[2])
			result.sqlPart = data[3]
		} else {

			if data := DetailTypeSqlRegexp.FindSubmatch(source); len(data) == 4 {
				result.sqlType = string(data[1])
				result.stmtName = string(data[2])
				// parse parameters: $1 = 'xxx', $2 = 'yyy'
				result.parameters = make(map[int]string, 0)
				for _, params := range SplitParamRegexp.FindAllStringSubmatch(string(data[3]), -1) {
					idx, _ := strconv.Atoi(params[1])
					result.parameters[idx] = params[2]
				}
			} else {
				// found LOG,DETAIL in messageHeader, but can't have valid regexp
				return nil, TypeSqlRegexpError
			}
		}

	} else {
		if UnusedLevelRegexp.Match(messageHeader) {
			return nil, nil
		} else {
			// is multiline part
			result.multiLinePart = true
			// all message is part of sql
			result.sqlPart = source
		}
		// else ignore UnusedLevel
	}
	return result, nil
}

func (l *Line) ToString() string {
	if l.multiLinePart {
		return fmt.Sprintf("type=multiline: %s", l.sqlPart)
	} else {
		return fmt.Sprintf("type=%s: date='%v' pid=%s,db=%s,user=%s,params=%v %s", l.sqlType, l.beginTime, l.workerId, l.dbName, l.userName, l.parameters, l.sqlPart)
	}
}
