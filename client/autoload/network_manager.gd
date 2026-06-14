extends Node



var current_user = {
	"username" : null,
	"user_id" : null,
}
# Called when the node enters the scene tree for the first time.
func _ready() -> void:
	pass # Replace with function body.


# Called every frame. 'delta' is the elapsed time since the previous frame.
func _process(delta: float) -> void:
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
