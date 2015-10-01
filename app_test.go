package poll

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"appengine/aetest"
)

func TestListPoll(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatalf("failed to create instance: %v", err)
	}
	defer inst.Close()

	r, err := inst.NewRequest("GET", "http://localhost:8080/_ah/api/sparck/v1/polls", nil)
	res, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Error(err)
	}

	assert := assert.New(t)
	assert.Equal(res.StatusCode, http.StatusOK, "Bad request status !")
}
