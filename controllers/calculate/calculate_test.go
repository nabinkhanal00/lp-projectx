package calculate

import "testing"

func TestIsValid(t *testing.T) {
	expression := "12.34 + 56 * !@10.5"
	err := isValid(expression)
	t.Log("Hello world")
	t.Fatal("HELLO")
	t.Log(err)
}
