# Requirements
- go get github.com/go-pg/pg
- go get github.com/gin-gonic/gin
# Install
- go get github.com/juankis/x-men
# Run
- In x-men folder run in the console: go run src/main.go
# Run test
- In folder x-men/src/analizer/ run in the console: go test -cover
# Demo Api
- Consult if it is mutant adn: POST http://18.219.147.241:8080/mutant
- Body example Post request:
```json
{ "dna":["ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"] } 
```
- Consult reports adn: GET http://18.219.147.241:8080/stats



