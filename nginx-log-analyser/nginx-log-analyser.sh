#!/bin/bash

# Set colors for output
RED='\e[1;31m'
GREEN='\e[1;32m'
YELLOW='\e[1;33m'
BLUE='\e[1;34m'
NC='\e[0m' # No Color

print_section() {
    echo -e "\n${BLUE}$1${NC}"
}

get_top_5_IP_addresses() {
  print_section "Top 5 IP addresses with the most requests:"
  COMMAND=$(cat nginx-access.log | awk '{print $1}' | uniq -c | sort -nrk1 | head -5 | awk '{print $2 " - " $1 " requests"}')
  echo -e "$COMMAND\n\n"
}

get_top_5_paths() {
  print_section "Top 5 most requested paths:"
  COMMAND=$(cat nginx-access.log | awk -F'"' '{print $2}' | awk '{print $2}' | sed 's/[[:space:]]*$//'  | sort -n | uniq -c | sort -nrk1 | head -5 | awk '{print $2 " - " $1 " requests"}')
  echo -e "$COMMAND\n\n"
}

get_top_5_status_codes() {
  print_section "Top 5 response status codes:"
  COMMAND=$(cat nginx-access.log| awk -F'"' '{print $3}' | awk '{print $1}' | sed 's/[[:space:]]*$//' | sort -n | uniq -c | sort -nrk1 | head -5 | awk '{print $2 " - " $1 " requests"}')
  echo -e "$COMMAND\n\n"
}


get_top_5_user_agents() {
  print_section "Top 5 user agents:"
  COMMAND=$(cat nginx-access.log| awk -F'"' '{print $6}' | sort | uniq -c | sort -nrk1 | head -5 | awk '{printf "%s - ", $1; for (i=2; i<=NF; i++) printf "%s ", $i; print ""}' | awk -F'-' '{print $2 " - " $1 " requests"}')
  echo -e "$COMMAND\n\n"
}

check_is_root() {
  if [ "$(id -u)" -ne 0 ]; then
    echo -e "${YELLOW}Warning: Some information might require root privileges${NC}"
  fi
}

teardown() {
  echo -e "\n${GREEN}System check completed at $(date)${NC}"
}


check_is_root
get_top_5_IP_addresses
get_top_5_paths
get_top_5_status_codes
get_top_5_user_agents
teardown
