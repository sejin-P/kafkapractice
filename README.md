# Kafka

## 목표
1. grpc server 두개를 띄워 하나는 producer, 하나는 consumer로써 역할을 하게 한다.
2. data는 별 다른 의미 없이 1초마다 '1'이라는 값만 보낸다.
3. go-kafka를 활용하여 queue역할을 한다. (container를 띄울지 안띄울지는 미지수)

## 목표가 아닌 것
1. 고차원적인 기능
2. kafka의 기능을 다양하게 활용하는 것. 지금은 queue 역할 그자체로만 활용한다.