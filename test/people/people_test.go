package people

import "testing"

func TestSerialize(t *testing.T) {
	people := &People{
		Name: "关",
		Age:  24,
	}
	res := people.Serialize()
	if !res {
		t.Fatalf("people.Serialize() 错误，希望为=%v 实际为=%v", true, res)
	}
	t.Logf("people.Serialize() test success!")
}

func TestUnserialize(t *testing.T) {
	people := &People{}
	res := people.Unserialize()
	if !res {
		t.Fatalf("people.Unserialize() 错误，希望为=%v 实际为=%v", true, res)
	}
	if people.Name != "关" {
		t.Fatalf("people.Unserialize() 错误，希望为=%v 实际为=%v", "关", people.Name)
	}
	t.Logf("people.Unserialize() test success")
}
