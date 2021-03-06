package graphicsbackend

type Interface interface {
	GetRendererName() string
	SetWindowIcon(fileName string)
	IsFullScreen() bool
	SetFullScreen(fullScreen bool)
	SetVSyncEnabled(vsync bool)
	GetVSyncEnabled() bool
	GetCursorPos() (int, int)
	CurrentFPS() float64
	// NewSurface(width, height int) Surface
	Render() error
}
