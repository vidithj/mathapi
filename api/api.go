package api

import (
	"errors"
	"math"
	"sort"
)

func validateLength(list []int) bool {
	return len(list) != 0
}
func validateQuantifier(quantifier int) bool {
	return quantifier >= 0
}
func CalcMin(list []int, quantifier int) ([]int, error) {
	if !validateLength(list) || !validateQuantifier(quantifier) {
		return nil, errors.New("validate the list and quantifier")
	}
	sort.Ints(list)
	return list[:quantifier], nil
}
func CalcMax(list []int, quantifier int) ([]int, error) {
	if !validateLength(list) || !validateQuantifier(quantifier) {
		return nil, errors.New("validate the list and quantifier")
	}

	sort.Sort(sort.Reverse(sort.IntSlice(list)))

	return list[:quantifier], nil
}
func CalcAvg(list []int) (float64, error) {
	if !validateLength(list) {
		return 0, errors.New("validate the list is not empty")
	}

	sum := 0
	for _, x := range list {
		sum += x
	}

	avg := (float64(sum)) / (float64(len(list)))
	return avg, nil
}
func CalcMedian(list []int) (float64, error) {
	if !validateLength(list) {
		return 0, errors.New("validate the list is not empty")
	}
	sort.Ints(list)

	middle := len(list) / 2
	var medianValue float64 = 0
	if len(list)%2 == 1 {
		medianValue = float64(list[middle])
	} else {
		medianValue = (float64(list[middle-1]) + float64(list[middle])) / float64(2)
	}

	return medianValue, nil
}
func CalcPercentile(list []int, quantifier int) (int, error) {
	if !validateLength(list) {
		return 0, errors.New("validate the list is not empty")
	}
	if quantifier < 0 || quantifier > 100 {
		return 0, errors.New("invalid value of quantifier")
	}
	sort.Ints(list)

	if quantifier == 100 {
		return list[len(list)-1], nil
	}

	ordinal := int(math.Ceil(float64(len(list)) * float64(quantifier) / 100))

	if ordinal == 0 {
		return list[0], nil
	}
	return list[ordinal-1], nil
}
