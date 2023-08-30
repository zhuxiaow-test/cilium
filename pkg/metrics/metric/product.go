package metric

type iterator struct {
	d []string
	i int
	n int
}

func (i *iterator) reset() {
	i.i = 0
}

func (i *iterator) Next() ([]string, bool) {
	index := i.i
	if index >= len(i.d) {
		return nil, false
	}

	r := i.d[index : index+i.n]
	i.i += i.n
	return r, true
}

func (i *iterator) len() int {
	return len(i.d)
}

func product(a []string, b []string) *iterator {
	if len(a) == 0 {
		return nil
	}
	if len(b) == 0 {
		return nil
	}
	p := make([]string, 0, len(a)*len(b))
	for _, x := range a {
		for _, y := range b {
			p = append(p, x)
			p = append(p, y)
		}
	}
	return &iterator{
		d: p,
		i: 0,
		n: 2,
	}
}
func product2(a *iterator, b []string) *iterator {
	p := make([]string, 0, a.len()*len(b))
	for {
		us, ok := a.Next()
		if !ok {
			break
		}
		for _, v := range b {
			p = append(p, us...)
			p = append(p, v)
		}
	}
	return &iterator{
		d: p,
		i: 0,
		n: a.n + 1,
	}
}

func product3(a *iterator, b *iterator) *iterator {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	p := make([]string, 0, a.len()*b.len())
outer:
	for {
		us, ok := a.Next()
		if !ok {
			break outer
		}
	inner:
		for {
			vs, ok := b.Next()
			if !ok {
				b.reset()
				break inner
			}
			p = append(p, us...)
			p = append(p, vs...)
		}
	}
	return &iterator{
		d: p,
		i: 0,
		n: a.n + 1,
	}
}

func (i *iterator) ForEach(fn func([]string)) {
	for {
		us, ok := i.Next()
		if !ok {
			break
		}
		fn(us)
	}
}

func Product(vs ...*iterator) *iterator {
	if len(vs) == 0 {
		return nil
	}

	if len(vs) == 1 {
		return vs[0]
	}

	lhs := product3(vs[0], vs[1])
	rhs := Product(vs[2:]...)
	return product3(lhs, rhs)
}

func vec1(d []string) *iterator {
	return &iterator{
		d: d,
		i: 0,
		n: 1,
	}
}
