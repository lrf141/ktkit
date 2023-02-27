cmake_minimum_required(VERSION 3.8)

project("ktkit-libs" VERSION 0.0.1)

FILE(GLOB ktkit_libs_src "include/*.hpp" "*.cpp")

add_library(ktkit STATIC ${ktkit_libs_src})
target_include_directories(ktkit PRIVATE ${Boost_INCLUDE_DIR})
target_link_libraries(ktkit ${Boost_LIBRARIES})
target_include_directories(ktkit PRIVATE /usr/local/lib/mysql-connector-cpp-8-0-30/include)
target_link_libraries(ktkit /usr/local/lib/mysql-connector-cpp-8-0-30/lib64/libmysqlcppconn.so)
