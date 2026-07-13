extends Node

var server_url = "127.0.0.1"
var http_server_port = "8080"
var mqtt_broker_port = "8089"
var http_server_url : String = "http://" + server_url + ":" + http_server_port 
var mqtt_broker_url : String = "tcp://" + server_url + ":" + mqtt_broker_port

const SIGNUP_PATH = "/auth/signup"
const SIGNIN_PATH = "/auth/signin"

const ROOMS_PATH = "/api/v1/rooms"

@onready var httpRequest : HTTPRequest = $HTTPRequest
@onready var mqtt_client : MQTTClient = $MQTT

var headers = ["Content-Type: application/json"] 


var current_user = null

var mqtt_user = {
	"username" : null,
	"password" : null,
}

func set_server_url(new_ip : String) -> void : 
	server_url = new_ip
	
func _get_api_path(path:String) -> String :
	return http_server_url + path
	
func parse_result(response : Array) -> Dictionary :
	var response_code = response[1]
	var res_headers = response[2]
	var body = JSON.parse_string(response[3].get_string_from_utf8())  
	
	var result_dict = {}
	if response_code == 200 :
		result_dict = {
			"result" : true,
			"body" : body
		}
	else :
		result_dict = {
			"result" : false,
			"body" : body
		}
		print(response)
	
	return result_dict
		

func signin(username : String) -> Dictionary :
	var signin_url= _get_api_path(SIGNIN_PATH)
	var request_body = { "username" : username }
	
	httpRequest.request(signin_url, headers, HTTPClient.METHOD_POST, JSON.stringify(request_body))
		
	var response = await httpRequest.request_completed

	return parse_result(response)
	
	
	
func signout() -> void :
	current_user = null

# MQTT Method
func connect_mqtt_broker() -> void :
	mqtt_client.connect_to_broker(mqtt_broker_url)
	mqtt_client.broker_disconnected.connect(reconnect_broker)

func reconnect_broker() -> void :
	print("broker disconnect. try to reconnect...")
	await get_tree().create_timer(0.5).timeout
	mqtt_client.connect_to_broker(mqtt_broker_url)

func connect_receive_topic(room_id : int ) -> void :
	var sub_chat_topic = "chat/" + str(room_id) + "/receive" 
	print(sub_chat_topic)
	mqtt_client.subscribe(sub_chat_topic, 1)

func disconnect_receive_topic(room_id : int ) -> void :
	var sub_chat_topic = "chat/" + str(room_id) + "/receive" 
	mqtt_client.unsubscribe(sub_chat_topic)
	
func send_message(room_id : int , content : String ) -> void :
	var send_chat_topic = "chat/" + str(room_id) + "/send"
	var message_data = {
		"sender_id" : current_user["id"],
		"room_id" : room_id,
		"sender" : current_user,
		"content" : content
	}
	
	var test = JSON.stringify(message_data)
	print(test)
	mqtt_client.publish( send_chat_topic , test )
	 

func get_chatroom_list() -> Dictionary : 
	var request_url = _get_api_path(ROOMS_PATH)
	
	httpRequest.request(request_url , headers, HTTPClient.METHOD_GET)
	var response = await httpRequest.request_completed
	
	return parse_result(response)
	
func create_chatroom(chatroom_name : String) -> Dictionary :
	var request_url = _get_api_path(ROOMS_PATH)
	var request_body = { "room_name" : chatroom_name } 
	
	httpRequest.request(request_url , headers, HTTPClient.METHOD_POST, JSON.stringify(request_body))
	var response = await httpRequest.request_completed
	
	return parse_result(response)

func get_message_log(room_id : int, start_id : int) -> Dictionary :
	var request_url = http_server_url + "/api/v1/rooms/" + str(room_id) +"/messages"
	
	request_url += "?" + "start_id=" + str(start_id)
	
	httpRequest.request(request_url, headers, HTTPClient.METHOD_GET, "")
	var response = await httpRequest.request_completed
	return parse_result(response)

func set_user( new_user : Dictionary) -> void :
	current_user = {
		"username" : new_user["username"],
		"id" : int(new_user["id"])
	}

func reset_user() -> void :
	current_user["username"] =  null
	current_user["id"] = null
	
