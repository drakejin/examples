# Kafka를 처음 다루는 사람을 위해

kafka의 설치는 `docker-compose.yml` 의 내용을 참고하세요.
- 처음에는 docker-compose.yml의 내용을 전부 지운다음에, docker-compose.yml의 내용을 주석(`========`) block단위로 복사 붙여넣기 한 다음

> docker-compose up --build -d && docker-compose logs --follow

로, 어떤 동작을 하는지 로그의 내용과 expose된 docker container의 결과물을 확인해주세요.


# 각 container 설명

#### 핵심 컨테이너
- zookeeper: kafka Cluster관리를 위한 zookeeper cluster.
- kafka: 해당 예제에서 메인이 되는 kafka

#### 핵심 서드파티 컨테이너

- ksql-server: kafka API를 sql로 컨트롤하기위한 API서버
- kafka-schema-registry: kafka topic에 대한 메세지를 schemable하게 전송하게 돕는녀석
- kafka-connect: kafka의 connect, 서드 파티 컨슈머. 여러가지 도구들을 설치해서 컨슈밍해줄 수 있다.

#### 부수적인 서드파티 컨테이너
- zoonavigator-api: zookeeper를 컨트롤하기위한 http web api
- kafka-rest-proxy: kafka를 http로 컨트롤하기위한 API

#### Dashboard 서드파티 컨테이너
- zoonavigator-web: zoonavigator-api를 컨트롤하기위한 web UI Tool
- kafka-schema-registry-ui: kafka-schema-registry를 위한 Web UI Tool
- kafka-topics-ui: kafka의 topic에 대한 Web UI Tool
- kafka-connect-ui: kafka-connect를 컨트롤 하기위 한 Web UI Tool
- kafka-cmak, kafka-kafdrop: kafka를 제어하기위한 web용 Dashboard

# Dashboard 서드 파티 컨테이너 접속
- zoonavigator-web: http://localhost:8000
- kafka-topics-ui: http://localhost:8001
- kafka-connect-ui: http://localhost:8002
- kafka-schema-registry-ui: http://localhost:8003
- kafka-cmak: http://localhost:8004
- kafka-kafdrop: http://localhost:8005
