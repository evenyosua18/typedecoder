package typedecoder

import (
	"errors"
	"reflect"
	"time"
)

const (
	TimeDataType = "time.Time"

	timeFormat = "2006-01-02 15:04:05"
)

func TimeToStringManipulation(in interface{}) (interface{}, error) {
	// get value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	if inVal.Type().String() == TimeDataType {
		switch v := in.(type) {
		case time.Time:
			return v.Format(timeFormat), nil
		case *time.Time:
			if v != nil {
				return v.Format(timeFormat), nil
			}
		}
		return "", nil
	}

	return nil, errors.New("invalid type of time when trying to manipulate input")
}

func TimeToUnixManipulation(in interface{}) (interface{}, error) {
	// get value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	if inVal.Type().String() == TimeDataType {
		switch v := in.(type) {
		case time.Time:
			return v.Unix(), nil
		case *time.Time:
			if v != nil {
				return v.Unix(), nil
			}
		}
		return "", nil
	}

	return nil, errors.New("invalid type of time when trying to manipulate input")
}
