//
// Created by lrf141 on 23/02/27.
//

#ifndef KTKIT_KENTSU_TOOL_KIT_UTIL_HPP
#define KTKIT_KENTSU_TOOL_KIT_UTIL_HPP

#include <iostream>

#define MYSQLX_URL_FORMAT "mysqlx://%s@%s:%s"

namespace util {
    class Util {
    public:
        static std::string createMySQLUrl(std::string user, std::string host, std::string port);
    };
}


#endif //KTKIT_KENTSU_TOOL_KIT_UTIL_HPP
