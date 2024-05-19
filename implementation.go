package typedecoder

import "reflect"

func Decode(in, out interface{}) error {
	decoder := Decoder{
		Output:          out,
		UseManipulation: true,

		listManipulationFunction: listManipulationFunc,
	}

	return decoder.Decode(in)
}

func DecodeWithoutManipulation(in, out interface{}) error {
	decoder := Decoder{
		Output:          out,
		UseManipulation: false,

		listManipulationFunction: listManipulationFunc,
	}

	return decoder.Decode(in)
}

func NewDecoder() *Decoder {
	return new(Decoder)
}

func (d *Decoder) Decode(in interface{}) error {
	return d.decode(in, reflect.ValueOf(d.Output).Elem())
}

// SetManipulation Function
func (d *Decoder) SetManipulation(name string, function ManipulationFunc) *Decoder {
	if d.listManipulationFunction == nil {
		d.listManipulationFunction = make(map[string]ManipulationFunc)
	}

	d.listManipulationFunction[name] = function
	d.UseManipulation = true

	return d
}

func (d *Decoder) Out(out interface{}) *Decoder {
	d.Output = out

	return d
}

// AddManipulationFunction use for configuration when program started
func AddManipulationFunction(name string, function ManipulationFunc) {
	// create if empty
	if listManipulationFunc == nil {
		listManipulationFunc = make(map[string]ManipulationFunc)
	}

	// add to map
	listManipulationFunc[name] = function
}
