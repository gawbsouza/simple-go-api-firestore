FROM node:19-alpine
RUN apk update && apk add bash
RUN apk add openjdk11

RUN npm install -g firebase-tools

WORKDIR /app
COPY firebase.json .

EXPOSE 4242
EXPOSE 4343
EXPOSE 4400

CMD ["sh", "-c", "firebase --project=library-emulator emulators:start --only firestore" ]