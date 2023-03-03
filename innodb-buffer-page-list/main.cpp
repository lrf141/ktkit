#include <iostream>
#include <boost/program_options.hpp>
#include "cmd.h"
#include "main.h"
#include "mysql.hpp"

int main(int argc, char **argv) {
    using namespace boost::program_options;
    cmd::Command command(CMD_CAPTION);
    command.add_options()
            ("help,H", "show help")
            ("version,v", "version")
            ("user,u", value<std::string>(),"mysql user name (required)")
            ("password,p", value<std::string>()->default_value(""),"mysql user password (optional)")
            ("host,h", value<std::string>()->default_value(DEFAULT_HOST), "mysql host (optional)")
            ("port,P", value<unsigned int>()->default_value(DEFAULT_PORT),"mysql port number (optional)")
            ("database,D", value<std::string>()->default_value(""), "mysql database name (optional)");
    try {
        command.parseCommandLineOptions(argc, argv);
    } catch (const error_with_option_name& e) {
        std::cout << e.what() << std::endl;
        return -1;
    }

    if (command.isEmptyOptions() || command.contains("help")) {
        std::cout << command.getOptionsDescription() << std::endl;
        return 0;
    }

    if (command.contains("version")) {
        std::cout << VERSION << std::endl;
        return 0;
    }

    variables_map variablesMap = command.getVariablesMap();
    try {
        const std::string user = variablesMap["user"].as<std::string>();
        const unsigned int port = variablesMap["port"].as<unsigned int>();
        const std::string host = variablesMap["host"].as<std::string>();
        const std::string pass = variablesMap["password"].as<std::string>();
        const std::string database = variablesMap["database"].as<std::string>();

        mysqlx::SessionSettings sessionSettings = myconn::MySQL::createSessionSettings(user, pass, port, host, database);
        mysqlx::Session session(sessionSettings);
    } catch (const boost::bad_any_cast& e) {
        std::cout << e.what() << std::endl;
        return 1;
    }

	return 0;
}

