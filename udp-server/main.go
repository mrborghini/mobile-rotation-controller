package main

import (
	"encoding/json"
	"fmt"
	"mobile-controller-udp-server/types"
	"math"
	"net"
	"os"
	"time"
)

func estimateW(x float64, y float64, z float64) float64 {
	return float64(math.Sqrt((1.0 - (x*x + y*y + z*z))))
}

func main() {
	// Resolve the string address to a UDP address
	hostUrl := "0.0.0.0:8080"
	udpAddr, err := net.ResolveUDPAddr("udp", "0.0.0.0:8080")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Start listening for UDP packages on the given address
	conn, err := net.ListenUDP("udp", udpAddr)

	fmt.Printf("Listening on %s\n", hostUrl)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var godot *net.UDPConn
	var godotAddr *net.UDPAddr

	start := time.Now() // Capture the current time as the start time

	var rotationVector *types.RotationVector3

	// Read from UDP listener in endless loop
	for {
		end := time.Now()                 // Capture the current time again as the end time
		delta := end.Sub(start).Seconds() // Calculate the real-time delta

		var buf [512]byte
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			fmt.Println(err)
			return
		}

		message := string(buf[:n])

		if message == "godot" {
			fmt.Println("Godot client connected")
			godot = conn
			godotAddr = addr
			start = time.Now()
			continue
		}

		if godot == nil {
			continue
		}

		if delta >= 10.0 {
			start = time.Now()
			godot = nil
			godotAddr = nil
			fmt.Println("Godot client disconnected due inactivity")
			continue
		}

		var androidData types.AndroidData

		err = json.Unmarshal([]byte(message), &androidData)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(androidData.Values) < 2 {
			continue
		}

		ad := androidData.Values

		var x float64 = ad[0]
		var y float64 = ad[1]
		var z float64 = ad[2]
		var w float64

		if len(ad) == 3 {
			w = estimateW(x, y, z)
		} else {
			w = ad[3]
		}

		if w == 0 {
			w = estimateW(x, y, z)
		}

		switch androidData.Type {
		case "android.sensor.rotation_vector":
			rotationVector = types.NewRotationVector3(x, y, z, w)
			break
		case "android.sensor.game_rotation_vector":
			rotationVector = types.NewRotationVector3(x, y, z, w)
			break
		}

		if rotationVector == nil {
			continue
		}

		coordinateString := fmt.Sprintf("%f,%f,%f,%f",
			rotationVector.X, rotationVector.Y, rotationVector.Z, rotationVector.W,
		)

		_, err = godot.WriteToUDP([]byte(coordinateString), godotAddr)
		if err != nil {
			fmt.Println("Error sending data to Godot:", err)
		}

		rotationVector = nil
		
		// fmt.Println(message)
		// fmt.Println(coordinateString)
	}
}
