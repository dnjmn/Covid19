# Covid19

https://dry-hamlet-22926.herokuapp.com/

##### near me: returns total positive cases in the state
curl --location --request POST 'https://dry-hamlet-22926.herokuapp.com/covid/case/count/nearme' \
--header 'Content-Type: application/json' \
--data-raw '{
    "latitude":12.919226,
    "longitude":77.591523
}'

##### update data: updates the current data to our database
curl --location --request GET 'https://dry-hamlet-22926.herokuapp.com/covid/case/update'
