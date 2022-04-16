FROM golang:1.18

WORKDIR /usr/src/app

COPY . .
RUN make build
RUN mkdir /usr/local/bin/app/

RUN cp /usr/src/app/dist/.env /usr/local/bin/app/
RUN cp /usr/src/app/dist/category-rest-service-runner /usr/local/bin/app/
RUN chmod 777 /usr/local/bin/app/.env

WORKDIR /usr/local/bin/app/
EXPOSE 8080

CMD ["./category-rest-service-runner"]