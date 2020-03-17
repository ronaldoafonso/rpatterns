package factory

import "testing"

func TestCreateObject(t *testing.T) {
	tests := []struct {
		objType ObjType
		objMsg  string
		want    string
	}{
		{type1, "msg1", "msg1 returned"},
		{type2, "msg2", "msg2 returned"},
	}

	for _, test := range tests {
		obj, err := GetObjectMethod(test.objType)
		if err != nil {
			t.Fatalf("GetObjectMethod error: %v.", err)
		}

		msg := obj.Method(test.objMsg)
		if msg != test.want {
			t.Errorf("Method error. Want %v, got: %v.", test.want, msg)
		}
	}

	if _, err := GetObjectMethod(nonExistentType); err == nil {
		t.Error("GetObjectMethod should return an error, but returned nil.")
	}
}
