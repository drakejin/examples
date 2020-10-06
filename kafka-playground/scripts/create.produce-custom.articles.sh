docker-compose exec kafka-connect-datagen \
  curl -X POST -H "Content-Type: application/json" \
  --data @/usr/share/confluent-hub-components/confluentinc-kafka-connect-datagen/etc/custom/articles.config.json \
  http://localhost:8083/connectors

cat /usr/share/confluent-hub-components/confluentinc-kafka-connect-datagen/etc/custom/articles.config.json


docker-compose exec kafka-connect-datagen \
  kafka-console-consumer  \
  --topic articles \
  --bootstrap-server kafka:9092 \
  --property print.key=true \
  --max-messages 10 \
  --from-beginning
