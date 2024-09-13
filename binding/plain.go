package binding

import (
	"fmt"
	"io"
	"reflect"

	"github.com/gin-gonic/gin/internal/bytesconv"
)

type plainBinding struct{}

func (plainBinding) Name() string {
	return "plain"
}

func (b plainBinding) Bind(c context, obj interface{}) error {
	return b.mapping(c, obj)
}

func (plainBinding) mapping(c context, obj interface{}) error {
	all, err := io.ReadAll(c.GetRequest().Body)
	if err != nil {
		return err
	}

	return decodePlain(all, obj)
}

func (plainBinding) BindBody(body []byte, obj any) error {
	return decodePlain(body, obj)
}

func decodePlain(data []byte, obj any) error {
	if obj == nil {
		return nil
	}

	v := reflect.ValueOf(obj)

	for v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return nil
		}
		v = v.Elem()
	}

	if v.Kind() == reflect.String {
		v.SetString(bytesconv.BytesToString(data))
		return nil
	}

	if _, ok := v.Interface().([]byte); ok {
		v.SetBytes(data)
		return nil
	}

	return fmt.Errorf("type (%T) unknown type", v)
}
