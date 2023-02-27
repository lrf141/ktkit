//
// Created by lrf141 on 23/02/27.
//

#ifndef KTKIT_KENTSU_TOOL_KIT_MYSQL_HPP
#define KTKIT_KENTSU_TOOL_KIT_MYSQL_HPP

#include <iostream>
#include <mysqlx/xdevapi.h>

namespace myconn {
    class MySQL {
    private:
        static mysqlx::SessionSettings createSessionSettings(
                const std::string& user, const std::string& pass,
                const unsigned int port, const std::string& host, const std::string& db
        );
    public:
        static mysqlx::Session createSession(
                const std::string& user, const std::string& pass,
                const unsigned int port, const std::string& host, const std::string& db
                );
    };
}

#endif //KTKIT_KENTSU_TOOL_KIT_MYSQL_HPP
