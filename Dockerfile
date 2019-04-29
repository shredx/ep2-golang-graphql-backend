FROM melvinodsa/go-web-application:latest
LABEL maintainer="melvinodsa@gmail.com"

# go get the dependencies and clone the repo
COPY . $GOPATH/src/github.com/shredx/ep2-golang-graphql-backend
WORKDIR $GOPATH/src/github.com/shredx/ep2-golang-graphql-backend
RUN cd $GOPATH/src/github.com/shredx/ep2-golang-graphql-backend \
    && dep ensure

EXPOSE 9090/tcp
EXPOSE 9090/udp

ENTRYPOINT ["revel", "run", "-v", "github.com/shredx/ep2-golang-graphql-backend"]