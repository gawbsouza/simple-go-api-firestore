FROM node:19-alpine

RUN apk update && apk add bash && apk add openjdk11
RUN npm install -g firebase-tools

WORKDIR /app
COPY firebase.json .

ENV FIREBASE_PROJECT_ID="library-emulator"

EXPOSE 4242
EXPOSE 4343

CMD ["sh", "-c", "firebase --project=$FIREBASE_PROJECT_ID emulators:start --only firestore" ]