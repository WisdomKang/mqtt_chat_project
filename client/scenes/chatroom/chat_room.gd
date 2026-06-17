extends MarginContainer

@onready var title_text = $VBoxContainer/TopLayout/Title
@onready var message_list = $VBoxContainer/MessageLayout
@onready var line_edit = $VBoxContainer/BottomLayout/LineEdit

var room_name : String
var room_id : int

func _ready() -> void:
	title_text.text = room_name

func init_chatroom( new_room_name : String , new_room_id : int) -> void :
	room_name = new_room_name
	room_id = new_room_id
	
func _on_line_edit_text_submitted(new_text: String) -> void:
	print("입력 텍스트 : " , new_text)
	if new_text == "" :
		return
		
	line_edit.text = ""
	message_list.append_text( _decorate_message("테스트" , 0 , new_text) + "\n")
	line_edit.grab_focus()
	_scroll_to_bottom()
	
func _decorate_message(user_name : String, user_id : int, message : String) -> String :
	var deco_text = "[color=red][" + user_name + "][/color]:" + message
	return deco_text


func _scroll_to_bottom() -> void :
	# RichTextLabel 내부에 숨어있는 세로 스크롤바 오브젝트를 찾아옵니다.
	var scroll_bar = message_list.get_v_scroll_bar()
	# 스크롤바의 값을 최대치로 변경하여 맨 아래로 내립니다.
	scroll_bar.value = scroll_bar.max_value
