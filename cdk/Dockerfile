# Base Image: Start with the official Go image
FROM docker:dind

RUN apk add bash

# Install Go (using apk)
RUN apk add --no-cache go

# Install required system packages
RUN apk add --no-cache git

# Install Node.js and NPM (using apk)
RUN apk add --no-cache nodejs npm

# Install Yarn globally (remove --global if you want it project-specific)
RUN npm install -g yarn

RUN npm install -g esbuild

# Install AWS CLI
RUN apk add --no-cache \
    python3 \
    py3-pip \
    groff \
    less \
    curl \
    jq  # jq is a useful tool for working with JSON in AWS CLI

RUN apk add --no-cache aws-cli

RUN npm install -g aws-cdk

# Set the working directory (adjust if your project is elsewhere)
WORKDIR /workspace
