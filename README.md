# publisher

COMMAND : go run booking.go <br />

POST : /booking <br />
JSON BODY example : {
	"username": "acep",
	"destination": "Jakarta"
} <br />

# consumer
go run notification.go
