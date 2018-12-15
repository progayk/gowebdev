package main

import "testing"

const (
	value = "I'm still breathing, I'm still breathing. I'm aliiiiiiiive. I'm aliiiiiiive..."
	expectedEncode = "SSdtIHN0aWxsIGJyZWF0aGluZywgSSdtIHN0aWxsIGJyZWF0aGluZy4gSSdtIGFsaWlpaWlpaWl2ZS4gSSdtIGFsaWlpaWlpaXZlLi4u"
)

func TestEncode(t *testing.T) {
	encoded := encode(value)
	if encoded != expectedEncode {
		t.Errorf("expected %s got %s", expectedEncode, encoded)
	}
}

func TestDecode(t *testing.T) {
	decoded := decode(expectedEncode)
	if decoded != value {
		t.Errorf("expected %s got %s", value, decoded)
	}
}
