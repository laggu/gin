// Copyright 2018 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (b uriBinding) Bind(c context, obj interface{}) error {
	err := b.mapping(c, obj)
	if err != nil {
		return err
	}
	return validate(obj)
}

func (uriBinding) mapping(c context, obj any) error {
	return mapURI(obj, c.GetParams())
}
