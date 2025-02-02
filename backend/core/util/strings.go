package util

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToInt64(s string) (int64, error) {
	if s == "" {
		return 0, errors.New("input string is empty")
	}

	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse to int64: %v", err)
	}

	return num, nil
}

func StringToFloat64(s string) (float64, error) {
	if s == "" {
		return 0, errors.New("input string is empty")
	}

	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse to float64: %v", err)
	}

	return num, nil
}
