package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	// Список переменных окружения
	for _, s := range os.Environ() {
		kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
		fmt.Printf("%q = %q\n", kv[0], kv[1])
	}

	fmt.Println("--------------------------------------")

	// Модификация переменных окружения
	fmt.Printf("%q\n", os.Getenv("ARTEM"))
	os.Setenv("ARTEM", "/bin/dash")
	fmt.Printf("%q\n", os.Getenv("ARTEM"))
	os.Unsetenv("ARTEM")
	fmt.Printf("%q\n", os.Getenv("ARTEM"))

	fmt.Println("--------------------------------------")

	// Запуск стороннего процесса
	binary, lookErr := exec.LookPath("curl")
	if lookErr != nil {
		fmt.Println(lookErr)
	} else {
		args := []string{"http://yandex.ru"}
		env := os.Environ()
		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			fmt.Println(execErr)
		}
	}

	// Запуск дочернего процесса
	output, err := exec.Command("curl", "http://yandex.ru").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit code: ", e.ExitCode())
		default:
			fmt.Println(err)
		}
	} else {
		fmt.Println(output)
	}

	// Общение с дочерним процессом
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println(string(grepBytes))

	fmt.Println("--------------------------------------")

	// Получение аргументов командной строки
	if len(os.Args) != 3 {
		var programName = filepath.Base(os.Args[0])
		fmt.Println("Usage:", programName, "PATTERN", "FILE")
	} else {
		pattern := os.Args[1]
		file := os.Args[2]
		fmt.Printf("%s: %s\n", file, pattern)
	}

	fmt.Println("--------------------------------------")

	// Получение сигналов от ОС
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done

	fmt.Println("--------------------------------------")

	// Принудительный выход из приложения
	defer fmt.Println("!!!!!!!!!!!!!")
	fmt.Println("exiting...")
	os.Exit(3) // <-- deferred код не выполняется!
	fmt.Println("exited!")
}
