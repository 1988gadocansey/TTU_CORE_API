package services

import "strings"

func IndexNumFormater(index string) string {

	return strings.ReplaceAll(index, "/", "")
}
