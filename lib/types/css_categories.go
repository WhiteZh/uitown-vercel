package types

import (
	"errors"
	"log"
)

type CssCategoryType int

const (
	CssCategoryButton CssCategoryType = iota
	CssCategoryCheckBox
	CssCategoryToggleSwitch
	CssCategoryCard
	CssCategoryLoader
	CssCategoryInput
	//CssCategoryTransition
	CssCategorySpecialEffect
)

func ConvertStringToCssCategory(s string) (CssCategoryType, error) {
	switch s {
	case "button":
		return CssCategoryButton, nil
	case "checkbox":
		return CssCategoryCheckBox, nil
	case "toggle switch":
		return CssCategoryToggleSwitch, nil
	case "card":
		return CssCategoryCard, nil
	case "loader":
		return CssCategoryLoader, nil
	case "input":
		return CssCategoryInput, nil
	//case "transition":
	//	return CssCategoryTransition, nil
	case "special effect":
		return CssCategorySpecialEffect, nil
	default:
		return -1, errors.New("invalid css category")
	}
}

func ConvertCssCategoryToString(c CssCategoryType) string {
	switch c {
	case CssCategoryButton:
		return "button"
	case CssCategoryCheckBox:
		return "checkbox"
	case CssCategoryToggleSwitch:
		return "toggle switch"
	case CssCategoryCard:
		return "card"
	case CssCategoryLoader:
		return "loader"
	case CssCategoryInput:
		return "input"
	//case CssCategoryTransition:
	//	return "transition"
	case CssCategorySpecialEffect:
		return "special effect"
	default:
		log.Panic("uncovered `CssCategoryType`")
		return ""
	}
}

func (c CssCategoryType) ToString() string {
	return ConvertCssCategoryToString(c)
}
