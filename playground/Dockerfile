FROM --platform=linux/amd64 debian:buster-slim AS base
RUN apt-get update && apt-get install -y --no-install-recommends \
	curl ca-certificates tar xz-utils make git libtinfo5 libxml2 \
	libgcc-8-dev binutils gcc-8-base libc6-dev-i386 \
	libgcc-8-dev-arm64-cross binutils-aarch64-linux-gnu gcc-8-aarch64-linux-gnu-base libc6-dev-arm64-cross \
	libgcc-8-dev-armhf-cross binutils-arm-linux-gnueabihf gcc-8-arm-linux-gnueabihf-base libc6-dev-armhf-cross \
	&& rm -rf /var/lib/apt/lists/*

ENV CLANG_ARCHIVE clang+llvm-8.0.0-x86_64-linux-gnu-ubuntu-18.04
# this is probably stripping too much from llvm, but it works for wasi-sysroot + hello world.
RUN curl -LOSs https://releases.llvm.org/8.0.0/${CLANG_ARCHIVE}.tar.xz && \
	tar xvf ${CLANG_ARCHIVE}.tar.xz -C /usr/local/ --strip-components=1 \
		${CLANG_ARCHIVE}/bin/clang \
		${CLANG_ARCHIVE}/bin/clang-8 \
		${CLANG_ARCHIVE}/bin/llvm-ar \
		${CLANG_ARCHIVE}/bin/llvm-nm \
		${CLANG_ARCHIVE}/bin/wasm-ld \
		${CLANG_ARCHIVE}/bin/lld \
		${CLANG_ARCHIVE}/lib/clang/8.0.0/include \
	&& rm ${CLANG_ARCHIVE}.tar.xz

WORKDIR /src
RUN git clone https://github.com/CraneStation/wasi-sysroot wasi-sysroot
WORKDIR wasi-sysroot
RUN make
RUN mkdir -p /usr/local/lib/clang/8.0.0/lib/wasi && echo '!<arch>' > /usr/local/lib/clang/8.0.0/lib/wasi/libclang_rt.builtins-wasm32.a


FROM base AS build

WORKDIR /src
COPY sock_send.c source.c


FROM build AS build-wasm
RUN clang --target=wasm32-unknown-wasi --sysroot=/src/wasi-sysroot/sysroot -o /program source.c

FROM build AS build-amd64
RUN clang -o /program source.c

FROM build AS build-arm64
RUN clang --target=aarch64-linux-gnu -o /program source.c

FROM build AS build-armhf
RUN clang --target=arm-linux-gnueabihf -o /program source.c

FROM build-$TARGETARCH AS build-final

# arm-linux-gnueabihf
# aarch-linux-gnu

FROM scratch
COPY --from=build-final /program /
ENTRYPOINT ["/program"]
