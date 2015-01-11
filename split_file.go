package main

import "fmt"
import "os"
import "log"
import "strconv"
import "bufio"

func SplitName(fileName string, MapJob int) string {
	fstr := "mytest-mapreduce." + fileName + "-" + strconv.Itoa(MapJob)
	//fmt.Printf("Mapname: %s.\n",fstr)
	return fstr
}

func Split(fileName string) {
	fmt.Printf("Split %s\n", fileName)
	infile, err := os.Open(fileName)
	if err != nil {
		log.Fatal("Split: ", err)
	}
	defer infile.Close()
	fi, err := infile.Stat()
	if err != nil {
		log.Fatal("Split: ", err)
	}
	size := fi.Size()
	nchunk := size / 3
	nchunk += 1

	outfile, err := os.Create(SplitName(fileName, 0))
	if err != nil {
		log.Fatal("Split: ", err)
	}
	writer := bufio.NewWriter(outfile)
	m := 1
	i := 0

	scanner := bufio.NewScanner(infile)
	for scanner.Scan() {
		if int64(i) > nchunk*int64(m) {
			writer.Flush()
			outfile.Close()
			outfile, err = os.Create(SplitName(fileName, m))

			writer = bufio.NewWriter(outfile)
			m += 1
		}
		line := scanner.Text() + "\n"
		writer.WriteString(line)
		i += len(line)
	}
	writer.Flush()
	outfile.Close()
}

func main() {
	Split("testmr.txt")
}
