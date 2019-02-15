package util_superchecker

import (
	"fmt"
	"testing"
)

type User struct {
	Id       int
	//Username string `validate:"regex,username"`
	Username string `validate:"regex,^[\u4E00-\u9FA5a-zA-Z0-9_.]{0,40}$"`

	//Password string `validate:"func,lengthBetween6And15"`
	Password string

	Money int `validate:"int,0:"`
	State int `validate:"range,[0,1,2,3,-1]"`
}

func (o User) ValidateSVBCreate() (bool, string, error) {
	if o.Password == "" {
		return false, "password should not be empty", nil
	}
	return true, "", nil
}

func (o User) ValidateSVBCreateSVSUpdate() (bool, string, error) {
	if o.Id != 0 {
		return false, "id can't be modified", nil
	}
	return true, "success", nil
}
func TestSPC(t *testing.T) {
	u := User{
		//Id:       2,
		Username: "fengtao",
		// Username: "fengtao?",
		Password: "123456",
		// Password: "12345",
		// Password: "",
		Money: 9000,
		// Money: -2,
		State: 3,
		// State: -3
	}
	Spc.AddFunc(func(in interface{}, fieldName string) (bool, string, error) {
		field := in.(string)
		if len(field) < 6 || len(field) > 15 {
			return false, fmt.Sprintf("field '%s' should be length beweent 6 and 15 but got %d", fieldName, len(field)), nil
		}
		return true, "success", nil
	}, "lengthBetween6And15")

	ok, msg, e := Spc.Validate(u)
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(ok, msg)

	ok, msg, e = Spc.ValidateMethods(u, "create", "update")
	if e != nil {
		fmt.Println(e.Error())
		t.Fail()
		return
	}
	fmt.Println(ok, msg)
}
