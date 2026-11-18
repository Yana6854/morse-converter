package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)
// AutoConvert — функция, принимающая строку, и возвращает результат и ошибку
func AutoConvert(input string) (string, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return "", errors.New("входная строка пуста")
	}

	// Если есть точки или тире — это код Морзе
	if strings.ContainsAny(input, ".-") {
		return morse.ToText(input), nil
	}

	// Если точек и тире нет, то это просто текст, его переводим в Морзе
	return morse.ToMorse(input), nil
}
