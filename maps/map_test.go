package ezgo

import (
	"testing"
)

func TestSetNestedMapValue(t *testing.T) {
	// Create a two-level nested map
	nestedMap := make(map[string]map[int]int)

	// Test: Set a value in the nested map
	SetNestedMapValue(nestedMap, "firstKey", 100, 42)

	// Check if the value is correctly set
	if nestedMap["firstKey"][100] != 42 {
		t.Errorf("Expected nestedMap[firstKey][100] = 42, but got %v", nestedMap["firstKey"][100])
	}

	// Test: Set another value in the map
	SetNestedMapValue(nestedMap, "secondKey", 200, 99)

	// Check if the second value is correctly set
	if nestedMap["secondKey"][200] != 99 {
		t.Errorf("Expected nestedMap[secondKey][200] = 99, but got %v", nestedMap["secondKey"][200])
	}

	// Test: Set a value in an existing outer map key
	SetNestedMapValue(nestedMap, "firstKey", 101, 43)
	if nestedMap["firstKey"][101] != 43 {
		t.Errorf("Expected nestedMap[firstKey][101] = 43, but got %v", nestedMap["firstKey"][101])
	}

	// Verify that a non-existent outer key is correctly initialized
	if nestedMap["thirdKey"] != nil {
		t.Errorf("Expected nestedMap[thirdKey] to be nil, but got %v", nestedMap["thirdKey"])
	}
}

func TestSetNestedMapValue3(t *testing.T) {
	// Create a three-level nested map
	nestedMap3 := make(map[string]map[int]map[float64]string)

	// Test: Set a value in the three-level nested map
	SetNestedMapValue3(nestedMap3, "firstKey", 100, 1.23, "Value A")

	// Check if the value is correctly set
	if nestedMap3["firstKey"][100][1.23] != "Value A" {
		t.Errorf("Expected nestedMap3[firstKey][100][1.23] = 'Value A', but got %v", nestedMap3["firstKey"][100][1.23])
	}

	// Test: Set another value in the map
	SetNestedMapValue3(nestedMap3, "secondKey", 200, 4.56, "Value B")

	// Check if the second value is correctly set
	if nestedMap3["secondKey"][200][4.56] != "Value B" {
		t.Errorf("Expected nestedMap3[secondKey][200][4.56] = 'Value B', but got %v", nestedMap3["secondKey"][200][4.56])
	}

	// Test: Set a value in an existing outer map key
	SetNestedMapValue3(nestedMap3, "firstKey", 101, 7.89, "Value C")
	if nestedMap3["firstKey"][101][7.89] != "Value C" {
		t.Errorf("Expected nestedMap3[firstKey][101][7.89] = 'Value C', but got %v", nestedMap3["firstKey"][101][7.89])
	}

	// Verify that non-existent levels are properly initialized
	if nestedMap3["thirdKey"] != nil {
		t.Errorf("Expected nestedMap3[thirdKey] to be nil, but got %v", nestedMap3["thirdKey"])
	}
}

func TestSetOrUpdateNestedMapValue(t *testing.T) {
	nestedMap := make(map[string]map[int]int)
	SetOrUpdateNestedMapValue(nestedMap, "firstKey", 100, 42, func(v int) int {
		return v + 1
	})
	if nestedMap["firstKey"][100] != 42 {
		t.Errorf("Expected nestedMap[firstKey][100] = 42, but got %v", nestedMap["firstKey"][100])
	}
	SetOrUpdateNestedMapValue(nestedMap, "firstKey", 100, 42, func(v int) int {
		return v + 1
	})
	if nestedMap["firstKey"][100] != 43 {
		t.Errorf("Expected nestedMap[firstKey][100] = 43, but got %v", nestedMap["firstKey"][100])
	}
}
