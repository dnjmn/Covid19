# Covid19

##### request:

curl --location --request POST 'localhost:1323/covid/case/count/nearme' \
--header 'Content-Type: application/json' \
--data-raw '{
    "latitude":12.919226,
    "longitude":77.591523
}'

##### response:
"total positive covid count cases in Karnatakais: 2988333"

##### request:
curl --location --request GET 'localhost:1323/covid/case/update'

##### response:
"updated cases: 37"
