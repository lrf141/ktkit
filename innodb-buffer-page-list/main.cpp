#include <iostream>
#include <boost/program_options.hpp>
#include <mysql/jdbc.h>
#include "cmd.hpp"

int main(int argc, char **argv) {
    using namespace boost::program_options;
    options_description description(CMD_CAPTION);
    description.add_options()
            ("help,H", "show help")
            ("version,v", "version")
            ("user,u", "mysql user name(required)")
            ("host,h", "mysql host(required)")
            ("password,p", "mysql user password")
            ("port,P", "mysql port number");

    variables_map variablesMap;
    store(parse_command_line(argc, argv, description), variablesMap);
    notify(variablesMap);

    if (variablesMap.count("help")) {
        std::cout << description << std::endl;
    } else if (variablesMap.count("version")) {
        std::cout << VERSION << std::endl;
    } else {
        std::cout << MYSQL_CONCPP_VERSION_NUMBER << std::endl;
        std::cout << description << std::endl;
    }

	return 0;
}

