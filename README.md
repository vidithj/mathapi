# mathapi
simple REST API for math functions.This API calculates min,max,avg,median and percentile of a list of numbers. This REST API was created using gokit . 

# API Documentation
Swagger file included with API route information.

# Testing
All the api logic is present in api.go .Test cases around each function has been added in api_test.go

mathapi/api	3.767s	coverage: 97.3% of statements

## Base Path
mathapi/v1/percentile
## Sample Request
localhost:8080/mathapi/v1/percentile
#### POST /min
```
{
	"list":[9, 10,2,12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25],
    "quantifier":3
}
```

#### POST /max
```
{
	"list":[9, 10,2,12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25],
    "quantifier":3
}
```

#### POST /median
```
{
	"list":[9, 10,2,12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25]
   
}
```

#### POST /avg
```
{
	"list":[9, 10,2,12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25]
    
}
```

#### POST /percentile
```
{
    "list":[9, 10,2,12, 13, 13, 13, 15, 15, 16, 16, 18, 22, 23, 24, 24, 25],
    "quantifier":100
}
```

