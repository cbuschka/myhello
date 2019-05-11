FROM node:10.15-stretch-slim

WORKDIR /app
COPY / /app
RUN npm -g install yarn && \
	yarn install

EXPOSE 8080

CMD [ "bash", "-c", "export PORT=8080; cd /app/ && exec node ./src/index.js" ]

