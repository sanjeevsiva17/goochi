# bookish-meme

For running tests, following docker containers are required:

```
docker run --name goochi-postgres -e POSTGRES_PASSWORD=pass123 -p 2005:5432 -d postgres:latest
docker run --name goochi-mysql -e MYSQL_ROOT_PASSWORD=password -p 2001:3306 -d mysql:latest
```