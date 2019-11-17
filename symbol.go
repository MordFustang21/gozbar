// Package gozbar symbol bindings.
// Read the ZBar documents for details
package gozbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

// Symbol is a wrapper around a zbar symbol.
type Symbol struct {
	symbol *C.zbar_symbol_t
}

// Next returns the next symbol or nil if there is none.
func (s *Symbol) Next() *Symbol {
	n := C.zbar_symbol_next(s.symbol)

	if n == nil {
		return nil
	}

	return &Symbol{
		symbol: n,
	}
}

// Data returns the scanned data for this symbol.
func (s *Symbol) Data() string {
	sym := C.zbar_symbol_get_data(s.symbol)

	if sym == nil {
		return ""
	}

	return C.GoString(sym)
}

// Type returns the symbol type.
// Compare it with types in constants to get the accurate symbol type.
func (s *Symbol) Type() C.zbar_symbol_type_t {
	return C.zbar_symbol_get_type(s.symbol)
}

// Each will iterate over all symbols after this symbol.
// passing them into the provided callback
func (s *Symbol) Each(f func(string)) {
	t := s

	for {
		f(t.Data())

		t = t.Next()

		if t == nil {
			break
		}
	}
}
