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
	
	
	
	
