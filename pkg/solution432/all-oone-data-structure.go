package solution432

type AllOne struct {
	data map[string]int
	keys []string
}

func Constructor() AllOne {
	return AllOne{data: make(map[string]int), keys: make([]string, 0)}
}

func (this *AllOne) Inc(key string) {
	if _, ok := this.data[key]; !ok {
		this.data[key] = 1
		this.keys = append(this.keys, key)
	} else {
		this.data[key]++
		this.forward(key)
	}
}

func (this *AllOne) Dec(key string) {
	if _, ok := this.data[key]; !ok {
		return
	}
	this.data[key]--
	if this.data[key] == 0 {
		delete(this.data, key)
		this.remove(key)
	} else {
		this.backward(key)
	}
}

func (this *AllOne) GetMaxKey() string {
	if len(this.keys) == 0 {
		return ""
	}
	return this.keys[0]
}

func (this *AllOne) GetMinKey() string {
	if len(this.keys) == 0 {
		return ""
	}
	return this.keys[len(this.keys)-1]
}

func (this *AllOne) forward(key string) {
	for i, s := range this.keys {
		if s == key {
			if i == 0 {
				return
			}
			if this.data[this.keys[i-1]] < this.data[key] {
				this.keys[i-1], this.keys[i] = this.keys[i], this.keys[i-1]
				this.forward(this.keys[i-1])
			}
		}
	}
}

func (this *AllOne) backward(key string) {
	for i, s := range this.keys {
		if s == key {
			if i == len(this.keys)-1 {
				return
			}
			if this.data[this.keys[i+1]] > this.data[key] {
				this.keys[i+1], this.keys[i] = this.keys[i], this.keys[i+1]
				this.backward(this.keys[i+1])
			}
		}
	}
}

func (this *AllOne) remove(key string) {
	for i, v := range this.keys {
		if v == key {
			this.keys = append(this.keys[:i], this.keys[i+1:]...)
			break
		}
	}
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
