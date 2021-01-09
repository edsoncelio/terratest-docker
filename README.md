# Labs using Terratest with Docker
>Work in progress 

## Requirements
* Go +1.15.2
* Terraform +1.13.4
* Docker +19.03.13

## How to use
Inside the `/test` directory, run:   
`go test -v -run <TEST_FUNCTION>`

where `TEST_FUNCTION` can be (see `dockerfile_test`):
* `TestImageAlpine`
* `TestImageNginx`

or run `go test` to run ALL the tests.

