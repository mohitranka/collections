package skiplist

import "testing"

func TestCreateSkipList(t *testing.T) {
	sl := NewSkipList(3, 0.5)
	if sl.MaxLevel != 3 {
		t.Fatalf("Expected '3' as MaxLevel, found %d\n", sl.MaxLevel)
	}

	if sl.P != 0.5 {
		t.Fatalf("Expected	'0.5' as P, found %v\n", sl.P)
	}
}

func TestItems(t *testing.T) {
	sl := NewSkipList(3, 0.5)
	sl.Insert(10)
	sl.Insert(20)

	if sl.Search(10) != false {
		t.Fatal("Expected 10 in the skip list, did not find")
	}

	if sl.Search(20) != false {
		t.Fatal("Expected 20 in the skip list, did not find")
	}

	if sl.Search(30) != false {
		t.Fatal("Did not expect 30 in the skip list, still found it")
	}

	sl.Delete(10)
	sl.Delete(20)

	if sl.Search(10) != false {
		t.Fatal("Did not expect 10 in the skip list, still found it")
	}

	if sl.Search(20) != false {
		t.Fatal("Did not expect 20 in the skip list, still found it")
	}
}
