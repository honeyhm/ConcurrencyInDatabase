package main

import (
	"bufio"
	"github.com/kataras/golog"
	"log"
	"os"
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


	readFile.Close()

	// make and open output file for Conservative2PL algorithm
	fo, err := os.Create("C:\\Users\\ASUS\\Desktop\\Conservative2PL.txt")
	if err != nil {
		panic(err)
	}

	// loop for iterating throw all schedules
	for i := 0; i < len(fileTextLines); i++ {

		tempLines := fileTextLines[i]
		tempLines2 := tempLines

		//var tempLines2 []uint8
		//copy(tempLines2[:], tempLines[:])
		//
		//for m :=0 ; m< len(tempLines) ; m++ {
		//	tempLines2 += tempLines[m]
		//}
		//tempLines2 := fileTextLines[i]
		//
		//string variable for holding result
		result := ""

		//number of each transaction orders in each schedule
		TransOrderNums := make([]int, 8)
		//number of locks given by each transaction
		TransOrderLockNums := make([]int, 8)

		LockComplete := make([]bool, 8)
		//abortArr := make([]int, 0)


		for n := 1; n < 8; n++ {

			TransOrderNums[n] = strings.Count(tempLines, strconv.Itoa(n)) - 1
			if TransOrderNums[n] < 0 {
				TransOrderNums[n] += 1
			}

		}

		golog.Info("TransOrderNums : ", TransOrderNums)

		// a table which has transactions as columns and variables(v-z) as rows
		helpTable := make([][]string, 5) // making 5 rows of variables
		for i := range helpTable {
			helpTable[i] = make([]string, 8) //making 7 (ignoring index 0) columns for each row
		}


		help:=1

		for help > 0 {

			golog.Info("*****************************************")

			abortArr := make([]int, 0)

			golog.Info("abortArr * : ",abortArr)
			step := 2
			for len(tempLines2) > 1 {

				golog.Info("step : ", step)
				if step > len(tempLines2){/////??
					break
				}

				index := step
				t:= 2
				//t:= index

				if LockComplete[tempLines[index]-48] == false {
					golog.Info("1")

					for index+4 < len(tempLines) && index != -1{

						index = strings.Index(tempLines, strconv.Itoa(int(tempLines[2]-48)))

						golog.Info("index : ",index)
						if index == -1 {
							break
						}


						golog.Info("string(tempLines[index-2]) : ",string(tempLines[index-2]))
						if tempLines[index-2] == 'w' {

							flag1 := 0
							if helpTable[tempLines[index+2]-118][tempLines[index]-48] == "" { //lock

								for k := 1; k < 8; k++ {
									if helpTable[tempLines[index+2]-118][k] != "" {
										flag1 = 1
										break
									}
								}

								if flag1 == 0 {

									TransOrderLockNums[tempLines[index]-48] += 1
									helpTable[tempLines[index+2]-118][tempLines[index]-48] = "w"
								}

							} else if !strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "w") && strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "r") { // lock upgrade

								for k := 1; k < 8; k++ {
									if helpTable[tempLines[index+2]-118][k] != "" {
										flag1 += 1
									}
								}

								if flag1 == 1 {

									TransOrderLockNums[tempLines[index]-48] += 1
									helpTable[tempLines[index+2]-118][tempLines[index]-48] += "w"
								}

							} else if strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "w") {

								TransOrderLockNums[tempLines[index]-48] += 1
								helpTable[tempLines[index+2]-118][tempLines[index]-48] += "w"
							}

						} else if tempLines[index-2] == 'r' {

							flag1 := 0
							if helpTable[tempLines[index+2]-118][tempLines[index]-48] == "" {

								for k := 1; k < 8; k++ {
									if helpTable[tempLines[index+2]-118][k] == "w" {
										flag1 = 1
										break
									}
								}

								if flag1 == 0 {

									TransOrderLockNums[tempLines[index]-48] += 1
									helpTable[tempLines[index+2]-118][tempLines[index]-48] = "r"
								}

							} else if helpTable[tempLines[index+2]-118][tempLines[index]-48] != "" {

								TransOrderLockNums[tempLines[index]-48] += 1
								helpTable[tempLines[index+2]-118][tempLines[index]-48] += "r"
							}
						}

						if tempLines[index-2] != 'a' && tempLines[index-2] != 'c'{
							tempLines = tempLines[:index-2] + tempLines[index+4:]
						}else {
							tempLines = tempLines[:index-2] + tempLines[index+2:]
						}

					}

				}



				if TransOrderLockNums[tempLines2[t]-48] == TransOrderNums[tempLines2[t]-48] {
					golog.Info("2")
					LockComplete[tempLines2[t]-48] = true
				}

				golog.Info("tempLines : ",tempLines)
				golog.Info("tempLines2 : ",tempLines2)
				golog.Info("tempLines2[t]-48 : ",tempLines2[t]-48)

				if LockComplete[tempLines2[t]-48] == true {
					golog.Info("3")
					//golog.Info("string(fileTextLines[i][t-2]) : ",string(fileTextLines[i][t-2]))

					if tempLines2[t-2] == 'w' || tempLines2[t-2] == 'r' {

						result += tempLines2[t-2:t-1] + "l(" + tempLines2[t:t+4]
						result += tempLines2[t-2 : t+4]
						result += "ul(" + tempLines2[t:t+4]

						//golog.Info(" helpTable : ",helpTable)
						//golog.Info("TransOrderLockNums : ",TransOrderLockNums)
						//golog.Info("TransOrderNums : ",TransOrderNums)
						//golog.Info("helpTable[tempLines2[t+2]-118][tempLines[t]-48] : ",helpTable[tempLines2[t+2]-118][tempLines[t]-48])
						//golog.Info("tempLines2[t+2]-118 : ",tempLines2[t+2]-118)
						//golog.Info("tempLines[t]-48 : ",tempLines[t]-48)
						//golog.Info("tempLines2[t-2:t-1] : ",tempLines2[t-2:t-1])
						//x:=strings.Index(helpTable[tempLines2[t+2]-118][tempLines[t]-48], tempLines2[t-2:t-1])
						//golog.Info("len(helpTable[tempLines2[t+2]-118][tempLines[t]-48]) : ",len(helpTable[tempLines2[t+2]-118][tempLines[t]-48]))
						//if x < len(helpTable[tempLines2[t+2]-118][tempLines[t]-48]){
							//helpTable[tempLines2[t+2]-118][tempLines[t]-48] = helpTable[tempLines2[t+2]-118][tempLines[t]-48][:x]+helpTable[tempLines2[t+2]-118][tempLines[t]-48][x+1:]
						//}
						helpTable[tempLines2[t+2]-118][tempLines[t]-48] = ""

						tempLines2 = tempLines2[:t-2] + tempLines2[t+4:]

					} else if tempLines2[t-2] == 'a' || tempLines2[t-2] == 'c'{

						result += tempLines2[t-2 : t+2]
						tempLines2 = tempLines2[:t-2] + tempLines2[t+2:]
					}

					//golog.Info("tempLines : ",tempLines)
					//golog.Info("tempLines2 : ",tempLines2)
					//golog.Info("result : ", result)

				}else{

					flagTemp:=0
					for l:=0 ; l<len(abortArr) ;l++{
						//golog.Info("int(tempLines2[t]-48) : ",int(tempLines2[t]-48))
						if int(tempLines2[t]-48) == abortArr[l]{
							flagTemp = 1
						}
					}

					if flagTemp == 0 {
						abortArr = append(abortArr, int(tempLines2[t]-48))
					}

/////////////////////////////////////////////////////
					if tempLines2[t-2] == 'w' || tempLines2[t-2] == 'r' {

						tempLines2 = tempLines2[:t-2] + tempLines2[t+4:]

					} else if tempLines2[t-2] == 'a' || tempLines2[t-2] == 'c'{

						tempLines2 = tempLines2[:t-2] + tempLines2[t+2:]

					}

					golog.Info("abortArr : ",abortArr)
					golog.Info("result *: ", result)


				}

			}


			help = len(abortArr)
			//help = 0

			golog.Info("tempLines* : ",tempLines)
			golog.Info("tempLines2* : ",tempLines2)

			//tempLines = tempLines2

			temp:=0
			tempLines = ""
			golog.Info("fileTextLines[i] : ",fileTextLines[i])
			for a:=0 ; a+2<len(fileTextLines[i]) ; a+=temp  {
				golog.Info("777")
				golog.Info("a : ",a)
				//golog.Info("temp : ",temp)
				updFlag := false
				for b:=0 ; b < len(abortArr) ; b++ {

					golog.Info("abortArr : ",abortArr)
					golog.Info("result **: ", result)


					//golog.Info("fileTextLines[i][a+2]-48 : ",fileTextLines[i][a+2]-48)

					if int(fileTextLines[i][a+2]-48) == abortArr[b]{
						golog.Info("888")
						updFlag = true

						if fileTextLines[i][a] == 'w' || fileTextLines[i][a] == 'r'{
							tempLines += fileTextLines[i][a:a+6]
							temp=6
						}else{
							tempLines += fileTextLines[i][a:a+4]
							temp=4
						}

						break
					}

					//if fileTextLines[i][a] == 'w' || fileTextLines[i][a] == 'r'{
					//	//tempLines += fileTextLines[i][a:a+6]
					//	temp=6
					//}else{
					//	//tempLines += fileTextLines[i][a:a+4]
					//	temp=4
					//}
				}

				if updFlag == false{

					if fileTextLines[i][a] == 'w' || fileTextLines[i][a] == 'r'{
						temp=6
					}else{
						temp=4
					}

				}

			}


			if len(abortArr)==1 { // not sure if it is needed

				golog.Info("999")
				for c:=0 ; c+2<len(fileTextLines[i]) ; c+=temp  {
					if fileTextLines[i][c] == 'w' || fileTextLines[i][c] == 'r'{
						result += fileTextLines[i][c:c+6]
						temp=6
					}else{
						result += fileTextLines[i][c:c+4]
						temp=4
					}
				}
				help =0
			}

			tempLines2=tempLines

			//golog.Info("tempLines* *: ",tempLines)
			//golog.Info("tempLines2* *: ",tempLines2)

		}

		if _, err := fo.Write([]byte(result + "\r\n")); err != nil {
			panic(err)
		}

		//step := 2
		//for len(tempLines) > 1 {
		//
		//	golog.Info("step : ", step)
		//	if step > len(tempLines){/////??
		//		break
		//	}
		//
		//	index := step
		//	t:= index
		//
		//	if LockComplete[tempLines[index]-48] == false {
		//		golog.Info("1")
		//
		//		for index+4 < len(tempLines) && index != -1{
		//
		//			//if tempLines[index-2] != 'a' && tempLines[index-2] != 'c'{
		//			//	tempLines = tempLines[:index-2] + tempLines[index+4:]
		//			//}else{
		//			//	tempLines = tempLines[:index-2] + tempLines[index+2:]
		//			//
		//			index = strings.Index(tempLines, strconv.Itoa(int(tempLines[2]-48)))
		//
		//			golog.Info("index : ",index)
		//			if index == -1 {
		//				break
		//			}
		//
		//			golog.Info("tempLines[index-2] : ",tempLines[index-2])
		//			golog.Info(string(tempLines[index-2]))
		//			if tempLines[index-2] == 'w' {
		//
		//				flag1 := 0
		//				if helpTable[tempLines[index+2]-118][tempLines[index]-48] == "" { //lock
		//
		//					for k := 1; k < 8; k++ {
		//						if helpTable[tempLines[index+2]-118][k] != "" {
		//							flag1 = 1
		//							break
		//						}
		//					}
		//
		//					if flag1 == 0 {
		//
		//						TransOrderLockNums[tempLines[index]-48] += 1
		//						helpTable[tempLines[index+2]-118][tempLines[index]-48] = "w"
		//					}
		//
		//				} else if !strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "w") && strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "r") { // lock upgrade
		//
		//					for k := 1; k < 8; k++ {
		//						if helpTable[tempLines[index+2]-118][k] != "" {
		//							flag1 += 1
		//						}
		//					}
		//
		//					if flag1 == 1 {
		//
		//						TransOrderLockNums[tempLines[index]-48] += 1
		//						helpTable[tempLines[index+2]-118][tempLines[index]-48] += "w"
		//					}
		//
		//				} else if strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "w") {
		//
		//					TransOrderLockNums[tempLines[index]-48] += 1
		//					helpTable[tempLines[index+2]-118][tempLines[index]-48] += "w"
		//				}
		//
		//			} else if tempLines[index-2] == 'r' {
		//
		//				flag1 := 0
		//				if helpTable[tempLines[index+2]-118][tempLines[index]-48] == "" {
		//
		//					for k := 1; k < 8; k++ {
		//						if helpTable[tempLines[index+2]-118][k] == "w" {
		//							flag1 = 1
		//							break
		//						}
		//					}
		//
		//					if flag1 == 0 {
		//
		//						TransOrderLockNums[tempLines[index]-48] += 1
		//						helpTable[tempLines[index+2]-118][tempLines[index]-48] = "r"
		//					}
		//
		//				} else if helpTable[tempLines[index+2]-118][tempLines[index]-48] != "" {
		//
		//					TransOrderLockNums[tempLines[index]-48] += 1
		//					helpTable[tempLines[index+2]-118][tempLines[index]-48] += "r"
		//				}
		//			}
		//
		//			if tempLines[index-2] != 'a' && tempLines[index-2] != 'c'{
		//				tempLines = tempLines[:index-2] + tempLines[index+4:]
		//			}else {
		//				tempLines = tempLines[:index-2] + tempLines[index+2:]
		//			}
		//
		//		}
		//
		//	}
		//
		//
		//	if TransOrderLockNums[tempLines[t]-48] == TransOrderNums[tempLines[t]-48] {
		//		golog.Info("2")
		//		LockComplete[tempLines[t]-48] = true
		//	}
		//
		//
		//	if LockComplete[tempLines[t]-48] == true {
		//		golog.Info("3")
		//
		//		if tempLines[t-2]-118 == 'w' || tempLines[t-2]-118 == 'r' {
		//			//step += 6
		//			result += tempLines[t-2 : t+4]
		//			tempLines = tempLines[:index-2] + tempLines[index+4:]
		//
		//		} else {
		//			//step += 4
		//			result += tempLines[t-2 : t+2]
		//			tempLines = tempLines[:index-2] + tempLines[index+2:]
		//		}
		//
		//		golog.Info("result : ", result)
		//
		//	}else{
		//		abortArr = append(abortArr, int(tempLines[t]-48))
		//		golog.Info("abortArr : ",abortArr)
		//		//tempLines =  fileTextLines[i]
		//	}
		//
		//
		//}
		//
		//
		//if _, err := fo.Write([]byte(result + "\r\n")); err != nil {
		//	panic(err)
		//}
	}

	//close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

}















//package main
//
//import (
//	"bufio"
//	"github.com/kataras/golog"
//	"log"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//
//	//read , modify with spaces and separating file line by line
//	readFile, err := os.Open("C:\\Users\\ASUS\\Desktop\\test.txt")
//	if err != nil {
//		log.Fatalf("failed to open file: %s", err)
//	}
//
//	fileScanner := bufio.NewScanner(readFile)
//	fileScanner.Split(bufio.ScanLines)
//	var fileTextLines []string
//
//	for fileScanner.Scan() {
//		fileTextLines = append(fileTextLines, fileScanner.Text())
//	}
//
//	//var rr []string
//	//tempLines := fileTextLines
//	//copy(tempLines, tempLines)
//
//	readFile.Close()
//
//	// make and open output file for Conservative2PL algorithm
//	fo, err := os.Create("C:\\Users\\ASUS\\Desktop\\Conservative2PL.txt")
//	if err != nil {
//		panic(err)
//	}
//
//	// loop for iterating throw all schedules
//	//for i := 0; i < len(tempLines); i++ {
//	for i := 0; i < len(fileTextLines); i++ {
//
//		tempLines := fileTextLines[i]
//
//		//string variable for holding result
//		result := ""
//
//		//number of each transaction orders in each schedule
//		TransOrderNums := make([]int, 8)
//		//number of locks given by each transaction
//		TransOrderLockNums := make([]int, 8)
//
//		LockComplete := make([]bool, 8)
//
//
//		for n := 1; n < 8; n++ {
//
//			TransOrderNums[n] = strings.Count(tempLines, strconv.Itoa(n)) - 1
//			if TransOrderNums[n] < 0 {
//				TransOrderNums[n] += 1
//			}
//
//		}
//
//		golog.Info("TransOrderNums : ", TransOrderNums)
//
//		// a table which has transactions as columns and variables(v-z) as rows
//		helpTable := make([][]string, 5) // making 5 rows of variables
//		for i := range helpTable {
//			helpTable[i] = make([]string, 8) //making 7 (ignoring index 0) columns for each row
//		}
//
//
//
//
//
//
//
//		step := 2
//		for len(tempLines) > 1 {
//
//			golog.Info("step : ", step)
//			if step > len(tempLines){/////??
//				break
//			}
//
//			//index := strings.Index(tempLines, strconv.Itoa(int(tempLines[step]-48)))
//			//
//			//t:=index
//			//golog.Info("t : ",t)
//			index := step
//			t:= index
//
//			if LockComplete[tempLines[index]-48] == false {
//				golog.Info("1")
//
//				for index+4 < len(tempLines) && index != -1{
//
//					if tempLines[index-2] != 'a' && tempLines[index-2] != 'c'{
//						tempLines = tempLines[:index-2] + tempLines[index+4:]
//					}else{
//						tempLines = tempLines[:index-2] + tempLines[index+2:]
//					}
//
//
//					index = strings.Index(tempLines, strconv.Itoa(int(tempLines[index]-48)))
//
//					golog.Info("index : ",index)
//					if index == -1 {
//						break
//					}
//					golog.Info("tempLines[index-2] : ",tempLines[index-2])
//					golog.Info(string(tempLines[index-2]))
//					if tempLines[index-2] == 'w' {
//
//						flag1 := 0
//						if helpTable[tempLines[index+2]-118][tempLines[index]-48] == "" { //lock
//
//							for k := 1; k < 8; k++ {
//								if helpTable[tempLines[index+2]-118][k] != "" {
//									flag1 = 1
//									break
//								}
//							}
//
//							if flag1 == 0 {
//
//								TransOrderLockNums[tempLines[index]-48] += 1
//								helpTable[tempLines[index+2]-118][tempLines[index]-48] = "w"
//							}
//
//						} else if !strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "w") && strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "r") { // lock upgrade
//
//							for k := 1; k < 8; k++ {
//								if helpTable[tempLines[index+2]-118][k] != "" {
//									flag1 += 1
//								}
//							}
//
//							if flag1 == 1 {
//
//								TransOrderLockNums[tempLines[index]-48] += 1
//								helpTable[tempLines[index+2]-118][tempLines[index]-48] += "w"
//							}
//
//						} else if strings.Contains(helpTable[tempLines[index+2]-118][tempLines[index]-48], "w") {
//
//							TransOrderLockNums[tempLines[index]-48] += 1
//							helpTable[tempLines[index+2]-118][tempLines[index]-48] += "w"
//						}
//
//					} else if tempLines[index-2] == 'r' {
//
//						flag1 := 0
//						if helpTable[tempLines[index+2]-118][tempLines[index]-48] == "" {
//
//							for k := 1; k < 8; k++ {
//								if helpTable[tempLines[index+2]-118][k] == "w" {
//									flag1 = 1
//									break
//								}
//							}
//
//							if flag1 == 0 {
//
//								TransOrderLockNums[tempLines[index]-48] += 1
//								helpTable[tempLines[index+2]-118][tempLines[index]-48] = "r"
//							}
//
//						} else if helpTable[tempLines[index+2]-118][tempLines[index]-48] != "" {
//
//							TransOrderLockNums[tempLines[index]-48] += 1
//							helpTable[tempLines[index+2]-118][tempLines[index]-48] += "r"
//						}
//					}
//				}
//
//			}
//
//			if TransOrderLockNums[tempLines[t]-48] == TransOrderNums[tempLines[t]-48] {
//				golog.Info("2")
//				LockComplete[tempLines[t]-48] = true
//			}
//
//
//			if LockComplete[tempLines[t]-48] == true {
//				golog.Info("3")
//
//				if fileTextLines[i][t-2]-118 == 'w' || fileTextLines[i][t-2]-118 == 'r' {
//
//					step += 6
//					result += fileTextLines[i][t-2 : t+4]
//				} else {
//
//					step += 4
//					result += fileTextLines[i][t-2 : t+2]
//				}
//
//				golog.Info("result : ", result)
//
//			}
//
//
//		}
//
//
//		if _, err := fo.Write([]byte(result + "\r\n")); err != nil {
//			panic(err)
//		}
//	}
//
//	//close fo on exit and check for its returned error
//	defer func() {
//		if err := fo.Close(); err != nil {
//			panic(err)
//		}
//	}()
//
//}
