FROM golang:1.15.3-alpine
WORKDIR /wiki-random
ADD . /wiki-random
RUN cd /wiki-random && go build -o wiki-random
EXPOSE 9936
ENTRYPOINT ./wiki-random