#!bin/bash
curl --location --request DELETE 'http://localhost:8080/api/v1/student/1'

curl --location --request PUT 'http://localhost:8080/api/v1/student' \
--header 'Content-Type: application/json' \
--data '{
    "Id":2,
    "name":"xiaoming",
    "sex":"男",
    "age":20
}'

curl --location 'http://localhost:8080/api/v1/student' \
--header 'Content-Type: application/json' \
--data '{
    "name":"laowang",
    "sex":"男",
    "age":40
}'

curl --location 'http://localhost:8080/api/v1/student?sex=男&age=60&name=lao'