FROM golang:latest
RUN mkdir -p /go/src/github.com/chazapp/tigbra
COPY . /go/src/github.com/chazapp/tigbra
RUN ls && pwd
RUN cd /go/src/github.com/chazapp/tigbra/cmd/cli && go build -o main && mv main /bin && ls
WORKDIR /go/src/github.com/chazapp/tigbra
ENTRYPOINT ["main", "print", "--file", "/go/src/github.com/chazapp/tigbra/datasets/dataset.txt"]
