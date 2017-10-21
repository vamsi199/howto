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
	Right_of_Way   string
	Survey_Date   string
	Notes        string
	File_Name     string
}

type dataDsin struct {
	Road_ID int
	Start_Chain float32
	End_Chain float32
	Admin_Unit int
	WSC_Extent string
	WSC_Severity string
	EB_Extent string
	EB_Severity string
	RV_Extent string
	RV_Severity string
	PH_Extent string
	PH_Severity string
	Survey_Date string
	File_Name string

}

type dataRoug struct {
	Road_ID      int
	Start_Chain  float32
	End_Chain    float32
	Admin_Unit   int
	IRI_Average  string
	Survey_Date   string
	Year_of_Survey string
	Notes        string
	File_Name     string
}

type dataOutput struct {
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

func readCway() {
	fileCWAY, err := os.Open("CWAY.csv")
	if err != nil {
		panic(err)
	}
	defer fileCWAY.Close()
	slice := []*dataCway{}
	err = gocsv.Unmarshal(fileCWAY, &slice)
	if err != nil {
		panic(err)
	}

	for _, dataCway := range slice {
		fmt.Println(dataCway)
	}
}

func readDsin() {
	fileDSIN, err := os.Open("DSIN.csv")
	if err != nil {
		panic(err)
	}
	defer fileDSIN.Close()
	slice := []*dataDsin{}
	err = gocsv.Unmarshal(fileDSIN, &slice)
	if err != nil {
		panic(err)
	}

	for _, dataDsin := range slice {
		fmt.Println(dataDsin)
	}
}

func readRoug() {
	fileROUG, err := os.Open("ROUG.csv")
	if err != nil {
		panic(err)
	}
	defer fileROUG.Close()
	slice := []*dataRoug{}
	err = gocsv.Unmarshal(fileROUG, &slice)
	if err != nil {
		panic(err)
	}

	for _, dataRoug := range slice {
		fmt.Println(dataRoug)
	}
}

func main() {
	readCway()
	readDsin()
	readRoug()
}
