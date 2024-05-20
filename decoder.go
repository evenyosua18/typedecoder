package typedecoder

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

// decode will be call recursively
func (d *Decoder) decode(in interface{}, out reflect.Value) (err error) {
	// get value for input
	var v reflect.Value

	// make sure type of input is not nil
	if in != nil {
		v = reflect.ValueOf(in)

		if v.Kind() == reflect.Ptr && v.IsNil() {
			in = nil
		}
	}

	// if input is nil or value is invalid, set zero output
	if in == nil || !v.IsValid() {
		out.Set(reflect.Zero(out.Type()))
		return
	}

	// check manipulation function
	if d.UseManipulation && d.listManipulationFunction != nil {
		inVal := reflect.Indirect(reflect.ValueOf(in))

		if d.listManipulationFunction[inVal.Type().String()] != nil {
			if in, err = d.listManipulationFunction[inVal.Type().String()](in); err != nil {
				d.AddError(err)
			}
		}
	}

	// decode based on kind of value
	switch d.parseKind(out.Kind()) {
	case reflect.Bool:
		err = d.decodeBool(in, out)
	case reflect.String:
		err = d.decodeString(in, out)
	case reflect.Int:
		err = d.decodeInt(in, out)
	case reflect.Uint:
		err = d.decodeUint(in, out)
	case reflect.Float32:
		err = d.decodeFloat(in, out)
	case reflect.Struct:
		err = d.decodeStruct(in, out)
	case reflect.Interface:
		err = d.decodeInterface(in, out)
	case reflect.Ptr:
		err = d.decodePtr(in, out)
	case reflect.Map:
		err = d.decodeMap(in, out)
	case reflect.Array:
		err = d.decodeArray(in, out)
	case reflect.Slice:
		err = d.decodeSlice(in, out)
	case reflect.Func:
		err = d.decodeFunc(in, out)

	default:
		return fmt.Errorf("unsupported type: %s", v.Kind())
	}

	if len(d.Errors) != 0 {
		return d.Errors[0]
	}

	return
}

// decodeBool 0 false else true
func (d *Decoder) decodeBool(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	switch inKind {
	case reflect.Bool:
		out.SetBool(inVal.Bool())
	case reflect.Int:
		out.SetBool(inVal.Int() != 0)
	case reflect.Uint:
		out.SetBool(inVal.Uint() != 0)
	case reflect.Float32:
		out.SetBool(inVal.Float() != 0)
	case reflect.String:
		if b, err := strconv.ParseBool(inVal.String()); err == nil {
			out.SetBool(b)
		} else if inVal.String() == "" {
			out.SetBool(false)
		} else {
			return fmt.Errorf("error when parsing bool: %s", err)
		}
	default:
		err = fmt.Errorf("got %s type, expected %s", inKind, out.Type())
	}

	return
}

func (d *Decoder) decodeString(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	switch inKind {
	case reflect.String:
		out.SetString(inVal.String())
	case reflect.Bool:
		if inVal.Bool() {
			out.SetString("1")
		} else {
			out.SetString("0")
		}
	case reflect.Int:
		out.SetString(strconv.FormatInt(inVal.Int(), 10))
	case reflect.Uint:
		out.SetString(strconv.FormatUint(inVal.Uint(), 10))
	case reflect.Float32:
		out.SetString(strconv.FormatFloat(inVal.Float(), 'f', -1, 64))
	default:
		err = fmt.Errorf("got %s type, expected %s", inKind, out.Type())
	}

	return
}

func (d *Decoder) decodeInt(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	switch inKind {
	case reflect.Int:
		out.SetInt(inVal.Int())
	case reflect.String:
		if inVal.String() == "" {
			inVal.SetString("0")
		}

		if res, err := strconv.ParseInt(inVal.String(), 0, out.Type().Bits()); err != nil {
			return fmt.Errorf("error when parsing int: %s", err)
		} else {
			out.SetInt(res)
		}
	case reflect.Bool:
		if inVal.Bool() {
			out.SetInt(1)
		} else {
			out.SetInt(0)
		}
	case reflect.Uint:
		out.SetInt(int64(inVal.Uint()))
	case reflect.Float32:
		out.SetInt(int64(inVal.Float()))
	default:
		err = fmt.Errorf("got %s type, expected %s", inKind, out.Type())
	}

	return
}

func (d *Decoder) decodeUint(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	switch inKind {
	case reflect.Uint:
		out.SetUint(inVal.Uint())
	case reflect.Int:
		if inVal.Int() < 0 {
			out.SetUint(0)
		} else {
			out.SetUint(uint64(inVal.Int()))
		}
	case reflect.String:
		if inVal.String() == "" {
			inVal.SetString("0")
		}

		if res, err := strconv.ParseUint(inVal.String(), 0, out.Type().Bits()); err != nil {
			return fmt.Errorf("error when parsing int: %s", err)
		} else {
			out.SetUint(res)
		}
	case reflect.Bool:
		if inVal.Bool() {
			out.SetUint(1)
		} else {
			out.SetUint(0)
		}
	case reflect.Float32:
		if inVal.Float() < 0 {
			out.SetUint(0)
		} else {
			out.SetUint(uint64(inVal.Float()))
		}
	default:
		err = fmt.Errorf("got %s type, expected %s", inKind, out.Type())
	}

	return
}

func (d *Decoder) decodeFloat(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	switch inKind {
	case reflect.Float32:
		out.SetFloat(inVal.Float())
	case reflect.Int:
		out.SetFloat(float64(inVal.Int()))
	case reflect.Uint:
		out.SetFloat(float64(inVal.Uint()))
	case reflect.String:
		if inVal.String() == "" {
			inVal.SetString("0")
		}

		if res, err := strconv.ParseFloat(inVal.String(), out.Type().Bits()); err != nil {
			return fmt.Errorf("error when parsing data from string type: %s", err)
		} else {
			out.SetFloat(res)
		}
	case reflect.Bool:
		if inVal.Bool() {
			out.SetFloat(1)
		} else {
			out.SetFloat(0)
		}
	}

	return
}

func (d *Decoder) decodeInterface(in interface{}, out reflect.Value) (err error) {
	// check if interface is pointer or not
	// if pointer then copy the value first
	if out.IsValid() && out.Elem().IsValid() {
		elem := out.Elem()

		if !elem.CanAddr() {
			c := reflect.New(elem.Type())
			c.Elem().Set(elem)
			elem = c
		}

		if err = Decode(in, elem); err != nil && !elem.CanAddr() {
			return err
		} else {
			out.Set(elem.Elem())
			return
		}
	}

	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	// check if input is pointer
	if inKind == reflect.Ptr && inVal.Type().Elem() == out.Type() {
		inVal = reflect.Indirect(inVal)
	}

	if !inVal.IsValid() {
		inVal = reflect.Zero(inVal.Type())
	}

	if !inVal.Type().AssignableTo(out.Type()) {
		return fmt.Errorf("got %s type, expected %s", inKind, out.Type())
	}

	out.Set(inVal)
	return
}

func (d *Decoder) decodePtr(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	// check if the input is nil
	isInputNil := in == nil

	if !isInputNil {
		switch inVal.Kind() {
		case reflect.Chan, reflect.Ptr, reflect.Func, reflect.Map, reflect.Slice, reflect.Interface:
			isInputNil = inVal.IsNil()
		}
	}

	if isInputNil {
		// if output is nil too
		if !out.IsNil() && out.CanSet() {
			out.Set(reflect.New(out.Type()).Elem())
		}

		return
	}

	// if input is not nil and the output is a pointer and have same data type with the input
	outType := out.Type()
	outElemType := outType.Elem()

	if out.Kind() == reflect.Ptr && inVal.Type() == outElemType {
		out.Set(reflect.ValueOf(in))
		return
	}

	// input value is not nil
	// create element if output is nil
	if out.CanSet() {
		realOut := out
		if realOut.IsNil() {
			realOut = reflect.New(outElemType)
		}

		err = d.decode(in, reflect.Indirect(realOut))

		out.Set(realOut)
	} else {
		err = d.decode(in, reflect.Indirect(out))
	}

	return
}

func (d *Decoder) decodeStruct(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	// have same type between input and output
	if inVal.Type() == out.Type() {
		out.Set(inVal)
		return
	}

	switch inKind {
	case reflect.Map:
		return d.decodeMapToStruct(inVal, out)
	case reflect.Struct:
		// create map to store input value
		mapVal := reflect.MakeMap(reflect.TypeOf((map[string]interface{})(nil)))

		if err = d.decodeStructToMap(inVal, mapVal); err != nil {
			return
		}

		// decode map from input value to struct
		if err = d.decodeMapToStruct(mapVal, out); err != nil {
			return
		}
	default:
		err = fmt.Errorf("got %s type, expected %s", inKind, out.Type())
	}

	return
}

func (d *Decoder) decodeFunc(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	// make sure input value has same type
	if inVal.Type() != out.Type() {
		return fmt.Errorf("got %s type, expected %s", inVal.Type(), out.Type())
	}

	// set value
	out.Set(inVal)

	return
}

func (d *Decoder) decodeMap(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))

	// if output is nil, create new map to the output value first
	if out.IsNil() {
		out.Set(reflect.MakeMap(reflect.MapOf(out.Type().Key(), out.Type().Elem())))
	}

	// check input
	switch inVal.Kind() {
	case reflect.Map:
		err = d.decodeMapToMap(inVal, out)
	case reflect.Struct:
		err = d.decodeStructToMap(inVal, out)
	default:
		err = fmt.Errorf("got %s type, expected %s", inVal.Kind(), out.Type())
	}

	return
}

func (d *Decoder) decodeArray(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	// if the input not array or slice
	switch inKind {
	case reflect.Int, reflect.Uint, reflect.Float32, reflect.String, reflect.Bool:
		return d.decodeArray([]interface{}{in}, out)
	case reflect.Array, reflect.Slice:
		valArray := reflect.New(reflect.ArrayOf(out.Type().Len(), out.Type().Elem())).Elem()

		for i := 0; i < inVal.Len(); i++ {
			//currentData := inVal.Index(i).Interface()
			currentField := valArray.Index(i)
			if err = d.decode(inVal.Index(i).Interface(), currentField); err != nil {
				return
			}
		}

		out.Set(valArray)
	default:
		return fmt.Errorf("got %s type, expected %s", inVal.Type(), out.Type())
	}

	return
}

func (d *Decoder) decodeSlice(in interface{}, out reflect.Value) (err error) {
	// get input value
	inVal := reflect.Indirect(reflect.ValueOf(in))
	inKind := d.parseKind(inVal.Kind())

	// if the input not array or slice
	switch inKind {
	case reflect.Int, reflect.Uint, reflect.Float32, reflect.String, reflect.Bool:
		return d.decodeSlice([]interface{}{in}, out)
	case reflect.Array, reflect.Slice:
		valSlice := out
		if valSlice.IsNil() {
			valSlice = reflect.MakeSlice(reflect.SliceOf(out.Type().Elem()), inVal.Len(), inVal.Len())
		}

		for i := 0; i < inVal.Len(); i++ {
			currentField := valSlice.Index(i)

			if err = d.decode(inVal.Index(i).Interface(), currentField); err != nil {
				return
			}
		}

		out.Set(valSlice)
	default:
		return fmt.Errorf("got %s type, expected %s", inVal.Type(), out.Type())
	}

	return
}

// decodeMapToStruct input should be map
func (d *Decoder) decodeMapToStruct(inVal, out reflect.Value) (err error) {
	// input validation
	inValType := inVal.Type()
	outValType := out.Type()

	if kind := inValType.Key().Kind(); kind != reflect.String && kind != reflect.Interface {
		return fmt.Errorf(
			"keys of map should be string, has '%s' keys",
			inValType.Key().Kind())
	}

	// loop output value
	for i := 0; i < out.NumField(); i++ {
		mapVal := inVal.MapIndex(reflect.ValueOf(outValType.Field(i).Name))

		if !mapVal.IsValid() {
			nilVal := reflect.New(out.Field(i).Type())
			out.Field(i).Set(reflect.Indirect(nilVal))
		} else if err = d.decode(mapVal.Interface(), out.Field(i)); err != nil {
			d.AddError(err)
			continue
		}
	}

	return
}

func (d *Decoder) decodeStructToMap(inVal, out reflect.Value) (err error) {
	inType := inVal.Type()

	for i := 0; i < inType.NumField(); i++ {
		// get field
		field := inType.Field(i)

		// get value
		value := inVal.Field(i)

		// check if type of value can be assigned to type of map value
		if !value.Type().AssignableTo(out.Type().Elem()) {
			return fmt.Errorf("got %s type, expected %s", value.Type(), out.Type())
		}

		// set value to map
		if value.Kind() == reflect.Struct {
			// copy value
			x := reflect.New(value.Type())
			x.Elem().Set(value)

			// embed struct, create new addressable map
			// this decode will hit decodeMap function to decode struct to map
			vMap := reflect.MakeMap(reflect.MapOf(out.Type().Key(), out.Type().Elem()))
			addrVal := reflect.New(vMap.Type())
			reflect.Indirect(addrVal).Set(vMap)

			if err = d.decode(x.Interface(), addrVal); err != nil {
				return err
			} else {
				// add to map
				out.SetMapIndex(reflect.ValueOf(field.Name), reflect.Indirect(addrVal))
			}
		} else {
			// add to map
			out.SetMapIndex(reflect.ValueOf(field.Name), value)
		}
	}

	return
}

func (d *Decoder) decodeMapToMap(inVal, out reflect.Value) (err error) {
	// store multiple errors
	var errs []error

	// if input is nil
	if inVal.Len() == 0 && inVal.IsNil() && !out.IsNil() {
		out.Set(inVal)
	}

	// loop maps
	for _, k := range inVal.MapKeys() {
		key := reflect.Indirect(reflect.New(out.Type().Key()))

		// decode key
		if err = d.decode(k.Interface(), key); err != nil {
			errs = append(errs, err)
			continue
		}

		// decode value
		value := reflect.Indirect(reflect.New(out.Type().Elem()))
		if err = d.decode(inVal.MapIndex(k).Interface(), value); err != nil {
			errs = append(errs, err)
			continue
		}

		// set to output
		out.SetMapIndex(key, value)
	}

	if len(errs) != 0 {
		errMsg := errs[0].Error()

		for i := 1; i < len(errs); i++ {
			errMsg += " | " + errs[i].Error()
		}

		return errors.New(errMsg)
	}

	return
}

/* helper */

func (d *Decoder) structToMap(in interface{}) (map[string]reflect.Value, error) {
	v := reflect.ValueOf(in)
	t := reflect.TypeOf(in)
	result := make(map[string]reflect.Value)

	for i := 0; i < v.NumField(); i++ {
		result[t.Field(i).Name] = v.Field(i)
	}

	return result, nil
}

func (d *Decoder) AddError(err error) {
	d.Errors = append(d.Errors, err)
}

func (d *Decoder) parseKind(kind reflect.Kind) reflect.Kind {
	if kind >= reflect.Int && kind <= reflect.Int64 {
		return reflect.Int
	} else if kind >= reflect.Uint && kind <= reflect.Uint64 {
		return reflect.Uint
	} else if kind >= reflect.Float32 && kind <= reflect.Float64 {
		return reflect.Float32
	} else {
		return kind
	}
}
