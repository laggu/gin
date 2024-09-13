// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"encoding/xml"
	"io"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (b xmlBinding) Bind(c context, obj any) error {
	if err := b.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (xmlBinding) mapping(c context, obj any) error {
	return decodeXML(c.GetRequest().Body, obj)
}

func (xmlBinding) BindBody(body []byte, obj any) error {
	if err := decodeXML(bytes.NewReader(body), obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeXML(r io.Reader, obj any) error {
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
