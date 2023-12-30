FROM node:16.20-slim

ARG GIT_SHA

ENV GIT_SHA=${GIT_SHA}

WORKDIR /app

COPY . /app

RUN npm install

EXPOSE 3000

CMD ["npm", "start"]