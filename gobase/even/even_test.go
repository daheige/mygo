package even

import (
	"testing"
)

//必须Test开头
func TestEven(t *testing.T) {
	if !Even(10) {
		t.Log("10 must be even")
		t.Fail()
	}

	if Even(5) {
		t.Log("7 is not even")
		t.Fail()
	}

	// if Even(10) {
	//  t.Log("10 is even")
	//  t.Fail()
	// }
}
