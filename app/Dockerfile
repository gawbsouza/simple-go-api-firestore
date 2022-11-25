FROM golang:latest

ENV FIREBASE_PROJECT_ID="library-emulator"
ENV FIRESTORE_EMULATOR_HOST="firebase:4242"

WORKDIR /app
ADD . .
EXPOSE 4042

RUN go build
ENTRYPOINT ./library
