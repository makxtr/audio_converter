package converter

import (
	"fmt"
	"os/exec"
	_ "path/filepath"
	"sync"
)

// Интерфейс для команд
type CommandRunner interface {
	Run(cmd string, args ...string) error
}

// Реальная реализация (использует exec.Command)
type RealCommandRunner struct{}

func (r RealCommandRunner) Run(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	return command.Run()
}

// Функция для конвертации файла в WAV
func ConvertToWav(inputFile string, runner CommandRunner, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	outputFile := inputFile[:len(inputFile)-4] + ".wav"

	// Используем переданный CommandRunner
	err := runner.Run("ffmpeg", "-i", inputFile, outputFile, "-y")
	if err != nil {
		ch <- fmt.Sprintf("Ошибка обработки %s: %v", inputFile, err)
		return
	}

	ch <- fmt.Sprintf("✅ Успешно сконвертирован: %s → %s", inputFile, outputFile)
}

//func main() {
//	// Получаем список файлов из текущей директории
//	files, err := filepath.Glob("*.mp3") // Ищем все MP3-файлы
//	if err != nil {
//		fmt.Println("Ошибка поиска файлов:", err)
//		return
//	}
//
//	if len(files) == 0 {
//		fmt.Println("⚠ Нет MP3-файлов для обработки.")
//		return
//	}
//
//	// Создаём WaitGroup для синхронизации горутин
//	var wg sync.WaitGroup
//	// Канал для получения статуса
//	ch := make(chan string, len(files))
//
//	runner := RealCommandRunner{}
//
//	// Запускаем обработку в горутинах
//	for _, file := range files {
//		wg.Add(1)
//		go convertToWav(file, runner, &wg, ch)
//	}
//
//	// Закрываем канал после завершения всех горутин
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	// Выводим статусы из канала
//	for msg := range ch {
//		fmt.Println(msg)
//	}
//
//	fmt.Println("🎵 Все файлы обработаны!")
//}
