extends Control

@onready var username_label = $MarginContainer/VBoxContainer/Username
@onready var user_id_label = $MarginContainer/VBoxContainer/UserId

@onready var MQTTClient = $MQTT

# Called when the node enters the scene tree for the first time.
func _ready() -> void:
	print(NetworkManager.current_user)
	username_label.text = NetworkManager.current_user["username"]
	
	MQTTClient.connect_to_broker("tcp://127.0.0.1:8089")


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta: float) -> void:
	pass


func _on_mqtt_broker_connected() -> void:
	MQTTClient.publish("chat/test", "hello!! test message!")
