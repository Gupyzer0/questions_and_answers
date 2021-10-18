FROM golang:1.17

#installing postgres
RUN apt-get update
RUN apt-get install -y postgresql-13 postgresql-client-13

#api rest http port
EXPOSE 4000

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

#setting up postgresql

USER postgres
RUN    /etc/init.d/postgresql start &&\
    psql --command "CREATE USER docker WITH SUPERUSER PASSWORD '123456';" &&\
    createdb -O docker questions_and_answers

RUN echo "host all  all    0.0.0.0/0  md5" >> /etc/postgresql/13/main/pg_hba.conf
RUN echo "listen_addresses='*'" >> /etc/postgresql/13/main/postgresql.conf

EXPOSE 5432

#postgres db persistence
VOLUME  ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]
#starting db server
CMD ["/usr/lib/postgresql/13/bin/postgres", "-D", "/var/lib/postgresql/13/main", "-c", "config_file=/etc/postgresql/13/main/postgresql.conf"]

#RUN go run cmd/main.go -migrate -seed
