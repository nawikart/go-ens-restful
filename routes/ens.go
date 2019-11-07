package routes

import (
	c "../controllers/ens"
	"github.com/gorilla/mux"
)

func Ens(r *mux.Router) {

	r.HandleFunc("/ens/resolve", c.Resolve).Methods("POST")
	r.HandleFunc("/ens/reverse-resolve", c.ReverseResolve).Methods("POST")
	r.HandleFunc("/ens/newname", c.NewName).Methods("POST")
	r.HandleFunc("/ens/set-owner", c.SetOwner).Methods("POST")
	r.HandleFunc("/ens/create-subdomain", c.CreateSubdomain).Methods("POST")
	r.HandleFunc("/ens/set-subdomain-owner", c.SetSubdomainOwner).Methods("POST")
	r.HandleFunc("/ens/set-resolver", c.SetResolver).Methods("POST")
	r.HandleFunc("/ens/set-address", c.SetAddress).Methods("POST")
	r.HandleFunc("/ens/set-contenthash", c.SetContenthash).Methods("POST")
	r.HandleFunc("/ens/set-text", c.SetText).Methods("POST")
	r.HandleFunc("/ens/set-abi", c.SetABI).Methods("POST")
	r.HandleFunc("/ens/set-name", c.SetName).Methods("POST")
}
