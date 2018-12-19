# iron/go is the alpine image with only ca-certificates added
FROM iron/go

# Change the timezone to TX (Central Standard Time)
RUN apk add --no-cache tzdata
RUN unlink /etc/localtime
RUN ln -s /usr/share/zoneinfo/Etc/GMT+6 /etc/localtime

WORKDIR /app

# copy env file
COPY .env /app/
COPY data.xlsx /app/
COPY data3.xlsx /app/
COPY names.txt /app/

# Now just add the binary
ADD main /app/
ENTRYPOINT ./main -d2 ./data3.xlsx ./names.txt