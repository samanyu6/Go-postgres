# Creates new database user
createdb user

# Creates Table name "users" for the database user
psql -d "user" -c 'CREATE TABLE "users" (Id SERIAL PRIMARY KEY, Name TEXT NOT NULL, Age INT NOT NULL);'

# Prints out all the relations in the database. Make sure the table name "users" is present in the table printed
psql -d "user" -c '\dt'
