package subsciption

import (
	"fmt"
	"testing"
)

func TextNewSubsciption(t *testing.T) {

	ok_addr := "alice@example.com"
	ok_email := fmt.Sprintf("Alice <%s>", ok_addr)

	bunk_email := "dev/null"

	_, err := NewSubsciption(bunk_email)

	if err == nil {
		t.Error("Bunk address passed muster (when it shouldn't")
	}

	sub, err := NewSubsciption(ok_email)

	if err != nil {
		t.Error(err)
	}

	if sub.Address != ok_addr {
		t.Error("Unexpected 'okay' address")
	}

}
