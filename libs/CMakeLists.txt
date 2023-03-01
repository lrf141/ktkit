cmake_minimum_required(VERSION 3.8)

project("ktkit-libs" VERSION 0.0.1)

FILE(GLOB ktkit_libs_src "include/*.hpp" "*.cpp")

add_library(ktkit STATIC ${ktkit_libs_src})

link_libraries(${MYSQL_CONNECTOR_CPP_DIR}/lib64)
target_include_directories(ktkit PUBLIC ${MYSQL_CONNECTOR_INCLUDE_DIR})
target_link_libraries(ktkit ${MYSQL_CONN_CPP_LIBS})
target_include_directories(ktkit PRIVATE ${Boost_INCLUDE_DIR})
target_link_libraries(ktkit ${Boost_LIBRARIES})
