cmake_minimum_required(VERSION 3.8)

project(innodb_buffer_page_list VERSION 0.0.1)

file(GLOB innodb_buffer_page_list_src "*.h" "*.cpp")

add_executable(kt-innodb-buffer-page-list ${innodb_buffer_page_list_src})

target_include_directories(kt-innodb-buffer-page-list PRIVATE ${MYSQL_CONNECTOR_INCLUDE_DIR})
target_link_libraries(kt-innodb-buffer-page-list ${MYSQL_CONN_CPP_LIBS})
target_include_directories(kt-innodb-buffer-page-list PRIVATE ${KTKIT_INCLUDE_DIR})
target_link_libraries(kt-innodb-buffer-page-list ktkit)
