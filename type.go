package typedecoder

import "reflect"

type (
	// Decoder main structure
	Decoder struct {
		Output interface{}
		Errors []error

		UseManipulation bool

		listManipulationFunction map[string]ManipulationFunc
		listDecodeFunction       map[string]DecodeFunc
	}

	ManipulationFunc func(in interface{}) (interface{}, error)
	DecodeFunc       func(in interface{}, out reflect.Value) error
)

var (
	listManipulationFunc map[string]ManipulationFunc
	listDecodeFunc       map[string]DecodeFunc
)
