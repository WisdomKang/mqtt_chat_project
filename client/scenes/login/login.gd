extends Control

@onready var name_input: LineEdit = $MarginContainer/VBoxContainer/NameInput
@onready var enter_button: Button = $MarginContainer/VBoxContainer/SignInButton

func _ready() -> void:
	# 버튼 클릭 이벤트 연결
	enter_button.pressed.connect(_on_enter_button_pressed)

func _on_enter_button_pressed() -> void:
	var user_name: String = name_input.text.strip_edges()
	
	if user_name.is_empty():
		print("이름을 입력해 주세요!")
		return
		
	print("입장 시도 유저 이름: ", user_name)
	
	NetworkManager.signin(user_name)
	
func _on_signin_request_completed(response : Array ) :
	print( response )
	if response[1] == 200:
		var json = JSON.parse_string(response[3].get_string_from_utf8())
		NetworkManager.set_user(json["user"])
		SceneManager.load_chatroom_list()
	else:
		print("서버 에러 발생! 상태 코드: ", response[1])
		GlobalUI.show_modal("로그인에 실패했습니다.")
	
	
	
