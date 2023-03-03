//
// Created by lrf141 on 23/02/27.
//

#ifndef KTKIT_KENTSU_TOOL_KIT_MYSQL_H
#define KTKIT_KENTSU_TOOL_KIT_MYSQL_H

#include <iostream>
#include <mysqlx/xdevapi.h>

namespace myconn {
    class Util {
    public:
        static mysqlx::SessionSettings createSessionSettings(
                const std::string& user, const std::string& pass,
                unsigned int port, const std::string& host, const std::string& db
        );

        static mysqlx::Client createClient(
		const std::string& user, const std::string& pass,
		const std::string& host, unsigned int port, const std::string& db
        	);
        mysqlx::Session createSession(mysqlx::Client client);
        mysqlx::Session createSession(mysqlx::SessionSettings sessionSettings);
    };

    class Config {
    private:
        std::string user;
        unsigned int port;
        std::string host;
        std::string pass;
        std::string database;
    public:
        Config(std::string, unsigned int, std::string, std::string, std::string);
        mysqlx::Client createClient();
    };
}

#endif //KTKIT_KENTSU_TOOL_KIT_MYSQL_H
