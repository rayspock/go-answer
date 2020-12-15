# use multi-stage build to reduce the size of final image
# build stage
FROM golang:alpine as builder

# enable module support
ENV GO111MODULE=on

RUN mkdir /build 
ADD . /build/
WORKDIR /build 

COPY go.mod .
COPY go.sum .

# get dependancies - will also be cached if we won't change mod/sum
RUN go mod download

COPY . .

# build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./main

# final stage
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
# copy our static executable.
COPY --from=builder /build/main  /app/
COPY --from=builder /build/.env.local /app/
WORKDIR /app/

# Environment Variables
ENV GO_ENV=production

CMD ["./main"]