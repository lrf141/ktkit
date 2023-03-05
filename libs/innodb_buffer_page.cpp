//
// Created by lrf141 on 23/03/05.
//

#include <iostream>
#include "include/innodb_buffer_page.h"

namespace idb_buffer_page {
    void InnoDBBufferPage::convert(mysqlx::Value& value, std::string colName) {
        if (colName == "POOL_ID") {
            this->POOL_ID = value.operator uint64_t();
        }
        if (colName == "BLOCK_ID") {
            this->BLOCK_ID = value.operator uint64_t();
        }
        if (colName == "SPACE") {
            this->SPACE = value.operator uint64_t();
        }
        if (colName == "PAGE_NUMBER") {
            this->PAGE_NUMBER = value.operator uint64_t();
        }
        if (colName == "PAGE_TYPE") {
            this->PAGE_TYPE = value.operator std::string();
        }
        if (colName == "FLUSH_TYPE") {
            this->FLUSH_TYPE = value.operator uint64_t();
        }
        if (colName == "FIX_COUNT") {
            this->FIX_COUNT = value.operator uint64_t();
        }
        if (colName == "IS_HASHED") {
            this->IS_HASHED = value.operator std::string();
        }
        if (colName == "NEWEST_MODIFICATION") {
            this->NEWEST_MODIFICATION = value.operator uint64_t();
        }
        if (colName == "OLDEST_MODIFICATION") {
            this->OLDEST_MODIFICATION = value.operator uint64_t();
        }
        if (colName == "ACCESS_TIME") {
            this->ACCESS_TIME = value.operator uint64_t();
        }
        if (colName == "TABLE_NAME") {
            //this->TABLE_NAME = value.operator mysqlx::string();
        }
        if (colName == "INDEX_NAME") {
            //this->INDEX_NAME = value.operator mysqlx::string();
        }
        if (colName == "NUMBER_RECORDS") {
            this->NUMBER_RECORDS = value.operator uint64_t();
        }
        if (colName == "DATA_SIZE") {
            this->DATA_SIZE = value.operator uint64_t();
        }
        if (colName == "COMPRESSED_SIZE") {
            this->COMPRESSED_SIZE = value.operator uint64_t();
        }
        if (colName == "PAGE_STATE") {
            this->PAGE_STATE = value.operator mysqlx::string();
        }
        if (colName == "IO_FIX") {
            this->IO_FIX = value.operator mysqlx::string();
        }
        if (colName == "IS_OLD") {
            this->IS_OLD = value.operator mysqlx::string();
        }
        if (colName == "FREE_PAGE_CLOCK") {
            this->FREE_PAGE_CLOCK = value.operator uint64_t();
        }
        if (colName == "IS_STALE") {
            this->IS_STALE = value.operator std::string();
        }
    }

    unsigned int InnoDBBufferPage::getDataSize() {
        return this->DATA_SIZE;
    }

    unsigned int InnoDBBufferPage::getNumberRecords() {
        return this->NUMBER_RECORDS;
    }

    std::string InnoDBBufferPage::getPageType() {
        return this->PAGE_TYPE;
    }
}
