//
// Created by lrf141 on 23/02/27.
//
#include <boost/format.hpp>
#include "include/mysql.h"

namespace myconn {

    Util::Util() {
        throw std::bad_exception();
    }

    mysqlx::SessionSettings Util::createSessionSettings(const std::string &user, const std::string &pass,
                                                        const std::string &host, const unsigned int port,
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

    std::string Util::createSelectAllQuery(std::string table) {
        boost::basic_format<char> formattedQuery = boost::format(SELECT_ALL_FORMAT) % table;
        return formattedQuery.str();
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

    mysqlx::SessionSettings Config::createSessionSettings() {
        return Util::createSessionSettings(user, pass, host, port, database);
    }

    std::string Config::getDatabase() {
        return this->database;
    }
}