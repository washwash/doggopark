# Build
FROM golang:1.14 AS build

RUN rm -rf /app /build
COPY . /build

WORKDIR /build
RUN CGO_ENABLED=0 GOBIN=/app go install


# Run
EXPOSE 8080
CMD ["/app/doggo"]
