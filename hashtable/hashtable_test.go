package hashtable

import "testing"

func TestHashtable_InsertEmpty(t *testing.T) {
	ht := NewHashtable()
	ht.Insert("yo", 5)
	num, ok := ht.Get("yo")
	if !ok {
		t.Errorf("hashtable.Insert() expected valid insert")
	}

	if num != 5 {
		t.Errorf("hashtable.Insert() got %v, expected %v", num, 5)
	}
}

func TestHashtable_InsertDuplicate(t *testing.T) {
	ht := NewHashtable()
	ht.Insert("yo", 5)
	ht.Insert("yo", 10)

	num, ok := ht.Get("yo")
	if !ok {
		t.Errorf("hashtable.Insert() expected valid insert")
	}

	if num != 10 {
		t.Errorf("hashtable.Insert() got %v, expected %v", num, 10)
	}
}

func TestHashtable_InsertMany(t *testing.T) {
	// Make small Hashtable so we can check that collisions chain properly
	table := make([]*node, 3)
	ht := Hashtable{table}

	tests := []struct {
		key   string
		value int
	}{
		{"ephraim", 0},
		{"is", 1},
		{"a", 2},
		{"pretty", 3},
		{"new", 4},
		{"Go", 5},
		{"programmer", 6},
		{"and", 7},
		{"needs", 8},
		{"practice", 9},
	}

	for _, test := range tests {
		ht.Insert(test.key, test.value)
	}

	// Read backwards just to make sure
	for i := len(tests) - 1; i >= 0; i-- {
		num, ok := ht.Get(tests[i].key)
		if !ok {
			t.Errorf("hashtable.Insert() expected valid insert")
		}

		if num != tests[i].value {
			t.Errorf("hashtable.Insert() got %v, expected %v", num, tests[i].value)
		}
	}
}

func TestHashtable_GetFake(t *testing.T) {
	// Fail to find in top level bucket
	ht := NewHashtable()
	_, ok := ht.Get("yo")
	if ok {
		t.Errorf("hashtable.Insert() expected invalid Get")
	}

	ht = Hashtable{make([]*node, 1)} // Fail to find in linked list
	ht.Insert("Yo", 5)
	_, ok = ht.Get("ephraim")
	if ok {
		t.Errorf("hashtable.Insert() expected invalid Get")
	}
}

func TestHashtable_RemoveAtFront(t *testing.T) {
	// Make small Hashtable so we can check that collisions chain properly
	table := make([]*node, 3)
	ht := Hashtable{table}

	ht.Insert("Yo", 5)
	ok := ht.Remove("Yo")
	if !ok {
		t.Errorf("hashtable.Insert() expected valid Remove")
	}

	ok = ht.Remove("Yo")
	if ok {
		t.Errorf("hashtable.Insert() expected invalid Remove")
	}
}

func TestHashtable_RemoveFake(t *testing.T) {
	// Make small Hashtable so we can check that collisions chain properly
	table := make([]*node, 3)
	ht := Hashtable{table}

	ok := ht.Remove("Yo")
	if ok {
		t.Errorf("hashtable.Insert() expected invalid Remove")
	}
}

func TestHashtable_InsertAndRemoveMany(t *testing.T) {
	// Make small Hashtable so we can check that collisions chain properly
	table := make([]*node, 3)
	ht := Hashtable{table}

	tests := []struct {
		key   string
		value int
	}{
		{"ephraim", 0},
		{"is", 1},
		{"a", 2},
		{"pretty", 3},
		{"new", 4},
		{"Go", 5},
		{"programmer", 6},
		{"and", 7},
		{"needs", 8},
		{"practice", 9},
	}

	for _, test := range tests {
		ht.Insert(test.key, test.value)
	}

	// Remove randomly so we get in the middle of some chains
	ht.Remove(tests[5].key)
	for i := 0; i < len(tests); i++ {
		num, ok := ht.Get(tests[i].key)

		if i == 5 {
			if ok {
				t.Errorf("hashtable.Insert() expected item %d to not exist", i)
			}
		} else {
			if !ok {
				t.Errorf("hashtable.Insert() expected item %d to exist", i)
			}

			if num != tests[i].value {
				t.Errorf("hashtable.Insert() got %v, expected %v", num, tests[i].value)
			}
		}
	}
}
