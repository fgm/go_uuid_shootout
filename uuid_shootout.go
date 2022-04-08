package uuid_shootout

import "strconv"

// NilWriter is an io.StringWriter with near-zero cost.
type NilWriter struct{}

func (w NilWriter) WriteString(string) (int, error) {
	return 0, nil
}

// UseString pretends to use its argument at a minimal cost.
func UseString(x string) {
	w := &NilWriter{}
	w.WriteString(x)
}

// UseInt pretends to use its argument at a minimal cost.
func UseInt(n int) {
	w := &NilWriter{}
	w.WriteString(strconv.Itoa(n))
}
