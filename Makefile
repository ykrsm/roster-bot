FILE=prod_ip
IP=`cat $(FILE)`

run:
	./main -p /mnt/ns/Schedule/current/[^~].xlxs

dev:
	env GOOS=linux GOARCH=386 go build -o main && docker-compose up --build

ssh:
	ssh slack@$(IP)
	
deploy:
	git checkout master
	git pull
	git status
	env GOOS=linux GOARCH=386 go build -o main
	scp ./main slack@$(IP):/home/slack/roster-bot  
	scp ./names_prod.txt slack@$(IP):/home/slack/roster-bot  
	echo Deployed source to IP: $(IP)

test:
	go clean -testcache
	go test -v ./...
