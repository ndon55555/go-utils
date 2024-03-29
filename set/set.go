package set

type Set interface {
	Add(...interface{})
	Remove(...interface{})
	Contains(...interface{}) bool
	Values() []interface{}
}

type HashSet struct {
	internal map[interface{}]bool
}

func NewHashSet(vals ...interface{}) *HashSet {
	h := &HashSet{internal: map[interface{}]bool{}}
	h.Add(vals...)
	return h
}

func (h *HashSet) Add(vals ...interface{}) {
	for _, v := range vals {
		h.internal[v] = true
	}
}

func (h *HashSet) Remove(vals ...interface{}) {
	for _, v := range vals {
		h.internal[v] = false
	}
}

func (h *HashSet) Contains(vals ...interface{}) bool {
	for _, v := range vals {
		if !h.internal[v] {
			return false
		}
	}

	return true
}

func (h *HashSet) Values() []interface{} {
	list := []interface{}{}
	for k, _ := range h.Values() {
		list = append(list, k)
	}

	return list
}
