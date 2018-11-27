FILE=prod_ip
IP=`cat $(FILE)`

run:
	./main -p /mnt/ns/Schedule/current/[^~].xlxs

dev:
	env GOOS=linux GOARCH=386 go build -o main && docker-compose up --build

deploy:
	scp ./main slack@$(IP):/home/slack/roster-bot  
	echo Deployed source to IP: $(IP)
