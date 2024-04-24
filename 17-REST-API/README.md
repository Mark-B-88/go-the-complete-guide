# Project Description

A Go Powered **_Event Booking_** REST API

- GET `/events` Get a list of available events
- GET `/events/<id>` Get a list of avaialble events
- POST `/events` Create a new bookable event **_Auth Required_**
- PUT `/events/<id>` Update an event **_Auth Required - Only By Creator_**
- DELETE `/events/<id>` Delete an event **_Auth Required - Only By Creator_**
- POST `/signup` Create new user
- POST `/login` Authenticate user **_Auth Token - (JWT)_**
- POST `/events/<id>register` Register user for event **_Auth Required_**
- DELETE `/events/<id>/register` Cancel registration **_Auth Required_**

---

# Setting up 1st route & handling our 1st request

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // Configures HTTP Server

	server.GET("/api/v1/events", getEvents)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}
```

---

# Setting up an event model

Setting **_struct tags_** with the **_binding_** field will make those fields required.

```go
package models

import "time"

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

func (e Event)Save() {
	// later: add it to a database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
```

---

# Setting up a POST route

Asides from our initial GET route, we have a POST route that for now just has a hard-coded value for events, and also basic error handling.

```go
package main

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // Configures HTTP Server

	server.GET("/api/v1/events", getEvents)
	server.POST("/api/v1/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	context.JSON(http.StatusCreated, gin.H{"message": "event created", "event": event})
}
```

---

# Adding SQLite3

```go
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10) // pool of on-going connections
	DB.SetMaxIdleConns(5) // how many kept connections are open
}
```

---
