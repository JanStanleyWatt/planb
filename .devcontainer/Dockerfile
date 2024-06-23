FROM mcr.microsoft.com/devcontainers/go:1-1.22-bookworm

USER vscode
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
USER root