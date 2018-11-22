FROM golang:1.11 as build

WORKDIR /
COPY . /
ENV CGO_ENABLED 0
RUN go build -tags 'static' -ldflags '-extldflags "-static"' .

FROM scratch
LABEL maintainer="https://github.com/Ilyes512/boilr"

COPY --from=build /boilr /
ENTRYPOINT ["/boilr"]
CMD ["-help"]
