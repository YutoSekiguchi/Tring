FROM node:15.8 as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install

COPY . .
RUN npm run build

# Nginx
FROM nginx:stable-alpine
COPY --from=build-stage /app/dist /var/www/html
COPY --from=build-stage /app/nginx.conf /etc/nginx/nginx.conf
COPY --from=build-stage /app/default.conf /etc/nginx/conf.d/default.conf
EXPOSE 80

CMD [ "nginx", "-g", "daemon off;" ]