cmake_minimum_required(VERSION 3.8)

project("ktkit-libs" VERSION 0.0.1)

FILE(GLOB ktkit_libs_src "include/*.hpp" "*.cpp")

add_library(ktkit STATIC ${ktkit_libs_src})
target_include_directories(ktkit PRIVATE ${Boost_INCLUDE_DIR})
target_link_libraries(ktkit ${Boost_LIBRARIES})
