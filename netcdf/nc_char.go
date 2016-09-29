// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.go
// DO NOT EDIT (except nc_double.go).

package netcdf

import (
	"fmt"
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteBytes writes data as the entire data for variable v.
func (v Var) WriteBytes(data []byte) error {
	if err := okData(v, CHAR, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_text(C.int(v.ds), C.int(v.id), (*C.char)(unsafe.Pointer(&data[0]))))
}

// ReadBytes reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadBytes(data []byte) error {
	if err := okData(v, CHAR, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_text(C.int(v.ds), C.int(v.id), (*C.char)(unsafe.Pointer(&data[0]))))
}

// WriteBytes sets the value of attribute a to val.
func (a Attr) WriteBytes(val []byte) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_text(C.int(a.v.ds), C.int(a.v.id), cname,
		C.size_t(len(val)), (*C.char)(unsafe.Pointer(&val[0]))))
}

// ReadBytes reads the entire attribute value into val.
func (a Attr) ReadBytes(val []byte) (err error) {
	if err := okData(a, CHAR, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_text(C.int(a.v.ds), C.int(a.v.id), cname,
		(*C.char)(unsafe.Pointer(&val[0]))))
	return
}

// BytesReader is a interface that allows reading a sequence of values of fixed length.
type BytesReader interface {
	Len() (n uint64, err error)
	ReadBytes(val []byte) (err error)
}

// GetBytes reads the entire data in r and returns it.
func GetBytes(r BytesReader) (data []byte, err error) {
	n, err := r.Len()
	if err != nil {
		return
	}
	data = make([]byte, n)
	err = r.ReadBytes(data)
	return
}

// testReadBytes writes somes data to v. N is v.Len().
// This function is only used for testing.
func testWriteBytes(v Var, n uint64) error {
	data := make([]byte, n)
	for i := 0; i < int(n); i++ {
		data[i] = byte(i + 10)
	}
	return v.WriteBytes(data)
}

// testReadBytes reads data from v and checks that it's the same as what
// was written by testWriteDouble. N is v.Len().
// This function is only used for testing.
func testReadBytes(v Var, n uint64) error {
	data := make([]byte, n)
	if err := v.ReadBytes(data); err != nil {
		return err
	}
	for i := 0; i < int(n); i++ {
		if val := byte(i + 10); data[i] != val {
			return fmt.Errorf("data at position %d is %v; expected %v\n", i, data[i], val)
		}
	}
	return nil
}

// GetBytes reads the entire data in r and returns it.
func (v Var) GetBytes() (data []byte, err error) {
	return GetBytes(v)
}

// GetBytes reads the entire data in r and returns it.
func (a Attr) GetBytes() (data []byte, err error) {
	return GetBytes(a)
}
