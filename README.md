# :heart: NyaaPantsu Manga :heart:
NyaaPantsu is intended as a possible replacement for Bato.to written in the Go programming language. It may be best described as a manga content management system, permitting third-parties to upload and manage manga translation, to be made available to the world (or not, thats on them).

## Installing
NyaaPantu Manga requires the following software:

  * PostgreSQL
  * Go Programming Language

### Setting up the database
The installation of the software is not covered here, please consult the appropriate operating system or software document for software installation.

By default, NyaaPantsu attempts to connect to PostgreSQL utilizing the configuration stored in /models/default.go, located within the root project folder. It is __highly__ recommended to alter these values if deploying a private server (checking out the code is covered below). By default the connection configuration is as follows:

    Port: 5432
    Username: manga
    Password: manga
    Database: manga

Presuming that Postgresql is properly installed and available on the system path, this may be done utilizing the following instructions:

    pg_ctl -o "-p 5432" start
    createuser -p 5432 -U postgres -D -P -W manga
    createdb -p 5432 -U postgres -O manga manga

Please note that the previous instructions assume the default configuration of Postgresql, if you have altered the superuser username and password, please use the appropriate information in the previous instructions (`-U superUsername -p superPassword`). Should some other problem arise, please check the FAQ below or consult the appropriate PostgreSQL documentation.

### Getting the source
 Execute `go get https://github.com/NyaaPantsu/manga` to retrieve the code from the github repo. By default the project will be checked out into GOPATH/src, it is possible to determine the exact value of GOPATH by executing `go env GOPATH`. Once having determined the appropriate value for GOPATH, execute the following instructions replacing GOPATH for the value previously determined:

    psql -p 5432 -d manga -U manga -f GOPATH/schema.sql
    go install GOPATH/src/NyaaPantsu/manga

It is assumed that the PostgreSQL server was set-up with the configuration given above, if this is not the case, please update the `psql` command above to reflect the appropriate configuration.

Once completed the above steps, the NyaaPantsu Manga server executable may be found in GOPATH/bin as projectName.exe where projectName is the name of the root source folder. By default this is manga, so the executable will by default be named manga.exe. To check that all is running as appropriate simply open a browser and visit `localhost:Port` where Port is the value for httpport stored in /conf in the project root folder. By default this value is 8080.

## FAQ

__I get an error stating `pg_ctl: no database directory specified and environment variable PGDATA unset`.__

  Try setting the PGDATA system variable using `setx PGDATA <DIR>` on Windows and `export PGDATA=<DIR>` where DIR is the absolute path to some directory where the server configuration and state will be stored.

__I get an error stating `pg_ctl: another server might be running; trying to start server anyway`.__

  If on Windows, try running `pg_ctl stop`. If that fails check the PGDATA variable using `echo %PGDATA%` on Windows or `echo $PGDATA` on Linux and then run `pg_ctl -D <PGDATA_DIR> stop` where <PGDATA_DIR> is the value returned from `echo`. Assuming no error occurs, attempt to execute the instruction that reported the error again.

## Contributing

### Project Structure

## Contact

__Issue Tracker:__ https://github.com/NyaaPantsu/manga/projects/1

__Discord:__ https://discord.gg/QvQPQS

__Irc:__

    Server: irc.rizon.net
    Channel: #manga-dev