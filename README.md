# web_log_analytic

# get the dependency
- go mod init github.com/tmluthfiana/web_log_analytic
- go get -u github.com/tmluthfiana/web_log_analytic or checkout https://github.com/tmluthfiana/web_log_analytic.git

# usage
- run it on local go run main.go
- use this command line input : analytics -t <mins>m -d <dir>
- example : analytics -t 3m -d /Users/triasluthfiana/go/src/github.com/tmluthfiana/web_log_analytic/http-log

# testing
- running file test in test folder with : go test -v -run (function name). example : go test -v -run TestProcessDir


		