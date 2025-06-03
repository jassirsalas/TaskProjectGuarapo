# Task Project API - Guarapo Labs
A small and simple API for task management. Users can add, update, delete or read all their tasks.

## Features
- User login
- Add, delete, update, read tasks by user
- Bearer Token authentication.

##  Requirements

- Go 1.20+
- Gin Gonic 1.10+
- Gorm 1.30+
- Go driver SQLite 1.5+
- (Optional) Go Air (https://github.com/air-verse/air) for live reload during development

## Endpoints
### Public
| Method |  Route   | Description      |
|--------|----------|------------------|
| POST   | `/login` | Log in the user  |

### Protected
| Method | Route         | Description                    |
|--------|---------------|--------------------------------|
| GET    | `/tasks`      | List all tasks for the user    |
| GET    | `/tasks/:id`  | Get a task by ID               |
| POST   | `/tasks`      | Create a new task              |
| PUT    | `/tasks/:id`  | Update an existing task        |
| DELETE | `/tasks/:id`  | Delete a task                  |

## How to test
Send a POST requets to the `/login` route for authentication. Example in the body add:
```
{
  "username": "user1"
}
```
"user1" can be change to any user you want, e.g., "user2"

Once you're authenticated you can acces to the protected routes. Make GET request to `/tasks` to see all the tasks available for `user1`


## How to run
The project can be executed simply by running:
```bash
go run ./cmd/api/main.go
```

Or using the make command:
```bash
make run
```
The server will run in port 8080
`http://localhost:8080`

For an easy, use make commands are available
Run build make command with tests
```bash
make all
```

Build the application
```bash
make build
```

Run the test suite:
```bash
make test
```

Clean up binary from the last build:
```bash
make clean
```

## Database
Simple SQLite database is already created for persistency in the API. Users `"user1"` and `"user2"` are already created with their corresping `tasks` for testing, but you can created and log in all the users you want.

`ID: int`
`Title: str`
`Completed: bool`
`Owner: str`

| ID | Title     | Completed | Owner  |
|----|-----------|-----------|--------|
| 1  | food      | 1         | user1  |
| 2  | laundry   | 0         | user1  |
| 3  | beer      | 1         | user2  |
| 4  | homework  | 0         | user2  |


Created with ♥️ by Jass
