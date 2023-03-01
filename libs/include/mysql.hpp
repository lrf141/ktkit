//
// Created by lrf141 on 23/02/27.
//

#ifndef KTKIT_KENTSU_TOOL_KIT_MYSQL_HPP
#define KTKIT_KENTSU_TOOL_KIT_MYSQL_HPP

#include <iostream>
#include <mysqlx/xdevapi.h>

namespace myconn {
    class MySQL {
    public:
        static mysqlx::SessionSettings createSessionSettings(
                const std::string& user, const std::string& pass,
                unsigned int port, const std::string& host, const std::string& db
        );

        static mysqlx::Client createClient(
		const std::string& user, const std::string& pass,
		unsigned int port, const std::string& host, const std::string& db
        	);
        mysqlx::Session createSession(mysqlx::Client client);
        mysqlx::Session createSession(mysqlx::SessionSettings sessionSettings);
    };
}

#endif //KTKIT_KENTSU_TOOL_KIT_MYSQL_HPP
