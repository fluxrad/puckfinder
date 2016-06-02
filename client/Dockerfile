FROM ubuntu:xenial

VOLUME ["/app"]

RUN apt-get -yqq update && \
    apt-get -qq install curl && \
    apt-get clean

RUN curl -sL https://deb.nodesource.com/setup_4.x | bash - &&  apt-get install -yq nodejs

WORKDIR /app

# Seriously though, I hate JavaScript
RUN npm -g install typings && typings install es6-promise --ambient 
