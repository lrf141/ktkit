#include <iostream>
#include <boost/program_options.hpp>
#include "cmd.hpp"
#include "util.hpp"

int main(int argc, char **argv) {
    using namespace boost::program_options;
    options_description description(CMD_CAPTION);
    description.add_options()
            ("help,H", "show help")
            ("version,v", "version")
            ("user,u", value<std::string>(),"mysql user name (required)")
            ("password,p", value<std::string>()->default_value(""),"mysql user password (optional)")
            ("host,h", value<std::string>()->default_value(DEFAULT_HOST), "mysql host (optional)")
            ("port,P", value<std::string>()->default_value(DEFAULT_PORT),"mysql port number (optional)");

    variables_map variablesMap;
    try {
        store(parse_command_line(argc, argv, description), variablesMap);
    } catch (const error_with_option_name& e) {
        std::cout << e.what() << std::endl;
        return 1;
    }
    notify(variablesMap);

    if (variablesMap.count("help")) {
        std::cout << description << std::endl;
        return 0;
    } else if (variablesMap.count("version")) {
        std::cout << VERSION << std::endl;
        return 0;
    }

    try {
        const std::string user = variablesMap["user"].as<std::string>();
        const std::string port = variablesMap["port"].as<std::string>();
        const std::string host = variablesMap["host"].as<std::string>();
        const std::string pass = variablesMap["password"].as<std::string>();
        const std::string url = util::Util::createMySQLUrl(user, host, port);
        std::cout << url << std::endl;
    } catch (const boost::bad_any_cast& e) {
        std::cout << e.what() << std::endl;
        return 1;
    }

	return 0;
}

