#include <iostream>
#include <memory>
#include <boost/program_options.hpp>
#include "cmd.h"
#include "main.h"
#include "mysql.h"
#include "innodb_buffer_page.h"

int main(int argc, char **argv) {
    using namespace boost::program_options;
    cmd::Command command(CMD_CAPTION);
    command.add_options()
            ("help,H", "show help")
            ("version,v", "version")
            ("user,u", value<std::string>(),"mysql user name (required)")
            ("password,p", value<std::string>()->default_value(""),"mysql user password (optional)")
            ("host,h", value<std::string>()->default_value(DEFAULT_HOST), "mysql host (optional)")
            ("port,P", value<unsigned int>()->default_value(DEFAULT_PORT),"mysql port number (optional)");

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
    std::unique_ptr<myconn::Config> config;
    try {
        const std::string user = variablesMap["user"].as<std::string>();
        const unsigned int port = variablesMap["port"].as<unsigned int>();
        const std::string host = variablesMap["host"].as<std::string>();
        const std::string pass = variablesMap["password"].as<std::string>();

        config = std::unique_ptr<myconn::Config>(new myconn::Config(user, port, host, pass, I_S));
    } catch (const boost::bad_any_cast& e) {
        std::cout << e.what() << std::endl;
        return 1;
    }

    try {
        mysqlx::SessionSettings sessionSettings = config->createSessionSettings();
        mysqlx::Session session(sessionSettings);
        mysqlx::Table table = session.getSchema(config->getDatabase(), true).getTable(INNODB_BUFFER_PAGE);
        mysqlx::TableSelect tableSelect = table.select();
        mysqlx::RowResult rowResult = tableSelect.execute();
        std::vector<idb_buffer_page::InnoDBBufferPage> innodbBufferPages;
        unsigned long int rowCount = 0;
        unsigned long int numberRecords = 0;
        unsigned long int dataFreePageCount = 0;
        for (auto row : rowResult) {
            unsigned long int colCount = row.colCount();
            idb_buffer_page::InnoDBBufferPage innoDbBufferPage;
            for (int j = 0; j < colCount; j++) {
                mysqlx::Value value = row.get(j);
                std::string columnName = rowResult.getColumn(j).getColumnName();
                innoDbBufferPage.convert(value, columnName);
            }
            rowCount++;
            numberRecords += innoDbBufferPage.getNumberRecords();
            if (innoDbBufferPage.getDataSize() == 0) {
                dataFreePageCount++;
            }
        }
        std::cout << "Total pages: " << rowCount << std::endl;
        std::cout << "Total number records: " << numberRecords << std::endl;
        std::cout << "Total data free pages: " << dataFreePageCount << std::endl;
    } catch (const mysqlx::Error &err) {
        std::cout << "MySQL Error: " << err << std::endl;
        return 1;
    } catch (const std::exception &e) {
        std::cout << "Std Exception: " << e.what() << std::endl;
        return 1;
    } catch (const char *e) {
        std::cout << "Exception: " << e << std::endl;
        return 1;
    }

	return 0;
}

