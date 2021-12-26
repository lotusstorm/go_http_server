FROM node:16.13.1

WORKDIR /app
COPY package.json package-lock.json /app/
RUN npm ci
COPY . .

# ENTRYPOINT ["npm", "run", "dev"]
