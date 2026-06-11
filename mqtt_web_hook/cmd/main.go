package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var targetURL string
var brokerURL string

func main() {
	// 💡 팁: 환경변수가 없으면 로컬 테스트용(localhost)을 기본값으로 사용
	brokerURL = getEnv("MQTT_BROKER_URL", "tcp://localhost:1883")
	targetURL = getEnv("WEBHOOK_TARGET_URL", "http://localhost:8080/api/v1/message")

	log.Printf("[Bridge] 브로커 연결 시도 중: %s", brokerURL)
	log.Printf("[Bridge] 웹훅 타겟 주소 설정: %s", targetURL)

	// MQTT 클라이언트 옵션 설정
	opts := mqtt.NewClientOptions()
	opts.AddBroker(brokerURL)
	opts.SetClientID("webhook_bridge_client")
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(time.Second * 2)

	// 브로커와 연결이 끊겼을 때 로그
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Printf("[⚠️ Bridge] MQTT 브로커와 연결이 끊어졌습니다: %v", err)
	}

	// 브로커와 연결되었을 때 실행할 핸들러
	opts.OnConnect = func(client mqtt.Client) {
		log.Println("[✨ Bridge] MQTT 브로커 연결 성공!")

		// chat/# 토픽 구독 시작 (QoS 1로 안정적으로 수신)
		topic := "chat/#"
		token := client.Subscribe(topic, 1, func(c mqtt.Client, msg mqtt.Message) {
			// 💡 중요: HTTP 요청이 블로킹되어 전체 수신이 멈추는 것을 방지하기 위해 고루틴으로 분기합니다.
			go messagePost(msg.Topic(), msg.Payload())
		})

		if token.Wait() && token.Error() != nil {
			log.Fatalf("[❌ Bridge] 토픽 구독 실패: %v", token.Error())
		}
		log.Printf("[📡 Bridge] 토픽 구독 중... (%s)", topic)
	}

	// 클라이언트 생성 및 연결 실행
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("[❌ Bridge] 초기 연결 실패: %v", token.Error())
	}

	// 프로그램이 바로 종료되지 않고 시그널을 기다리도록 대기 (ctrl+c 시 안전하게 종료)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Println("[Bridge] 브릿지 프로그램을 안전하게 종료합니다.")
	client.Disconnect(250)
}

func messagePost(topic string, payload []byte) {
	log.Printf("[📥 Received] Topic: %s | Payload: %s", topic, string(payload))

	// 메인 HTTP 서버로 웹훅 발송
	resp, err := http.Post(targetURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("[❌ Webhook Error] HTTP 서버로 전송 실패: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Printf("[🚀 Webhook Sent] HTTP 응답 상태 코드: %d", resp.StatusCode)
}

// 환경변수 읽기용 헬퍼 함수
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
