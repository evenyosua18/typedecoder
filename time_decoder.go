package typedecoder

import (
	"fmt"
	"reflect"
	"time"
)

func TimeDecoder(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	switch inVal.Kind() {
	case reflect.Int64:
		// unix to time.Time
		t := time.Unix(inVal.Int(), 0)
		out.Set(reflect.ValueOf(t))
	case reflect.String:
		// string to time.Time
		if t, err := time.Parse(timeFormat, inVal.String()); err == nil {
			out.Set(reflect.ValueOf(t))
		}
	default:
		err = fmt.Errorf("got %s type, expected %s", inVal.Kind(), out.Type())
	}

	return
}
