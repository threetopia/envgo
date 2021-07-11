package envgo

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
)

const defaultDelimiter string = ":"

func LoadDotEnv(filenames ...string) error {
	err := godotenv.Load(filenames...)
	if err != nil {
		return err
	}
	return nil
}

func GetString(path string, predefined string) string {
	res := os.Getenv(path)
	if res != "" {
		return res
	}
	return predefined
}

func GetStringSlice(path string, delimiter string) []string {
	res := make([]string, 0)
	getString := GetString(path, "")
	if getString == "" {
		return res
	}
	// set default delimiter if empty
	if delimiter == "" {
		delimiter = defaultDelimiter
	}
	return strings.Split(getString, delimiter)
}

func GetInt(path string, predefined int) int {
	res := os.Getenv(path)
	if res != "" {
		ires, err := strconv.Atoi(res)
		if err == nil {
			return ires
		}
	}
	return predefined
}

func GetIntSlice(path string, delimiter string) []int {
	res := make([]int, 0)
	getInt := os.Getenv(path)
	if getInt == "" {
		return res
	}
	// set default delimiter if empty
	if delimiter == "" {
		delimiter = defaultDelimiter
	}
	getIntSlice := strings.Split(getInt, delimiter)
	if len(getIntSlice) > 0 {
		for _, v := range getIntSlice {
			ires, err := strconv.Atoi(v)
			if err == nil {
				res = append(res, ires)
			}
		}
	}
	return res
}

func GetInt64(path string, predefined int64) int64 {
	res := os.Getenv(path)
	if res != "" {
		ires, err := strconv.ParseInt(res, 10, 64)
		if err == nil {
			return ires
		}
	}
	return predefined
}

func GetInt64Slice(path string, delimiter string) []int64 {
	res := make([]int64, 0)
	getInt := os.Getenv(path)
	if getInt == "" {
		return res
	}
	// set default delimiter if empty
	if delimiter == "" {
		delimiter = defaultDelimiter
	}
	getIntSlice := strings.Split(getInt, delimiter)
	if len(getIntSlice) > 0 {
		for _, v := range getIntSlice {
			ires, err := strconv.ParseInt(v, 10, 64)
			if err == nil {
				res = append(res, ires)
			}
		}
	}
	return res
}

func GetFloat32(path string, predefined float32) float32 {
	res := os.Getenv(path)
	if res != "" {
		if f64, err := strconv.ParseFloat(res, 32); err == nil {
			return float32(f64)
		}
	}
	return predefined
}

func GetFloat64(path string, predefined float64) float64 {
	res := os.Getenv(path)
	if res != "" {
		if f64, err := strconv.ParseFloat(res, 64); err == nil {
			return f64
		}
	}
	return predefined
}

func GetBool(path string, predefined bool) bool {
	res := os.Getenv(path)
	if res != "" {
		return res == "true"
	}
	return predefined
}

func GetPort(path string, predefined int) string {
	res := GetInt(path, predefined)
	if res > 0 {
		return fmt.Sprintf(":%d", res)
	}
	return fmt.Sprintf(":%d", predefined)
}
