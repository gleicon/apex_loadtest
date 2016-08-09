## A braindump on Apex golang

Vanilla HTTP load test using lambda functions.

### build/run/etc

Install and Configure [apex](http://apex.run/)

	- Create a project:
	$ mkdir lambda_loadtest
	$ cd lambda_loadtest
	$ apex init
	- Answer the questions, etc

Copy this repo's functions/loadtest directory to lambda_loadtest/functions

	$ apex deploy
	$ echo '{"event":{"value":"http://www.example"}}' | apex invoke loadtest | python -m json.tool
	$ apex logs
	$ apex metrics


Be careful, try to parallelize it with xargs etc.