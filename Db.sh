
createdb user
psql user
CREATE TABLE "users"(Id SERIAL PRIMARY KEY, Name TEXT NOT NULL, Age INT NOT NULL);
\dt