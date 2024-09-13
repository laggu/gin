// Copyright 2018 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io"
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (b yamlBinding) Bind(c context, obj any) error {
	if err := b.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (yamlBinding) mapping(c context, obj any) error {
	return decodeYAML(c.GetRequest().Body, obj)
}

func (yamlBinding) BindBody(body []byte, obj any) error {
	if err := decodeYAML(bytes.NewReader(body), obj); err != nil {
		return err
	}
	return validate(obj)
}

func decodeYAML(r io.Reader, obj any) error {
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
