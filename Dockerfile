FROM trustwallet/wallet-core

WORKDIR /var/www/html/crypto-wallet

#Some configurations to avoid errors and apt update
RUN wget -O - https://apt.kitware.com/keys/kitware-archive-latest.asc 2>/dev/null | \
            gpg --dearmor - | \
            tee /etc/apt/trusted.gpg.d/kitware.gpg >/dev/null && \
    apt-add-repository 'deb https://apt.kitware.com/ubuntu/ bionic main' && \
    apt-get update

#Install go on wallet-core image
RUN wget https://golang.org/dl/go1.18.3.linux-amd64.tar.gz -P /usr/local && \
    tar -xvf /usr/local/go1.18.3.linux-amd64.tar.gz -C /usr/local && \
    chown -R root:root /usr/local/go && \
    rm /usr/local/go1.18.3.linux-amd64.tar.gz

# Golang config
ENV GO111MODULE=on
ENV GOROOT "/usr/local/go"
ENV GOPATH "/root/go"
ENV PATH "$PATH:$GOROOT/bin"
ENV PATH "$PATH:$GOPATH/bin"

# Install go packages
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download
RUN go mod tidy

# Copy source code
COPY . .
