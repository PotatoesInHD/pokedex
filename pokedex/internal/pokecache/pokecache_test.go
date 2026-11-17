package pokeapi
import "testing"
import "bytes"
import "fmt"
import "time"

func TestCache(t *testing.T) {
	t.Run("add and get existing key", func(t *testing.T) {
		c := NewCache(5 * time.Second)
		c.AddCache("banana", []byte{1, 2, 3})
		val, ok := c.GetCache("banana")
		expected := []byte{1, 2, 3}
		fmt.Printf("teesssst %v \n",c.data["banana"].val)
		if !ok || !bytes.Equal(expected, val) {
			t.Errorf("Expected: %v Got: %v", expected, val)
		}
	})

	t.Run("add an existing key", func(t *testing.T) {
		c := NewCache(5 * time.Second)
		c.AddCache("banana", []byte{1, 2, 3})
		c.AddCache("banana", []byte{3, 2, 1})
		expected := []byte{3, 2, 1}
		val, ok := c.GetCache("banana")
		if !ok || !bytes.Equal(expected, val) {
			t.Errorf("Expected: %v Got: %v", expected, val)
		}
	})

	t.Run("get key that doesnt exist", func(t *testing.T) {
		c := NewCache(5 * time.Second)
		val, ok := c.GetCache("banana")
		if !ok && !bytes.Equal(val, nil) {
			t.Errorf("Expected: %v Got: %v\nExpected: false Got: %v", nil, val, ok)
		}
	})

	t.Run("add 2 values get 2 values", func(t *testing.T) {
		c := NewCache(5 * time.Second)
		c.AddCache("banana", []byte{1, 2, 3})
		c.AddCache("apple", []byte{3, 2, 1})
		val1, _ := c.GetCache("banana")
		val2, _ := c.GetCache("apple")
		expected1 :=[]byte{1, 2, 3}
		expected2 :=[]byte{3, 2, 1}
		if !bytes.Equal(val1, expected1) ||
		!bytes.Equal(val2, expected2) {
			t.Errorf("\n1- Expected: %v Got: %v\n2- Expected: %v Got: %v",expected1, val1, expected2, val2)
		}
	})

	t.Run("add value make sure stays for interval", func(t *testing.T) {
		c := NewCache(200 * time.Millisecond)
		c.AddCache("banana", []byte{1, 2, 3})
		time.Sleep(100 * time.Millisecond)
		val, _ := c.GetCache("banana")
		expected := []byte{1, 2, 3}
		if !bytes.Equal(val, expected) {
			t.Errorf("Expected: %v Got: %v", expected, val)
		}
	})

	t.Run("add value make gets deleted after interval", func(t *testing.T) {
		c := NewCache(200 * time.Millisecond)
		c.AddCache("banana", []byte{1, 2, 3})
		time.Sleep(400 * time.Millisecond)
		val, _ := c.GetCache("banana")
		if !bytes.Equal(val, nil) {
			t.Errorf("Expected: %v Got: %v", nil, val)
		}
	})
}
