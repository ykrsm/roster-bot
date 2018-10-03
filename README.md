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


# Deployment
```
git clone

scp .env ./roster-bot
scp ./data.xlsx ./roster-bot

crontab -e 
59 9 * * * git pull && env GOOS=linux GOARCH=386 go build -o main && docker-compose up --build
```