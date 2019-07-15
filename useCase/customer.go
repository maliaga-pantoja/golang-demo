package useCase

import (
	"go-struct01/interfaces"
	"strings"
)

func SayHello(fn string, birthday string) (error, string) {
	var customer interfaces.CustomerInterface
	var sfE = customer.SetFullname(fn)
	var sbE = customer.SetBirthday(birthday)
	if sfE != nil {
		return sfE, ""
	}
	if sbE != nil {
		return sbE, ""
	}
	var s = []string{"hi", customer.GetFullname(), "u are", customer.GetBirthday()}
	return nil, strings.Join(s, " ")
}