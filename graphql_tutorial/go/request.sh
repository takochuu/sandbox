#!/bin/bash
curl -X POST -H "Content-Type: application/json" -d 'query { hello(id:4) }' http://localhost:8080/
