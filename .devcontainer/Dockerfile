# copy from https://github.com/microsoft/vscode-dev-containers/blob/master/containers/go/.devcontainer/Dockerfile
# [Choice] Go version: 1, 1.16, 1.15
ARG VARIANT=1
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler

USER vscode
RUN  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26.0 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
