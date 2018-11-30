# roster-bot

# Compile for linux
```
dep ensure
env GOOS=linux GOARCH=386 go build -o main
```
# Start docker
```
docker-compose build
docker-compose up
```
# Feeling lazy? (one liner)
```
env GOOS=linux GOARCH=386 go build -o main && docker-compose up --build
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
59 23 * * * cd /home/slack/roster-bot && ./main
```

# Deployment
```
scp ./.env slack@<PROD_SERVER_IP>:/home/slack/roster-bot  
scp ./data.xlsx slack@<PROD_SERVER_IP>:/home/slack/roster-bot  
scp ./main slack@<PROD_SERVER_IP>:/home/slack/roster-bot  
```
