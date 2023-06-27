FROM golang:1.20 as builder

RUN apt update && apt install -y \
        g++ \
        git \
        cmake \
        patch && \
    rm -rf /var/lib/apt/lists/*

ADD . /build

WORKDIR /build

RUN git clone --recurse-submodules https://github.com/go-skynet/go-llama.cpp && \
    cd go-llama.cpp && \
    git checkout --recurse-submodules 35a3c99f08b5dfedc4bb92c00a7f4c8450821d45 && \
    cd .. && \
    make -C go-llama.cpp/ libbinding.a && \
    C_INCLUDE_PATH=/build/go-llama.cpp CGO_CXXFLAGS="-I/build/go-llama.cpp/llama.cpp/examples -I/build/go-llama.cpp/llama.cpp/" LIBRARY_PATH=/build/go-llama.cpp go build -o ./aichatllama ./cmd/aichatllama

FROM debian:12

RUN apt update && \
    apt install -y ca-certificates && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/aichatllama /usr/local/bin/aichatllama