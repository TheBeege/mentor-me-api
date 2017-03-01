# Mentor Me API
Mentor-Me is a simple mentor matchmaking app. Users can sign up as a mentor and/or mentee. Mentors can list a skill they'd like to mentor in along with their level of expertise in said skill. Mentees can search for mentors by skill and level. A mentee can submit a request for mentorship upon finding a desired mentor. The mentor can then accept or reject the request. Upon acceptance, the mentor and mentee will receive each other's contact information.

This is the API for the Mentor Me project, this time using Beego. You can find the [front-end here](https://github.com/TheBeege/mentor-me).

## Setup
1. [Install Go](https://golang.org/doc/install)
    * You may need to add $GOPATH/bin to your $PATH
2. [Install Beego](https://beego.me/docs/install/)
3. [Install dep](https://github.com/golang/dep) - `go get -u github.com/golang/dep/...`
4. Run `dep ensure` to resolve dependencies
5. [Download](https://www.postgresql.org/download/) and [setup](https://wiki.postgresql.org/wiki/Detailed_installation_guides) PostgreSQL
6. Create the database user with `psql --username=postgres --password -c "CREATE USER <name> PASSWORD <password>"`
7. Create the database with `psql  --username=postgres --password -c "CREATE DATABASE <name> OWNER <user_name>"`
8. Setup the database schema with `psql --user=<user> --password -d <database_name> -f schema.sql`
9. Set environment variables for `dbhost`, `dbuser`, `dbpass`, and `dbdatabase`.
10. Run `bee run -downdoc=true -gendoc=true` to generate documentation and run the server