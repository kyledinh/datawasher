package util

import (
	"testing"
)

func Test_IsRawbodyArray(t *testing.T) {
	js := []byte(`{"first_name":"Dallas","last_name":"Krueger","email":"Dallas.Krueger@mail.com","phone_number":"555-395-7681","street_address":"56 Silver Street","city":"Pawnee","state":"IN"}`)
	ar := []byte(`[{"first_name":"Dallas", "first_name":"Hobo", "first_name":"Black"}]`)

	if IsRawbodyArray(js) == true {
		t.Error("false positive for js")
	}

	if IsRawbodyArray(ar) == false {
		t.Error("did not check correctly ar")
	}

}

func Test_WrapJsonAsRoot(t *testing.T) {
	ar := []byte(`[{"first_name":"Dallas"}, {"first_name":"Hobo"}]`)
	expected := []byte(`{ "root" : [{"first_name":"Dallas"}, {"first_name":"Hobo"}]}`)

	if string(WrapJsonAsRoot(ar)) != string(expected) {
		t.Error("wrapped incorrect ", string(WrapJsonAsRoot(ar)))
	}
}
