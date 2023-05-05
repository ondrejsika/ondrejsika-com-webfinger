FROM golang:1.19 as build
WORKDIR /build
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=build /build/ondrejsika-com-webfinger /
CMD ["/ondrejsika-com-webfinger"]
EXPOSE 8000
