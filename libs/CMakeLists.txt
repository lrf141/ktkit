cmake_minimum_required(VERSION 3.8)

project("ktkit-libs" VERSION 0.0.1)

FILE(GLOB ktkit_libs_src "include/*.hpp" "*.cpp")

set(MYSQL_CONNECTOR_CPP_DIR "/usr/local/lib/mysql-connector-cpp-8-0-30")

add_library(ktkit STATIC ${ktkit_libs_src})

find_library(
        MYSQL_CONN_CPP_LIB
        NAMES mysqlcppconn8 libmysqlconncpp8
        HINTS "${MYSQL_CONNECTOR_CPP_DIR}/lib64"
        NO_DEFAULT_PATH
)

link_libraries(${MYSQL_CONNECTOR_CPP_DIR}/lib64)
target_include_directories(ktkit PUBLIC ${MYSQL_CONNECTOR_CPP_DIR}/include)
target_link_libraries(ktkit ${MYSQL_CONN_CPP_LIB})
target_include_directories(ktkit PRIVATE ${Boost_INCLUDE_DIR})
target_link_libraries(ktkit ${Boost_LIBRARIES})
