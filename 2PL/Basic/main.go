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

	// make and open output file for Basic2PL algorithm
	fo, err := os.Create("C:\\Users\\ASUS\\Desktop\\Basic2PL.txt")
	if err != nil {
		panic(err)
	}


	// loop for iterating throw all schedules
	for i:=0 ; i<len(fileTextLines) ; i++ {

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
		golog.Info("helpTable : ",helpTable)


		step:=0

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
							result = result + tempLines[i][j:j+1] + "l(" + tempLines[i][j+2:j+6]
							result = result + tempLines[i][j:j+6]
							tempLines[i] = tempLines[i][:j] + "*" +tempLines[i][j+1:]
							golog.Info("tempLines[i]  : ",tempLines[i])
							//result = result + "ul" + tempLines[i][j+1:j+6]

						}else {

							helpArray[fileTextLines[i][j+2]-48]+=1
							TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
							helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"
						}
						//TransOrderLockNums[fileTextLines[i][j+2]-48] += 1
						//helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"

					}else {

						helpArray[fileTextLines[i][j+2]-48]+=1
					}

				//}else if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] == "r" { // lock upgrade
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
							result = result + tempLines[i][j:j+1] + "l(" + tempLines[i][j+2:j+6]
							result = result + tempLines[i][j:j+6]
							tempLines[i] = tempLines[i][:j] + "*" +tempLines[i][j+1:]
							golog.Info("tempLines[i]  : ",tempLines[i])
							//result = result + "ul" + tempLines[i][j+1:j+6]

						}else {

							helpArray[fileTextLines[i][j+2]-48]+=1
							TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
							helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
						}
						//TransOrderLockNums[fileTextLines[i][j+2]-48]+=1
						//helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
					}

				//}else {// if w exists
				}else if strings.Contains(helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48], "w") {

					if helpArray[fileTextLines[i][j+2]-48] == 0 {

						TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
						helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
						result = result + tempLines[i][j:j+1] + "l(" + tempLines[i][j+2:j+6]
						result = result + tempLines[i][j:j+6]
						tempLines[i] = tempLines[i][:j] + "*" +tempLines[i][j+1:]
						golog.Info("tempLines[i]  : ",tempLines[i])
						//result = result + "ul" + tempLines[i][j+1:j+6]

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
							result = result + tempLines[i][j:j+1] + "l(" + tempLines[i][j+2:j+6]
							result = result + tempLines[i][j:j+6]
							tempLines[i] = tempLines[i][:j] + "*" +tempLines[i][j+1:]
							golog.Info("tempLines[i]*  : ",tempLines[i])
							//result = result + "ul" + tempLines[i][j+1:j+6]

						}else {

							helpArray[fileTextLines[i][j+2]-48]+=1
							TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
							helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "r"
						}

					}else {

						helpArray[fileTextLines[i][j+2]-48]+=1
					}
				//}else if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] == "r" {// if r exists////modify

				}else if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] != "" {

					if helpArray[fileTextLines[i][j+2]-48] == 0 {

						TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
						helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "r"
						result = result + tempLines[i][j:j+1] + "l(" + tempLines[i][j+2:j+6]
						result = result + tempLines[i][j:j+6]
						tempLines[i] = tempLines[i][:j] + "*" +tempLines[i][j+1:]
						golog.Info("tempLines[i]**  : ",tempLines[i])
						//result = result + "ul" + tempLines[i][j+1:j+6]

					}else {

						helpArray[fileTextLines[i][j+2]-48]+=1
						TransOrderLockNums[fileTextLines[i][j+2]-48] +=1
						helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "r"
					}
				}
				//else if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] == "w"{// if w exists
				//
				//	if helpArray[j+2] == 0 {
				//
				//		result = result + tempLines[i][j:j+1] + "l(" + tempLines[i][j+1:j+6]
				//		result = result + tempLines[i][j:j+6]
				//		result = result + "ul" + tempLines[i][j:j+6]
				//
				//	}else {
				//		helpArray[j+2]+=1
				//		helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "r"
				//	}
				//	//TransOrderLockNums[fileTextLines[i][j+2]-48]+=1
				//	//helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] += "w"
				//}

			}else{
				step = 4
			}

			golog.Info("tempLines[i]& : ",tempLines[i])
			golog.Info("helpArray* : ",helpArray)

			//
			//golog.Info("j+2 : ",j+2)
			//golog.Info(len(fileTextLines[i]))
			//golog.Info("fileTextLines[i][j+2]-48 : ",fileTextLines[i][j+2]-48)
			//golog.Info("v : ", strconv.Itoa(int(fileTextLines[i][j])))
			//golog.Info("fileTextLines[i][j]-48 : ",fileTextLines[i][j]-48)
			//golog.Info("fileTextLines[i][j] == 'w' : ",fileTextLines[i][j] == 'w')
			//golog.Info("fileTextLines[i][j] == 'r' : ",fileTextLines[i][j] == 'r')
			//golog.Info("fileTextLines[i][j] == 'a' : ",fileTextLines[i][j] == 'a')
			//golog.Info("fileTextLines[i][j] == 'c' : ",fileTextLines[i][j] == 'c')
			//golog.Info("TransOrderNums[fileTextLines[i][j+2]-48]* : ",TransOrderNums[fileTextLines[i][j+2]-48])


			//if fileTextLines[i][j] == 'w' || fileTextLines[i][j] == 'r' {
			//
			//}
			l:= fileTextLines[i][j+2]-48
			golog.Info("l : ",l," , ",reflect.TypeOf(l))

			golog.Info("TransOrderLockNums[l] == TransOrderNums[l] : ",TransOrderLockNums[l] == TransOrderNums[l])
			if TransOrderLockNums[l] == TransOrderNums[l] {
				//golog.Info("len(tempLines[i]) : ",len(tempLines[i]))
				//for m := 0 ; m < TransOrderNums[l] && len(tempLines[i])>0 ; m++ {

				golog.Info("*("+strconv.Itoa(int(l)))
				t := strings.Index(tempLines[i], "*("+strconv.Itoa(int(l)))
				golog.Info("t *: ", t)

				for true  {
					golog.Info("********")
					golog.Info("tempLines[i] before : ",tempLines[i])
					//if tempLines[i][h+2] != 'a' && tempLines[i][h+2] != 'c' {
					//
					//}
					result = result + "ul" + tempLines[i][t+1:t+6]
					tempLines[i] = tempLines[i][:t] + tempLines[i][t+6:]
					golog.Info("tempLines[i] after  : ",tempLines[i])
					golog.Info("res : ",result)
					//if t != 2 && t != -1{
					//
					//	h:=2
					//
					//	for  h < t && tempLines[i][h+2] != 'a' && tempLines[i][h+2] != 'c'{
					//		golog.Info("h : ",h)
					//
					//		golog.Info("tempLines[i][h]-48 : ", tempLines[i][h]-48, " , ", reflect.TypeOf(tempLines[i][h]-48))
					//		golog.Info("tempLines[i][h+2]-118 : ", tempLines[i][h+2]-118, " , ", reflect.TypeOf(tempLines[i][h+2]-118))
					//
					//		if helpTable[tempLines[i][h+2]-118][tempLines[i][h]-48] == "" {
					//
					//			TransOrderLockNums[tempLines[i][h]-48]+=1
					//		}
					//
					//		if TransOrderLockNums[tempLines[i][h]-48] == TransOrderNums[tempLines[i][h]-48]  && TransOrderLockNums[tempLines[i][h]-48]>0{//////////////
					//
					//			result = result + tempLines[i][h-2:h-1] + "l(" + tempLines[i][h:h+4]
					//			result = result + tempLines[i][h-2:h+4]
					//			result = result + "ul" + tempLines[i][h-1:h+4]
					//
					//			golog.Info("helpTable before : ",helpTable)
					//			golog.Info("result ; ",result)
					//
					//
					//			golog.Info("len(helpTable[tempLines[i][h+2]-118][l]) : ",len(helpTable[tempLines[i][h+2]-118][l]))
					//			if len(helpTable[tempLines[i][h+2]-118][l]) <= 1 {
					//
					//				helpTable[tempLines[i][h+2]-118][l] = "" //not sure
					//
					//			}else{
					//
					//				golog.Info("tempLines[i][h-2:h-1] : ",tempLines[i][h-2:h-1] , " , ",reflect.TypeOf(tempLines[i][h-2:h-1]))
					//				index := strings.Index(helpTable[tempLines[i][h+2]-118][l], tempLines[i][h-2:h-1])
					//				if index != -1 {
					//					helpTable[tempLines[i][h+2]-118][l] = helpTable[tempLines[i][h+2]-118][l][:index]+helpTable[tempLines[i][h+2]-118][l][index+1:]
					//				}
					//			}
					//
					//
					//			golog.Info("helpTable after : ",helpTable)
					//
					//
					//			TransOrderLockNums[tempLines[i][h]-48]-=1 //not sure
					//			TransOrderNums[tempLines[i][h]-48]-=1//not sure
					//
					//			tempLines[i] = tempLines[i][:h-2] + tempLines[i][h+4:]
					//
					//		}
					//
					//
					//
					//		if tempLines[i][h-2]  == 'w' || tempLines[i][h-2]  == 'r'{
					//
					//			h += 6
					//
					//		}else{
					//
					//			h += 4
					//
					//		}
					//
					//	}
					//
					//}else if t == 2 {
					//
					//	result = result + tempLines[i][t-2:t-1] + "l(" + tempLines[i][t:t+4]
					//	result = result + tempLines[i][t-2:t+4]
					//	result = result + "ul" + tempLines[i][t-1:t+4]
					//
					//
					//	golog.Info("helpTable before* : ",helpTable)
					//	golog.Info("result* ; ",result)
					//
					//
					//	golog.Info("len(helpTable[tempLines[i][h+2]-118][l]) : ",len(helpTable[tempLines[i][t+2]-118][l]))
					//	if len(helpTable[tempLines[i][t+2]-118][l]) <= 1 {
					//
					//		helpTable[tempLines[i][t+2]-118][l] = "" //not sure
					//
					//	}else{
					//
					//		golog.Info("tempLines[i][h-2:h-1] : ",tempLines[i][t-2:t-1] , " , ",reflect.TypeOf(tempLines[i][t-2:t-1]))
					//		index := strings.Index(helpTable[tempLines[i][t+2]-118][l], tempLines[i][t-2:t-1])
					//		if index != -1 {
					//			helpTable[tempLines[i][t+2]-118][l] = helpTable[tempLines[i][t+2]-118][l][:index]+helpTable[tempLines[i][t+2]-118][l][index+1:]
					//		}
					//	}
					//
					//
					//	golog.Info("helpTable after* : ",helpTable)
					//
					//
					//
					//	golog.Info("tempLines[i][t-2]-118 : ", tempLines[i][t-2]-118, " , ", reflect.TypeOf(tempLines[i][t-2]-118))
					//	tempLines[i] = tempLines[i][:t-2] + tempLines[i][t+4:]
					//
					//}

					t := strings.Index(tempLines[i], "*("+strconv.Itoa(int(l)))
					golog.Info("t : ", t)
					if t == -1{
						break
					}
				}

				end := strings.Index(tempLines[i], "("+strconv.Itoa(int(l))+")")
				golog.Info("end : ",end)
				if end != -1{
					result = result + tempLines[i][end-1:end+3]
					tempLines[i] = tempLines[i][:end-1] + tempLines[i][end+3:]
				}

			}

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