# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS build

RUN apk update && apk add --no-cache git

# Create appuser.
ENV USER=appuser
ENV UID=10001 

# See https://stackoverflow.com/a/55757473/12429735RUN 
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /fiber-crud

FROM scratch

# Import the user and group files from the builder.
COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /etc/group /etc/group

COPY --from=build /fiber-crud /fiber-crud

EXPOSE 3000

USER appuser:appuser

ENTRYPOINT ["/fiber-crud"]