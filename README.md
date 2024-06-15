# go-todo-api
This is a REST API for ToDo application using Gin and Gorm.

## Usage
### Run Server
```
$ git clone git@github.com:koizumi7010/go-todo-api.git
$ cd go-todo-api
$ docker-compose up -d 
$ docker exec -d go-todo-api go run cmd/go-todo-api/main.go
```

### End points
| Method  | Path | Description |
| ------------- | ------------- | ------------- |
| GET  | /todo  | Get all task list |
| GET  | /todo/{id}  | Get a task |
| POST  | /todo | Create a new task |
| PUT  | /todo/{id}  | Update a task |
| DELETE  | /todo/{id}  | Delete a task |

### API call samples
```
# Get all task list
$ curl -i -XGET localhost/todo

# Get a task
$ curl -i -XGET localhost/todo/1

# Create a new task
$ curl -i localhost/todo -H "Content-Type: application/json" -X POST -d '{"task": "test1"}' 

# Update a task
$ curl -i localhost/todo/1 -H "Content-Type: application/json" -X PUT -d '{"task": "test1","status": "done"}'

# Delete a task
$ curl -i localhost/todo/1 -X DELETE

```
