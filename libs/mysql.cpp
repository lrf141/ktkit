//
// Created by lrf141 on 23/02/27.
//
#include "include/mysql.h"

namespace myconn {

    mysqlx::SessionSettings Util::createSessionSettings(const std::string &user, const std::string &pass,
                                                         const unsigned int port, const std::string &host,
                                                         const std::string &db) {
        mysqlx::SessionSettings sessionSettings = mysqlx::SessionSettings(host, port, user, pass, db);
        return sessionSettings;
    }

    mysqlx::Client Util::createClient(const std::string &user, const std::string &pass, const std::string &host,
                                      const unsigned int port, const std::string &db) {
    	mysqlx::Client client(
    		mysqlx::SessionOption::USER, user,
    		mysqlx::SessionOption::PWD, pass,
    		mysqlx::SessionOption::HOST, host,
    		mysqlx::SessionOption::PORT, port,
    		mysqlx::SessionOption::DB, db
    		);
    	return client;
    }

    mysqlx::Session Util::createSession(mysqlx::Client client) {
    	mysqlx::Session session(client);
    	return session;
    }

    mysqlx::Session Util::createSession(mysqlx::SessionSettings sessionOption) {
    	mysqlx::Session session(sessionOption);
    	return session;
    }

    Config::Config(std::string user, unsigned int port, const std::string host, std::string pass, std::string database) {
        this->user = user;
        this->port = port;
        this->host = host;
        this->pass = pass;
        this->database = database;
    }

    mysqlx::Client Config::createClient() {
        return Util::createClient(user, pass, host, port, database);
    }
}