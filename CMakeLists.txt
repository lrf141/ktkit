cmake_minimum_required(VERSION 3.8)

project("ktkit - kentsu tool kit" VERSION 0.0.1)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_COMPILER "/usr/bin/g++" CACHE STRING "gnu c++ compiler" FORCE)
set(CMAKE_C_COMPILER "/usr/bin/gcc" CACHE STRING "gnu cc compiler" FORCE)
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -Wall -O3")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wall -O3 -std=c++11")

find_package(Boost 1.74 COMPONENTS program_options REQUIRED)

if (UNIX AND NOT APPLE)
    set(MYSQL_CONNECTOR_CPP_DIR "/usr/lib/x86_64-linux-gnu")
    set(MYSQL_CONNECTOR_INCLUDE_DIR "/usr/include/mysql-cppconn-8")
elseif (APPLE)
    set(MYSQL_CONNECTOR_CPP_DIR "/usr/local/lib/mysql-connector-cpp-8-0-30/lib64")
    set(MYSQL_CONNECTOR_INCLUDE_DIR "/usr/local/lib/mysql-connector-cpp-8-0-30/include")
else()
    message("-- not supported os")
    set(MYSQL_CONNECTOR_CPP_DIR "/usr/lib/x86_64-linux-gnu")
    set(MYSQL_CONNECTOR_INCLUDE_DIR "/usr/include/mysql-cppconn-8")
endif()

find_library(
        MYSQL_CONN_CPP_LIBS
        NAMES mysqlcppconn8 libmysqlconncpp8
        HINTS "${MYSQL_CONNECTOR_CPP_DIR}"
        NO_DEFAULT_PATH
)
message("-- mysql connector libs - ${MYSQL_CONN_CPP_LIBS}")

set(KTKIT_INCLUDE_DIR ${CMAKE_SOURCE_DIR}/libs/include)
message("-- ktkit include dir - ${KTKIT_INCLUDE_DIR}")

add_subdirectory(libs)
add_subdirectory(innodb-buffer-page-list)
