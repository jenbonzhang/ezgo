package ezgo

// SetNestedMapValue handles nested maps with different key types at each level
func SetNestedMapValue[K1 comparable, K2 comparable, V any](
	m map[K1]map[K2]V,
	key1 K1,
	key2 K2,
	value V,
) {
	// Check if the inner map exists
	if m[key1] == nil {
		// If not, initialize it
		m[key1] = make(map[K2]V)
	}

	// Set the value in the inner map
	m[key1][key2] = value
}

// SetNestedMapValue3 sets a value in a three-level nested map with different key types at each level
func SetNestedMapValue3[K1 comparable, K2 comparable, K3 comparable, V any](
	m map[K1]map[K2]map[K3]V,
	key1 K1,
	key2 K2,
	key3 K3,
	value V,
) {
	// Check if the first level map exists
	if m[key1] == nil {
		m[key1] = make(map[K2]map[K3]V)
	}

	// Check if the second level map exists
	if m[key1][key2] == nil {
		m[key1][key2] = make(map[K3]V)
	}

	// Set the value in the third level map
	m[key1][key2][key3] = value
}

func SetOrUpdateNestedMapValue[K1 comparable, K2 comparable, V any](
	m map[K1]map[K2]V,
	key1 K1,
	key2 K2,
	value V,
	f func(v V) V,
) {
	// Check if the inner map exists
	if m[key1] == nil {
		// If not, initialize it
		m[key1] = make(map[K2]V)
	}
	if v, ok := m[key1][key2]; ok {
		// Set the value in the inner map
		m[key1][key2] = f(v)
	} else {
		// Set the value in the inner map
		m[key1][key2] = value
	}
}
