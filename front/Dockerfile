FROM node

RUN mkdir -p /usr/src/app
WORKDIR /usr/src/app

COPY ./ /usr/src/app/
RUN yarn && yarn build

ENTRYPOINT [ "yarn", "serve", "--host=0.0.0.0", "--port=8032" ]