extends Control

const CHATROOM_BUTTON =  preload("res://scenes/chatroom_list/ChatRoomButton.tscn")
const CHATROOM_SCENE = preload("res://scenes/chatroom/ChatRoom.tscn")
var chatroom_list : Array = []
@onready var list_box = $MarginContainer/VBoxContainer/ScrollContainer/ListVboxContainer
@onready var line_edit = $MarginContainer/VBoxContainer/VBoxContainer/LineEdit

func _ready() -> void:
	_load_chatroom_list()


func _load_chatroom_list() :
	var response = await NetworkManager.get_chatroom_list()
	if response["result"] :
		for child in list_box.get_children() :
			child.queue_free()
		
		chatroom_list = response["body"]
		
		for chatroom in chatroom_list :
			var button = CHATROOM_BUTTON.instantiate()
			button.text = chatroom["room_name"]
			button.pressed.connect(_on_room_button_pressed.bind(chatroom["room_name"] , chatroom["id"]))
			list_box.add_child(button)
			

func _on_room_button_pressed(room_name : String, room_id : int) -> void :
	var chatroom_instance = CHATROOM_SCENE.instantiate()
	chatroom_instance.init_chatroom(room_name, room_id)
	get_tree().change_scene_to_node(chatroom_instance)
	


func _on_button_pressed() -> void:
	_load_chatroom_list()


func _on_room_create_button_pressed() -> void:
	if line_edit.text != "":
		await NetworkManager.create_chatroom(line_edit.text)
		line_edit.text = "";
		await _load_chatroom_list()
		
