// Copyright 2017 Manu Martinez-Almeida. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (b queryBinding) Bind(c context, obj any) error {
	if err := b.mapping(c, obj); err != nil {
		return err
	}
	return validate(obj)
}

func (queryBinding) mapping(c context, obj any) error {
	values := c.GetRequest().URL.Query()
	if err := mapForm(obj, values); err != nil {
		return err
	}
	return nil
}
