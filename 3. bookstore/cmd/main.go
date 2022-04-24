package main

/**
This main file has the responsibility of creating and starting the server,
and inform the application where the router resides.
**/
import (
	"log"
	"net/http"

	"github.com/abhisenberg/bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/**
The flow of this backend webapp goes like this:
1. User makes some request, it reaches the routes module
2. From routes, the appropriate handler is called in the controller layer
3. The controller layer holds the business logic of the request, and accordingly informs
	the model layer about the changes that need to be done in database
4. The control goes to the model layer which talks directly to the DB and makes the required
	changes in the DB.

	Routes -> Controller -> Model
**/
func main() {
	r := mux.NewRouter()                                //Initializing a new router here
	routes.RegisterBookStoreRoutes(r)                   //Registering our declared routes using the Router object
	http.Handle("/", r)                                 //❗❗ Why this?
	log.Fatal(http.ListenAndServe("localhost:9010", r)) //Attempt to start and listen to the server on the port that is mentioned i.e. 9010, and log any errors.
}
