package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kirillmc/data_filler/pkg/filler"
	fillerPb "github.com/kirillmc/data_filler/pkg/filler_pb"
)

/**
Time to create JSON: 102.876µs  Time to create Protobuff: 83.454µs
Пустая запись:          Размер формата JSON: 23.000 байт                Размер формата Protobuff: 0.000 байт
Time to create JSON: 31.121µs   Time to create Protobuff: 94.979µs
Одна запись[1]:         Размер формата JSON: 705.000 байт               Размер формата Protobuff: 495.000 байт
Time to create JSON: 708.692µs  Time to create Protobuff: 465.517µs
Маленький[5]:           Размер формата JSON: 60.220 килобайт            Размер формата Protobuff: 37.690 килобайт
Time to create JSON: 11.792219ms        Time to create Protobuff: 4.915919ms
Средний[11]:            Размер формата JSON: 900.317 килобайт   Размер формата Protobuff: 495.047 килобайт
Time to create JSON: 405.141665ms       Time to create Protobuff: 152.287426ms
Большой[31]:            Размер формата JSON: 40.992 мегабайт            Размер формата Protobuff: 19.138 мегабайт
Time to create JSON: 3.370918681s       Time to create Protobuff: 1.408141365s
Меге[55]:               Размер формата JSON: 373.294 мегабайт           Размер формата Protobuff: 163.746 мегабайт
Time to create JSON: 11.427827523s      Time to create Protobuff: 5.289707675s
Ультра[75]:             Размер формата JSON: 1252.037 мегабайт          Размер формата Protobuff: 535.690 мегабайт
*/

func main() {

	printSizes("", 0)
	printSizes("", 1)
	printSizes("", 5)
	printSizes("", 11)
	printSizes("", 31)
	printSizes("Меге", 55)
	printSizes("Ультра", 75)

}

func getSizeInFormattedString(byteSize int64) string {
	if byteSize < 1024 {
		return fmt.Sprintf("%.3f байт\t", float64(byteSize))
	}
	if byteSize < 1024*1024 {
		return fmt.Sprintf("%.3f килобайт\t", float64(byteSize)/1024)
	} else {
		return fmt.Sprintf("%.3f мегабайт\t", float64(byteSize)/(1024*1024))
	}
}

func printSizes(typeName string, count int) {
	switch count {
	case 0:
		start := time.Now()

		emptySizeJSON, err := filler.FindByteSizeOfJson(filler.CreateEmptySetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create JSON: %v\t", time.Now().Sub(start))

		start = time.Now()

		emptySizePb, err := fillerPb.FindByteSizeOfProto(fillerPb.CreateEmptySetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create Protobuff: %v\n", time.Now().Sub(start))
		fmt.Printf("Пустая запись:\t\tРазмер формата JSON: %s\tРазмер формата Protobuff: %s\n", getSizeInFormattedString(emptySizeJSON), getSizeInFormattedString(emptySizePb))
	case 1:
		start := time.Now()

		oneSizeJSON, err := filler.FindByteSizeOfJson(filler.CreateSetOfOneProgram())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create JSON: %v\t", time.Now().Sub(start))

		start = time.Now()

		oneSizePb, err := fillerPb.FindByteSizeOfProto(fillerPb.CreateSetOfOneProgram())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create Protobuff: %v\n", time.Now().Sub(start))
		fmt.Printf("Одна запись[1]:\t\tРазмер формата JSON: %s\tРазмер формата Protobuff: %s\n", getSizeInFormattedString(oneSizeJSON), getSizeInFormattedString(oneSizePb))
	case 5:
		start := time.Now()

		smallSizeJSON, err := filler.FindByteSizeOfJson(filler.CreateSmallSetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create JSON: %v\t", time.Now().Sub(start))

		start = time.Now()

		smallSizePb, err := fillerPb.FindByteSizeOfProto(fillerPb.CreateSmallSetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create Protobuff: %v\n", time.Now().Sub(start))
		fmt.Printf("Маленький[5]:\t\tРазмер формата JSON: %s\tРазмер формата Protobuff: %s\n", getSizeInFormattedString(smallSizeJSON), getSizeInFormattedString(smallSizePb))
	case 11:
		start := time.Now()

		normalSizeJSON, err := filler.FindByteSizeOfJson(filler.CreateNormalSetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create JSON: %v\t", time.Now().Sub(start))

		start = time.Now()

		normalSizePb, err := fillerPb.FindByteSizeOfProto(fillerPb.CreateNormalSetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create Protobuff: %v\n", time.Now().Sub(start))
		fmt.Printf("Средний[11]:\t\tРазмер формата JSON: %sРазмер формата Protobuff: %s\n", getSizeInFormattedString(normalSizeJSON), getSizeInFormattedString(normalSizePb))
	case 31:
		start := time.Now()

		bigSizeJSON, err := filler.FindByteSizeOfJson(filler.CreateBigSetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create JSON: %v\t", time.Now().Sub(start))

		start = time.Now()

		bigSizePb, err := fillerPb.FindByteSizeOfProto(fillerPb.CreateBigSetOfPrograms())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create Protobuff: %v\n", time.Now().Sub(start))
		fmt.Printf("Большой[31]:\t\tРазмер формата JSON: %s\tРазмер формата Protobuff: %s\n", getSizeInFormattedString(bigSizeJSON), getSizeInFormattedString(bigSizePb))
	default:
		start := time.Now()

		customSizeJSON, err := filler.FindByteSizeOfJson(filler.CreateOwnSetOfPrograms(count))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create JSON: %v\t", time.Now().Sub(start))

		start = time.Now()

		customSizePb, err := fillerPb.FindByteSizeOfProto(fillerPb.CreateOwnSetOfPrograms(count))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Time to create Protobuff: %v\n", time.Now().Sub(start))
		fmt.Printf("%s[%d]:\t\tРазмер формата JSON: %s\tРазмер формата Protobuff: %s\n", typeName, count, getSizeInFormattedString(customSizeJSON), getSizeInFormattedString(customSizePb))

	}

}
