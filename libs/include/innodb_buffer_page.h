//
// Created by lrf141 on 23/03/05.
//

#ifndef KTKIT_KENTSU_TOOL_KIT_INNODB_BUFFER_PAGE_H
#define KTKIT_KENTSU_TOOL_KIT_INNODB_BUFFER_PAGE_H

#include <iostream>
#include <list>
#include <mysqlx/xdevapi.h>

namespace idb_buffer_page {
    class InnoDBBufferPage {
    private:
        std::list<std::string> UINT64_COL_NAMES = {
                "POOL_ID", "BLOCK_ID", "SPACE", "PAGE_NUMBER",
                "FLUSH_TYPE", "FIX_COUNT", "NEWEST_MODIFICATION", "OLDEST_MODIFICATION",
                "ACCESS_TIME", "NUMBER_RECORDS", "DATA_SIZE", "COMPRESSED_SIZE", "FREE_PAGE_CLOCK"
        };

        std::list<std::string> STRING_COL_NAMES = {
                "PAGE_TYPE", "IS_HASHED", "TABLE_NAME", "INDEX_NAME",
                "PAGE_STATE", "IO_FIX", "IS_OLD", "IS_OLD"
        };
        unsigned int POOL_ID;
        unsigned int BLOCK_ID;
        unsigned int SPACE;
        unsigned int PAGE_NUMBER;
        std::string PAGE_TYPE;
        unsigned int FLUSH_TYPE;
        unsigned int FIX_COUNT;
        std::string IS_HASHED;
        unsigned int NEWEST_MODIFICATION;
        unsigned int OLDEST_MODIFICATION;
        unsigned int ACCESS_TIME;
        std::string TABLE_NAME;
        std::string INDEX_NAME;
        unsigned int NUMBER_RECORDS;
        unsigned int DATA_SIZE;
        unsigned int COMPRESSED_SIZE;
        std::string PAGE_STATE;
        std::string IO_FIX;
        std::string IS_OLD;
        unsigned int FREE_PAGE_CLOCK;
        std::string IS_STALE;
    public:
        void convert(mysqlx::Value&, std::string);
        std::string getPageType();
        unsigned int getNumberRecords();
        unsigned int getDataSize();
    };
}

#endif //KTKIT_KENTSU_TOOL_KIT_INNODB_BUFFER_PAGE_H
