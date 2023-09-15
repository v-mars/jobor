package notify

import (
	"net/http"
	"testing"
)

func Test_JSONPost(t *testing.T) {
	_, err := JSONPost(http.MethodPost, "http://webhook.test", nil, http.DefaultClient)
	if err != nil {
		t.Fatal(err)
	}
}
