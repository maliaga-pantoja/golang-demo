package domain

import (
	"go-struct01/interfaces"
	"testing"
)

func TestCustomerFullname(t *testing.T) {
	t.Run("testing fullname", func (t *testing.T) {
		var c interfaces.CustomerInterface
		var expect string = "Alan Damian"
		c.SetFullname("Alan Damian")
		if c.GetFullname() != expect { // expects fullname be equals to 'Alan Damian'
			t.Fatalf("Expect customer.GetFullname equals to %s, got %s", c.GetFullname(), expect) // throw error if exists
		}
	})
	t.Run("testing fullname, expect to throw", func (t *testing.T) {
		var c interfaces.CustomerInterface
		var ok, err = c.SetFullname("")
		if err != nil && ok == true{ // expects fullname be equals to 'Alan Damian'
			t.Fatalf("Expect customer.GetFullname to throw and no throw") // throw error if exists
		}
	})
}