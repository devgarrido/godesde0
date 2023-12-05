package variables

import "fmt"

// Constantes a nivel de package.
const SO_WINDOWS = "Windows"

// Const with agrupation.
const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
)

// Const iota (autoincremental only for constants) -> iota starts with 0
const (
	Jan = iota + 1 // 1
	Feb            // 2
	Mar            // 3
)

func ShowConstants() {
	// Definition nivel de funcion.
	const SO_LINUX = "Linux"
	const SO_MAC, SO_ANDROID = "Mac", "Android"
	fmt.Printf("System: %s\n", SO_LINUX)
	fmt.Printf("System: %s\n", SO_WINDOWS)
	fmt.Printf("La Capacidad de 1KB: %d bytes\n", KB)
	fmt.Printf("Mac tiene el sistema operativo: %s\n", SO_MAC)
	fmt.Printf("Android tiene el sistema operativo: %s\n", SO_ANDROID)

	fmt.Printf("Enero el mes número del año: %d\n", Jan)

}
