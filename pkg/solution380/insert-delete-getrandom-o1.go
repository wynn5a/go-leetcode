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

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.m[val]; ok {
		return false
	}
	this.m[val] = len(this.a)
	this.a = append(this.a, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.m[val]; !ok {
		return false
	}
	last := len(this.a) - 1
	if val != this.a[last] {
		this.a[this.m[val]] = this.a[last]
		this.m[this.a[last]] = this.m[val]
	}
	this.a = this.a[:last]
	delete(this.m, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.a[rand.Intn(len(this.a))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
