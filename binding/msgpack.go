// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !nomsgpack

package binding

import (
	"bytes"
	"github.com/ugorji/go/codec"
	"io"
)

type msgpackBinding struct{}

func (msgpackBinding) Name() string {
	return "msgpack"
}

func (b msgpackBinding) Bind(c context, obj any) error {
	if err := b.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (msgpackBinding) mapping(c context, obj any) error {
	return decodeMsgPack(c.GetRequest().Body, obj)
}

func (msgpackBinding) BindBody(body []byte, obj any) error {
	if err := decodeMsgPack(bytes.NewReader(body), obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeMsgPack(r io.Reader, obj any) error {
	cdc := new(codec.MsgpackHandle)
	if err := codec.NewDecoder(r, cdc).Decode(&obj); err != nil {
		return err
	}
	return nil
}
