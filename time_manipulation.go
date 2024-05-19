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
		return in.(time.Time).Format(timeFormat), nil
	}

	return nil, errors.New("invalid type of time when trying to manipulate input")
}

func TimeToUnixManipulation(in interface{}) (interface{}, error) {
	// get value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	if inVal.Type().String() == TimeDataType {
		return in.(time.Time).Unix(), nil
	}

	return nil, errors.New("invalid type of time when trying to manipulate input")
}
