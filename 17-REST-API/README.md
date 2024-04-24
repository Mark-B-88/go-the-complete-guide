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

# Creating a SQL Table

```go
package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10) // pool of on-going connections
	DB.SetMaxIdleConns(5) // how many kept connections are open

	creteTables()
}

func creteTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table")
	}
}
```

---

# Changing Go Versions

Due to an issue with gcc complier and issues with MinGW configuration on Windows, I have decided to continue this project in Linux. During installation of Go on Linux, I am now running **_go version go1.19.8 linux/amd64_**, for this reason that is why go.mod was changed from version 1.22.2 to version 1.19

---

# Storing Data in the database

```go
package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID int
}

var events = []Event{}

func (e Event)Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
```

---

# Preparing Statements vs Directly Executing Queries (Prepare() vs Exec()/Query())

- `DB.Exec()` (when we created the tables)
- `Prepare() + stmt.Exec()` (when we inserted data into the database)
- `DB.Query()` (when we fetched data from the database)

Using `Prepare()` is 100% optional! You could send all your commands directly via `Exec()` or `Query()`.

The difference between those two methods then just is whether you're fetching data from the database (=> use `Query()`) or your manipulating the database / data in the database (=> use `Exec()`).

## But what's the advantage of using `Prepare()`?

`Prepare()` prepares a **_SQL statement_** - this can lead to better performance if the same statement is executed multiple times (potentially with different data for its placeholders).

This is only true, if the prepared statement is not closed `stmt.Close()` in between those executions. In that case, there wouldn't be any advantages.

And, indeed, in this application, we are calling `stmt.Close()` directly after calling `stmt.Exec()`. So here, it really wouldn't matter which approach you're using.

But in order to show you the different ways of using the sql package, I decided to also include this preparation approach in this course.

---
