package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"
	"strings"
)

type PhoneReader string

func (phone PhoneReader) Read(p []byte) (int, error) {
	count := 0
	for i := 0; i < len(phone); i++ {
		if phone[i] >= '0' && phone[i] <= '9' {
			p[count] = phone[i]
			count++
		}
	}
	return count, io.EOF
}

type PhoneWriter struct{}

func (writer PhoneWriter) Write(bs []byte) (int, error) {
	if len(bs) == 0 {
		return 0, nil
	}
	for i := 0; i < len(bs); i++ {
		if bs[i] >= '0' && bs[i] <= '9' {
			fmt.Print(string(bs[i]))
		}
	}
	fmt.Println()
	return len(bs), nil
}

type Person struct {
	Name string
	Age  int
}

func main() {
	// Кастомный reader
	phone1 := PhoneReader("+1(234)567 9010")
	phone2 := PhoneReader("+2-345-678-12-35")

	buffer := make([]byte, len(phone1))
	phone1.Read(buffer)
	fmt.Println(string(buffer))

	buffer = make([]byte, len(phone2))
	phone2.Read(buffer)
	fmt.Println(string(buffer))

	fmt.Println("--------------------------------------")

	// Кастомный writer
	bytes1 := []byte("+1(234)567 9010")
	bytes2 := []byte("+2-345-678-12-35")

	pw := PhoneWriter{}
	pw.Write(bytes1)
	pw.Write(bytes2)

	fmt.Println("--------------------------------------")

	// открытие файла для чтения
	f1, err := os.OpenFile("hello.txt", os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	f1.Close()

	if exists, _ := FileExists("hello.txt"); !exists {
		// Создание файла
		f3, err := os.Create("hello.txt")
		if err != nil {
			fmt.Println(err)
		}
		f3.Close()
	}

	// открытие файла для записи
	f2, err := os.OpenFile("hello.txt", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println(err)
	}
	f2.Close()

	fmt.Println("--------------------------------------")

	// Запись в файл
	wFile, err := os.OpenFile("hello.txt", os.O_WRONLY, 0777)
	defer wFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		wFile.WriteString("Hello, World!\n")

		var person = Person{Name: "Artem", Age: 20}
		data, _ := json.Marshal(person)
		wFile.Write(data)
	}

	fmt.Println("--------------------------------------")

	// Чтение из файла
	rFile, err := os.Open("hello.txt")
	defer rFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		data := make([]byte, 4)
		rFile.Seek(4, io.SeekStart) // сдвигаем позицию для чтения
		for {
			n, err := io.ReadAtLeast(rFile, data, 2)
			//n, err := rFile.Read(data)
			if err == io.EOF { // если конец файла
				break
			}
			fmt.Print(string(data[:n]))
		}
	}
	fmt.Println()

	// Чтение файла целиком
	dat, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print(string(dat))
	}

	fmt.Println("\n--------------------------------------")

	// Копирование потоков
	io.Copy(os.Stdout, phone1)
	fmt.Println()

	rFile, err = os.Open("hello.txt")
	defer rFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		io.Copy(os.Stdout, rFile)
	}

	fmt.Println("\n--------------------------------------")

	// Форматированная запись в файл
	wFile, err = os.OpenFile("format.txt", os.O_WRONLY|os.O_CREATE, 0777)
	defer wFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Fprintf(wFile, "%0.2f\n", 3.14)

		var person1 = Person{Name: "Tom", Age: 30}
		fmt.Fprintf(wFile, "%s %d\n", person1.Name, person1.Age)

		var person2 = Person{Name: "Artem", Age: 20}
		fmt.Fprintf(wFile, "%+v\n", person2)

		fmt.Fprintln(wFile, "Hello, World!")
	}

	fmt.Println("\n--------------------------------------")

	// Форматированное чтение из файла
	rFile, err = os.Open("format.txt")
	defer rFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		var num float64
		_, err = fmt.Fscanf(rFile, "%g", &num)
		fmt.Println(num)

		var person1 = Person{}
		fmt.Fscanf(rFile, "%s %d\n", &person1.Name, &person1.Age)
		fmt.Printf("%+v\n", person1)

		// TODO: Десериализация объектов так не взлетает пока
		//var person2 = Person{}
		//fmt.Fscanf(rFile, "%+v\n", &person2)
		//fmt.Printf("%+v", person2)

		// Так хрень какая-то получается - Fscanln считывает через пробелы, а не через перевод каретки
		var str string
		for {
			_, err := fmt.Fscanln(rFile, &str)
			fmt.Print(str)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	}

	fmt.Println("\n--------------------------------------")

	// Буферизованная запись в файл
	wFile, err = os.OpenFile("buffer.txt", os.O_WRONLY|os.O_CREATE, 0777)
	defer wFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		writer := bufio.NewWriter(wFile)
		writer.WriteString("Hi, people!")
		writer.WriteString("\n")

		writer.WriteString(fmt.Sprintf("%0.2f", 2.71))
		writer.WriteString("\n")

		var person = Person{Name: "Artem", Age: 30}
		pJson, _ := json.Marshal(person)
		writer.Write(pJson)

		writer.Flush()
	}

	fmt.Println("\n--------------------------------------")

	// Буферизованное чтение из файла
	rFile, err = os.Open("buffer.txt")
	defer rFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		reader := bufio.NewReader(rFile)
		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(line))
		}
	}

	fmt.Println("\n--------------------------------------")

	// Буферизованно-форматированное чтение из файла
	rFile, err = os.Open("buffer.txt")
	defer rFile.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		reader := bufio.NewReader(rFile)
		// TODO: Обработка ошибок
		line /* строка */, _, _ := reader.ReadLine()
		strings.NewReader(string(line))
		fmt.Println(string(line))

		line /* число */, _, _ = reader.ReadLine()
		sr := strings.NewReader(string(line))
		var num float64
		fmt.Fscanf(sr, "%f", &num)
		fmt.Println(num)
	}

	fmt.Println("\n--------------------------------------")

	// Буфер байтов
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "Size: %d MB.", 85)
	str := buf.String()
	fmt.Println(str)

	fmt.Println("\n--------------------------------------")

	CreateImage("image.png")

	fmt.Println("\n--------------------------------------")

	// Ввод с консоли
	var name string
	var age int
	fmt.Print("Введите имя и возраст: ")
	fmt.Scan(&name, &age)
	fmt.Println(name, age)

}

func FileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func CreateImage(filePath string) {
	width := 200
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Colors are defined by Red, Green, Blue, Alpha uint8 values.
	cyan := color.RGBA{100, 200, 200, 0xff}

	// Set color for each pixel.
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/2 && y < height/2: // upper left quadrant
				img.Set(x, y, cyan)
			case x >= width/2 && y >= height/2: // lower right quadrant
				img.Set(x, y, color.White)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create(filePath)
	png.Encode(f, img)
}
