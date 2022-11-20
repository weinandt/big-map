# big-map
Dictionary for strings backed by anonymous pages using mmap.


## To Build/Deploy the Lambda
NOTE: the lambda handler must be configured to 'main'.

From the 'lambda' directory:
1. `GOOS=linux GOARCH=amd64 go build -o main main.go`
1. `zip main.zip main`
1. "Upload from" -> "zip" in the AWS console.


## TODO:
- Make sure when strings are inserted/read, they are used as runes (utf-8 encoding) and not assumed that
every character is a single byte
- Checks around the string fitting in the allocated size before it is inserted.
- When converting string to byte, are their two copies of the string in real memory at once?
     - Should the string be converted character by character?