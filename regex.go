package regex

import (
	"errors"
	"regexp"
)

// Group - Retorna o macth em grupos
// Retorno: []string, error
func Group(text, expression string) ([]string, error) {
	rgx, err := regexp.Compile(expression)
	if err != nil {
		return nil, err
	}

	if !rgx.MatchString(text) {
		return nil, errors.New("Conteudo não está no padrão especificado")
	}

	var matchs []string

	for _, match := range rgx.FindStringSubmatch(text)[1:] {
		if match != "" {
			matchs = append(matchs, match)
		}
	}

	return matchs, nil
}
