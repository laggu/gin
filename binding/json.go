// Copyright 2014 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin/internal/json"
	"io"
)

// EnableDecoderUseNumber is used to call the UseNumber method on the JSON
// Decoder instance. UseNumber causes the Decoder to unmarshal a number into an
// any as a Number instead of as a float64.
var EnableDecoderUseNumber = false

// EnableDecoderDisallowUnknownFields is used to call the DisallowUnknownFields method
// on the JSON Decoder instance. DisallowUnknownFields causes the Decoder to
// return an error when the destination is a struct and the input contains object
// keys which do not match any non-ignored, exported fields in the destination.
var EnableDecoderDisallowUnknownFields = false

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return "json"
}

func (b jsonBinding) Bind(c context, obj any) error {
	if c == nil || c.GetRequest().Body == nil {
		return errors.New("invalid request")
	}

	if err := b.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (jsonBinding) mapping(c context, obj any) error {
	return decodeJSON(c.GetRequest().Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj any) error {
	if err := decodeJSON(bytes.NewReader(body), obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeJSON(r io.Reader, obj any) error {
	decoder := json.NewDecoder(r)
	if EnableDecoderUseNumber {
		decoder.UseNumber()
	}
	if EnableDecoderDisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
