package solution380

import "math/rand"

type RandomizedSet struct {
	m map[int]int
	a []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		m: make(map[int]int),
		a: make([]int, 0),
	}
}

func (s *RandomizedSet) Insert(val int) bool {
	if _, ok := s.m[val]; ok {
		return false
	}
	s.m[val] = len(s.a)
	s.a = append(s.a, val)
	return true
}

func (s *RandomizedSet) Remove(val int) bool {
	if _, ok := s.m[val]; !ok {
		return false
	}
	last := len(s.a) - 1
	if val != s.a[last] {
		s.a[s.m[val]] = s.a[last]
		s.m[s.a[last]] = s.m[val]
	}
	s.a = s.a[:last]
	delete(s.m, val)
	return true
}

func (s *RandomizedSet) GetRandom() int {
	return s.a[rand.Intn(len(s.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
