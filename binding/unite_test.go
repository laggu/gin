package binding

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_uniteBinding_Name(t *testing.T) {
	b := NewUniteBinding(Query, JSON)
	assert.Equal(t, "unite", b.Name())
}

func Test_uniteBinding_Bind(t *testing.T) {
	var s struct {
		Body   string `json:"body"`
		Query  string `form:"query"`
		Header string `header:"header"`
		Param  string `uri:"param"`
	}

	req := requestWithBody("POST", "/?query=bar", `{"body": "foo"}`)
	req.Header.Add("header", "ping")
	c := testContext{req: req, params: map[string][]string{"param": {"pong"}}}

	b := NewUniteBinding(JSON, Query, Header, Uri)
	err := b.Bind(c, &s)
	require.NoError(t, err)

	assert.Equal(t, "foo", s.Body)
	assert.Equal(t, "bar", s.Query)
	assert.Equal(t, "ping", s.Header)
	assert.Equal(t, "pong", s.Param)
}

func Test_uniteBinding_Bind_invalid_type(t *testing.T) {
	var s struct {
		Body   int `json:"body"`
		Query  int `form:"query"`
		Header int `header:"header"`
		Param  int `uri:"param"`
	}

	req := requestWithBody("POST", "/?query=bar", `{"body": "foo"}`)
	req.Header.Add("header", "ping")
	c := testContext{req: req, params: map[string][]string{"param": {"pong"}}}

	b := NewUniteBinding(JSON, Query, Header, Uri)
	err := b.Bind(c, &s)
	require.Error(t, err)
}
