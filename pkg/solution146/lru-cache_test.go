package solution146

import (
	"testing"
)

func TestCase(t *testing.T) {
	cmd := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	args := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}
	out := []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4}

	t.Run("Case 1", func(t *testing.T) {
		var cache LRUCache
		for i, c := range cmd {
			switch c {
			case "LRUCache":
				cache = Constructor(args[i][0])
			case "put":
				cache.Put(args[i][0], args[i][1])
			case "get":
				o := cache.Get(args[i][0])
				if o != out[i] {
					t.Errorf("case failed of cmd at %d: %s, expected: %d, actual: %d \n", i, cmd[i], out[i], o)
					return
				}
			}
		}
	})
}
