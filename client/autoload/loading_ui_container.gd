extends PanelContainer

@onready var loading_timer = $LoadingTimer
@onready var loading_text = $VBoxContainer/LoadingText

var BASE_TEXT = "처리중"
var dot_count = 0
# Called when the node enters the scene tree for the first time.
func _ready() -> void:
	loading_timer.timeout.connect(_on_timer_timeout)

func _on_timer_timeout() : 
	dot_count = (dot_count + 1) % 4 
	loading_text.text = BASE_TEXT + ".".repeat(dot_count)
	
#func reset 
