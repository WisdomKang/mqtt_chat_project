extends CanvasLayer

@onready var loading_ui = $CenterContainer/LoadingUIContainer
@onready var modal_ui = $CenterContainer/ModalUIContainer
@onready var modal_content = $CenterContainer/ModalUIContainer/ModalContent

func _ready() -> void:
	hide_loading()
	_hide_modal()

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
	
	await get_tree().create_timer(1.0).timeout
	
	_hide_modal()
	
func _hide_modal() -> void :
	visible = false
	modal_ui.visible = false
	
