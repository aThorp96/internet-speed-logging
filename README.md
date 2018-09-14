#Internet Speed Logger

Internet speed logger takes an your internet speed and logs the speed into a google sheet.
For best results run with Cron or the Systemd scheduler. 

Usage: `go build speedLog.go jsonUtils.go requests.go && ./speedLog`

* TODO: Refactor to eliminate need for external script
