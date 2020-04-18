 #### Steps to run:
 1. Clone the repo
 2. Run `go mod download` to download all dependencies
 3. Fill the `.env` file. The required variables are:
    * `DB_USER`
    * `DB_PASSWORD`
    * `DB_HOST`
    * `DB_PORT`
    * `DB_NAME`
    * `JWT_SECRET`
 4. Database should be postgresql
 5. Execute `go run ./server.go` from the root directory
 ------------------------------------------------------------------------

## GraphQL in Golang

The was done following the tutorial by [EQuimper](https://github.com/EQuimper). <br/>
The complete tutorial can be found [here](https://www.youtube.com/playlist?list=PLzQWIQOqeUSNwXcneWYJHUREAIucJ5UZn). <br/>
The original repo can be found [here](https://github.com/EQuimper/youtube-golang-graphql-tutorial).
------------------------------------
Backend for a meetup app using GraphQL in Go. <br/>
The following operations have been covered:
 * The user can register
 * The user can log in
 * The logged in user can create meetups
 * The authorized user (owner of meetup) can update/delete meetups
 * Meetups can be fetched with filter, limit and offset
 * User can be fetched for a given userID
 --------------------------------------------------
 * The graphQL server was build using [gqlgen](https://gqlgen.com/)
 * [gorm](https://gorm.io/) used for DB operations
 * [gorilla/mux](https://github.com/gorilla/mux) used for routing
 -------------------------------------------------------