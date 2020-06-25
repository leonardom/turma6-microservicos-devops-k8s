package greeting

import "testing"

func TestGreeting(t *testing.T) {
	message := "Code.education Rocks!"
	got := Greeting(message)
	expected := "<b>" + message + "</b>"

	if (got != expected) {
		t.Errorf("Greeting(\"%v\") got: \"%v\", expected: \"%v\"", message, got, expected)
	}
}
