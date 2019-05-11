FROM golang:latest AS build_base
WORKDIR /go/src/github.com/chazapp/dante
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod vendor

FROM build_base AS dante
RUN mkdir -p /go/src/github.com/chazapp/dante
COPY . /go/src/github.com/chazapp/dante
RUN cd /go/src/github.com/chazapp/dante/cmd/cli \
        && go build -o dante-cli \
        && mv dante-cli /bin && ls
WORKDIR /go/src/github.com/chazapp/dante
ENTRYPOINT ["dante-cli", "elastic", \
            "--file", "/go/src/github.com/chazapp/dante/datasets/dataset.txt", \
            "--name", "annees-annie", \
            "--db", "http://elasticsearch:9200"]
