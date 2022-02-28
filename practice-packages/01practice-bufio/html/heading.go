package main

import "errors"

func CheckHeadings(line *string) (string, error) {
	switch *line {
	case "<" + H1 + ">":
		return H1, nil
	case "<" + H2 + ">":
		return H2, nil
	case "<" + H3 + ">":
		return H3, nil
	case "<" + H4 + ">":
		return H4, nil
	case "<" + H5 + ">":
		return H5, nil
	case "<" + H6 + ">":
		return H6, nil
	default:
		return "", errors.New("there is no heading")
	}
}
