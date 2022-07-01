#!/usr/bin/env bash

git submodule update --init --recursive

#echo $(dirname $(which emcc))/cmake/Modules/Platform/Emscripten.cmake

mkdir js_build
cd js_build

cmake ../ -DCMAKE_TOOLCHAIN_FILE=$(dirname $(which emcc))/cmake/Modules/Platform/Emscripten.cmake
cmake --build . --
