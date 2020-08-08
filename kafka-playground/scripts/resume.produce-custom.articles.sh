# https://docs.confluent.io/current/connect/references/restapi.html
docker-compose exec kafka-connect-datagen \
  curl -X PUT -H "Content-Type: application/json" \
  http://localhost:8083/connectors/datagen-articles/resume


docker-compose exec kafka-connect-datagen \
  curl -X GET -H "Content-Type: application/json" \
  http://localhost:8083/connectors/datagen-articles/status

