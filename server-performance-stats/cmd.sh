# CPU usage
cat /proc/stat |grep cpu |head -n 1|awk '{print ($5*100)/($2+$3+$4+$5+$6+$7+$8+$9+$10)}'|awk '{print "CPU Usage: " 100-$1}'

# Memory usage (free, used, total)
cat /proc/meminfo | grep Mem |head -n 3 | awk 'NR==1{total=$2} NR==2{free=$2} END {used=total-free; print "Memory Usage: "used "kB ("(used*100)/total" %)"; print "Total Memory: "total"kB ("(total*100)/total" %)"; print "Free Memory: "free"kB ("(free*100)/total" %)"}'

# Disk usage (free, used, total)
df -h --total | grep total | awk '{print "Total: " $2 " ("100"%)"; print "Used: " $3 " ("$5")"; print "Free: " $4 " ("100-$5"%)"}'

# Top 5 processes by CPU usage
ps aux --sort=-%cpu | head -n 6

# Top 5 processes by memory usage
ps aux --sort=-%mem | head -n 6

# OS version
cat /etc/os-release | grep PRETTY_NAME | cut -d '=' -f2 | tr -d '"'

# Logged in users
users

# Failed login attempts 
lastb

