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
6. Set environment variables for `dbhost`, `dbuser`, `dbpass`, and `dbdatabase`. If you're on Mac or Linux, you can run `source setup_dev_db_env_vars.sh` after changing the values within
7. Create the database user 
    * Mac/Linux: `psql --username=postgres --password -c "CREATE USER ${dbuser} PASSWORD '${dbpass}'"`
    * Windows: `psql --username=postgres --password -c "CREATE USER %dbuser% PASSWORD '%dbpass%'"`
8. Create the database
    * Mac/Linux: `psql --username=postgres --password -c "CREATE DATABASE ${dbdatabase} OWNER ${dbuser}"`
    * Windows: `psql --username=postgres --password -c "CREATE DATABASE %dbdatabase% OWNER %dbuser%"`
9. Setup the database schema 
    * Mac/Linux: `psql --username=${dbuser} --password -d ${dbdatabase} -f schema.sql`
    * Windows: `psql --username=%dbuser% --password -d %dbdatabase% -f schema.sql`
10. Fix up permissions
    * Mac/Linux: `psql --username=postgres --password -c "REASSIGN OWNED BY postgres TO ${dbuser}" ${dbdatabase}`
    * Windows:`psql --username=postgres --password -c "REASSIGN OWNED BY postgres TO %dbuser%" %dbdatabase%`
11. Run `bee run -downdoc=true -gendoc=true -vendor=true` to generate documentation and run the server

## Database
### How to Export Database Changes
Mac/Linux: `pg_dump -U ${dbuser} -W -s -O --clean ${dbdatabase} > schema.sql`
Windows: `pg_dump -U %dbuser% -W -s -O --clean %dbdatabase% > schema.sql`

### How to Import Database Schema
Mac/Linux: `psql -U ${dbuser} -W < schema.sql`
Windows: `psql -U %dbuser% -W < schema.sql`
