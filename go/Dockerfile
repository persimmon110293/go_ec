FROM golang:1.22-alpine
ARG PROTOBUF_VERSION=27.1

RUN apk add curl unzip

WORKDIR /tmp/protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip && \
    unzip protoc.zip && \
    mv bin/* /usr/local/bin/ && \
    mv include/* /usr/local/include/


WORKDIR /app
COPY ./go/src .

ENV GOBIN=/usr/local/bin
RUN go get -u google.golang.org/grpc \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest \
    && go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest \
    && go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc \
    && go mod tidy

CMD ["go", "run", "server/main.go"]