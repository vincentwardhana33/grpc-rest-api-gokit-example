# GRPC-REST-API-GOKIT-EXAMPLE

**by: Vincent Wardhana**

This example is written in GoLang. Install Go LTS Version from [here](https://golang.org/dl/).

```gRPC``` and ```Go Kit``` example can be found in folder  ```service```

### Setup

1. copy ```api_key_example.go``` in ```config``` directory to ```api_key.go```
2. uncomment ```const``` variables
3. set ```OMDB API KEY``` in ```OmdbAPIKey``` variable


### Unit testing

$ ```go test ./... -v```


### Run application

$ ```go run ./cmd```


### Command

$ ```curl -XGET "localhost:8080/search?searchword=Spiderman&pagination=1"```
