FROM golang AS build

COPY --from=app go.mod ./
COPY --from=app go.sum ./

COPY --from=app ./src/proto/ ./proto/

COPY *.go ./

RUN go mod download

RUN CGO_ENABLED=0 GARCH=amd64 GOOS=linux go build -o user .

FROM alpine

COPY --from=build /user /user

EXPOSE 0000

CMD ["./user"]



RUN go mod tidy