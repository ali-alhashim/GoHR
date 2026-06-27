package controllers

import (
	"net/http"
	"gohr/core"
)

func Dashboard(w http.ResponseWriter, r *http.Request) {
	 

	
    

	

	data := map[string]interface{}{
		"Title": "Dashboard",
	
	}

	core.RenderPage(w,r, "apps/users/views/dashboard.html", data)
}