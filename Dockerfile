FROM golang:latest as build

COPY . /emulator
WORKDIR /emulator

RUN go build -mod=vendor

FROM debian:testing as production

COPY --from=build /emulator/docker-bigtable-emulator /emulator/docker-bigtable-emulator

EXPOSE 8086

ENTRYPOINT ["/emulator/docker-bigtable-emulator"]
CMD ["--port", "8086"]
