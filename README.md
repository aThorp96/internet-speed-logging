# Internet Speed Logger

Internet speed logger takes an your internet speed and logs the speed into a google sheet.
For best results run with Cron or the Systemd scheduler. 

Dependancies: 
- [speedtest](https://github.com/sivel/speedtest-cli)

Setup: 
1. Follow [Google Sheets' API setup step 1](https://developers.google.com/sheets/api/quickstart/go) and save `credentials.json` to `./resources/sheetID`. 
2. From the source diriectory, run `echo "<spreadsheet ID>" > ./resources/credentials.json`.
3. Run `go build speedLog.go jsonUtils.go requests.go && ./speedLog`
4. The first time you run it will print a url on standard out. Enter the url into a browser, login and create the oauth2 token. Copy the token ID and paste it into the terminal running the program. 
5. Run `go install` to install the binary.
6. Run program using `internet-speed-logging`

Scheduling
- Copy `./resources/*` to `/etc/speedlog/`
- Copy `speedlog.timer` and `speedlog.service` to `/etc/systemd/system/`
- Change variable `resourcePath` in `requests.go` to `/etc/speedlog/`
- To change frequency of timer change `speedlog.timer` variable `OnCalendar` to desired frequency **See `man systemd.timer` for details**
