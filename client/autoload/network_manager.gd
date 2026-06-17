extends Node

signal request_start
signal request_completed
signal signin_completed(userinfo : Dictionary)
signal signin_failed

signal get_chatroom_completed(chatroom_list : Array)
signal get_chatroom_failed

signal create_chatroom_completed(chatroom_info : Dictionary)
signal create_chatroom_failed

signal get_chatlog_completed
signal get_chatlog_failed

var server_url = "127.0.0.1"
var http_server_port = "8080"
var mqtt_broker_port = "1883"
var http_server_url : String = "http://" + server_url + ":" + http_server_port 
var mqtt_broker_url : String = "tcp://" + server_url + ":" + mqtt_broker_port

@onready var httpRequest : HTTPRequest = $HTTPRequest
@onready var mqtt_client : MQTTClient = $MQTT

var headers = ["Content-Type: application/json"] 


var current_user = {
	"username" : null,
	"user_id" : null,
}

var mqtt_user = {
	"username" : null,
	"password" : null,
}

func signin(username : String) -> void :
	var signin_url= "http://127.0.0.1:8080/auth/signin"
	var request_body = { "username" : username }
	
	request_start.emit()
	httpRequest.request(signin_url, headers, HTTPClient.METHOD_POST, JSON.stringify(request_body))
		
	var response = await httpRequest.request_completed
	request_completed.emit()
	_on_signin_request_completed(response)
	
func signout() -> void :
	current_user = null
	
func _on_signin_request_completed(response : Array ) :
	print( response )
	if response[1] == 200:
		var json = JSON.parse_string(response[3].get_string_from_utf8())
		set_user(json["user"])
		get_tree().change_scene_to_file("res://scenes/chatroom_list/ChatList.tscn")
	else:
		print("서버 에러 발생! 상태 코드: ", response[1])
		

func get_chatroom_list() -> Array : 
	var request_url = "http://127.0.0.1:8080/api/v1/rooms"
	
	request_start.emit()
	
	httpRequest.request(request_url , headers, HTTPClient.METHOD_GET)
	var response_body = await httpRequest.request_completed
	
	request_completed.emit()
	
	return response_body
	

	
func create_chatroom(chatroom_name : String) -> Array :
	var request_url = "http://127.0.0.1:8080/api/v1/rooms"
	var request_body = { "room_name" : chatroom_name } 
	
	request_start.emit()
	httpRequest.request(request_url , headers, HTTPClient.METHOD_POST, JSON.stringify(request_body))
	var response_body = await httpRequest.request_completed
	request_completed.emit()
	
	return response_body
	
func _on_get_chatroom_list_completed(response : Array) -> void  :
	pass
	
	
func _on_create_chatroom(response : Array) -> void :
	pass

func set_user( new_user : Dictionary) -> void :
	current_user = {
		"username" : new_user["username"],
		"user_id" : int(new_user["user_id"])
	}

func reset_user() -> void :
	current_user = {
		"username" : null ,
		"user_id" : null,
	}
