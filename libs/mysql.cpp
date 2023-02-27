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

    mysqlx::Client MySQL::createClient(const std::string &user, const std::string &pass, const unsigned int port,
    					const std::string &host, const std::string &db) {
    	mysqlx::Client client(
    		mysqlx::SessionOption::USER, user,
    		mysqlx::SessionOption::PWD, pass,
    		mysqlx::SessionOption::HOST, host,
    		mysqlx::SessionOption::PORT, port,
    		mysqlx::SessionOption::DB, db
    		);
    	return client;
    }

    mysqlx::Session MySQL::createSession(mysqlx::Client client) {
    	mysqlx::Session session(client);
    	return session;
    }

    mysqlx::Session MySQL::createSession(mysqlx::SessionSettings sessionOption) {
    	mysqlx::Session session(sessionOption);
    	return session;
    }
}