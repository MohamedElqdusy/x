# Myhttp

## System Requirements
- Golang 1.16

## Usage
#### 1- Build
- `go build`  
####  2- Run
- `./myhttp -parallel=3 adjust.com google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com`
#### 3- Test
- ` go test ./...`
#### 4- Concurrency test
- To test the routines' saftey from data races

    `go run -race main.go google.com facebook.com yahoo.com yandex.com twitter.com`
