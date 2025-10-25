package converter

import (
	"fmt"
	"os/exec"
	_ "path/filepath"
	"sync"
)

type CommandRunner interface {
	Run(cmd string, args ...string) error
}

type RealCommandRunner struct{}

func (r RealCommandRunner) Run(cmd string, args ...string) error {
	command := exec.Command(cmd, args...)
	return command.Run()
}

func ConvertToWav(inputFile string, runner CommandRunner, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done()

	outputFile := inputFile[:len(inputFile)-4] + ".wav"

	err := runner.Run("ffmpeg", "-i", inputFile, outputFile, "-y")
	if err != nil {
		ch <- fmt.Sprintf("Ошибка обработки %s: %v", inputFile, err)
		return
	}

	ch <- fmt.Sprintf("✅ Успешно сконвертирован: %s → %s", inputFile, outputFile)
}

//func main() {
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
//	var wg sync.WaitGroup
//	ch := make(chan string, len(files))
//
//	runner := RealCommandRunner{}
//
//	for _, file := range files {
//		wg.Add(1)
//		go convertToWav(file, runner, &wg, ch)
//	}
//
//	go func() {
//		wg.Wait()
//		close(ch)
//	}()
//
//	for msg := range ch {
//		fmt.Println(msg)
//	}
//
//	fmt.Println("🎵 Все файлы обработаны!")
//}
