package interfaces

import (
	"github.com/pkg/errors"
	"go-struct01/domain"
)

type CustomerInterface struct {
	customer domain.CustomerStruct
}

func (c *CustomerInterface) SetFullname(fn string) error {
	if fn == "" || len(fn)<1 {
		return errors.Wrapf(nil, "new fullname must no to be empty")
	} else {
		c.customer.SetFullname(fn)
	}
	return nil
}

func (c *CustomerInterface) GetFullname() string {
	return c.customer.GetFullname()
}

func (c *CustomerInterface) SetBirthday(bd string) error {
	if bd == "" || len(bd)<1 {
		return errors.Wrapf(nil, "birthday must no to be empty")
	} else {
		c.customer.SetBirthday(bd)
	}
	return nil
}

func (c *CustomerInterface) GetBirthday() string {
	return c.customer.GetBirthday()
}
