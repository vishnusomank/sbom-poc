FROM golang:1.18.3 AS builder
WORKDIR /home/sbom-service
COPY . .
RUN export GIN_MODE=release
RUN go mod tidy
RUN go build -o SBOMservice .

FROM ubuntu:22.10
WORKDIR /home/sbom-service
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
RUN export GIN_MODE=release
ENV ZONEINFO /zoneinfo.zip
RUN apt-get update && apt-get install -y apt-transport-https ca-certificates sqlite3 curl
RUN curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b /usr/local/bin
RUN curl -sSfL https://raw.githubusercontent.com/anchore/grype/main/install.sh | sudo sh -s -- -b /usr/local/bin
COPY --from=builder /home/sbom-service/ /home/sbom-service/
CMD ["./SBOMservice"]