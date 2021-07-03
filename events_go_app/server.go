// Services
// Web UI
// Search
// Bookings
// Events
//   -Searching for events via...
//     -ID: /events/id/3434 (GET), No data expected in HTTP body
//     -Name: /events/name/jazz_concert (GET), No data expected in HTTP body
//   -Retrieving all events at once
//      - /events (GET), No data expected in HTTP body
//   -Creating a new event
//     -/events (POST), Expected data in HTTP body is the JSON for the new Event to add.
//     -HTTP body is just JSON
// Using Gorilla web toolkit, we can create a subrouter for the /events relative URL

import "github.com/gorilla/mux"

type eventServiceHandler struct{}

func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {}
func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request)  {}
func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request)  {}

func ServeAPI(endpoint string) error {
	handler := &eventservicehandler{}
	r := mux.NewRouter()
	// Create a subrouter for URLs prefixed with /events...
	eventsrouter := r.PathPrefix("/events").Subrouter()
	// Task - Searching for events via ID and name:
	eventsrouter.Methods("GET").Path("/{SearchCritera}/{search}").HandlerFunc(handler.findEventHandler)
	// Task - Retrieving all events at one - Relative URL is /events, method is GET, no data
	// expected in the HTTP body:
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	// Task - Creating a new event - Relative URL is /events, method POST, expected data
	// in HTTP body is the JSON representation of the new event we are adding...
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}

type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}

/* MONGO DB
wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
sudo apt-get install gnupg
wget -qO - https://www.mongodb.org/static/pgp/server-4.2.asc | sudo apt-key add -
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu bionic/mongodb-org/4.2 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-4.2.list

sudo apt-get update
sudo apt-get install -y mongodb-org
sudo nano /etc/init.d/mongod

#give permissions
sudo chmod +x /etc/init.d/mongod

#start the service
sudo service mongod start
Now, you can run mongo to reach the database.
*/

// MONGO DB DOCUMENT COLLECTIONS NEEDED FOR OUR EVENTS APP
// 1. Bookings collection
// 2. Events collection
// 3. Users collection

a