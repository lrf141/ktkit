#include <iostream>
#include <boost/program_options.hpp>

int main(int argc, char **argv) {
    using namespace boost::program_options;
    options_description description("Options");
    description.add_options()
            ("help,H", "command helps")
            ("version,v", "version");

    variables_map variablesMap;
    store(parse_command_line(argc, argv, description), variablesMap);
    notify(variablesMap);
    if (variablesMap.count("help"))
        std::cout << description << std::endl;
	return 0;
}

