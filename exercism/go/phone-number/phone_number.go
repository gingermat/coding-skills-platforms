package phonenumber

import (
	"errors"
	"fmt"
)

type PhoneNumber struct {
	number   string
	areaCode string
}

func (pn PhoneNumber) Format() string {
	return fmt.Sprintf("(%s) %s-%s", pn.areaCode, pn.number[:3], pn.number[3:])
}

func parsePhoneNumber(input string) (*PhoneNumber, error) {
	var (
		pos      int
		filtered []rune
	)

	for _, chr := range input {
		if pos > 10 {
			return nil, errors.New("number more than 11 digits")
		}

		if chr >= '0' && chr <= '9' {
			filtered = append(filtered, chr)
			pos++

			continue
		}

		if !(chr == '+' || chr == '(' || chr == ')' || chr == '-' || chr == ' ' || chr == '.') {
			return nil, errors.New("invalid character")
		}
	}

	if len(filtered) < 10 {
		return nil, errors.New("invalid number length")
	}

	if pos == 11 && filtered[0] != '1' {
		return nil, errors.New("11 digits number must start with 1")
	}

	if len(filtered) == 11 {
		filtered = filtered[1:]
	}

	for _, pos := range []int{0, 3} {
		if !(filtered[pos] >= '2' && filtered[pos] <= '9') {
			return nil, errors.New("invalid number on position")
		}
	}

	return &PhoneNumber{
		areaCode: string(filtered[0:3]),
		number:   string(filtered[3:]),
	}, nil
}

func Number(input string) (string, error) {
	pn, err := parsePhoneNumber(input)
	if err != nil {
		return "", err
	}

	return pn.areaCode + pn.number, nil
}

func AreaCode(input string) (string, error) {
	pn, err := parsePhoneNumber(input)
	if err != nil {
		return "", err
	}

	return pn.areaCode, nil
}

func Format(input string) (string, error) {
	pn, err := parsePhoneNumber(input)
	if err != nil {
		return "", err
	}

	return pn.Format(), nil
}
