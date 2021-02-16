# Go - Postgres setup

Server built with Go and Postgres following a somewhat MVC Structure with Unit Tests included. Use this either for boilerplate/ reference or if you're planning to learn Go and want a simple project to familiarise yourself with the language.

# Project Directory
  1. ### database/ 
     Consists of the model and database connector files.
  
  2. ### middleware/
     Consists of all database operations that can be called.

  3. ### routers/
     Consists of all the routes available for API endpoints.

# Initial Setup

## Postgres
- Install postgres on your machine. (The following setup works on MacOS).
- Set up the database and table in postgres using the command shown below
    ```bash
    chmod +x Db.sh
    ./Db.sh 
    ```
- You should see the following output on your terminal 
   ```      
            List of relations

    Schema | Name  | Type  |  Owner  
    --------+-------+-------+--------- 

    public | users | table | << owner >>
    ```

## Environment variables
- Create a file named ".env" in the main folder and add the following variables stated below.
-   DB_USER = {Enter postgres user name}
    DB_PASS = {Enter the password for the above user}
    DB_NAME = user

## Running tests 

 ```bash
    chmod +x run_tests.sh
    ./run_tests.sh
```
- code coverage upto 80%.

## Running the server
- ```Go
    go run main.go
  ```
- runs on localhost:8000/

## Api Endpoints

- ### /api/user (POST):
     Creates a new user and enters data in the PSQL database. Post a json of type {name: "", Age: 0}.
- ### /api/all (GET): 
    Gets all values from the db.
- ### /api/usr/{id} (GET) : 
Gets user with specific id associated. (ID's in PSQL are serial (autoincremented), so each user is given a specific ID.)
- ### /api/update/{id} (POST): 
    Updates user with specific id. Post a json with updated values of type {name:"", age:0}.
- ### /api/delete/{id} (DELETE): 
    Deletes user with specific Id.
