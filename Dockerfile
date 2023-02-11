FROM node:18 as webbuilder

WORKDIR /app
COPY web/ ./
RUN npm install
RUN npm run build-only

FROM golang:1.19-alpine as builder
WORKDIR /app
COPY . .
COPY --from=webbuilder /app/dist/ ./web/dist/
RUN go mod tidy
RUN go build -o ssp cmd/ssp.go

FROM alpine
RUN apk add --no-cache tzdata
COPY --from=builder /app/ssp /usr/local/bin/ssp
EXPOSE 3000

VOLUME ["/app/data.db"]
ENTRYPOINT ["ssp" , "-db=/app/data.db", "-listen=:3000"]
