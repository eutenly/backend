# Eutenly Backend Services
Written in Golang and Echo (for now)

### How to run this:
1) Make sure you have the Golang toolchain installed.
2) In the directory with the Golang files, run `go get -d ./...` to get all the dependencies.
3) Make sure to sort out your .env (see below)
3) Then run `go build` and Golang will shit out a binary for you to run. 

### Environment template
To run the app you'll need a .env file in the same directory which the binary runs in.

`PORT=`

`WEBSERVER_URL=`

`DISCORD_CLIENT_ID=`

`DISCORD_CLIENT_SECRET=`

`SESSION_SECRET=`

`GITHUB_CLIENT_ID=`
`GITHUB_CLIENT_SECRET=`

`TWITTER_KEY=`
`TWITTER_SECRET=`

`TOPGG_WEBHOOK_SECRET=`

`MONGO_URI=`
`MONGO_DATABASE=`

`INFLUXDB_URL=http://server.eutenly.com:8086`
`INFLUXDB_TOKEN=root:PASSWORD`
`INFLUXDB_BUCKET=dbname/autogen`

---

`PORT` is which port the app should run on, for example, 8080.

`WEBSERVER_URL` is the URL of the web server root, for example, http://localhost:8080

The `DISCORD` stuff is self-explanatory. 

`SESSION_SECRET` is what's used to encrypt session storage. Set it to whatever you want. In production this should be a long randomly generated password.

`TOPGG_WEBHOOK_SECRET` is the webhook secret for Top.GG