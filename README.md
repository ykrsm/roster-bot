[![CircleCI](https://circleci.com/gh/ykrsm/roster-bot.svg?style=svg)](https://circleci.com/gh/ykrsm/roster-bot)

# roster-bot

# Compile for linux
```
dep ensure
env GOOS=linux GOARCH=386 go build -o main
```
# Test run
```
make dev
```
# Run on the production server
```
make run
```

# Set up for production server

## Set timezone to CST

```
as slack

$ sudo unlink /etc/localtime 
$ sudo ln -s /usr/share/zoneinfo/Etc/GMT+6 /etc/localtime
# add below to .bash_profile
export TZ="/usr/share/zoneinfo/Etc/GMT+6"
$ date
# should be like  GMT+6 
```

## Set crontab
```
crontab -e 
59 23 * * * cd /home/slack/roster-bot && make run
```

# Deployment
```
make deploy
```
