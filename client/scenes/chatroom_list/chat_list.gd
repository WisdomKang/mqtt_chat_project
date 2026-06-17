extends Control

const CHATROOM_BUTTON =  preload("res://scenes/chatroom_list/ChatRoomButton.tscn")
const CHATROOM_SCENE = preload("res://scenes/chatroom/ChatRoom.tscn")
@onready var list_box = $MarginContainer/VBoxContainer/ScrollContainer/ListVboxContainer
var chatroom_list : Array = []
# Called every frame. 'delta' is the elapsed time since the previous frame.
func _ready() -> void:
	_load_chatroom_list()


func _load_chatroom_list() :
	var response = await NetworkManager.get_chatroom_list()
	if response[1] == 200 :
		chatroom_list = JSON.parse_string(response[3].get_string_from_utf8())
		
		for chatroom in chatroom_list :
			var button = CHATROOM_BUTTON.instantiate()
			button.text = chatroom["room_name"]
			button.pressed.connect(_on_room_button_pressed.bind(chatroom["room_name"] , chatroom["room_id"]))
			list_box.add_child(button)
			

func _on_room_button_pressed(room_name : String, room_id : int) -> void :
	var chatroom_instance = CHATROOM_SCENE.instantiate()
	chatroom_instance.init_chatroom(room_name, room_id)
	get_tree().change_scene_to_node(chatroom_instance)
	
