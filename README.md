# MQTT 기반 채팅 서버

## 목적
MQTT 기반으로 채팅서버를 구현하여 현재 알고 있는 기술들을 응용하여 프로젝트를 구성하는 목적

## 사용 기술
- ### Server
    - 관리 서버 - Golang : 병렬 작업적합으로 관리서버 구성
    - MQTT Broker - Mosquitto : 저전력 경량 MQTT Broker로 구독형태로 빠르고 가볍게 기능할 브로커 선정
    - RDB - PostgreSQL : 영구 저장 RDB 자체 기능으로 메세지 내용으로 검색 기능까지 구현에 적합
    - 캐싱 DB - Redis : 캐싱으로 빠르게 쓰고 읽기위한 On Memory DB

- 배포 계획
    - Docker-compose로 MQTT 브로커 서버, RDB서버, 관리 서버 구성
    
---

- ### Client 
    - Application - Godot : 게임 엔진 기반으로 여러 OS포팅에 이점이 있는 엔진으로 MQTT Client 라이브러리도 확인 간단하게 동작 할 UI만 구현

