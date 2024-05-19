package typedecoder

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestDecoder_Boolean(t *testing.T) {
	var result, expected bool

	// value: false (bool) | expected: false (bool)
	expected = false
	if err := Decode(false, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %t want %t", result, expected)
	}

	// value: true (bool) | expected: false (bool)
	expected = true
	if err := Decode(true, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %t want %t", result, expected)
	}

	// value: 1 (int) | expected: true (bool)
	expected = true
	if err := Decode(1, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %t want %t", result, expected)
	}

	// value: 100 (int) | expected: true (bool)
	expected = true
	if err := Decode(100, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %t want %t", result, expected)
	}

	// value: 0 (int) | expected: false (bool)
	expected = false
	if err := Decode(0, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %t want %t", result, expected)
	}

	// value: "true" (string) | expected: true (bool)
	expected = true
	if err := Decode("true", &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %t want %t", result, expected)
	}

	// value: "FALSE" (string) | expected: false (bool)
	expected = false
	if err := Decode("FALSE", &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %t want %t", result, expected)
	}
}

func TestDecoder_String(t *testing.T) {
	var result, expected string

	// value: "test" (string) | expected: "test"
	expected = "TEST"
	if err := Decode("TEST", &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %s want %s", result, expected)
	}

	// value: 10 (int) | expected: "10"
	expected = "10"
	if err := Decode(10, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %s want %s", result, expected)
	}

	// value: 100.123 (float32) | expected: "100.123"
	expected = "100.123"
	if err := Decode(100.123, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %s want %s", result, expected)
	}

	// value:
}

func TestDecoder_Int(t *testing.T) {
	var result, expected int

	// value: 10 (int) | expected: 10 (int)
	expected = 10
	if err := Decode(10, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: true (bool) | expected: 1 (int)
	expected = 1
	if err := Decode(true, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: false (bool) | expected: 0 (int)
	expected = 0
	if err := Decode(false, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: "10" (string) | expected: 10 (int)
	expected = 10
	if err := Decode("10", &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: 3.14 (float32) | expected: 3 (int)
	expected = 3
	if err := Decode(3.14, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: 10 (uint) | expected: 10 (int)
	expected = 10
	if err := Decode(uint(10), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}
}

// TestDecoder_DifferentInt to test simple case between different type of int (int32 to int, int64 to int32, etc)
func TestDecoder_DifferentInt(t *testing.T) {
	var result32, expected32 int32

	// value: 10 (int) | expected: 10 (int32)
	expected32 = 10
	if err := Decode(10, &result32); err != nil {
		t.Error(err)
	} else if result32 != expected32 {
		t.Errorf("got %d want %d", result32, expected32)
	}

	// value: 10 (int64) | expected: 10 (int32)
	expected32 = 10
	if err := Decode(int64(10), &result32); err != nil {
		t.Error(err)
	} else if result32 != expected32 {
		t.Errorf("got %d want %d", result32, expected32)
	}

	var result64, expected64 int64

	// value: 10 (int) | expected: 10 (int64)
	expected64 = 10
	if err := Decode(10, &result64); err != nil {
		t.Error(err)
	} else if result64 != expected64 {
		t.Errorf("got %d want %d", result64, expected64)
	}

	// value: 10 (int32) | expected: 10 (int64)
	expected64 = 10
	if err := Decode(int32(10), &result64); err != nil {
		t.Error(err)
	} else if result64 != expected64 {
		t.Errorf("got %d want %d", result64, expected64)
	}

	var result, expected int

	// value: 10 (int64) | expected: 10 (int)
	expected = 10
	if err := Decode(int64(10), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: 10 (int32) | expected: 10 (int)
	expected = 10
	if err := Decode(int32(10), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}
}

func TestDecoder_Uint(t *testing.T) {
	var result, expected uint

	// value: 10 (uint) | expected: 10 (uint)
	expected = 10
	if err := Decode(uint(10), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: true (bool) | expected: 1 (uint)
	expected = 1
	if err := Decode(true, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: false (bool) | expected: 0 (int)
	expected = 0
	if err := Decode(false, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: "10" (string) | expected: 10 (uint)
	expected = 10
	if err := Decode("10", &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: 3.14 (float32) | expected: 3 (uint)
	expected = 3
	if err := Decode(3.14, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: -3.14 (float32) | expected: 0 (uint)
	expected = 0
	if err := Decode(-3.14, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: 10 (int) | expected: 10 (uint)
	expected = 10
	if err := Decode(10, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: -10 (int) | expected: 0 (uint)
	expected = 0
	if err := Decode(-10, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}
}

// TestDecoder_DifferentInt to test simple case between different type of int (int32 to int, int64 to int32, etc)
func TestDecoder_DifferentUint(t *testing.T) {
	var result32, expected32 uint32

	// value: 10 (uint) | expected: 10 (uint32)
	expected32 = 10
	if err := Decode(uint(10), &result32); err != nil {
		t.Error(err)
	} else if result32 != expected32 {
		t.Errorf("got %d want %d", result32, expected32)
	}

	// value: 10 (uint64) | expected: 10 (uint32)
	expected32 = 10
	if err := Decode(uint64(10), &result32); err != nil {
		t.Error(err)
	} else if result32 != expected32 {
		t.Errorf("got %d want %d", result32, expected32)
	}

	var result64, expected64 uint64

	// value: 10 (uint) | expected: 10 (uint64)
	expected64 = 10
	if err := Decode(uint(10), &result64); err != nil {
		t.Error(err)
	} else if result64 != expected64 {
		t.Errorf("got %d want %d", result64, expected64)
	}

	// value: 10 (uint32) | expected: 10 (uint64)
	expected64 = 10
	if err := Decode(uint32(10), &result64); err != nil {
		t.Error(err)
	} else if result64 != expected64 {
		t.Errorf("got %d want %d", result64, expected64)
	}

	var result, expected int

	// value: 10 (uint64) | expected: 10 (uint)
	expected = 10
	if err := Decode(uint64(10), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}

	// value: 10 (uint32) | expected: 10 (uint)
	expected = 10
	if err := Decode(uint32(10), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %d want %d", result, expected)
	}
}

func TestDecoder_Float(t *testing.T) {
	var result, expected float32

	// value: 3.14 (float32) | expected: 3.14 (float32)
	expected = 3.14
	if err := Decode(float32(3.14), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %f want %f", result, expected)
	}

	// value: 10 (int) | expected: 10 (float32)
	expected = 10
	if err := Decode(10, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %f want %f", result, expected)
	}

	// value: 10 (uint) | expected: 10 (float32)
	expected = 10
	if err := Decode(uint(10), &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %f want %f", result, expected)
	}

	// value: "3.14" (string) | expected: 3.14 (float32)
	expected = 3.14
	if err := Decode("3.14", &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %f want %f", result, expected)
	}

	// value: true (bool) | expected: 1 (float32)
	expected = 1
	if err := Decode(true, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %f want %f", result, expected)
	}

	// value: false (bool) | expected: 0 (float32)
	expected = 0
	if err := Decode(false, &result); err != nil {
		t.Error(err)
	} else if result != expected {
		t.Errorf("got %f want %f", result, expected)
	}
}

// TestDecoder_DifferentFloat to test simple case between different type of int (float32 to float, float64 to float32, etc)
func TestDecoder_DifferentFloat(t *testing.T) {
	var result32, expected32 float32

	// value: 3.14 (float64) | expected: 3.14 (float32)
	expected32 = 3.14
	if err := Decode(3.14, &result32); err != nil {
		t.Error(err)
	} else if result32 != expected32 {
		t.Errorf("got %f want %f", result32, expected32)
	}

	var result64, expected64 float64

	// value: 3.14 (float32) | expected: 3.14 (float64)
	// special case (float32 to float64 will be added some decimal numbers)
	// in this case will become 3.140000104904175, so this test only compare two number after comma
	expected64 = 3.14
	if err := Decode(float32(3.14), &result64); err != nil {
		t.Error(err)
	} else if fmt.Sprintf("%.2f", result64) != fmt.Sprintf("%.2f", expected64) {
		t.Errorf("got %.2f want %.2f", result64, expected64)
	}
}

func TestDecoder_Interface(t *testing.T) {
	var expected interface{}

	// value: 10 (int) | expected: 10 (int)
	{
		var result int
		expected = 10
		if err := Decode(10, &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// value: "10" (string) | expected: "10" (string)
	{
		var result string
		expected = "10"
		if err := Decode("10", &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// value: 3.14 (float32) | expected: 3.14 (float32)
	{
		var result float32
		expected = float32(3.14)
		if err := Decode(float32(3.14), &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// value: 3.14 (float64) | expected: 3.14 (float64)
	{
		var result float64
		expected = 3.14
		if err := Decode(3.14, &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// value: true (bool) | expected: true (bool)
	{
		var result bool
		expected = true
		if err := Decode(true, &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// value: 10 (uint) | expected: 10 (uint)
	{
		var result uint
		expected = uint(10)
		if err := Decode(uint(10), &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// value: 10 (int) give the address value | expected: 10 (int)
	{
		var result int
		expected = 10
		input := 10
		if err := Decode(&input, &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}
}

func TestDecoder_Pointer(t *testing.T) {
	var expected interface{}

	// value: nil | expected: nil

	// value: 10 (int) | expected: 10 (int, pointer interface)
	{
		var result interface{}
		resultPtr := &result
		expected = 10

		if err := Decode(10, &resultPtr); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", resultPtr, reflect.TypeOf(resultPtr), expected, reflect.TypeOf(expected))
		}
	}

	// value: "TEST" (string) | expected: "TEST" (string, pointer interface)
	{
		var result interface{}
		resultPtr := &result
		expected = "TEST"

		if err := Decode("TEST", &resultPtr); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", resultPtr, reflect.TypeOf(resultPtr), expected, reflect.TypeOf(expected))
		}
	}

	// value: 3.14 (float32) | expected: 3.14 (float32, interface)
	{
		var result interface{}
		resultPtr := &result
		expected = float32(3.14)

		if err := Decode(float32(3.14), &resultPtr); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", resultPtr, reflect.TypeOf(resultPtr), expected, reflect.TypeOf(expected))
		}
	}

	// value: 3.14 (float64) | expected: 3.14 (float64, interface)
	{
		var result interface{}
		resultPtr := &result
		expected = 3.14

		if err := Decode(3.14, &resultPtr); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", resultPtr, reflect.TypeOf(resultPtr), expected, reflect.TypeOf(expected))
		}
	}

	// value: 10 (uint) | expected: 10 (uint, interface)
	{
		var result interface{}
		resultPtr := &result
		expected = uint(10)

		if err := Decode(uint(10), &resultPtr); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("got %v with type %s want %v with type %s", resultPtr, reflect.TypeOf(resultPtr), expected, reflect.TypeOf(expected))
		}
	}
}

func TestDecoder_MapToMap(t *testing.T) {
	// val: map[string]string | expected: map[string]string
	{
		var result map[string]string
		expected := map[string]string{
			"TEST":  "TEST",
			"TEST2": "TEST2",
		}
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else if expected["TEST"] != result["TEST"] && expected["TEST2"] != result["TEST2"] {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// val: map[int]int | expected: map[int]int
	{
		var result map[int]int
		expected := map[int]int{
			1: 1,
			2: 2,
		}
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else if expected[1] != result[1] || expected[2] != result[2] {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}

	// val: map[int]int | expected: map[string]string
	{
		var result map[string]string
		expected := map[string]string{
			"1": "1",
			"2": "3",
		}
		if err := Decode(map[int]int{
			1: 1,
			2: 3,
		}, &result); err != nil {
			t.Error(err)
		} else if expected["1"] != result["1"] || expected["2"] != result["2"] {
			t.Errorf("got %v with type %s want %v with type %s", result, reflect.TypeOf(result), expected, reflect.TypeOf(expected))
		}
	}
}

func TestDecoder_StructToMap(t *testing.T) {
	// val: struct (embed) | expected: map
	{
		var result map[string]interface{}
		expected := map[string]interface{}{
			"Name":     "TEST",
			"Age":      0,
			"NickName": "TEST",
			"Phone": map[string]interface{}{
				"Code":        "+62",
				"PhoneNumber": "123",
			},
		}

		if err := Decode(struct {
			Name     string
			NickName string
			Age      int
			Phone    struct {
				Code        string
				PhoneNumber string
			}
		}{
			Name:     "TEST",
			NickName: "TEST",
			Age:      0,
			Phone: struct {
				Code        string
				PhoneNumber string
			}{Code: "+62", PhoneNumber: "123"},
		}, &result); err != nil {
			t.Error(err)
		} else {
			for key, val := range expected {
				if reflect.ValueOf(expected[key]).Kind() == reflect.Map {
					for key2, _ := range val.(map[string]interface{}) {
						if expected[key2] != result[key2] {
							t.Errorf("got %v with type %s want %v with type %s", result[key2], reflect.TypeOf(result[key2]), expected[key2], reflect.TypeOf(expected[key2]))
							break
						}
					}
				} else {
					if expected[key] != result[key] {
						t.Errorf("got %v with type %s want %v with type %s", result[key], reflect.TypeOf(result[key]), expected[key], reflect.TypeOf(expected[key]))
						break
					}
				}
			}
		}
	}
}

func TestDecoder_Func(t *testing.T) {
	// value: func | expected: func
	var result func(x int) int
	f := func(x int) int {
		return x + 5
	}

	if err := Decode(f, &result); err != nil {
		t.Error(err)
	} else if result == nil {
		t.Error("result is empty")
	} else if result(5) != 10 {
		t.Errorf("expected function with return value is %d, but got value %d", 10, result(5))
	}
}

func TestDecoder_Array(t *testing.T) {
	// value: []int | expected: []int
	{
		var expected = [3]int{1, 2, 3}
		var result [3]int
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: []string | expected: []string
	{
		var expected = [3]string{"TEST_1", "TEST_2", "TEST_3"}
		var result [3]string
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: []bool | expected: []bool
	{
		var expected = [3]bool{true, false, true}
		var result [3]bool
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: int | expected: []int
	{
		var expected = [3]int{1}
		var result [3]int
		if err := Decode(1, &result); err != nil {
			t.Error(err)
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: int | expected: []string
	{
		var expected = [3]string{"1"}
		var result [3]string
		if err := Decode(1, &result); err != nil {
			t.Error(err)
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: bool | expected: []int
	{
		var expected = [3]int{1}
		var result [3]int
		if err := Decode(true, &result); err != nil {
			t.Error(err)
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}
}

func TestDecoder_Slice(t *testing.T) {
	// value: []int | expected: []int
	{
		expected := []int{1, 2, 3}
		var result []int
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else if len(expected) != len(result) {
			t.Errorf("expected length of value is %d, but got %d", len(expected), len(result))
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: []string | expected: []string
	{
		expected := []string{"1", "2", "3"}
		var result []string
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else if len(expected) != len(result) {
			t.Errorf("expected length of value is %d, but got %d", len(expected), len(result))
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: []bool | expected: []bool
	{
		expected := []bool{true, false, true}
		var result []bool
		if err := Decode(expected, &result); err != nil {
			t.Error(err)
		} else if len(expected) != len(result) {
			t.Errorf("expected length of value is %d, but got %d", len(expected), len(result))
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: int | expected: []int
	{
		expected := []int{1}
		var result []int
		if err := Decode(1, &result); err != nil {
			t.Error(err)
		} else if len(expected) != len(result) {
			t.Errorf("expected length of value is %d, but got %d", len(expected), len(result))
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: int | expected: []string
	{
		expected := []string{"1"}
		var result []string
		if err := Decode(1, &result); err != nil {
			t.Error(err)
		} else if len(expected) != len(result) {
			t.Errorf("expected length of value is %d, but got %d", len(expected), len(result))
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}

	// value: bool | expected: []int
	{
		expected := []int{1}
		var result []int
		if err := Decode(true, &result); err != nil {
			t.Error(err)
		} else if len(expected) != len(result) {
			t.Errorf("expected length of value is %d, but got %d", len(expected), len(result))
		} else {
			for i := 0; i < len(expected); i++ {
				if expected[i] != result[i] {
					t.Errorf("expected value is %v, but got %v", expected[i], result[i])
				}
			}
		}
	}
}

func TestDecoder_MapToStruct(t *testing.T) {
	type student struct {
		GPA    float32
		Name   string
		Age    int
		IsPass bool
	}

	var result student

	expected := student{
		GPA:    3.14,
		Name:   "TEST",
		Age:    10,
		IsPass: true,
	}

	if err := Decode(map[string]interface{}{
		"GPA":    3.14,
		"Name":   "TEST",
		"Age":    10,
		"IsPass": true,
	}, &result); err != nil {
		t.Error(err)
	} else if expected != result {
		t.Errorf("expected value is %v, but got %v", expected, result)
	}
}

func TestDecoder_Struct(t *testing.T) {
	// struct to struct
	{
		type input struct {
			GPA    float32
			Name   string
			Age    int
			IsPass bool
		}

		type output struct {
			GPA    float32
			Name   string
			Age    int
			IsPass bool
			Score  float32
			Gender string
		}

		expected := output{
			GPA:    3.14,
			Name:   "TEST",
			Age:    10,
			IsPass: true,
			Score:  0,
			Gender: "",
		}

		var result output

		if err := Decode(input{
			GPA:    3.14,
			Name:   "TEST",
			Age:    10,
			IsPass: true,
		}, &result); err != nil {
			t.Error(err)
		} else if result != expected {
			t.Errorf("expected value is %v, but got %v", expected, result)
		}
	}

	// nested struct
}

func TestDecoder_Time(t *testing.T) {
	// value: time | expected: string
	{
		now := time.Now()
		expected := now.Format("2006-01-02 15:04:05")
		var result string

		if err := NewDecoder().SetManipulation(TimeDataType, TimeToStringManipulation).Out(&result).Decode(now); err != nil {
			t.Error(err)
		} else if expected != result {
			t.Errorf("expected value is %v, but got %v", expected, result)
		}
	}

	// value: time | expected: int64 (unix)
	{
		now := time.Now()
		expected := now.Unix()
		var result int64

		if err := NewDecoder().SetManipulation(TimeDataType, TimeToUnixManipulation).Out(&result).Decode(now); err != nil {
			t.Error(err)
		} else if expected != result {
			t.Errorf("expected value is %v, but got %v", expected, result)
		}
	}
}
