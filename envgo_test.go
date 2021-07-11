package envgo

import (
	"fmt"
	"testing"
)

func TestGoEnv_LoadDotEnv(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
}

func TestGoEnv_GetString(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	str := GetString("TEST_GET_STRING", "")
	if str != "Loaded Correctly" {
		t.Errorf("Not Match (%s is wrong value)", str)
	}
}

func TestGoEnv_GetStringSlice(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	slc := [3]string{"ValueOne", "ValueTwo", "ValueThree"}
	// test with default delimiter
	slice := GetStringSlice("TEST_GET_STRING_SLICE", "")
	for k, v := range slc {
		if slice[k] != v {
			t.Errorf("Not Match (%s is wrong value for %s)", slice[k], v)
		}
	}
	// test with custom delimiter
	slice = GetStringSlice("TEST_GET_STRING_SLICE_DELIMITER", ",")
	for k, v := range slc {
		if slice[k] != v {
			t.Errorf("Not Match (%s is wrong value for %s)", slice[k], v)
		}
	}
}

func TestGoEnv_GetInt(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	integer := GetInt("TEST_GET_NUMBER", 0)
	if integer != 1234567890 {
		t.Errorf("Not Match (%d is wrong value)", integer)
	}
}

func TestGoEnv_GetIntSlice(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	slc := [5]int{1, 2, 3, 4, 5}
	// test with default delimiter
	slice := GetIntSlice("TEST_GET_NUMBER_SLICE", "")
	for k, v := range slc {
		if slice[k] != v {
			t.Errorf("Not Match (%d is wrong value for %d)", slice[k], v)
		}
	}
	// test with custom delimiter
	slice = GetIntSlice("TEST_GET_NUMBER_SLICE_DELIMITER", ",")
	for k, v := range slc {
		if slice[k] != v {
			t.Errorf("Not Match (%d is wrong value for %d)", slice[k], v)
		}
	}
}

func TestGoEnv_GetInt64(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	integer := GetInt64("TEST_GET_NUMBER", 0)
	if integer != int64(1234567890) {
		t.Errorf("Not Match (%d is wrong value)", integer)
	}
}

func TestGoEnv_GetInt64Slice(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	slc := [5]int64{1, 2, 3, 4, 5}
	// test with default delimiter
	slice := GetInt64Slice("TEST_GET_NUMBER_SLICE", "")
	for k, v := range slc {
		if slice[k] != v {
			t.Errorf("Not Match (%d is wrong value for %d)", slice[k], v)
		}
	}
	// test with custom delimiter
	slice = GetInt64Slice("TEST_GET_NUMBER_SLICE_DELIMITER", ",")
	for k, v := range slc {
		if slice[k] != v {
			t.Errorf("Not Match (%d is wrong value for %d)", slice[k], v)
		}
	}
}

func TestGoEnv_GetFloat32(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	f64 := GetFloat32("TEST_GET_NUMBER", 0)
	if f64 != float32(1234567890) {
		t.Errorf("Not Match (%f is wrong value)", f64)
	}
}

func TestGoEnv_GetFloat64(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	f64 := GetFloat64("TEST_GET_NUMBER", 0)
	if f64 != float64(1234567890) {
		t.Errorf("Not Match (%f is wrong value)", f64)
	}
}

func TestGoEnv_GetBool(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	isTrue := GetBool("TEST_GET_BOOL", true)
	if isTrue == false {
		t.Errorf("Not Match (%v is wrong value)", isTrue)
	}
}

func TestGoEnv_GetPort(t *testing.T) {
	err := LoadDotEnv()
	if err != nil {
		t.Errorf("Error (%s)", err.Error())
	}
	port := GetPort("TEST_GET_PORT", 0)
	if port != fmt.Sprintf(":8080") {
		t.Errorf("Not Match (%v is wrong value)", port)
	}
}
