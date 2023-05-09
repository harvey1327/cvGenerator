package sensativedata

import (
	"flag"

	"github.com/harvey1327/resumehack/internal/model"
)

type data struct {
	email string
	phone string
}

func GetData() data {
	email := flag.String("email", "", "an email as a string")
	phone := flag.String("phone", "", "a phone number as a string")
	flag.Parse()
	return data{
		email: *email,
		phone: *phone,
	}
}

func (f data) Apply(pageData *model.PageData) {
	if f.email != "" {
		pageData.Contact.Email = f.email
	}
	if f.phone != "" {
		pageData.Contact.Phone = f.phone
	}
}
