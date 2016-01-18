#go-shopping

Just playing around here - nothing to see... :-)

##Running

###Product
```
$ cd product
$ go run main.go
```

###Store
```
$ cd store
$ export SERVICE_PORT=3000
$ export PRODUCT_SERVICE_ADDR=http://localhost:3010
$ go run main.go 
```
