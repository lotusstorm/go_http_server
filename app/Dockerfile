FROM node:lts-alpine

WORKDIR /app

RUN chown -R node:node .

USER node

COPY ["package.json", "package-lock.json*", "./"]

RUN npm ci

COPY . .

CMD ["npm", "run", "dev"]