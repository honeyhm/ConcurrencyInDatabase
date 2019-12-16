package main

import (
	"bufio"
	"github.com/kataras/golog"
	"log"
	"os"
	"reflect"
)

func main() {

	helpTable := make([][]string, 5)
	for i := range helpTable {
		helpTable[i] = make([]string, 8)
	}


	readFile, err := os.Open("C:\\Users\\ASUS\\Desktop\\test.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}


	tempLines := fileTextLines


	readFile.Close()

	//for _, eachline := range fileTextLines {
	//	fmt.Println(eachline)
	//}

	// open output file
	fo, err := os.Create("C:\\Users\\ASUS\\Desktop\\Conservative2PL.txt")
	if err != nil {
		panic(err)
	}


	step:=0
	row:=""
	col:=""

	golog.Info("len(fileTextLines) : ",len(fileTextLines))
	for i:=0 ; i<len(fileTextLines) ; i++ {

		for n:=0 ; n<5 ; n++  { // empty helpTable for new schedule
			for m:=0 ; m<8 ; m++  {
				helpTable[n][m]=""
			}
		}

		golog.Info("len(fileTextLines[i])  : ",len(fileTextLines[i]) )
		for j:=0 ; j<len(fileTextLines[i]) ; j+=step  {

			if j+4 >= len(fileTextLines[i]){
				break
			}

			if fileTextLines[i][j]  == 'w' || fileTextLines[i][j]  == 'r'{

				step = 6

				row = string(fileTextLines[i][j+4])
				golog.Info("row : ",row)

				col = string(fileTextLines[i][j+2])
				golog.Info("col : ",col)

			}else{

				step = 4
				col = string(fileTextLines[i][j+2])
			}


			//r(4,x)w(1,z)w(7,z)r(7,x)w(4,w)r(7,z)w(6,x)r(5,v)w(4,v)
			if fileTextLines[i][j] == 'w' {

				golog.Info("fileTextLines[i][j] == 'w' : ",fileTextLines[i][j] == 'w')
				helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"

				golog.Info("helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] : ",helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48])


			}else if fileTextLines[i][j] == 'r' {

				golog.Info("fileTextLines[i][j] == 'r' : ",fileTextLines[i][j] == 'r')
				helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "r"

				golog.Info("fileTextLines[i][j+4]-118 : ",fileTextLines[i][j+4]-118," , ",reflect.TypeOf(fileTextLines[i][j+4]-118))

				golog.Info("fileTextLines[i][j+2]-48 : ",fileTextLines[i][j+2]-48," , ",reflect.TypeOf(fileTextLines[i][j+2]-48))

			}

			//golog.Info(fileTextLines[0][5:12])

			//tempLines[i] = tempLines[i][:2] + "***" + tempLines[i][5:]
			tempLines[i] = tempLines[i][:2] + tempLines[i][5:]

			//input[:index] + string(replacement) + input[index+1:]

			//out := []rune(tempLines[i])
			//out[7] = "*"
			//x:= string(out)

		}

		golog.Info("helpTable line ",i," : ",helpTable)


		if _, err := fo.Write([]byte(tempLines[i]+"\r\n")); err != nil {
			panic(err)
		}

		//if _, err := fo.Write([]byte(fileTextLines[i]+"\r\n")); err != nil {
		//	panic(err)
		//}

	}




	//file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	log.Fatalf("failed creating file: %s", err)
	//}
	//for i:=0;i<len(fileTextLines) ; i++ {
	//	golog.Info("hhh")
	//
	//	//datawriter := bufio.NewWriter(fo)
	//	//
	//	//for _, data := range fileTextLines {
	//	//	_, _ = datawriter.WriteString(data + "\n")
	//	//}
	//	//
	//	//datawriter.Flush()
	//
	//	fmt.Println(fileTextLines[i])
	//	if _, err := fo.Write([]byte(fileTextLines[i]+"\r\n")); err != nil {
	//		panic(err)
	//	}
	//	//if _, err := fo.WriteString("\r\n"); err != nil {
	//	//	panic(err)
	//	//}
	//}


	//close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// an array with 5 rows and 2 columns
	//var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	//var i, j int
	//
	///* output each array element's value */
	//for  i = 0; i < 5; i++ {
	//	for j = 0; j < 2; j++ {
	//		fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
	//	}
	//}
}
