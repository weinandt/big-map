# big-map
Dictionary for strings backed by anonymous pages using mmap.


## To Build/Deploy the Lambda
NOTE: the lambda handler must be configured to 'main'.

From the 'lambda' directory:
1. `GOOS=linux GOARCH=amd64 go build -o main main.go`
1. `zip main.zip main`
1. "Upload from" -> "zip" in the AWS console.