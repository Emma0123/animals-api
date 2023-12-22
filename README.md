# Animals API Development Setup

## MySQL Workbench
1. Open MySQL Workbench.
2. Connect to one of the existing connections.
3. Create a new schema named "animals-api."

## VS Code
1. Open Visual Studio Code (VSCode).
2. Open the integrated terminal in VSCode.
   - Shortcut: `Ctrl + `` (backtick)` or navigate to `View -> Terminal`.
3. Clone the repository:
   ```bash
   git clone https://github.com/Emma0123/animals-api.git
4. Open the `main.go` file and locate the `dsn` variable.

   ```go
   dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
   (Replace "user," "pass," and "dbname" with your MySQL Workbench credentials.)
5. Return to the terminal in VS Code and run the server:
   `go run main.go`
