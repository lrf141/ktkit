cmake_minimum_required(VERSION 3.22.0)

project("ktkit - kentsu tool kit" VERSION 0.0.1)

find_package(Boost 1.74 COMPONENTS program_options REQUIRED)
add_subdirectory(innodb-buffer-page-list)

set(CMAKE_CXX_COMPILER "/usr/bin/g++")
set(CMAKE_C_COMPILER "/usr/bin/gcc")
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -O3")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -O3 -std=c++11")

