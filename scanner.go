// Package gozbar scanner bindings for golang.
// Read the ZBar documents for details
package gozbar

// #cgo LDFLAGS: -lzbar
// #include <zbar.h>
import "C"

import (
	"errors"
	"fmt"
)

// Scanner contains a reference to zbar scanner.
type Scanner struct {
	scanner *C.zbar_image_scanner_t
}

// NewScanner returns a new instance of Scanner.
func NewScanner() *Scanner {
	return &Scanner{
		scanner: C.zbar_image_scanner_create(),
	}
}

// SetConfig gives the zbar scanner the configuration to run..
// Read the ZBar docs for details.
func (s *Scanner) SetConfig(symbology C.zbar_symbol_type_t, config C.zbar_config_t, value int) error {
	// returns 0 for success, non-0 for failure
	resp := int(C.zbar_image_scanner_set_config(s.scanner, symbology, config, C.int(value)))
	if resp != 0 {
		return fmt.Errorf("error encountered when setting config status: %d", resp)
	}

	return nil
}

// Scan parses the given image and loads the scanner.
func (s *Scanner) Scan(img *Image) error {
	status := int(C.zbar_scan_image(s.scanner, img.image))
	switch status {
	// no symbols were found
	case 0:
		return errors.New("no symbols found")
	// an error occurred
	case -1:
		return errors.New("an error was encountered during scan")
	default:
		return nil
	}
}

// Destroy destroys the scanner.
func (s *Scanner) Destroy() {
	C.zbar_image_scanner_destroy(s.scanner)
}
