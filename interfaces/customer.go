package interfaces

import (
	"github.com/pkg/errors"
	"go-struct01/domain"
)

type CustomerInterface struct {
	customer domain.CustomerStruct
}

func (c *CustomerInterface) SetFullname(fn string) (bool, error) {
	if fn == "" {
		return false, errors.Wrapf(nil, "new fullname must no to be empty")
	} else {
		c.customer.SetFullname(fn)
	}
	return true, nil
}

func (c *CustomerInterface) GetFullname() string {
	return c.customer.GetFullname()
}

func (c *CustomerInterface) SetBirthday(bd string) (bool, error) {
	if bd == "" {
		return false, errors.Wrapf(nil, "birthday must no to be empty")
	} else {
		c.customer.SetBirthday(bd)
	}
	return true, nil
}

func (c *CustomerInterface) GetBirthday() string {
	return c.customer.GetBirthday()
}
