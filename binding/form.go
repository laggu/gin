// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"errors"
	"net/http"
)

const defaultMemory = 32 << 20

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

func (formBinding) Name() string {
	return "form"
}

func (b formBinding) Bind(c context, obj any) error {
	if err := b.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (formBinding) mapping(c context, obj any) error {
	if err := c.GetRequest().ParseForm(); err != nil {
		return err
	}
	if err := c.GetRequest().ParseMultipartForm(defaultMemory); err != nil && !errors.Is(err, http.ErrNotMultipart) {
		return err
	}
	if err := mapForm(obj, c.GetRequest().Form); err != nil {
		return err
	}
	return nil
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (f formPostBinding) Bind(c context, obj any) error {
	if err := f.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) mapping(c context, obj any) error {
	if err := c.GetRequest().ParseForm(); err != nil {
		return err
	}
	if err := mapForm(obj, c.GetRequest().PostForm); err != nil {
		return err
	}
	return nil
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (f formMultipartBinding) Bind(c context, obj any) error {
	if err := f.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (formMultipartBinding) mapping(c context, obj any) error {
	if err := c.GetRequest().ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := mappingByPtr(obj, (*multipartRequest)(c.GetRequest()), "form"); err != nil {
		return err
	}
	return nil
}
