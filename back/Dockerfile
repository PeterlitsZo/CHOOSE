FROM golang

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app
COPY ./ /usr/src/app/

RUN go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN go build -v -o /usr/local/bin/app .

ENTRYPOINT [ "sh", "-c", "bash /usr/src/app/migrate.sh up && /usr/local/bin/app" ]