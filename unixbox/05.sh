least_common_line=$(sort poem.txt | uniq -c | sort -n | head -n 1 | cut -c9-)
echo "$least_common_line"
