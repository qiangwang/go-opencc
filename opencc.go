package opencc

/*
#cgo LDFLAGS: -lopencc
#include<stdlib.h>
#include<opencc/opencc.h>
//FYI: https://github.com/BYVoid/OpenCC/blob/master/src/opencc.h#L70
opencc_t opencc_open_or_null(const char *configFileName) {
	opencc_t t = opencc_open(configFileName);
	if ((long)t == -1) {
		return NULL;
	}
	return t;
}
*/
import "C"

import (
	"errors"
	"unsafe"
)

// Converter cache opencc instance
type Converter struct {
	cache map[string]C.opencc_t
}

// Convert convert by format: t2s, s2t ..
func (converter *Converter) Convert(text string, format string) (string, error) {
	cc, ok := converter.cache[format]
	if !ok {
		path := "/usr/share/opencc/" + format + ".json"

		cpath := C.CString(path)
		defer C.free(unsafe.Pointer(cpath))

		cc = C.opencc_open_or_null(cpath)
		if cc == nil {
			return "", errors.New("invalid format")
		}

		if converter.cache == nil {
			converter.cache = map[string]C.opencc_t{}
		}
		converter.cache[format] = cc
	}

	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	cOut := C.opencc_convert_utf8(cc, ctext, C.size_t(len(text)))
	if cOut == nil {
		return "", errors.New("convert fail")
	}

	defer C.opencc_convert_utf8_free(cOut)

	return C.GoString(cOut), nil
}

// Close close instance and clean cache
func (converter *Converter) Close() {
	for _, cc := range converter.cache {
		C.opencc_close(cc)
	}
	converter.cache = map[string]C.opencc_t{}
}
