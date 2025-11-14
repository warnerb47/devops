#!/bin/bash

# Set colors for output
RED='\e[1;31m'
GREEN='\e[1;32m'
YELLOW='\e[1;33m'
BLUE='\e[1;34m'
NC='\e[0m' # No Color

LOG_DIRECTORY="."


create_archive_directory() {
    if [ ! -d "archives" ]; then
        mkdir archives
    fi
}

compress_logs() {
    tar cvzf $(echo "archives/logs_archive_$(date "+%Y%m%d_%H%M%S").tar.gz") *
}

move_to_directory_log() {
    cd "$1" || exit
}

set_log_directory() {
  if [ "$#" -eq 0 ]; then
    echo -e "${YELLOW}Warning: No log directory specified we will use current location${NC}"
  else
    LOG_DIRECTORY="$1"
  fi
}

teardown() {
  echo -e "\n${GREEN}Log compressed successfully at $(date)${NC}"
}

set_log_directory $1
move_to_directory_log $LOG_DIRECTORY
create_archive_directory
compress_logs
teardown

