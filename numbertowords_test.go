package numbertowords

import "testing"

func TestInvalidInput(t *testing.T) {
	result, err := Convert(-1)
	if result != "" || err == nil {
		t.Fatal("failed")
	}
	result, err = Convert(5)
	if result != "" || err == nil {

	}
}
