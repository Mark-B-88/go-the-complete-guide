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
