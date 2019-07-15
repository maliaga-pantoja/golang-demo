package domain

import (
	"go-struct01/interfaces"
	"testing"
)

func TestCustomerFullname(t *testing.T) {
	t.Run("testing fullname", func (t *testing.T) {
		var c interfaces.CustomerInterface
		var expect string = "Alan Damian"
		var setFullnameError = c.SetFullname("Alan Damian")
		if setFullnameError != nil {
			t.Errorf(setFullnameError.Error())
		}
		if c.GetFullname() != expect { // expects fullname be equals to 'Alan Damian'
			t.Fatalf("Expect customer.GetFullname equals to %s, got %s", c.GetFullname(), expect) // throw error if exists
		}
	})
	t.Run("testing fullname, expect to throw", func (t *testing.T) {
		var c interfaces.CustomerInterface
		var err = c.SetFullname("")
		if err != nil { // expects fullname be equals to 'Alan Damian'
			t.Fatalf("Expect customer.GetFullname to throw and no throw") // throw error if exists
		}
	})
}

func TestCustomerBirthday(t *testing.T) {
	t.Run("testing birthday", func (t *testing.T) {
		var c interfaces.CustomerInterface
		var expect string = "04/03/1991"
		var setBirthdayError = c.SetBirthday("04/03/1991")
		if setBirthdayError != nil {
			t.Errorf(setBirthdayError.Error())
		}
		if c.GetBirthday() != expect { // expects fullname be equals to 'Alan Damian'
			t.Fatalf("Expect customer.GetBirthday equals to %s, got %s", c.GetBirthday(), expect) // throw error if exists
		}
	})
	t.Run("testing fullname, expect to throw", func (t *testing.T) {
		var c interfaces.CustomerInterface
		var err = c.SetBirthday("")
		if err != nil { // expects fullname be equals to 'Alan Damian'
			t.Fatalf("Expect customer.GetFullname to throw and no throw") // throw error if exists
		}
	})
}