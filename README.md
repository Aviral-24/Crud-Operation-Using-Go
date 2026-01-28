main.go: Starts an HTTP server on port 8080 and sets up two routes:

users - for listing and creating users
users - for updating and deleting specific users
handler.go: Contains two handler functions:

UsersHandler - handles POST (create) and GET (read all) requests
UserHandler - handles PUT (update) and DELETE requests
db.go: Manages database connection

users.go: Defines the User data model

Technologies Used
Go and SQLite
