# Requirements
- go get github.com/go-pg/pg
- go get github.com/gin-gonic/gin
# Install
- go get github.com/juankis/x-men
# Run
- In x-men folder run in the console: go run src/main.go
# Run test
- In folder x-men/src/analizer/ run in the console: go test -cover
# Demo Api (En mantenimiento ...)
- Consult if it is mutant adn: POST http://ec2-18-219-147-241.us-east-2.compute.amazonaws.com:8080/mutant
- Body example Post request:
```json
{ "dna":["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"] } 
```
- Consult reports adn: GET http://ec2-18-219-147-241.us-east-2.compute.amazonaws.com:8080/stats
# Details
- The project is deployed on an amazon server with load balancer.
- In the repository there are 2 functional branches, one with a synchronous code (origin), another with an asynchronous code with go routines (async)
