package utils

import "github.com/agrison/go-commons-lang/stringUtils"

const lengthCondition = "Password should contain at least 4 signs"
const nonEmptyStringCondition = "Provided string is not empty"
const nonWhitespacesOnlyCondition = "Provided string doesnt consists of whitespaces only"
const notAllowedSymbolsCondition = "Provided string doesnt contain symbols that are not allowed"

func ValidateInput(password string) map[string]bool {

	var preconditions = map[string]bool{
		lengthCondition:             false,
		nonEmptyStringCondition:     false,
		nonWhitespacesOnlyCondition: false,
		notAllowedSymbolsCondition:  false}

	if len(password) >= 4 {
		preconditions[lengthCondition] = true
	}
	if !stringUtils.IsBlank(password) {
		preconditions[nonEmptyStringCondition] = true
	}
	if !stringUtils.IsWhitespace(password) {
		preconditions[nonWhitespacesOnlyCondition] = true
	}
	if !stringUtils.ContainsAny(password, "\\0") {
		preconditions[notAllowedSymbolsCondition] = true
	}

	return preconditions
}

func ArePreconditionsSatisfied(preconditions map[string]bool) bool {
	for _, v := range preconditions {
		if v == false {
			return false
		}
	}
	return true
}
