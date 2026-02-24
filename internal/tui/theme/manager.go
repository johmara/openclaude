package theme

import "sync"

var (
	mu      sync.RWMutex
	current Theme = DefaultTheme{}
	themes  []Theme
)

func init() {
	themes = []Theme{
		DefaultTheme{},
		CatppuccinTheme{},
		DraculaTheme{},
	}
}

// Current returns the active theme.
func Current() Theme {
	mu.RLock()
	defer mu.RUnlock()
	return current
}

// Set changes the active theme.
func Set(t Theme) {
	mu.Lock()
	defer mu.Unlock()
	current = t
}

// SetByIndex sets the active theme by index.
func SetByIndex(idx int) {
	mu.Lock()
	defer mu.Unlock()
	if idx >= 0 && idx < len(themes) {
		current = themes[idx]
	}
}

// All returns all registered themes.
func All() []Theme {
	mu.RLock()
	defer mu.RUnlock()
	return themes
}

// CurrentIndex returns the index of the current theme.
func CurrentIndex() int {
	mu.RLock()
	defer mu.RUnlock()
	for i, t := range themes {
		if t.Name() == current.Name() {
			return i
		}
	}
	return 0
}
