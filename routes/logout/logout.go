package logout

import (
	"net/http"
	"net/url"
	"os"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	logoutUrl, err := url.Parse("https://login.microsoftonline.com/" + os.Getenv("tenantid") + "/oauth2/logout?client_id=" + os.Getenv("clientid"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
}
