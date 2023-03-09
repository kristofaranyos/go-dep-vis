package main

type DependencyMap map[string][]string

func (m DependencyMap) Add(k string, v []string) {
	if _, ok := m[k]; !ok {
		m[k] = v
		return
	}

	m[k] = append(m[k], v...)
}

func (m DependencyMap) Merge(additional DependencyMap) {
	for k, v := range additional {
		m.Add(k, v)
	}
}
