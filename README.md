# Noti

Simply send notification to IM. Slack, Teams, Wechat Work, etc...

- [x] Wechat Work
- [ ] Slack
- [ ] Teams

## Getting start
Use Noti by Environment Variables
```shell script
NOTI_PROVIDER=wework
WEWORK_INFO_KEY=
WEWORK_WARN_KEY=
WEWORK_ERROR_KEY=
```
The KEY can find in Wechat Work group robot url.

Then use noti just like this:
```go
noti.Info("info")
noti.Warn("warn")
noti.Error("error")
```
