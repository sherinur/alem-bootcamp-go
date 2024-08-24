awk -F, '{OFS=","; print $1, $4, $3, $2, $5}' movies.csv | tr ',' ' '
