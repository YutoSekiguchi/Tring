FROM node:15.8


WORKDIR /app

COPY package*.json ./
RUN npm install


CMD [ "npm", "run", "serve" ]