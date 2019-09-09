FROM node:latest
ENV NPM_CONFIG_LOGLEVEL warn

COPY package.json package.json
RUN npm install -g serve
RUN npm install
COPY squawk-project-front .
RUN npm run build --production

CMD serve -s build

EXPOSE 5000
