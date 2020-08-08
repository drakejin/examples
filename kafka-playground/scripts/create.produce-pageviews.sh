## pageviews auto-gen 을 만듭니다.
docker-compose exec kafka-connect-datagen \
  curl -X POST -H "Content-Type: application/json" \
  --data @/usr/share/confluent-hub-components/confluentinc-kafka-connect-datagen/etc/connector_pageviews.config \
  http://localhost:8083/connectors

cat /usr/share/confluent-hub-components/confluentinc-kafka-connect-datagen/etc/connector_pageviews.config


## pageviews 데이터 잘 들어가고 있나요?
docker-compose exec kafka-connect-datagen \
  kafka-console-consumer  \
  --topic pageviews \
  --bootstrap-server kafka:9092 \
  --property print.key=true \
  --max-messages 10 \
  --from-beginning
