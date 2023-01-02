FROM golang:alpine AS build

ENV FIREBASE_PROJECT_ID="library-emulator"
ENV FIRESTORE_EMULATOR_HOST="firebase:4242"

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o main

FROM alpine:latest

COPY --from=build app/main .

EXPOSE 4042

CMD ["./main"]
