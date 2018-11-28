package core

import (
	// "fmt"
	"github.com/gorilla/mux"
	"net/http"
	// "os"
)

type Router struct {
	router *mux.Router
}

func (r *Router) newRouter() {
	r.router = mux.NewRouter()
}

func (r *Router) serv(port string) error {

	// router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	// router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	// router.HandleFunc("/api/contacts/new", controllers.CreateContact).Methods("POST")
	// router.HandleFunc("/api/me/contacts", controllers.GetContactsFor).Methods("GET") //  user/2/contacts

	// router.Use(app.JwtAuthentication) //attach JWT auth middleware

	//router.NotFoundHandler = app.NotFoundHandler

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8008" //localhost
	// }

	// fmt.Println(port)

	return http.ListenAndServe(":"+port, r.router) //Launch the app, visit localhost:8000/api
	// if err != nil {
	// 	fmt.Print(err)
	// }
}
