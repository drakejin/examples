# protobuf-on-http-practice
http 메세지위에서 protobuf 메세지 주고받고 있는데, protobuf 메세지의 프로퍼티를 추가해서 다시 서버프로그램을 재 시작하면 클라이언트에서, 언마샬링(deserialize)할 때에 올바르게 되지 않는 문제가 있다고 함, 때문에 문제재현을 위해 해당 examples를 만들었다.

### 작업상황
1. back 상황구현
2. front 상황구현
  - response 값을 deserialize를 못해봄 ㅠㅠ
3. back에서 Request D 메세지를 하나 더 추가해서 배포.
4. front에서 정상적인 desialize를 못하면 에러재현 성공

## back

#### 서버 실행하기
``` bash
cd back
go mod download
go run main.go
```

#### 서버 프로토파일 수정하기

``` bash
cd back
yarn install
./scripts/genproto.sh
```

## front

#### 프론트 실행하기
```
cd front
yarn install
yarn start
```

#### 서버에서 만든 프로토파일 가져오기

```
cd front
cp -rf ../back/proto src
```