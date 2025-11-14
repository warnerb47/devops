# Top 5 IP addresses with the most requests:
cat nginx-access.log | awk '{print $1}' | uniq -c | sort -nrk1 | head -5 | awk '{print $2 " - " $1 " requests"}'

# Top 5 most requested paths:
cat nginx-access.log | awk -F'"' '{print $2}' | awk '{print $2}' | sed 's/[[:space:]]*$//'  | sort -n | uniq -c | sort -nrk1 | head -5 | awk '{print $2 " - " $1 " requests"}'

# Top 5 response status codes:
cat nginx-access.log| awk -F'"' '{print $3}' | awk '{print $1}' | sed 's/[[:space:]]*$//' | sort -n | uniq -c | sort -nrk1 | head -5 | awk '{print $2 " - " $1 " requests"}'

# Top 5 user agents:
cat nginx-access.log| awk -F'"' '{print $6}' | sort | uniq -c | sort -nrk1 | head -5 | awk '{printf "%s - ", $1; for (i=2; i<=NF; i++) printf "%s ", $i; print ""}' | awk -F'-' '{print $2 " - " $1 " requests"}'

