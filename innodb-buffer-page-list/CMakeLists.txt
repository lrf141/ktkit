cmake_minimum_required(VERSION 3.8)

project(innodb_buffer_page_list VERSION 0.0.1)

file(GLOB innodb_buffer_page_list_src "*.hpp" "*.cpp")
add_executable(kt-innodb-buffer-page-list ${innodb_buffer_page_list_src})

target_include_directories(kt-innodb-buffer-page-list PRIVATE ${Boost_INCLUDE_DIR})
target_link_libraries(kt-innodb-buffer-page-list ${Boost_LIBRARIES})
target_include_directories(kt-innodb-buffer-page-list PRIVATE /usr/local/lib/mysql-connector-cpp-8-0-30/include)
if (UNIX AND NOT APPLE) 
    target_link_libraries(kt-innodb-buffer-page-list /usr/local/lib/mysql-connector-cpp-8-0-30/lib64/libmysqlcppconn.so)
elseif (UNIX AND APPLE) 
    target_link_libraries(kt-innodb-buffer-page-list /usr/local/lib/mysql-connector-cpp-8-0-30/lib64/libmysqlcppconn.dylib)
else() 
    target_link_libraries(kt-innodb-buffer-page-list /usr/local/lib/mysql-connector-cpp-8-0-30/lib64/libmysqlcppconn.so)
endif()
target_include_directories(kt-innodb-buffer-page-list PRIVATE ${CMAKE_SOURCE_DIR}/libs/include)
target_link_libraries(kt-innodb-buffer-page-list ktkit)
