#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d '{ hello(id:4) }' http://localhost:8080

curl -X POST -H "Content-Type: application/json" -d 'mutation { create(id:4){ id }}' http://localhost:8080

curl -X POST -H "Content-Type: application/json" -d 'mutation { create(id:4){ name }}' http://localhost:8080
curl -X POST -H "Content-Type: application/json" -d 'mutation { create(id:4){ id, name }}' http://localhost:8080
