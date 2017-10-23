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
		t := dataOutput{Common:c.Common}
		//t.OutputCway = c.OutputCway
		temp = append(temp, t)
	}
	for _, c := range dsin {
		t := dataOutput{Common:c.Common}
		temp = append(temp, t)
	}
	for _, c := range roug {
		t := dataOutput{Common:c.Common}
		temp = append(temp, t)
	}


	for _, c := range temp{
		fmt.Println(c)
	}


	// for each Road_ID

	//// extract the start and end chain values into []float32

	//// pass it to algo func

	//// to the resulting []chain set, join with the original set for the given Road_ID, to bring in the rest of the fields

	//// merge this resultset into the final output result set

	// save the final output set output file
}
