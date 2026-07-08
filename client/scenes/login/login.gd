extends Control

@onready var name_input: LineEdit = $MarginContainer/VBoxContainer/NameInput
@onready var enter_button: Button = $MarginContainer/VBoxContainer/SignInButton

func _ready() -> void:
	# 버튼 클릭 이벤트 연결
	enter_button.pressed.connect(_on_enter_button_pressed)

func _on_enter_button_pressed() -> void:
	var user_name: String = name_input.text.strip_edges()
	
	if user_name.is_empty():
		GlobalUI.show_modal("이름을 입력해주세요.")
		return
		
	GlobalUI.show_loading()
	
	var result = await NetworkManager.signin(user_name)
	
	print( result )
	if result["result"] :
		NetworkManager.set_user( result["body"]["user"] )
		GlobalUI.hide_loading()
		SceneManager.load_chatroom_list()
	else :
		GlobalUI.show_modal("로그인에 실패했습니다.")
	
	
	
	
	
