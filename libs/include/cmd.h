//
// Created by lrf141 on 23/03/04.
//

#ifndef KTKIT_KENTSU_TOOL_KIT_CMD_H
#define KTKIT_KENTSU_TOOL_KIT_CMD_H

#include <boost/program_options.hpp>

namespace cmd {
    using namespace boost::program_options;
    class Command {
    private:
        options_description *optionsDescription;
        variables_map *variablesMap;
        bool emptyOptionFlag;
    public:
        Command(std::string);
        ~Command();

        variables_map getVariablesMap();
        options_description getOptionsDescription();
        options_description_easy_init add_options();
        void parseCommandLineOptions(int, char **);
        bool isEmptyOptions();
        bool contains(std::string);
    };
}

#endif //KTKIT_KENTSU_TOOL_KIT_CMD_H
