package singleton

import (
	"log"
	"testing"
)

func TestGetInstance(t *testing.T) {

	counter1 := GetInstance()
	if counter1 == nil {
		//Test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after calling GetInstance(), not nil")
	}

	log.Println(counter1)

	expectedCounter := counter1
	cnt := counter1.AddOne()
	if cnt != 1 {
		t.Errorf("value should have been 1,since call is made for the first time but the value is %d", cnt)
	}

	counter2 := GetInstance()
	if expectedCounter != counter2 {
		t.Errorf("Both the instance should be same")
	}

	log.Println(counter2)

	cnt2 := counter2.AddOne()
	if cnt2 != 2 {
		t.Errorf("The count should have been 2,but found counter value is %d", cnt2)
	}

}
