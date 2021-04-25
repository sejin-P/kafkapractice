# Kafka

## 목표
1. grpc server 두개를 띄워 하나는 producer, 하나는 consumer로써 역할을 하게 한다.
2. data는 별 다른 의미 없이 1초마다 '1'이라는 값만 보낸다.
3. go-kafka를 활용하여 queue역할을 한다. (container를 띄울지 안띄울지는 미지수)

## 목표가 아닌 것
1. 고차원적인 기능
2. kafka의 기능을 다양하게 활용하는 것. 지금은 queue 역할 그자체로만 활용한다.

## 참고한 글
1. https://www.popit.kr/golang%EC%97%90%EC%84%9C-%EC%B9%B4%ED%94%84%EC%B9%B4-%EC%BB%A8%EC%8A%88%EB%A8%B8-%EA%B7%B8%EB%A3%B9%EA%B3%BC-%EC%9E%AC%EC%8B%9C%EB%8F%84%EB%A1%9C-%EA%B2%B0%EA%B3%BC%EC%A0%81-%EC%9D%BC%EA%B4%80/

## protoc command
protoc --go_out=. --go_opt=paths=source_relative \                                                                                                                 ✔
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
protos/v1/empty/number.proto

