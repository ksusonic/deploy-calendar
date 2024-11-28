package main

type Calendar struct {
	ConditionsBySign map[string]ConditionByDay `json:"conditionsBySign"`
}

type ConditionByDay map[int]string
