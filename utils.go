package poll

import (
	"errors"
	"net/url"

	"go-endpoints/endpoints"

	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

type QueryMarker struct {
	datastore.Cursor
}

func (qm *QueryMarker) MarshalJSON() ([]byte, error) {
	return []byte(`"` + qm.String() + `"`), nil
}

func (qm *QueryMarker) UnmarshalJSON(buf []byte) error {
	if len(buf) < 2 || buf[0] != '"' || buf[len(buf)-1] != '"' {
		return errors.New("QueryMarker: bad cursor value")
	}
	cursor, err := datastore.DecodeCursor(string(buf[1 : len(buf)-1]))
	if err != nil {
		return err
	}
	*qm = QueryMarker{cursor}
	return nil
}

func checkReferer(c context.Context) error {
	if appengine.IsDevAppServer() {
		return nil
	}

	request := endpoints.HTTPRequest(c)
	r := request.Referer()
	u, err := url.Parse(r)
	if err != nil {
		return endpoints.NewUnauthorizedError("couldn't extract domain from referer")
	}

	if u.Host != appengine.AppID(c)+".appspot.com" {
		return endpoints.NewUnauthorizedError("referer unauthorized")
	}

	return nil
}
