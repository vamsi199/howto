package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type Common struct {
	Road_ID     int
	Start_Chain float32
	End_Chain   float32
	Admin_Unit  int
}
type OutputCway struct {
	Cway_Type  string
	Cway_Width int
}

type dataCway struct {
	Common
	Terrain_Type string
	OutputCway
	Embakment    string
	Submergent   string
	Right_of_Way string
	Survey_Date  string
	Notes        string
	File_Name    string
}

type OutputDsin struct {
	WSC_Extent   string
	WSC_Severity string
	EB_Extent    string
	EB_Severity  string
	RV_Extent    string
	RV_Severity  string
	PH_Extent    string
	PH_Severity  string
}

type dataDsin struct {
	Common
	OutputDsin
	Survey_Date string
	File_Name   string
}

type OutputRoug struct {
	IRI_Average string
}

type dataRoug struct {
	Common
	OutputRoug
	Survey_Date    string
	Year_of_Survey string
	Notes          string
	File_Name      string
}

type dataOutput struct { // TODO: need some changes to this struct fields
	Common
	OutputCway
	OutputDsin
	OutputRoug
	fileType string
}

func readCway() []dataCway {
	fileCWAY, err := os.Open("/Users/muly/go/src/github.com/muly/howto/golang/project/csv-processing/data/CWAY.csv")
	if err != nil {
		panic(err)
	}
	defer fileCWAY.Close()

	slice := []dataCway{}
	err = gocsv.Unmarshal(fileCWAY, &slice)
	if err != nil {
		panic(err)
	}

	return slice
}

func readDsin() []dataDsin {
	fileDSIN, err := os.Open("/Users/muly/go/src/github.com/muly/howto/golang/project/csv-processing/data/DSIN.csv")
	if err != nil {
		panic(err)
	}
	defer fileDSIN.Close()

	slice := []dataDsin{}
	err = gocsv.Unmarshal(fileDSIN, &slice)
	if err != nil {
		panic(err)
	}

	return slice
}

func readRoug() []dataRoug {
	fileROUG, err := os.Open("/Users/muly/go/src/github.com/muly/howto/golang/project/csv-processing/data/ROUG.csv")
	if err != nil {
		panic(err)
	}
	defer fileROUG.Close()

	slice := []dataRoug{}
	err = gocsv.Unmarshal(fileROUG, &slice)
	if err != nil {
		panic(err)
	}

	return slice
}

func main() {

	//read all the 3 files
	cway := readCway()
	dsin := readDsin()
	roug := readRoug()

	//for _, c := range cway {
	//	fmt.Println(c)
	//}
	//for _, c := range dsin {
	//	fmt.Println(c)
	//}
	//for _, c := range roug {
	//	fmt.Println(c)
	//}

	// merge them into single dataset, get only the necessary fields from each of the structs
	temp := make([]dataOutput, 0, len(cway)+len(dsin)+len(roug))
	for _, c := range cway {
		t := dataOutput{Common: c.Common}
		t.OutputCway = c.OutputCway
		t.fileType = "cway"
		temp = append(temp, t)
	}
	for _, c := range dsin {
		t := dataOutput{Common: c.Common}
		t.OutputDsin = c.OutputDsin
		t.fileType = "dsin"
		temp = append(temp, t)
	}
	for _, c := range roug {
		t := dataOutput{Common: c.Common}
		t.OutputRoug = c.OutputRoug
		t.fileType = "roug"
		temp = append(temp, t)
	}

	for _, c := range temp {
		fmt.Println(c)
	}

	out := []dataOutput{} //TODO: use make

	// for each Road_ID
	dataMap := map[int][]dataOutput{}
	for _, v := range temp {
		c, exists := dataMap[v.Road_ID]
		if exists {
			c = append(c, v)
		}
		dataMap[v.Road_ID] = c
	}
	for roadId, val := range dataMap {

		//// extract the start and end chain values into []float32
		oneD := make([]float32, 0, len(val)*2)
		for _, c := range val {
			oneD = append(oneD, c.Start_Chain, c.End_Chain)
		}

		//// pass it to algo func
		chain := algo(oneD)

		//// to the resulting []chain set, join with the original set for the given Road_ID, to bring in the rest of the fields
		for _, c := range chain {
			d := dataOutput{}
			for _, v := range val {
				if c.Start_Chain >= v.Start_Chain && c.End_Chain <= v.End_Chain {
					switch v.fileType {
					case "cway":
						d.OutputCway = v.OutputCway
					case "dsin":
						d.OutputDsin = v.OutputDsin
					case "roug":
						d.OutputRoug = v.OutputRoug
					}
//TODO: need to review, looks like there is bug in getting all the output sets
				}
				d.Admin_Unit = v.Admin_Unit

			}
			d.Start_Chain = c.Start_Chain
			d.End_Chain = c.End_Chain
			d.Road_ID = roadId

			out = append(out, d)

		}

		// save the final output set output file //TODO
		fmt.Println("Final Output:::::::::::::::\n")
		for _, o:= range out {
			fmt.Println(o)
		}


	}
}
