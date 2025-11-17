#!/bin/bash

# CPU Load
stress_cpu() {
  dd if=/dev/urandom | bzip2 -9 >> /dev/null
}

# Memory Load
stress_mem() {
  memload=$(awk 'BEGIN{printf "%.0f", 0.8 * 1024 * 1024 * 1024}')  # 80% of 1GB
  head -c $memload < /dev/zero | gzip > /dev/null
}

# Disk I/O Load
stress_io() {
  dd if=/dev/urandom of=/tmp/loadtest.tmp bs=1M count=100 conv=fdatasync
}

wait_enter_key() {
    echo "System load started. Press Enter to stop..."
    while true; do
        read -n 1 -t 1 key
        if [[ $key == $'\x0a' ]]; then
            cleanup
            break
        fi
    done
}

cleanup(){
    killall dd gzip 2>/dev/null
    rm -f /tmp/loadtest.tmp   
}

wait_enter_key() {
  prompt="System load started. Press Q to stop..."

  while true; do
    echo "${prompt}"
    read -r key

    case "$key" in
      [Qq]*) cleanup;;
    esac
  done
}


run_stress(){
    for i in {1..$(nproc)}; do stress_cpu & done
    stress_mem &
    stress_io &
}

run_stress
wait_enter_key
