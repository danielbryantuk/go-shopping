#go-shopping

Just playing around here - nothing to see... :-)

##Running

###Product
```
cd product
export PRODUCT_SERVICE_PORT=3010
go run main.go
```

###Store
```
cd store
export STORE_SERVICE_PORT=3000
export PRODUCT_SERVICE_ADDR=http://localhost:3010
go run main.go 
```

##TODO
 * Set envs properly for local execution
 * Set envs in Dockerfile
