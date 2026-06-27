package controllers

import (
	"net/http"
	"gohr/core"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	 

	
    var User = core.GetCurrentUser(r)

	

	data := map[string]interface{}{
		"Title": "Dashboard",
		"User":User,
	
	}

	core.RenderPage(w,r, "apps/users/views/dashboard.html", data)
}