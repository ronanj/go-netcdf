// Copyright 2014 The Go-NetCDF Authors. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

// These files are autogenerated from nc_double.go using generate.sh

package netcdf

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <netcdf.h>
import "C"

// WriteInt writes data as the entire data for variable v.
func (v Var) WriteInt(data []int32) error {
	if err := okData(v, NC_INT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_put_var_int(C.int(v.f), C.int(v.id), (*C.int)(unsafe.Pointer(&data[0]))))
}

// ReadInt reads the entire variable v into data, which must have enough
// space for all the values (i.e. len(data) must be at least v.Len()).
func (v Var) ReadInt(data []int32) error {
	if err := okData(v, NC_INT, len(data)); err != nil {
		return err
	}
	return newError(C.nc_get_var_int(C.int(v.f), C.int(v.id), (*C.int)(unsafe.Pointer(&data[0]))))
}

// WriteInt sets the value of attribute a to val.
func (a Attr) WriteInt(val []int32) error {
	// We don't need okData here because netcdf library doesn't know
	// the length or type of the attribute yet.
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	return newError(C.nc_put_att_int(C.int(a.v.f), C.int(a.v.id), cname,
		C.nc_type(NC_INT), C.size_t(len(val)), (*C.int)(unsafe.Pointer(&val[0]))))
}

// ReadInt returns the attribute value.
func (a Attr) ReadInt(val []int32) (err error) {
	if err := okData(a, NC_INT, len(val)); err != nil {
		return err
	}
	cname := C.CString(a.name)
	defer C.free(unsafe.Pointer(cname))
	err = newError(C.nc_get_att_int(C.int(a.v.f), C.int(a.v.id), cname,
		(*C.int)(unsafe.Pointer(&val[0]))))
	return
}
