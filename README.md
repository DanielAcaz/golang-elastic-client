# Go Lang with ElasticSearch

This project has a simple program in Go Lang to create a connection with ElasticSearch and make a request with a simple filter.

The project use the elastic client in go go-elasticsearch. The client can be found in repository below:

<https://github.com/elastic/go-elasticsearch>

The mangement of project was doing tought Go Modules introduced in 1.13 Go version.

## Building Project

The initial thing to build this project is run elasticseatch from docker-compose. To run elasticsearch, apply this command in project root by terminal:

```bash
    $ docker-compose up -d
```

Now, also you need realize a POST http with the first document will be created in elasticsearch, with index *users_index*.  You can use *POSTMAN* to this request, but you can use the comnand in terminal bellow:

```bash
curl -X POST --header "Content-Type: application/json" --header "Accept: application/json" -d  {"id":100,"firstname":"Sharlene","lastname":"Moses","age":34,"isActive":false,"balance":16233,"email":"sharlenemoses@tropolis.com","phone":"+1 (952) 481-3178"}
```

Now, you can build the project and run with commands in project root:

```bash
$ go build
$ ./elastic-client
```

Your response should be as bellow:

```bash
2020/07/09 18:32:28 {
"id": 100,
"email": "sharlenemoses@tropolis.com",
"firstname": "Sharlene",
"lastname": "Moses",
"age": 34,
"isActive": false,
"balance": 16233,
"phone": "+1 (952) 481-3178"
}
```

---
***NOTE***

You need check the import of models in *main.go*!

---