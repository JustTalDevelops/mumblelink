package main

// mumblePosition contains position info, which is used for both the avatar and the camera.
type mumblePosition struct {
	// Position is the exact position of the player.
	Position [3]float32
	// Front is currently undocumented.
	Front [3]float32
	// Top is currently undocumented.
	Top [3]float32
}

// mumbleData is the data needed to be sent to Mumble every tick.
type mumbleData struct {
	// Version is always 2.
	Version uint32
	// Tick is used to make sure the connection is still alive, so it needs to be incremented every tick.
	Tick uint32
	// Avatar is the position of the avatar. This is the same as the Camera position for us, since this is Minecraft.
	Avatar mumblePosition
	// Name is the name of the link in runes, converted to uint16s.
	Name [256]uint16
	// Camera is the position of the camera. This is the same as the Avatar position for us, since this is Minecraft.
	Camera mumblePosition
	// Identity is used to track the identity of the mumble user.
	Identity [256]uint16
	// ContextLength is always 48.
	ContextLength uint32
	// Context is used to make sure the only voices that get broadcast are those of the same identity. This should be the connected Minecraft server.
	Context [256]byte
	// Description is the description of the link.
	Description [2048]uint16
}
