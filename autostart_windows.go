package myautostart

import (
	"golang.org/x/sys/windows/registry"
	"os"
	"path/filepath"
)

const (
	runRegistryKey = `Software\Microsoft\Windows\CurrentVersion\Run`
)

func (a *App) getExePath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(exe)
}

func (a *App) IsEnabled() bool {
	key, err := registry.OpenKey(registry.CURRENT_USER, runRegistryKey, registry.QUERY_VALUE)
	if err != nil {
		return false
	}
	defer func() {
		_ = key.Close()
	}()

	val, _, err := key.GetStringValue(a.Name)
	if err == registry.ErrNotExist {
		return false
	} else if err != nil {
		return false
	}

	exePath, err := a.getExePath()
	if err != nil {
		return false
	}

	return val == exePath
}

func (a *App) Enable() error {
	key, _, err := registry.CreateKey(registry.CURRENT_USER, runRegistryKey, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer key.Close()

	exePath, err := a.getExePath()
	if err != nil {
		return err
	}

	return key.SetStringValue(a.Name, exePath)
}

func (a *App) Disable() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, runRegistryKey, registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer func() {
		_ = key.Close()
	}()

	return key.DeleteValue(a.Name)
}
