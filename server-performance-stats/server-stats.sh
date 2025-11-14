#!/bin/bash

# Set colors for output
RED='\e[1;31m'
GREEN='\e[1;32m'
YELLOW='\e[1;33m'
BLUE='\e[1;34m'
NC='\e[0m' # No Color

# Function to print section headers
print_section() {
    echo -e "\n${BLUE}===== $1 =====${NC}"
}

get_cpu_usage() {
  print_section "CPU usage"
  CPU_USAGE=$(cat /proc/stat |grep cpu |head -n 1|awk '{print ($5*100)/($2+$3+$4+$5+$6+$7+$8+$9+$10)}'|awk '{print "CPU Usage: " 100-$1"%"}')
  echo -e "$CPU_USAGE\n\n"
}

get_memory_usage() {
  print_section "Memory usage"
  MEM_USAGE=$(cat /proc/meminfo | grep Mem |head -n 3 | awk 'NR==1{total=$2} NR==2{free=$2} END {used=total-free; print "Memory Usage: "used "kB ("(used*100)/total"%)"; print "Total Memory: "total"kB ("(total*100)/total"%)"; print "Free Memory: "free"kB ("(free*100)/total"%)"}')
  echo -e "$MEM_USAGE\n\n"
}

get_disk_usage() {
  print_section "Disk usage"
  DISK_USAGE=$(df -h --total | grep total | awk '{print "Total: " $2 " ("100"%)"; print "Used: " $3 " ("$5")"; print "Free: " $4 " ("100-$5"%)"}')
  echo -e "$DISK_USAGE\n\n"
}

get_top_process_by_cpu_usage() {
  print_section "Top 5 processes by CPU usage"
  PROCESS_CPU_USAGE=$(ps aux --sort=-%cpu | head -n 6)
  echo -e "$PROCESS_CPU_USAGE\n\n"
}

get_top_process_by_mem_usage() {
  print_section "Top 5 processes by memory usage"
  PROCESS_MEM_USAGE=$(ps aux --sort=-%mem | head -n 6)
  echo -e "$PROCESS_MEM_USAGE\n\n"
}

get_os_version() {
  print_section "OS version"
  OS_VERSION=$(cat /etc/os-release | grep PRETTY_NAME | cut -d '=' -f2 | tr -d '"')
  echo -e "$OS_VERSION\n\n"
}

get_logged_users() {
  print_section "Logged in users"
  LOGGED_USERS=$(users)
  echo -e "$LOGGED_USERS\n\n"
}

get_failed_login_attempts() {
  print_section "Failed login attempts"
  FAILED_LOGIN_ATTEMPTS=$(lastb)
  echo -e "$FAILED_LOGIN_ATTEMPTS\n\n"
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
get_cpu_usage
get_memory_usage
get_disk_usage
get_top_process_by_cpu_usage
get_top_process_by_mem_usage
get_os_version
get_logged_users
get_failed_login_attempts
teardown

