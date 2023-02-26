//
// Created by lrf141 on 23/02/27.
//
#include <boost/format.hpp>
#include "include/util.hpp"

namespace util {
    std::string Util::createMySQLUrl(std::string user, std::string host, std::string port) {
        const boost::basic_format<char> formatted_url = boost::format(MYSQLX_URL_FORMAT) % user % host % port;
        return formatted_url.str();
    }
}