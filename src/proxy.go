package extract

import (
	"net/http"
	"appengine"
	"appengine/urlfetch"
	"io/ioutil"
)

//
// Just forward the response from another site, to the user.
// Used to bypass the "Access-Control-Allow-Origin" for simple non-private resources.
// 
func serveDistant(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")

	c := appengine.NewContext(r)
	resp, err := urlfetch.Client(c).Get(url)
	if err != nil {
		c.Errorf("%v", err)
		return
	}
	defer resp.Body.Close()
	x, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Errorf("%v", err)
		return
	}
	_, err = w.Write( x )
	if err != nil {
		c.Errorf("%v", err)
		return
	}
	
}