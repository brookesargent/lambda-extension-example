`docker build -t myfunction:latest .`

`docker run -p 9000:8080  myfunction:latest `

`curl -XPOST "http://localhost:9000/2015-03-31/functions/function/invocations" -d '{}'.`