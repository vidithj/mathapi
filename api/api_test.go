package api

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValidateLength(t *testing.T) {

	tests := []struct {
		name     string
		request  []int
		response bool
	}{
		{
			name:     "valid list input",
			request:  []int{2, 3, 4, 5},
			response: true,
		},
		{
			name:     "empty list input",
			request:  []int{},
			response: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := validateLength(test.request)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}

func TestValidateQuantifier(t *testing.T) {

	tests := []struct {
		name     string
		request  int
		response bool
	}{
		{
			name:     "valid quantifier input",
			request:  20,
			response: true,
		},
		{
			name:     "quantifier input",
			request:  10,
			response: true,
		},
		{
			name:     "negative quantifier input",
			request:  -20,
			response: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response := validateQuantifier(test.request)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}

func TestCalcMin(t *testing.T) {

	tests := []struct {
		name              string
		requestlist       []int
		requestquantifier int
		response          []int
	}{
		{
			name:              "valid input",
			requestlist:       []int{9, 10, 2, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			requestquantifier: 1,
			response:          []int{2},
		},
		{
			name:              "Invalid quantifier input",
			requestlist:       []int{9, 10, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			requestquantifier: -1,
			response:          nil,
		},
		{
			name:              "empty list input",
			requestlist:       []int{},
			requestquantifier: 1,
			response:          nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, _ := CalcMin(test.requestlist, test.requestquantifier)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}
func TestCalcMax(t *testing.T) {

	tests := []struct {
		name              string
		requestlist       []int
		requestquantifier int
		response          []int
	}{
		{
			name:              "valid input",
			requestlist:       []int{9, 10, 2, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			requestquantifier: 1,
			response:          []int{25},
		},
		{
			name:              "Invalid quantifier input",
			requestlist:       []int{9, 10, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			requestquantifier: -1,
			response:          nil,
		},
		{
			name:              "empty list input",
			requestlist:       []int{},
			requestquantifier: 1,
			response:          nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, _ := CalcMax(test.requestlist, test.requestquantifier)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}

func TestCalcAvg(t *testing.T) {

	tests := []struct {
		name        string
		requestlist []int
		response    float64
	}{
		{
			name:        "valid input",
			requestlist: []int{9, 10, 2, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			response:    15.882352941176471,
		},
		{
			name:        "valid input",
			requestlist: []int{9, 10, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			response:    16.75,
		},
		{
			name:        "empty list input",
			requestlist: []int{},
			response:    0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, _ := CalcAvg(test.requestlist)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}

func TestCalcMedian(t *testing.T) {

	tests := []struct {
		name        string
		requestlist []int
		response    float64
	}{
		{
			name:        "valid input",
			requestlist: []int{9, 10, 2, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			response:    15,
		},
		{
			name:        "valid input",
			requestlist: []int{9, 10, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			response:    15.5,
		},
		{
			name:        "empty list input",
			requestlist: []int{},
			response:    0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, _ := CalcMedian(test.requestlist)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}

func TestCalcPercentile(t *testing.T) {

	tests := []struct {
		name              string
		requestlist       []int
		requestquantifier int
		response          int
	}{
		{
			name:              "valid input",
			requestlist:       []int{9, 10, 2, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			requestquantifier: 90,
			response:          24,
		},
		{
			name:              "valid input,quantifier 100",
			requestlist:       []int{9, 10, 2, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			requestquantifier: 100,
			response:          25,
		},
		{
			name:              "invalid quantifier",
			requestlist:       []int{9, 10, 12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25},
			requestquantifier: 101,
			response:          0,
		},
		{
			name:              "negative quantifier",
			requestlist:       []int{},
			requestquantifier: -10,
			response:          0,
		},
		{
			name:              "empty list input",
			requestlist:       []int{},
			requestquantifier: 100,
			response:          0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			response, _ := CalcPercentile(test.requestlist, test.requestquantifier)
			if diff := cmp.Diff(response, test.response); diff != "" {
				t.Errorf("%s failed - differs: (-want +got)\n %s", test.name, diff)
			}
		})
	}

}
