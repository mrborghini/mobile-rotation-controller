extends RichTextLabel

var time_passed: float = 0.0

func _process(delta: float) -> void:
	time_passed += delta
	if time_passed >= 1.0:
		# Reset the timer
		time_passed = 0.0
		update_fps(delta)

func update_fps(delta: float) -> void:
	var fps: float = 1 / delta
	self.text = "Fps: %d" % [fps]
