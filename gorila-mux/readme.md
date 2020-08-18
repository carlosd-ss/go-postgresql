# Crud Gorila-Mux and Postgres

## Dependencies

- [x] Go
- [x] Gorila Mux
- [x] Postgres



### Running the project


```bash
RUN main.go
```


### API

* Endpoints

| Version | HTTP | Path |
|:--:|:--:|:--|
| /api | GET | /user | 
| /api | GET | /user/1 | 
| /api | POST | /user |
| /api | PUT | /user/1 | 
| /api | DELETE | /user/1 |


### Test the project


```bash
# GET user
curl -v http://localhost:8080/api/user/{id}

# GET ALL Users
curl -v http://localhost:8080/api/user

# POST Create Users
curl -d '{"id":3,"name":"carlos","location":"Brasil","age":20}' http://localhost:8080/api/user

# PUT User
curl -d '{"id":3,"name":"carlos","location":"Brasil","age":20}' -X PUT http://localhost:8080/api/user/{id}

# DELETE User
curl -X DELETE http://localhost:8080/api/user/{id}

```


### PostgreSQL Table

    CREATE TABLE users (
      userid SERIAL PRIMARY KEY,
      name TEXT,
      age INT,
      location TEXT
    );


### See the project on browser

Access the browser on [http://localhost:8080](http://localhost:8080)

