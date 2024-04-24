# Stage 1
FROM golang:1.22 as build

WORKDIR /src

COPY go.mod ./
RUN go mod download && go mod verify

COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /radioanty .

# Stage 2
FROM scratch

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /radioanty .

EXPOSE 8088

CMD ["./radioanty"]