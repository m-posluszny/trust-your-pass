package passwords

import "github.com/agrison/go-commons-lang/stringUtils"

func validateInput(password string) []PreconditionDto {
	preconditions := make([]PreconditionDto, 0)

	preconditions = append(preconditions, PreconditionDto{lengthCondition, len(password) >= 4})
	preconditions = append(preconditions, PreconditionDto{nonEmptyStringCondition, !stringUtils.IsBlank(password)})
	preconditions = append(preconditions, PreconditionDto{nonWhitespacesOnlyCondition, !stringUtils.IsWhitespace(password)})
	preconditions = append(preconditions, PreconditionDto{notAllowedSymbolsCondition, !stringUtils.ContainsAny(password, "\\0")})

	return preconditions
}

func arePreconditionsSatisfied(preconditions []PreconditionDto) bool {
	for _, precond := range preconditions {
		if !precond.IsSatisfied {
			return false
		}
	}
	return true
}
