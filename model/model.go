package model

type MathRequest struct {
	Quantifier int   `json:"quantifier,omitempty"`
	List       []int `json:"list"`
}
