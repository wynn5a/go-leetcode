package solution432

type AllOne struct {
	data map[string]int
	keys []string
}

func Constructor() AllOne {
	return AllOne{data: make(map[string]int), keys: make([]string, 0)}
}

func (o *AllOne) Inc(key string) {
	if _, ok := o.data[key]; !ok {
		o.data[key] = 1
		o.keys = append(o.keys, key)
	} else {
		o.data[key]++
		o.forward(key)
	}
}

func (o *AllOne) Dec(key string) {
	if _, ok := o.data[key]; !ok {
		return
	}
	o.data[key]--
	if o.data[key] == 0 {
		delete(o.data, key)
		o.remove(key)
	} else {
		o.backward(key)
	}
}

func (o *AllOne) GetMaxKey() string {
	if len(o.keys) == 0 {
		return ""
	}
	return o.keys[0]
}

func (o *AllOne) GetMinKey() string {
	if len(o.keys) == 0 {
		return ""
	}
	return o.keys[len(o.keys)-1]
}

func (o *AllOne) forward(key string) {
	for i, s := range o.keys {
		if s == key {
			if i == 0 {
				return
			}
			if o.data[o.keys[i-1]] < o.data[key] {
				o.keys[i-1], o.keys[i] = o.keys[i], o.keys[i-1]
				o.forward(o.keys[i-1])
			}
		}
	}
}

func (o *AllOne) backward(key string) {
	for i, s := range o.keys {
		if s == key {
			if i == len(o.keys)-1 {
				return
			}
			if o.data[o.keys[i+1]] > o.data[key] {
				o.keys[i+1], o.keys[i] = o.keys[i], o.keys[i+1]
				o.backward(o.keys[i+1])
			}
		}
	}
}

func (o *AllOne) remove(key string) {
	for i, v := range o.keys {
		if v == key {
			o.keys = append(o.keys[:i], o.keys[i+1:]...)
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
