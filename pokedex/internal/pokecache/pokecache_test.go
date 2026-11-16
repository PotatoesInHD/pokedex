package pokeapi
import "testing"
import "bytes"

func TestCache(t *testing.T) {
	t.Run("add and get existing key", func(t *testing.T) {
		c := NewCache()
		c.AddCache("banana", []byte{1, 2, 3})
		val, ok := c.GetCache("banana")
		expected := []byte{1, 2, 3}
		if !ok || !bytes.Equal(expected, val) {
			t.Errorf("Expected: %v Got: %v", expected, val)
		}
	})

	t.Run("add an existing key", func(t *testing.T) {
		c := NewCache()
		c.AddCache("banana", []byte{1, 2, 3})
		c.AddCache("banana", []byte{3, 2, 1})
		expected := []byte{3, 2, 1}
		val, ok := c.GetCache("banana")
		if !ok || !bytes.Equal(expected, val) {
			t.Errorf("Expected: %v Got: %v", expected, val)
		}
	})

	t.Run("get key that doesnt exist", func(t *testing.T) {
		c := NewCache()
		val, ok := c.GetCache("banana")
		if !ok && !bytes.Equal(val, nil) {
			t.Errorf("Expected: %v Got: %v\nExpected: false Got: %v", nil, val, ok)
		}
	})

	t.Run("add 2 values get 2 values", func(t *testing.T) {
		c := NewCache()
		c.AddCache("banana", []byte{1, 2, 3})
		c.AddCache("apple", []byte{3, 2, 1})
		val1, _ := c.GetCache("banana")
		val2, _ := c.GetCache("apple")
		expected1 :=[]byte{1, 2, 3}
		expected2 :=[]byte{3, 2, 1}
		if !bytes.Equal(val1, expected1) ||
		!bytes.Equal(val2, expected2) {
			t.Errorf("\n1- Expected: %v Got: %v\n2- Expected: %v Got: %v",val1, expected1, val2, expected2,)
		}
	})

}
