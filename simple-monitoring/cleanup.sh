#!/bin/sh

# Set colors for output
RED='\e[1;31m'
GREEN='\e[1;32m'
YELLOW='\e[1;33m'
BLUE='\e[1;34m'
NC='\e[0m' # No Color

check_is_root() {
  if [ "$(id -u)" -ne 0 ]; then
    echo -e "${YELLOW}Warning: cleanup netdata might require root privileges${NC}"
  fi
}

uninstall_netdata() {
    wget -O /tmp/netdata-kickstart.sh https://get.netdata.cloud/kickstart.sh && sh /tmp/netdata-kickstart.sh --uninstall
}

check_is_root
uninstall_netdata

