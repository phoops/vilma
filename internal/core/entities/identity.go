package entities

import "fmt"

type InvalidIdentityErr struct {
	Field string
}

func (e *InvalidIdentityErr) Error() string {
	return fmt.Sprintf("identity not valid, invalid field: %s", e.Field)
}

func (e *InvalidIdentityErr) Is(target error) bool {
	t, ok := target.(*InvalidIdentityErr)
	if !ok {
		return false
	}
	return e.Field == t.Field
}

type Identity struct {
	Email     string
	FirstName string
	LastName  string
}

func (i *Identity) Valid() error {
	if i.Email == "" {
		return &InvalidIdentityErr{Field: "Email"}
	}

	if i.FirstName == "" {
		return &InvalidIdentityErr{Field: "FirstName"}
	}

	if i.LastName == "" {
		return &InvalidIdentityErr{Field: "LastName"}
	}

	return nil
}
