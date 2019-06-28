package my_request

import (
	"testing"
	"io/ioutil"
)

func TestNew_Client(t *testing.T) {
	cookieStr := "phpdisk_info=UG0EMg1pUWkCMAVsWwhTPVQNBDdbMQBgBTdXMFVrBzNYZVBlBmQAO1dcB28AalRrVWQHNFpiXGlSNwU0UmAANlA3BGINOlFsAjQFNltlUz5UYgRgWzIAMgVjV2JVZgdnWGVQYAZhADxXNwdeAGtUPVUyBzRaNFwzUmcFbFJjADJQYw%3D%3D"
	c := NewMyClient(false)
	resp, err := c.Request("https://up.woozooo.com/mydisk.php", "get", cookieStr, nil, nil, nil)
	if nil != err {
		t.Error(err)
	}
	b, e:= ioutil.ReadAll(resp.Body)
	if nil != e {
		t.Error(e)
	}
	t.Log(string(b), resp.Status, resp.Header.Get("Location"))
}
