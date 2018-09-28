# iron/go is the alpine image with only ca-certificates added
FROM iron/go

# Change the timezone to TX (Central Time)
RUN apk add --no-cache tzdata
ENV TZ America/Chicago

WORKDIR /app

# copy env file
COPY .env /app/
COPY data.xlsx /app/

# Now just add the binary
ADD main /app/
ENTRYPOINT ["./main"]