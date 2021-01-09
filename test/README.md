Tests

## Running the tests

### Prerequisites
 - Go +1.15.2
 - Terraform +1.13.4
 - Docker +19.03.13

### One time setup
Download Go dependencies:
```
cd test
go mod download
```

### Run all the tests
```
cd test
go test -v 
```

### Run a specific test
```
cd test
go test -v -run <TEST_NAME>
```

Available tests:
 - `TestImageAlpine`:
    - Test the alpine image timezone (UTC)
 - `TestImageNginx`:
    - Test the content from `/usr/share/nginx/html`
