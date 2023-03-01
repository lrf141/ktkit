cmake_minimum_required(VERSION 3.8)

project(innodb_buffer_page_list VERSION 0.0.1)

file(GLOB innodb_buffer_page_list_src "*.hpp" "*.cpp")
set(MYSQL_CONNECTOR_CPP_DIR "/usr/local/lib/mysql-connector-cpp-8-0-30")
find_library(
        MYSQL_CONN_CPP_LIB
        NAMES mysqlcppconn8 libmysqlconncpp8
        HINTS "${MYSQL_CONNECTOR_CPP_DIR}/lib64"
        NO_DEFAULT_PATH
)

add_executable(kt-innodb-buffer-page-list ${innodb_buffer_page_list_src})

include_directories(${CMAKE_SOURCE_DIR}/libs/include)
include_directories(${MYSQL_CONNECTOR_CPP_DIR}/include)
target_include_directories(kt-innodb-buffer-page-list PUBLIC ${MYSQL_CONNECTOR_CPP_DIR}/include)
target_link_libraries(kt-innodb-buffer-page-list ${MYSQL_CONN_CPP_LIB})
target_include_directories(kt-innodb-buffer-page-list PRIVATE ${MYSQL_CONNECTOR_CPP_DIR}/include)
link_directories(${CMAKE_BINARY_DIR}/libs)
target_link_libraries(kt-innodb-buffer-page-list ktkit)
