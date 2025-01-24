package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return defaultValue, fmt.Errorf("error found %w", err) // I dont know what is happening here.
	}
	return value, nil
}
