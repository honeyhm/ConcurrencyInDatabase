package main

import (
	"bufio"
	"github.com/kataras/golog"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	//read , modify with spaces and separating file line by line
	readFile, err := os.Open("C:\\Users\\ASUS\\Desktop\\Schedule.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}

	//tempLines := fileTextLines

	readFile.Close()

	// make and open output file for Basic2PL algorithm
	fo, err := os.Create("C:\\Users\\ASUS\\Desktop\\Basic2PL.txt")
	if err != nil {
		panic(err)
	}


	// loop for iterating throw all schedules
	for i:=0 ; i<len(fileTextLines) ; i++ {

		tempLines := fileTextLines[i]

		golog.Info("new schedule started : ************************************************************************")

		//string variable for holding result
		result := ""
		//number of each transaction orders in each schedule
		TransOrderNums := make([]int, 8)
		//number of locks given by each transaction
		TransOrderLockNums := make([]int, 8)

		helpArray := make([]int, 8)

		for n:=1 ; n<8 ; n++  {

			TransOrderNums[n] = strings.Count(fileTextLines[i], strconv.Itoa(n)) -1
			if TransOrderNums[n] < 0{
				TransOrderNums[n]+=1
			}

		}
		golog.Info("TransOrderNums : ",TransOrderNums)


		// a table which has transactions as columns and variables(v-z) as rows
		helpTable := make([][]string, 5)// making 5 rows of variables
		for i := range helpTable {
			helpTable[i] = make([]string, 8)//making 7 (ignoring index 0) columns for each row
		}


		step:=0

		counter:=0
		for len(fileTextLines[i])>1 {
			//deadlock checking
			counter+=1
			if counter>3{
				break
			}
			// loop for iterating throw each schedule content
			for j:=0 ; j<len(fileTextLines[i]) ; j+=step  {

				if j+6 >= len(fileTextLines[i]){
					break
				}

				if fileTextLines[i][j] == 'w' {

					step = 6
					flag1 := 0

					if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] == "" {//lock

						for k:=1 ; k<8 ; k++  {
							if helpTable[fileTextLines[i][j+4]-118][k] != ""{
								flag1 = 1
								break
							}
						}

						if flag1 == 0 {

							if helpArray[fileTextLines[i][j+2]-48] == 0 {

								TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
								helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"
								result = result + tempLines[j:j+1] + "l(" + tempLines[j+2:j+6]
								result = result + tempLines[j:j+6]
								tempLines = tempLines[:j] + "*" +tempLines[j+1:]

							}else {

								helpArray[fileTextLines[i][j+2]-48]+=1
								TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
								helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"
							}

						}else {

							helpArray[fileTextLines[i][j+2]-48]+=1
						}

					}else if !strings.Contains(helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48], "w") && strings.Contains(helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48], "r") { // lock upgrade

						for k:=1 ; k<8 ; k++  {
							if helpTable[fileTextLines[i][j+4]-118][k] != ""{
								flag1 += 1
							}
						}

						if flag1 == 1{

							if helpArray[fileTextLines[i][j+2]-48] == 0 {

								TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
								helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
								result = result + tempLines[j:j+1] + "l(" + tempLines[j+2:j+6]
								result = result + tempLines[j:j+6]
								tempLines = tempLines[:j] + "*" +tempLines[j+1:]

							}else {

								helpArray[fileTextLines[i][j+2]-48]+=1
								TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
								helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
							}
						}

					}else if strings.Contains(helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48], "w") {

						if helpArray[fileTextLines[i][j+2]-48] == 0 {

							TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
							helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
							result = result + tempLines[j:j+1] + "l(" + tempLines[j+2:j+6]
							result = result + tempLines[j:j+6]
							tempLines = tempLines[:j] + "*" +tempLines[j+1:]

						}else {

							helpArray[fileTextLines[i][j+2]-48]+=1
							TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
							helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
						}
					}



				}else if fileTextLines[i][j] == 'r' {

					step = 6
					flag1 := 0

					if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] == "" {

						for k:=1 ; k<8 ; k++  {
							if helpTable[fileTextLines[i][j+4]-118][k] == "w"{
								flag1 = 1
								break
							}
						}

						if flag1 == 0{

							if helpArray[fileTextLines[i][j+2]-48] == 0 {

								TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
								helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "r"
								result = result + tempLines[j:j+1] + "l(" + tempLines[j+2:j+6]
								result = result + tempLines[j:j+6]
								tempLines = tempLines[:j] + "*" +tempLines[j+1:]

							}else {

								helpArray[fileTextLines[i][j+2]-48]+=1
								TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
								helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "r"
							}

						}else {

							helpArray[fileTextLines[i][j+2]-48]+=1
						}

					}else if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] != "" {

						if helpArray[fileTextLines[i][j+2]-48] == 0 {

							TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
							helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "r"
							result = result + tempLines[j:j+1] + "l(" + tempLines[j+2:j+6]
							result = result + tempLines[j:j+6]
							tempLines = tempLines[:j] + "*" +tempLines[j+1:]

						}else {

							helpArray[fileTextLines[i][j+2]-48]+=1
							TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
							helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "r"
						}
					}

				}else{
					step = 4
				}

				golog.Info("string(fileTextLines[i][j]) ; ",string(fileTextLines[i][j]))

				if fileTextLines[i][j] == 'w' || fileTextLines[i][j] == 'r' {

					golog.Info("1*1")
					l:= fileTextLines[i][j+2]-48
					golog.Info("fileTextLines[i][j+2]-48 : ",l," , ",reflect.TypeOf(l))

					if TransOrderLockNums[l] == TransOrderNums[l] {

						golog.Info("2*2")

						for k:=0 ; k<5 ; k++ {
							helpTable[k][fileTextLines[i][j+2]-48] = ""
						}


						for true  {
							golog.Info("************///////////////////////////*************")
							golog.Info("tempLines before : ",tempLines)

							t := strings.Index(tempLines, "*("+strconv.Itoa(int(l)))
							golog.Info("t : ", t)
							if t == -1{
								break
							}
							result = result + "ul" + tempLines[t+1:t+6]

							tempLines = tempLines[:t] + tempLines[t+6:]
							fileTextLines[i] = fileTextLines[i][:t] + fileTextLines[i][t+6:]
							if t <= j {
								step = 0
							}

						}

						end := strings.Index(tempLines, "("+strconv.Itoa(int(l))+")")
						golog.Info("end : ",end)
						if end != -1{
							result = result + tempLines[end-1:end+3]
							tempLines = tempLines[:end-1] + tempLines[end+3:]
							fileTextLines[i] = fileTextLines[i][:end-1] + fileTextLines[i][end+3:]
						}

						golog.Info("tempLines after  : ",tempLines)
						golog.Info("res : ",result)

					}
				}
			}

		}


		golog.Info("fileTextLines[i] : ",fileTextLines[i])
		golog.Info("tempLines : ",tempLines)

		result+=";"
		if counter>3{
			result+="D"
		}
		if _, err := fo.Write([]byte(result+"\r\n")); err != nil {
			panic(err)
		}

	}


	//close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

}