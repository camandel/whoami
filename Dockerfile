FROM golang:1.14-alpine AS build
ADD . /app
WORKDIR /app
ARG color=red
RUN go build -ldflags "-X main.DefaultColor=${color}" -o http

FROM alpine:3.12
WORKDIR /app
ENV PORT=8000
EXPOSE 8000
COPY --from=build /app/http /app
CMD ["/app/http"]