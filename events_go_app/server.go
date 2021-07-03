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

type eventServiceHandler struct{
	dbhandler persistence.DatabaseHandler
}

// Constructor to initialize an eventServiceHandler object..
func newEventHandler (databaseHandler persistence.DatabaseHandler) *eventServiceHandler  {
	return &eventServiceHandler{
		dbhandler: databaseHandler,
	}	
}


func (eh *eventServiceHandler) findEventHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Returns a map of keys and values representing our request
	                    // URL variables and their values.
						// Resulting value is stored in the vars variable
						// first k/v pair: "SearchCriteria:name"
						// second k/v pair: "search"jazz concert"
	criteria, ok := vars["SearchCriteria"] // criteria variable will now have either
	                                       // name or id if the user sent the correct
										   // request URL. 
										   // The "ok" variable is type boolean, 
										   // if ok is true, then we will find a key
										   // called SearchCritera in our vars map.
										   // If it is false, then we know that the
										   // request URL we received is not valid.
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, `{error: No search criteria found, you can either search by id via /id/ 4
			to search by name via /name/coldplayconcert}`)
		return 
	}
	var event persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "name":
	event, err = eh.dbhandler.FindEventByName(searchkey)
	case "id":
	id, err := hex.DecodeString(searchkey)
	if err == nil {
		event, err = eh.dbhandler.FindEvent(id)
	}
	}
if err != nil {
	fmt.Fprintf(w, "{error %s", err)
	return 
}
w.Header().Set("Content-Type", "application/json;charset=utf8")
json.NewEncoder(w).Encode(&event)

}
// you just wrote an event handler call findEventHandler ^^^

func (eh *eventServiceHandler) allEventHandler(w http.ResponseWriter, r *http.Request)  {}
func (eh *eventServiceHandler) newEventHandler(w http.ResponseWriter, r *http.Request)  {}

func ServeAPI(endpoint string) error {
	handler := &eventservicehandler{}
	r := mux.NewRouter()
	// Create a subrouter for URLs prefixed with /events...
	eventsrouter := r.PathPrefix("/events").Subrouter()
	// Task - Searching for events via ID and name:
	// Define path and link to handler..
	// SearchCriteria and search are to variables in our path..
	// SearchCriteria can be replaced with id or name..
	eventsrouter.Methods("GET").Path("/{SearchCritera}/{search}").HandlerFunc(handler.findEventHandler)
	// Task - Retrieving all events at one - Relative URL is /events, method is GET, no data
	// expected in the HTTP body:
	eventsrouter.Methods("GET").Path("").HandlerFunc(handler.allEventHandler)
	// Task - Creating a new event - Relative URL is /events, method POST, expected data
	// in HTTP body is the JSON representation of the new event we are adding...
	eventsrouter.Methods("POST").Path("").HandlerFunc(handler.newEventHandler)
	return http.ListenAndServe(endpoint, r)
}

// Persistence layer
type DatabaseHandler interface {
	AddEvent(Event) ([]byte, error)
	FindEvent([]byte) (Event, error)
	FindEventByName(string) (Event, error)
	FindAllAvailableEvents() ([]Event, error)
}

// Persistence layer
package mongolayer
import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Persistence layer
// Create constants to represent the name of our database and the names
// of our mongodb collections involved in our persistence layer..
const (
	DB = "myevents"
	USERS = "users"
	EVENTS = "events"
)

// Persistence layer
// Database session object type *mgo.session
// Wrap the *mgo.session type in a struct type called MongoDBLayer..
type MongoDBLayer struct {
	session *mgo.Session
}

// Persistence layer
// Implement the DatabaseHandler interface to construct the concrete persistence
// layer..
// Make the implementor object type of the DatabaseHandler interface
// be a pointer to a MongoDBLayer struct object, or just simploy *MongoDBLayer..

// Create a constructor function called NewMongoDBLayer, which requires
// a single argument of type string..
// The argument represents the connection string with the information needed to
// establish the connection to the MongoDB database.
// Constructor function called NewMongoDBLayer

// Note, format of connection string needs to look like this..
// [mongodb://][user:pass@host1[:port1][,host2[:port2][/database][?options]
// or..
// mongodb://127.0.0.1
// port defaults to 27017
// Constructor function called NewMongoDBLayer
func NewMongoDBLayer(connection string) (*MongoDBLayer, error) { 
s, err:= mgo.Dial(connection) // mgo.Dial is the function in the mgo package
 if err!= nil{                // which will return a MongoDB connection session
	 return nil, err          // mgo.Dial(connection) returns a *mgo.Session object and
 }                            // an error object
return &MongoDBLayer{
	session: s,
},err
}

// Implement the methods of the DatabaseHandler interface..
// We have 4 methods..
// AddEvent(Event)
// FindEvent([]byte)
// FindEventByName(string)
// FindAllAvailableEvents()

// AddEvent method gives us a working *.mgo.Session object from the database
// connection pool to use in our code..
func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte,error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()  // Ensure this session gets returned back to the mgo database
	if !e.ID.Valid() {  // connection pool after the AddEvent() method exits.
		e.ID = bson.NewObjectId()
	}
	//let's assume the method below checks if the ID is valid for the location object of the event..
	if !e.Location.ID.Valid(){
		e.Location.ID = bson.NewObjectId()
	}
	// Return both the event ID of the added event, and an error object
	// representing the result of the event insertion operation..
	return []byte(e.ID), s.DB(DB).C(EVENTS).Insert(e) // e is the event object
}

// FindEvent() method.. Retrieve info of a certain event from the db..
func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{} // Create an empty event object
	err := s.DB(DB).C(EVENTS).FindId(bson.ObjectId(id)).One(&e) // access events collection
	                                                            // in the MongoDB.
	return e, err
}

// FindEventByName() method - retrieve an event by its name from the MongoDB database..
func (mgoLayer *MongoDBLayer) FindEventByName(name string)(persistence.Event, error){
	s := mgoLayer.getFreshSession()
	defer s.Close()
	e := persistence.Event{}
	err := s.DB(DB).C(EVENTS).Find(bson.M{"name": name}).One(&e)
	return e, err
}

// FindAllAvailableEvents() method..
// returns all available events in our db, i.e. returns entire events collection..
func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error){
  s := mgoLayer.getFreshSession()
  defer s.Close()
  events := []persistence.Event{}
  err := s.DB(DB).C(EVENTS).Find(nil).All(&events)
  return events, err
}

// Persistence layer
/* MONGO DB SETUP
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

// Our Mongodb database:
// MONGO DB DOCUMENT COLLECTIONS NEEDED FOR OUR EVENTS APP
// 1. Bookings collection
// 2. Events collection
// 3. Users collection

// Restful API handlers
// Implementing our Restful APIs handler functions..


