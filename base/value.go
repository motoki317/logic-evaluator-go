package base

type Bool struct {
	value bool
}

func (b *Bool) Set(newValue bool) {
	b.value = newValue
}

func (b *Bool) Get() bool {
	return b.value
}
