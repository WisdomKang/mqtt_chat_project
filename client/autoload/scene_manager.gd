extends Node

const CHATROOM_SCENE = "res://scenes/chatroom/ChatRoom.tscn"
const CHATROOM_LIST_SCENE = "res://scenes/chatroom_list/ChatList.tscn"
const LOGIN_SCENE = "res://scenes/login/Login.tscn"


func load_login() -> void :
	get_tree().change_scene_to_file(LOGIN_SCENE)

func load_chatroom_list() -> void :
	get_tree().change_scene_to_file(CHATROOM_LIST_SCENE)

func load_chatroom( new_room_name : String , new_room_id : int) -> void :
	var scene_resource = load(CHATROOM_SCENE)
	var chatroom_scene_instance = scene_resource.instantiate()
	
	chatroom_scene_instance.init_chatroom( new_room_name, new_room_id)
	
	get_tree().change_scene_to_node(chatroom_scene_instance)
	
func leave_chatroom() -> void :
	get_tree().change_scene_to_file(CHATROOM_LIST_SCENE)
