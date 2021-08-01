package mumblelink

import (
	"github.com/go-gl/mathgl/mgl32"
	"syscall"
	"time"
	"unsafe"
)

// MumblePosition contains position info, which is used for both the avatar and the camera.
type MumblePosition struct {
	// Position is the exact position of the player.
	Position mgl32.Vec3
	// Front is currently undocumented.
	Front mgl32.Vec3
	// Top is currently undocumented.
	Top mgl32.Vec3
}

// Mumble contains the necessary information to broadcast data on the Mumble link.
type Mumble struct {
	// Name is the name of the link.
	Name string
	// Description is the description of the link.
	Description string
	// Identity is the identity of the Mumble user.
	Identity string
	// Context is used to make sure the only voices that get broadcast are those of the same identity. This should be the connected Minecraft server.
	Context string
	// Position is the position data sent over the link.
	Position MumblePosition
}

// Start starts ticking the Mumble instance on a goroutine. Any changes made will get broadcast on the next tick.
func (m *Mumble) Start() {
	go func() {
		file, _ := syscall.UTF16PtrFromString("MumbleLink")
		handle, err := syscall.CreateFileMapping(0, nil, syscall.PAGE_READWRITE, 0, 1000, file)
		if err != nil {
			panic(err)
		}
		defer syscall.CloseHandle(handle)
		addr, err := syscall.MapViewOfFile(handle, syscall.FILE_MAP_WRITE, 0, 0, 0)
		if err != nil {
			panic(err)
		}

		ticker := time.NewTicker(time.Second / 20)
		defer ticker.Stop()
		for range ticker.C {
			//goland:noinspection GoVetUnsafePointer
			data := (*mumbleData)(unsafe.Pointer(addr))
			if data == nil {
				panic("mumble is not running")
			}
			data.Version = 2
			data.Tick++

			data.Avatar.Position = m.Position.Position
			data.Avatar.Front = m.Position.Front
			data.Avatar.Top = m.Position.Top

			data.Name = stringTo256Uint16Slice(m.Name)
			data.Description = stringTo2048Uint16Slice(m.Description)

			data.Camera.Position = m.Position.Position
			data.Camera.Front = m.Position.Front
			data.Camera.Top = m.Position.Top

			data.Identity = stringTo256Uint16Slice(m.Identity)

			data.ContextLength = 48 // Always 48 for some reason?
			data.Context = stringTo256ByteSlice(m.Context)
		}
	}()
}
