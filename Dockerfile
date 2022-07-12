FROM golang:1.18.3 AS builder
WORKDIR /home/sbom-service
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o SBOMservice .

FROM ubuntu:22.10
WORKDIR /home/sbom-service
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
RUN apt-get update && apt-get install -y apt-transport-https ca-certificates
#RUN apk update && apk add ca-certificates
COPY --from=builder /home/sbom-service/SBOMservice /home/sbom-service/SBOMservice
CMD ["./SBOMservice"]