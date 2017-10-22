package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type dataCway struct {
	Road_ID      int
	Start_Chain  float32
	End_Chain    float32
	Admin_Unit   int
	Terrain_Type string
	Cway_Type    string
	Cway_Width   int
	Embakment    string
	Submergent   string
	Right_of_Way string
	Survey_Date  string
	Notes        string
	File_Name    string
}

type dataDsin struct {
	Road_ID      int
	Start_Chain  float32
	End_Chain    float32
	Admin_Unit   int
	WSC_Extent   string
	WSC_Severity string
	EB_Extent    string
	EB_Severity  string
	RV_Extent    string
	RV_Severity  string
	PH_Extent    string
	PH_Severity  string
	Survey_Date  string
	File_Name    string
}

type dataRoug struct {
	Road_ID        int
	Start_Chain    float32
	End_Chain      float32
	Admin_Unit     int
	IRI_Average    string
	Survey_Date    string
	Year_of_Survey string
	Notes          string
	File_Name      string
}

type dataOutput struct { // TODO: need some changes to this struct fields
	dataCway
	dataDsin
	dataRoug
	Road_ID      int
	Start_Chain  float32
	End_Chain    float32
	Admin_Unit   int
	Cway_Type    string
	Cway_Width   int
	WSC_Extent   string
	WSC_Severity string
	EB_Extent    string
	EB_Severity  string
	RV_Extent    string
	RV_Severity  string
	PH_Extent    string
}

func readCway() []dataCway {
	fileCWAY, err := os.Open("data/CWAY.csv")
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
	fileDSIN, err := os.Open("data/DSIN.csv")
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
	fileROUG, err := os.Open("data/ROUG.csv")
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
	for _, c := range cway {
		fmt.Println(c)
	}
	//readDsin()
	//readRoug()

	// merge them into single dataset, get only the necessary fields from each of the structs

	// for each Road_ID

	//// extract the start and end chain values into []float32

	//// pass it to algo func

	//// to the resulting []chain set, join with the original set for the given Road_ID, to bring in the rest of the fields

	//// merge this resultset into the final output result set

	// save the final output set output file
}
