FROM node:latest as build-stage
WORKDIR /app
COPY package*.json ./
RUN npm install
RUN npm install -g @vue/cli
CMD ["npm", "run", "serve", "--", "--port", "8081"]
