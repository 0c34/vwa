#### VWA (Vulnerable Web Application) with GO 
VWA is a web application developed to help the pentester and programmers to learn the vulnerabilities that often occur in web applications which is developed using golang.

#### How To Install VWA

#### Installing golang
If you didn't have golang installed on your system. first, install it using automation script from https://github.com/canha/golang-tools-install-script.

Follow the instruction which is provided by the author and install golang depending on your Operating System Architecture.

If successfully installed you would have directory 'go' in your home directory. the go directory has three subdirectory (bin, pgk, src). switch to src directory then clone vwa repository. 

#### Installing Postgres & Restore Database
Install Postgres
```
sudo apt-get update
sudo apt-get install postgresql postgresql-contrib
```
Restore DB
```
passwd postgres //Change Password from user postgres
su postgres //Login to user postgres
createdb [database_name] //create database
psql [database_name] < db/vwa.sql
psql -d [database_name] //connect to that database, for check
\dt //To see table
```

#### Installing VWA
```
git clone https://github.com/0c34/vwa.git

```
#### Installing Golang Package
```
go get github.com/lib/pq
go get github.com/gorilla/sessions
go get github.com/julienschmidt/httprouter
```

#### VWA config
Open the file config.json which is located in config directory. Change the configuration according to your needs.

```
{
    "user": "root",
    "password": "toor",
    "dbname": "vwa",
    "sqlhost": "localhost",
    "sqlport": "5432",
    "webserver": "http://localhost",
    "webport": "8082"
}
```

#### Run VWA
```
vwa@ubuntu-server:~/go/src/vwa$ go run app.go
Server running at port :8082
Open this url http://localhost:8082/ on your browser to access VWA
```

#### VWA users

|		email		| password	|
|-------------------|-----------|
| eko@gmail.com		| testing	|
| andi@gmail.com	| testing	|
| attacker@gmail.com| testing	|

Explore the vulnerability.

#### To Do

* Reflected XSS & Stored XSS
* IDOR (Insecure Direct Object Reference)
* SQL Injection
* CSRF & Missing CORS Origin
