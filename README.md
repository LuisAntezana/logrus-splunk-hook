[![Build Status](https://travis-ci.org/LuisAntezana/logrus-splunk-hook.svg?branch=master)](https://travis-ci.org/LuisAntezana/logrus-splunk-hook)

# Splunk Hook for Logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>
Splunk hook for logrus 
## Install

```
go get github.com/LuisAntezana/logrus-splunk-hook
```

## New Features
- **Async logs:** Now logs can be async with a boolean parameter in the hook.
- **Retries:** If it fails to send a log to splunk, it can retry sending the log with the new "retries" parameter.

```go
splunk "github.com/LuisAntezana/logrus-splunk-hook"
...
splunkHook := splunk.NewHook(
		splunkClient, // your splunk client
		logrus.AllLevels,
		true, // enable async logs
		3, // add retries (it can retry up to 3 more times, if it fails)
	)
```

## Fixes
- Logs to splunk were not sent in json format. ([jakobkrein's fix](https://github.com/Franco-Poveda/logrus-splunk-hook/compare/master...jakobkrein:logrus-splunk-hook:master))
- Logs to splunk did not include milliseconds. (this fixes the problem of logs in different order)
