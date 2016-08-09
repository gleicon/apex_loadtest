## A braindump on Apex golang

Vanilla HTTP load test using lambda functions.

### build/run/etc

Install and Configure [apex](http://apex.run/)

	- Create a project:
	$ mkdir lambda_loadtest
	$ cd lambda_loadtest
	$ apex init
	- Answer the questions, etc
	- 
Install golang support and vegeta

	$ go get -u github.com/apex/go-apex
  	$ go get -u github.com/tsenart/vegeta/lib
  
Copy this repo's functions/loadtest directory to lambda_loadtest/functions

	$ apex deploy
	$ echo '{"event":{"value":"http://www.example"}}' | apex invoke loadtest | python -m json.tool
	$ apex logs
	$ apex metrics


Be careful not to flood the interweb, try to do multiple parallel calls with for/xargs etc. Tune timeout and memory to suit your test.
