#!/bin/bash

docker build --platform wasi/wasm --output . .
mv program sock_send.wasm
