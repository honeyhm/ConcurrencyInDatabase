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

	// make and open output file for Conservative2PL algorithm
	fo, err := os.Create("C:\\Users\\ASUS\\Desktop\\Conservative2PL.txt")
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

		for n:=1 ; n<8 ; n++  {

			TransOrderNums[n] = strings.Count(fileTextLines[i], strconv.Itoa(n)) -1
			if TransOrderNums[n] < 0{
				TransOrderNums[n]+=1
			}

		}
		golog.Info("TransOrderNums : ",TransOrderNums)


		////array for counting lost transaction members
		//helpArray := make([]int, 8)
		//golog.Info("helpArray : ",helpArray)
		//// empty helpTable for new schedule in each line
		//for n:=0 ; n<5 ; n++  {
		//	for m:=0 ; m<8 ; m++  {
		//		helpTable[n][m]=""
		//	}
		//}


		// a table which has transactions as columns and variables(v-z) as rows
		helpTable := make([][]string, 5)// making 5 rows of variables
		for i := range helpTable {
			helpTable[i] = make([]string, 8)//making 7 (ignoring index 0) columns for each row
		}
		golog.Info("helpTable : ",helpTable)

		step:=0
		// loop for iterating throw each schedule content
		for j:=0 ; j<len(fileTextLines[i]) ; j+=step  {

			if j+4 >= len(fileTextLines[i]){
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

						TransOrderLockNums[fileTextLines[i][j+2]-48]+=1
						helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"

					}
					//else{
					//	helpArray[fileTextLines[i][j+2]-48] += 1
					//	//nothing yet
					//}

				}else if helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] == "r" { // lock upgrade

					for k:=1 ; k<8 ; k++  {
						if helpTable[fileTextLines[i][j+4]-118][k] != ""{
							flag1 += 1
						}
					}

					if flag1 == 1{

						TransOrderLockNums[fileTextLines[i][j+2]-48]+=1
						helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"

					}else{
						//helpArray[fileTextLines[i][j+2]-48] += 1
						//nothing yet
					}

				}else{// if w exists
					TransOrderLockNums[fileTextLines[i][j+2]-48]+=1
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

						TransOrderLockNums[fileTextLines[i][j+2]-48]+=1
						helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "r"

					}else{
						//helpArray[fileTextLines[i][j+2]-48] += 1
						//nothing yet
					}

				}else{// if r or w exists
					TransOrderLockNums[fileTextLines[i][j+2]-48]+=1
				}

			}else{
				step = 4
			}



			golog.Info("TransOrderLockNums : *******",TransOrderLockNums)
			golog.Info("TransOrderLockNums[fileTextLines[i][j+2]-48] : ",TransOrderLockNums[fileTextLines[i][j+2]-48])
			golog.Info("TransOrderNums[fileTextLines[i][j+2]-48] : ",TransOrderNums[fileTextLines[i][j+2]-48])

			l:= fileTextLines[i][j+2]-48
			golog.Info("l : ",l," , ",reflect.TypeOf(l))

			//if TransOrderLockNums[fileTextLines[i][j+2]-48] == TransOrderNums[fileTextLines[i][j+2]-48]{
			if TransOrderLockNums[l] == TransOrderNums[l] {

				//for m:=0 ; m<TransOrderNums[fileTextLines[i][j+2]-48] ; m++  {
				for m := 0; m < TransOrderNums[l] && len(tempLines[i])>0 ; m++ {

					//t := strings.Index(tempLines[i], strconv.Itoa(int(fileTextLines[i][j+2]-48)))
					t := strings.Index(tempLines[i], strconv.Itoa(int(l)))
					golog.Info("t : ", t)

					if t != 2 && t != -1{

						h:=2
						//golog.Info("tempLines[i][h+2] != 'a' && tempLines[i][h+2] != 'c' : ",tempLines[i][h+2] != 'a' && tempLines[i][h+2] != 'c')
						for  h < t && tempLines[i][h+2] != 'a' && tempLines[i][h+2] != 'c'{
							golog.Info("h : ",h)

							golog.Info("tempLines[i][h]-48 : ", tempLines[i][h]-48, " , ", reflect.TypeOf(tempLines[i][h]-48))
							golog.Info("tempLines[i][h+2]-118 : ", tempLines[i][h+2]-118, " , ", reflect.TypeOf(tempLines[i][h+2]-118))

							if helpTable[tempLines[i][h+2]-118][tempLines[i][h]-48] == "" {

								TransOrderLockNums[tempLines[i][h]-48]+=1

							}

							if TransOrderLockNums[tempLines[i][h]-48] == TransOrderNums[tempLines[i][h]-48]  && TransOrderLockNums[tempLines[i][h]-48]>0{

								result = result + tempLines[i][h-2:h-1] + "l(" + tempLines[i][h:h+4]
								result = result + tempLines[i][h-2:h+4]
								result = result + "ul" + tempLines[i][h-1:h+4]

								golog.Info("helpTable before : ",helpTable)
								golog.Info("result ; ",result)


								helpTable[tempLines[i][h+2]-118][l] = "" //not sure

								golog.Info("helpTable after : ",helpTable)


								TransOrderLockNums[tempLines[i][h]-48]-=1 //not sure
								TransOrderNums[tempLines[i][h]-48]-=1//not sure

								tempLines[i] = tempLines[i][:h-2] + tempLines[i][h+4:]


							}



							if tempLines[i][h-2]  == 'w' || tempLines[i][h-2]  == 'r'{


								h += 6

							}else{

								h += 4

							}


						}

					}else if t == 2 {

						result = result + tempLines[i][t-2:t-1] + "l(" + tempLines[i][t:t+4]
						result = result + tempLines[i][t-2:t+4]
						result = result + "ul" + tempLines[i][t-1:t+4]

						golog.Info("result ;* ",result)

						helpTable[tempLines[i][t+2]-118][l] = "" //not sure
						golog.Info("tempLines[i][t-2]-118 : ", tempLines[i][t-2]-118, " , ", reflect.TypeOf(tempLines[i][t-2]-118))
						tempLines[i] = tempLines[i][:t-2] + tempLines[i][t+4:]

					}

				}

			}

				//for l:=1 ; l<8 ; l++ {
			//
			//	golog.Info("TransOrderLockNums[fileTextLines[i][j+2]-48] : ",TransOrderLockNums[fileTextLines[i][j+2]-48])
			//	golog.Info("TransOrderNums[fileTextLines[i][j+2]-48] : ",TransOrderNums[fileTextLines[i][j+2]-48])
			//
			//	//if TransOrderLockNums[fileTextLines[i][j+2]-48] == TransOrderNums[fileTextLines[i][j+2]-48]{
			//	if TransOrderLockNums[l] == TransOrderNums[l]{
			//
			//		//for m:=0 ; m<TransOrderNums[fileTextLines[i][j+2]-48] ; m++  {
			//		for m:=0 ; m<TransOrderNums[l] ; m++  {
			//
			//			//t := strings.Index(tempLines[i], strconv.Itoa(int(fileTextLines[i][j+2]-48)))
			//			t := strings.Index(tempLines[i], strconv.Itoa(l))
			//			golog.Info("t : ",t)
			//
			//
			//			if t!=2{
			//
			//				for h:=2 ; h<t ; h+=6  {
			//
			//					golog.Info("tempLines[i][h] : ",tempLines[i][h]," , ",reflect.TypeOf(tempLines[i][h]))
			//					golog.Info("tempLines[i][h]-48 : ",tempLines[i][h]-48," , ",reflect.TypeOf(tempLines[i][h]-48))
			//					if TransOrderLockNums[tempLines[i][h]-48] != TransOrderNums[tempLines[i][h]-48]{
			//
			//					}
			//
			//				}
			//
			//			}
			//
			//			if t == 2 {
			//
			//				result = result + tempLines[i][t-2:t-2] + "l(" + tempLines[i][t:t+4]
			//				result = result + tempLines[i][t-2:t+4]
			//				result = result + "ul" + tempLines[i][t-1:t+4]
			//				helpTable[tempLines[i][t-2]-118][l] = ""
			//				golog.Info("tempLines[i][t-2]-118 : ",tempLines[i][t-2]-118," , ",reflect.TypeOf(tempLines[i][t-2]-118))
			//				//helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = ""
			//				tempLines[i] = tempLines[i][:t-2] + tempLines[i][t+4:]
			//
			//			}
			//
			//
			//		}
			//
			//	}
			//
			//
			//}

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










//package main
//
//import (
//	"bufio"
//	"github.com/kataras/golog"
//	"log"
//	"os"
//	"reflect"
//)
//
//func main() {
//
//	helpTable := make([][]string, 5)
//	for i := range helpTable {
//		helpTable[i] = make([]string, 8)
//	}
//
//
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
//
//	tempLines := fileTextLines
//
//
//	readFile.Close()
//
//	//for _, eachline := range fileTextLines {
//	//	fmt.Println(eachline)
//	//}
//
//	// open output file
//	fo, err := os.Create("C:\\Users\\ASUS\\Desktop\\Conservative2PL.txt")
//	if err != nil {
//		panic(err)
//	}
//
//
//	step:=0
//	row:=""
//	col:=""
//
//	golog.Info("len(fileTextLines) : ",len(fileTextLines))
//	for i:=0 ; i<len(fileTextLines) ; i++ {
//
//		for n:=0 ; n<5 ; n++  { // empty helpTable for new schedule
//			for m:=0 ; m<8 ; m++  {
//				helpTable[n][m]=""
//			}
//		}
//
//		golog.Info("len(fileTextLines[i])  : ",len(fileTextLines[i]) )
//		for j:=0 ; j<len(fileTextLines[i]) ; j+=step  {
//
//			if j+4 >= len(fileTextLines[i]){
//				break
//			}
//
//			if fileTextLines[i][j]  == 'w' || fileTextLines[i][j]  == 'r'{
//
//				step = 6
//
//				row = string(fileTextLines[i][j+4])
//				golog.Info("row : ",row)
//
//				col = string(fileTextLines[i][j+2])
//				golog.Info("col : ",col)
//
//			}else{
//
//				step = 4
//				col = string(fileTextLines[i][j+2])
//			}
//
//
//			//r(4,x)w(1,z)w(7,z)r(7,x)w(4,w)r(7,z)w(6,x)r(5,v)w(4,v)
//			if fileTextLines[i][j] == 'w' {
//
//				golog.Info("fileTextLines[i][j] == 'w' : ",fileTextLines[i][j] == 'w')
//				helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "w"
//
//				golog.Info("helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] : ",helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48])
//
//
//			}else if fileTextLines[i][j] == 'r' {
//
//				golog.Info("fileTextLines[i][j] == 'r' : ",fileTextLines[i][j] == 'r')
//				helpTable[fileTextLines[i][j+4]-118][fileTextLines[i][j+2]-48] = "r"
//
//				golog.Info("fileTextLines[i][j+4]-118 : ",fileTextLines[i][j+4]-118," , ",reflect.TypeOf(fileTextLines[i][j+4]-118))
//
//				golog.Info("fileTextLines[i][j+2]-48 : ",fileTextLines[i][j+2]-48," , ",reflect.TypeOf(fileTextLines[i][j+2]-48))
//
//			}
//
//			//golog.Info(fileTextLines[0][5:12])
//
//			//tempLines[i] = tempLines[i][:2] + "***" + tempLines[i][5:]
//			tempLines[i] = tempLines[i][:2] + tempLines[i][5:]
//
//			//input[:index] + string(replacement) + input[index+1:]
//
//			//out := []rune(tempLines[i])
//			//out[7] = "*"
//			//x:= string(out)
//
//		}
//
//		golog.Info("helpTable line ",i," : ",helpTable)
//
//
//		if _, err := fo.Write([]byte(tempLines[i]+"\r\n")); err != nil {
//			panic(err)
//		}
//
//		//if _, err := fo.Write([]byte(fileTextLines[i]+"\r\n")); err != nil {
//		//	panic(err)
//		//}
//
//	}
//
//
//
//
//	//file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	//if err != nil {
//	//	log.Fatalf("failed creating file: %s", err)
//	//}
//	//for i:=0;i<len(fileTextLines) ; i++ {
//	//	golog.Info("hhh")
//	//
//	//	//datawriter := bufio.NewWriter(fo)
//	//	//
//	//	//for _, data := range fileTextLines {
//	//	//	_, _ = datawriter.WriteString(data + "\n")
//	//	//}
//	//	//
//	//	//datawriter.Flush()
//	//
//	//	fmt.Println(fileTextLines[i])
//	//	if _, err := fo.Write([]byte(fileTextLines[i]+"\r\n")); err != nil {
//	//		panic(err)
//	//	}
//	//	//if _, err := fo.WriteString("\r\n"); err != nil {
//	//	//	panic(err)
//	//	//}
//	//}
//
//
//	//close fo on exit and check for its returned error
//	defer func() {
//		if err := fo.Close(); err != nil {
//			panic(err)
//		}
//	}()
//
//	// an array with 5 rows and 2 columns
//	//var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
//	//var i, j int
//	//
//	///* output each array element's value */
//	//for  i = 0; i < 5; i++ {
//	//	for j = 0; j < 2; j++ {
//	//		fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
//	//	}
//	//}
//}
//
//
