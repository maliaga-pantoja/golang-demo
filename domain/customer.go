package domain

// CustomerStruct struct
type CustomerStruct struct {
	fullname string
	birthday string
}

// SetFullname unexported
func (c *CustomerStruct) SetFullname(fn string) {
	c.fullname = fn
}

// GetFullname unexported
func (c *CustomerStruct) GetFullname() string {
	var filenames string = c.fullname
	return filenames
}

// SetBirthday unexported
func (c *CustomerStruct) SetBirthday(bd string) {
	c.birthday = bd
}

// GetBirthday unexported
func (c *CustomerStruct) GetBirthday() string {
	var birthday string = c.birthday
	return birthday
}
