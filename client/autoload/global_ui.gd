extends CanvasLayer

@onready var LoadingText = $CenterContainer/PanelContainer/VBoxContainer/LoadingText

func _ready() -> void:
	NetworkManager.request_start.connect(show_loading)
	NetworkManager.request_completed.connect(hide_loading)

func show_loading() -> void :
	visible = true
	
func hide_loading() -> void :
	visible = false

func _process(delta: float) -> void:
	pass
