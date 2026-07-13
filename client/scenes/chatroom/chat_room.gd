extends MarginContainer

const CHATROOM_SCENE = preload("res://scenes/chatroom_list/ChatList.tscn")

@onready var title_text = $VBoxContainer/TopLayout/Title
@onready var message_log = $VBoxContainer/MessageLayout
@onready var line_edit = $VBoxContainer/BottomLayout/LineEdit
@onready var back_button = $VBoxContainer/TopLayout/ExitButton

var room_name : String
var room_id : int

var is_loading_history: bool = false
var top_index = -1

func _ready() -> void:
	title_text.text = room_name
	NetworkManager.connect_receive_topic(room_id)
	NetworkManager.mqtt_client.received_message.connect(_on_receive_message)
	await _load_message_logs(top_index)
	
	var v_scrollvar = message_log.get_v_scroll_bar()
	v_scrollvar.value_changed.connect(_on_scroll_changed)
	
	await get_tree().process_frame
	_scroll_to_bottom()

func init_chatroom( new_room_name : String , new_room_id : int) -> void :
	room_name = new_room_name
	room_id = new_room_id
	
func _on_line_edit_text_submitted(new_text: String) -> void:
	if new_text == "" :
		return
		
	line_edit.text = ""
	
	NetworkManager.send_message(room_id , new_text)
	line_edit.grab_focus()
	_scroll_to_bottom()
	
func _decorate_message(user_name : String, user_id : int, message : String) -> String :
	var color = "[color=green]["
	if user_id == NetworkManager.current_user["id"] :
		color = "[color=yellow]["
		
	var deco_text = color + user_name + "][/color]:" + message + "\n"
	return deco_text


func _load_message_logs( start_id : int) -> void :
	is_loading_history = true
	var response = await NetworkManager.get_message_log(room_id, start_id)
	
	if response["result"] :
		var message_list = response["body"]
		if len(message_list) > 0 :
			var old_text = message_log.text
			
			var previous_text : String = ""
			for message in message_list :
				var log_ling = _decorate_message( message["sender"]["username"], message["sender_id"] , message["content"])
				previous_text += log_ling
			
			message_log.text = previous_text + old_text			
			
				
				
			top_index = message_list[0]["id"]
	
	is_loading_history = false
	
	

func _on_receive_message( topic : String, message : String) -> void :
	var message_data = JSON.parse_string(message)
	
	var message_deco = _decorate_message(message_data["sender"]["username"] , message_data["sender_id"] , message_data["content"])
	message_log.append_text(message_deco)
	
	await get_tree().process_frame
	_scroll_to_bottom()


func _on_scroll_changed( current_value : float) -> void :
	if current_value == 0 and not is_loading_history:
		_load_message_logs(top_index)
	
func _scroll_to_bottom() -> void :
	# RichTextLabel 내부에 숨어있는 세로 스크롤바 오브젝트를 찾아옵니다.
	var scroll_bar = message_log.get_v_scroll_bar()
	# 스크롤바의 값을 최대치로 변경하여 맨 아래로 내립니다.
	scroll_bar.value = scroll_bar.max_value


func _on_exit_button_pressed() -> void:
	NetworkManager.disconnect_receive_topic(room_id)
	SceneManager.load_chatroom_list()
