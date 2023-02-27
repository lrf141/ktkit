//
// Created by lrf141 on 23/02/27.
//
#include "include/mysql.hpp"

namespace myconn {

    mysqlx::SessionSettings MySQL::createSessionSettings(const std::string &user, const std::string &pass,
                                                         const unsigned int port, const std::string &host,
                                                         const std::string &db) {
        mysqlx::SessionSettings sessionSettings = mysqlx::SessionSettings(host, port, user, pass, db);
        return sessionSettings;
    }

    mysqlx::Session MySQL::createSession(const std::string &user, const std::string &pass, const unsigned int port,
                                         const std::string &host, const std::string &db) {
        mysqlx::SessionSettings sessionSettings = createSessionSettings(user, pass, port, host, db);
        mysqlx::Session session(sessionSettings);
        return session;
    }
}