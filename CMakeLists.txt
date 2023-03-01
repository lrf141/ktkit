cmake_minimum_required(VERSION 3.8)

project("ktkit - kentsu tool kit" VERSION 0.0.1)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_COMPILER "/usr/bin/g++" CACHE STRING "gnu c++ compiler" FORCE)
set(CMAKE_C_COMPILER "/usr/bin/gcc" CACHE STRING "gnu cc compiler" FORCE)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -O3")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -O3 -std=c++14")

find_package(Boost 1.74 COMPONENTS program_options REQUIRED)

add_subdirectory(libs)
add_subdirectory(innodb-buffer-page-list)
