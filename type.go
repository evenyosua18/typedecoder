package typedecoder

type (
	// Decoder main structure
	Decoder struct {
		Output interface{}
		Errors []error

		UseManipulation bool

		listManipulationFunction map[string]ManipulationFunc
	}

	ManipulationFunc func(in interface{}) (interface{}, error)
)

var (
	listManipulationFunc map[string]ManipulationFunc
)
