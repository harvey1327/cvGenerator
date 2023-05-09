package flags

import (
	"flag"

	"github.com/harvey1327/resumehack/internal/model"
)

type flags struct {
	email string
	phone string
}

func GetFlags() flags {
	email := flag.String("email", "", "an email as a string")
	phone := flag.String("phone", "", "a phone number as a string")
	flag.Parse()
	return flags{
		email: *email,
		phone: *phone,
	}
}

func (f flags) Apply(pageData *model.PageData) {
	if f.email != "" {
		pageData.Contact.Email = f.email
	}
	if f.phone != "" {
		pageData.Contact.Phone = f.phone
	}
}
