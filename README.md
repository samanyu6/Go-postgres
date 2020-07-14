# Go - Postgres setup

A boilerplate web server built with Go and integrate Postgres as a database. 

# Initial Setup

## Postgres
- Install postgres on your machine. (The following setup works on MacOS).
- Set up the database and table in postgres using the command shown below
- ```bash
    chmod +x Db.sh
    ./Db.sh 
    ```
- The following output is successful if you're able to see a "users" table present after execution.

## Environment variables
- Create a file named ".env" in the main folder and add the following variables stated below.
-   DB_USER = {Enter postgres user}
    DB_PASS = {Enter the password}
    DB_NAME = {Name of the database}

## Running the server
- ```Go
    go run main.go
  ```

