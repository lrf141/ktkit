//
// Created by lrf141 on 23/03/04.
//
#include <iostream>
#include "include/cmd.h"

namespace cmd {
    Command::Command(std::string desc) {
        this->variablesMap = new variables_map();
        this->optionsDescription = new options_description(desc);
        this->emptyOptionFlag = false;
    }

    Command::~Command() {
        free(this->optionsDescription);
        free(this->variablesMap);
    }

    options_description Command::getOptionsDescription() {
        return *optionsDescription;
    }

    variables_map Command::getVariablesMap() {
        return *variablesMap;
    }

    options_description_easy_init Command::add_options() {
        return this->optionsDescription->add_options();
    }

    void Command::parseCommandLineOptions(int argc, char **argv) {
        try {
            basic_parsed_options<char> basicParsedOptions = parse_command_line(argc, argv, *optionsDescription);
            if (basicParsedOptions.options.empty()) {
                this->emptyOptionFlag = true;
            }
            store(basicParsedOptions, *variablesMap);
        } catch (const error_with_option_name& e) {
            throw e;
        }
        notify(*variablesMap);
    }

    bool Command::contains(std::string key) {
        return (*variablesMap).count(key);
    }

    bool Command::isEmptyOptions() {
        return this->emptyOptionFlag;
    }
}