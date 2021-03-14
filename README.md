# web_log_analytic

# get the dependency
- go mod init github.com/tmluthfiana/web_log_analytic
- go get -u github.com/tmluthfiana/web_log_analytic or checkout https://github.com/tmluthfiana/web_log_analytic.git

# usage
- run it on local go run main.go
- use this command line input : go run main.go -t <mins>m -d <dir>
- example : go run main.go -t 3m -d /Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log

# testing
- adjust dirname or filename for test based on your file path
- running file test in api folder with : go test -v -run (function name). example : go test -v -run TestProcessDir


		