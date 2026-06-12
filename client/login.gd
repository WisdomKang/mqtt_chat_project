extends Control

@onready var name_input: LineEdit = $MarginContainer/VBoxContainer/NameInput
@onready var enter_button: Button = $MarginContainer/VBoxContainer/SignInButton
@onready var http_request : HTTPRequest = $HTTPRequest

func _ready() -> void:
	# 버튼 클릭 이벤트 연결
	enter_button.pressed.connect(_on_enter_button_pressed)
	http_request.request_completed.connect(_on_request_completed)

func _on_enter_button_pressed() -> void:
	var user_name: String = name_input.text.strip_edges()
	
	if user_name.is_empty():
		print("이름을 입력해 주세요!")
		return
		
	print("입장 시도 유저 이름: ", user_name)
	
	var headers = ["Content-Type: application/json"]
	var url= "http://127.0.0.1:8080/api/v1/rooms"
	var error = http_request.request(url,headers, HTTPClient.METHOD_GET, "")
	
	if error != OK:
		print("HTTP 요청 실패")
	
	
func _on_request_completed(result:int , response_code:int, headers : PackedStringArray, body: PackedByteArray) -> void :
	if response_code == 200:
		var json = JSON.parse_string(body.get_string_from_utf8())
		print("서버 응답 성공! 방 정보: ", json)
		# TODO: 여기서 다음 화면(채팅방)으로 넘어가거나 MQTT 연결 시작!
	else:
		print("서버 에러 발생! 상태 코드: ", response_code)
