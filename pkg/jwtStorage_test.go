package pkg

import "testing"

func TestAddToMap(t *testing.T) {
	AddToMap("test", "test")
	if len(jwtMap) != 1 {
		t.Error("Expected 1, got ", len(jwtMap))
	}
}

func TestGetToken(t *testing.T) {
	AddToMap("test", "test")
	if GetToken("test") != "test" {
		t.Error("Expected test, got ", GetToken("test"))
	}
}

func TestListTokens(t *testing.T) {
	AddToMap("test", "test")
	if len(ListTokens()) != 1 {
		t.Error("Expected 1, got ", len(ListTokens()))
	}
}

func TestDeleteToken(t *testing.T) {
	AddToMap("test", "test")
	DeleteToken("test")
	if len(jwtMap) != 0 {
		t.Error("Expected 0, got ", len(jwtMap))
	}
}

func TestClearTokens(t *testing.T) {
	AddToMap("test", "test")
	ClearTokens()
	if len(jwtMap) != 0 {
		t.Error("Expected 0, got ", len(jwtMap))
	}
}

func TestGetTokenCount(t *testing.T) {
	AddToMap("test", "test")
	if GetTokenCount() != 1 {
		t.Error("Expected 1, got ", GetTokenCount())
	}
}

func TestGetTokenKeys(t *testing.T) {
	AddToMap("test", "test")
	if len(GetTokenKeys()) != 1 {
		t.Error("Expected 1, got ", len(GetTokenKeys()))
	}
}

func TestToJSON(t *testing.T) {
	AddToMap("test", "test")
	json := ToJSON()
	if json != "{\"test\":\"test\"}" {
		t.Error("Expected {\"test\":\"test\"}, got ", json)
	}

	AddToMap("test2", "test2")
	json = ToJSON()
	if json != "{\"test\":\"test\",\"test2\":\"test2\"}" {
		t.Error("Expected {\"test\":\"test\",\"test2\":\"test2\"}, got ", json)
	}
}

func TestFromJSON(t *testing.T) {
	FromJSON("{\"test\":\"test\"}")
	if len(jwtMap) != 1 {
		t.Error("Expected 1, got ", len(jwtMap))
	}
}
