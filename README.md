# go-auth-svc

#### Pull latest postgres image
```docker pull postgres```
#### Run Postgres DB 

```docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres```

#### Create Database 

If the auth_svc database does not exist the sh into postgres and create it 

```docker exec -it postgres sh```

The connect to the psql server

```psql --u postgres```

Run this command to create the database 

```CREATE DATABASE auth_svc;```

To view the list of databases run this command 
```\l```