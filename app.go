package main

import (
	"context"
	"fmt"
	"strings"
)

// App структура приложения
type App struct {
	ctx context.Context
}

// NewApp создаёт новый экземпляр приложения
func NewApp() *App {
	return &App{}
}

// startup вызывается при запуске приложения. Контекст сохраняется
// для возможности вызова runtime методов
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// TimeEntry представляет одну временную запись
type TimeEntry struct {
	Original string `json:"original"`
	Seconds  int    `json:"seconds"`
}

// TimeResult представляет результат вычисления
type TimeResult struct {
	Entries        []TimeEntry `json:"entries"`
	TotalSeconds   int         `json:"totalSeconds"`
	TotalFormatted string      `json:"totalFormatted"`
}

// ParseAndCalculateTime обрабатывает строки времени и вычисляет общую сумму
func (a *App) ParseAndCalculateTime(input string) TimeResult {
	// Нормализуем переносы строк
	input = strings.ReplaceAll(input, "\r\n", "\n")
	input = strings.ReplaceAll(input, "\r", "\n")

	lines := strings.Split(strings.TrimSpace(input), "\n")
	var entries []TimeEntry
	totalSeconds := 0

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		seconds := parseTimeString(line)
		if seconds >= 0 {
			entries = append(entries, TimeEntry{
				Original: line,
				Seconds:  seconds,
			})
			totalSeconds += seconds
		}
	}

	return TimeResult{
		Entries:        entries,
		TotalSeconds:   totalSeconds,
		TotalFormatted: formatSeconds(totalSeconds),
	}
}

// parseTimeString обрабатывает формат ЧЧ:ММ:СС и возвращает общее количество секунд
func parseTimeString(timeStr string) int {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return -1
	}

	var hours, minutes, seconds int
	n1, err1 := fmt.Sscanf(parts[0], "%d", &hours)
	n2, err2 := fmt.Sscanf(parts[1], "%d", &minutes)
	n3, err3 := fmt.Sscanf(parts[2], "%d", &seconds)

	// Проверяем, что все части были успешно распознаны
	if n1 != 1 || n2 != 1 || n3 != 1 || err1 != nil || err2 != nil || err3 != nil {
		return -1
	}

	// Проверяем корректность временных значений (часы могут быть > 24 для накопленного времени)
	if hours < 0 || minutes < 0 || minutes > 59 || seconds < 0 || seconds > 59 {
		return -1
	}

	return hours*3600 + minutes*60 + seconds
}

// formatSeconds преобразует секунды в формат ЧЧ:ММ:СС
func formatSeconds(totalSeconds int) string {
	hours := totalSeconds / 3600
	minutes := (totalSeconds % 3600) / 60
	seconds := totalSeconds % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

// GetClipboardText получает текст из буфера обмена (вызывается из фронтенда)
func (a *App) GetClipboardText() string {
	// Обрабатывается фронтенд JavaScript кодом
	return ""
}
