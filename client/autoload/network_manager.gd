extends Node

var server_url = "127.0.0.1"
var http_server_port = "8080"
var mqtt_broker_port = "1883"
var http_server_url : String = "http://" + server_url + ":" + http_server_port 
var mqtt_broker_url : String = "tcp://" + server_url + ":" + mqtt_broker_port

const SIGNUP_PATH = "/auth/signup"
const SIGNIN_PATH = "/auth/signin"

const GET_ROOMS_PATH = "/api/v1/rooms"

@onready var httpRequest : HTTPRequest = $HTTPRequest
@onready var mqtt_client : MQTTClient = $MQTT

var headers = ["Content-Type: application/json"] 


var current_user = null

var mqtt_user = {
	"username" : null,
	"password" : null,
}

func set_server_url( new_ip : String) -> void : 
	server_url = new_ip
	
func _get_api_path( path  : String) -> String :
	return "http://" + server_url + ":" + http_server_port + path

func signin(username : String) -> Dictionary :
	var signin_url= _get_api_path("/auth/signin")
	var request_body = { "username" : username }
	
	httpRequest.request(signin_url, headers, HTTPClient.METHOD_POST, JSON.stringify(request_body))
		
	var response = await httpRequest.request_completed
	
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
	
	
	
func signout() -> void :
	current_user = null
	
func connect_mqtt_broker() -> void :
	mqtt_client.connect_to_broker(mqtt_broker_url)
	
	
	

func get_chatroom_list() -> Array : 
	var request_url = "http://127.0.0.1:8080" + GET_ROOMS_PATH
	
	
	httpRequest.request(request_url , headers, HTTPClient.METHOD_GET)
	var response_body = await httpRequest.request_completed
	
	return response_body
	
func create_chatroom(chatroom_name : String) -> Array :
	var request_url = "http://127.0.0.1:8080/api/v1/rooms"
	var request_body = { "room_name" : chatroom_name } 
	
	httpRequest.request(request_url , headers, HTTPClient.METHOD_POST, JSON.stringify(request_body))
	var response_body = await httpRequest.request_completed
	
	return response_body

func get_message_log(room_id : int, start_id : int) -> Array :
	var request_url = http_server_url + "/api/v1/rooms/" + str(room_id) +"/messages"
	
	request_url += "?" + "start_id=" + str(start_id)
	
	print(request_url)
	httpRequest.request(request_url, headers, HTTPClient.METHOD_GET, "")
	var response_body = await httpRequest.request_completed
	return response_body

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
	
