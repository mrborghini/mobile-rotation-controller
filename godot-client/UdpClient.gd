extends Node3D


var udp_client: PacketPeerUDP

# Configure the server's IP and port.
const SERVER_IP = "127.0.0.1" # Replace with the actual server IP
const SERVER_PORT = 8080 # Replace with the actual server port


@export var scale_down: float = 1
@export var objects_to_rotate: Node3D
var timer: Timer

func _on_Timer_timeout() -> void:
	keep_alive()

# Initialize UDP client and connect
func _ready() -> void:
	udp_client = PacketPeerUDP.new()
	udp_client.set_dest_address(SERVER_IP, SERVER_PORT)

	# Example: Send a message to the server on start
	keep_alive()
	timer = $Timer
	timer.connect("timeout", _on_Timer_timeout)
	timer.wait_time = 8
	timer.start()

	# Start polling for incoming messages
	set_process(true)

func keep_alive() -> void:
	send_message("godot")

# Send a message to the server
func send_message(message: String) -> void:
	if udp_client:
		var byte_array: PackedByteArray = PackedByteArray()
		byte_array.append_array(message.to_utf8_buffer())
		udp_client.put_packet(byte_array)
		print("Message sent:", message)

# Process incoming messages
func _process(_delta: float) -> void:
	if udp_client and udp_client.get_available_packet_count() > 0:
		var packet: String = udp_client.get_packet().get_string_from_utf8()
		# print("Received from server:", packet)

		var parsed_data: PackedStringArray = packet.split(",")
		
		var x: float = float(parsed_data[0]) * scale_down
		var y: float = float(parsed_data[1]) * scale_down
		var z: float = float(parsed_data[2]) * scale_down
		var w: float = float(parsed_data[3]) * scale_down

		var rotation_vector: Quaternion = Quaternion(-w, -x, z, y)

		objects_to_rotate.quaternion = rotation_vector
