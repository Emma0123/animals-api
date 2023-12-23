# Animals API Development Setup

## MySQL Workbench

1. Open MySQL Workbench.
2. Connect to one of the existing connections.
3. Create a new schema named "animals-api."

## VS Code

1. Open Visual Studio Code (VSCode).
2. Open the integrated terminal in VSCode.
   - Shortcut: ` Ctrl + `` (backtick) ` or navigate to `View -> Terminal`.
3. Clone the repository:
   ```bash
   `git clone https://github.com/Emma0123/animals-api.git`
   ```
4. Open the `main.go` file and locate the `dsn` variable.
   (Replace "user," "pass," and "dbname" with your MySQL Workbench credentials.)

   ```go
   dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

   ```

5. Return to the terminal in VS Code and run the server:
   `go run main.go`

## Postman

1. Open Postman.
2. In "My Workspace," go to "Collections."
3. Create a new collection with the "REST API basics" template.

#### Add Data

Use the POST method to add data:

- Path: `http://localhost:8080/v1/animal/create-animal`
- Body (JSON) :
  `{
    "name": "octopus",
    "class": "cephalopoda",
    "legs": 8
}`

#### Retrieve Data

1. Retrieve all data from the database using the GET method:
   - Path : `http://localhost:8080/v1/animal`
2. Retrieve a specific data based on ID using the GET method:
   - Path: `http://localhost:8080/v1/animal/3`
     (where 3 is the ID of an animal data)

#### Update Data

Update a data based on its ID using the PUT method:

- Path: `http://localhost:8080/v1/animal/3`
  (where 3 is the ID of an animal data)
- Body (JSON): Update the content as needed. If the data doesn't exist in the database, it will be added automatically.

#### Delete Data

Delete a data based on its ID using the DELETE method:

- Path: `http://localhost:8080/v1/animal/3`
  (where 3 is the ID of an animal data)
- The data with ID 3 will be deleted from the database.
