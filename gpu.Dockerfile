FROM golang:1.20 as builder

RUN apt update && apt install -y \
        g++ \
        git \
        cmake \
        patch \
        ocl-icd-opencl-dev \
        opencl-headers \
        libclblast-dev && \
    rm -rf /var/lib/apt/lists/*

ADD . /build

WORKDIR /build

RUN git clone --recurse-submodules https://github.com/go-skynet/go-llama.cpp && \
    cd go-llama.cpp && \
    git checkout --recurse-submodules f104111358e8098aea69ce408b85b707528179ef && \
    cd .. && \
    BUILD_TYPE=clblas CLBLAS_DIR=... make -C go-llama.cpp/ libbinding.a && \
    CGO_LDFLAGS="-lOpenCL -lclblast -L/usr/local/lib64/" C_INCLUDE_PATH=/build/go-llama.cpp CGO_CXXFLAGS="-I/build/go-llama.cpp/llama.cpp/examples -I/build/go-llama.cpp/llama.cpp/" LIBRARY_PATH=/build/go-llama.cpp go build -o ./aichatllama ./cmd/aichatllama

FROM debian:12

RUN apt update && \
    apt install -y \
        ca-certificates  \
        xz-utils \
        ocl-icd-libopencl1 \
        opencl-headers \
        libclblast1 \
        clinfo && \
    update-ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /build/aichatllama /usr/local/bin/aichatllama

RUN mkdir -p /etc/OpenCL/vendors && \
    echo "libnvidia-opencl.so.1" > /etc/OpenCL/vendors/nvidia.icd
ENV NVIDIA_VISIBLE_DEVICES all
ENV NVIDIA_DRIVER_CAPABILITIES compute,utility