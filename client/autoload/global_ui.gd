extends CanvasLayer

@onready var loading_ui = $CenterContainer/LoadingUIContainer
@onready var modal_ui = $CenterContainer/ModalUIContainer
@onready var modal_content = $CenterContainer/ModalUIContainer/ModalContent

func _ready() -> void:
	NetworkManager.request_start.connect(show_loading)
	NetworkManager.request_completed.connect(hide_loading)
	
	hide_loading()
	hide_modal()

func show_loading() -> void :
	visible = true
	loading_ui.visible = true
	loading_ui.process_mode = Node.PROCESS_MODE_INHERIT
	
func hide_loading() -> void :
	visible = false
	loading_ui.visible = true
	loading_ui.process_mode = Node.PROCESS_MODE_DISABLED
	
func show_modal(content : String) -> void :
	visible = true
	modal_ui.visible = true
	modal_content.text = content
	
func hide_modal() -> void :
	visible = false
	modal_ui.visible = false
	
