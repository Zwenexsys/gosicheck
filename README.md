GoSiCheck - Go Site Checker
===

This is inspired by the code from [https://github.com/yelinaung/site_checker](https://github.com/yelinaung/site_checker).

Main Functionality
---
- Check a status of the given site by sending HEAD requests in pre-defined interval.
- Display the status of last 10 request on web page which request to server with AJAX in pre-defined interval.

Configurables
---
- The `interval` to send status check `HEAD` requests.
- The `interval` to refresh status web page in Javascript code.
- The `port` of the status heck web server to run.

Dependencies
---
Only standard libraries are used. Even AJAX code. No 3rd Party dependency.

How to run
---
Put two files in the same folder.
Just type 
`go run status_server.go`

TODO
---
- Add configuration file.
- Multiple sites checking at multiple interval.
- Make the web page more useful.
- Send email upon getting failed status check requests.
- Stop calling for `pre-defined interval` of status check request upon receiving certian number of failed requests.