package model

import (
	"embed"
	"sort"
)

//go:embed grades.csv
var gradesCsvFile embed.FS

type Grade struct {
	ID          int
	US          string
	BritishTech string
	BritishAdj  string
	French      string
	UIAA        string
	Saxony      string
	Ewbank      string
	Ewbank2     string
	Finnish     string
	NOR         string
	Brazilian   string
	Kurtyka     string
}

var GradesList = []*Grade{
	{ID: 1, US: "3–4", BritishTech: "1", BritishAdj: "M", French: "1", UIAA: "I", Saxony: "I", Ewbank: "1–2", Ewbank2: "1–2", Finnish: "1", NOR: "1", Brazilian: "I", Kurtyka: "I"},
	{ID: 2, US: "5.0", BritishTech: "1", BritishAdj: "M", French: "1", UIAA: "I", Saxony: "I", Ewbank: "3–4", Ewbank2: "3–4", Finnish: "1", NOR: "1", Brazilian: "I sup", Kurtyka: "I"},
	{ID: 3, US: "5.1", BritishTech: "2", BritishAdj: "M", French: "2", UIAA: "II", Saxony: "II", Ewbank: "5–6", Ewbank2: "5–6", Finnish: "2", NOR: "2", Brazilian: "II", Kurtyka: "II"},
	{ID: 4, US: "5.2", BritishTech: "2", BritishAdj: "D", French: "2", UIAA: "II", Saxony: "II", Ewbank: "7", Ewbank2: "7", Finnish: "2", NOR: "2", Brazilian: "II sup", Kurtyka: "II"},
	{ID: 5, US: "5.3", BritishTech: "3", BritishAdj: "D", French: "3", UIAA: "III", Saxony: "III", Ewbank: "8–9", Ewbank2: "8–9", Finnish: "3", NOR: "3", Brazilian: "II sup", Kurtyka: "III"},
	{ID: 6, US: "5.4", BritishTech: "3", BritishAdj: "VD", French: "4a", UIAA: "IV", Saxony: "IV", Ewbank: "10", Ewbank2: "10", Finnish: "3", NOR: "3", Brazilian: "III", Kurtyka: "IV"},
	{ID: 7, US: "5.5", BritishTech: "4a", BritishAdj: "S", French: "4b", UIAA: "IV+/V-", Saxony: "V", Ewbank: "11–12", Ewbank2: "11–12", Finnish: "4", NOR: "4", Brazilian: "III sup", Kurtyka: "IV"},
	{ID: 8, US: "5.6", BritishTech: "4b", BritishAdj: "HS", French: "4c", UIAA: "V", Saxony: "VI", Ewbank: "13", Ewbank2: "13", Finnish: "4", NOR: "4", Brazilian: "IV", Kurtyka: "IV+"},
	{ID: 9, US: "5.7", BritishTech: "4c", BritishAdj: "VS", French: "5a", UIAA: "V+", Saxony: "VI", Ewbank: "14–15", Ewbank2: "14–15", Finnish: "5-", NOR: "5-", Brazilian: "IV", Kurtyka: "V-"},
	{ID: 10, US: "5.8", BritishTech: "4c", BritishAdj: "HVS", French: "5b", UIAA: "VI-", Saxony: "VIIa", Ewbank: "16", Ewbank2: "16", Finnish: "5", NOR: "5", Brazilian: "IV sup", Kurtyka: "V"},
	{ID: 11, US: "5.9", BritishTech: "5a", BritishAdj: "HVS", French: "5c", UIAA: "VI", Saxony: "VIIb", Ewbank: "17", Ewbank2: "17–18", Finnish: "5+", NOR: "5+", Brazilian: "V", Kurtyka: "V+"},
	{ID: 12, US: "5.10a", BritishTech: "5a", BritishAdj: "E1", French: "6a", UIAA: "VI+", Saxony: "VIIc", Ewbank: "18", Ewbank2: "19", Finnish: "5+", NOR: "6-", Brazilian: "VI", Kurtyka: "VI"},
	{ID: 13, US: "5.10b", BritishTech: "5b", BritishAdj: "E2", French: "6a+", UIAA: "VII-", Saxony: "VIIc", Ewbank: "19", Ewbank2: "20", Finnish: "6-", NOR: "6-", Brazilian: "VI", Kurtyka: "VI+"},
	{ID: 14, US: "5.10c", BritishTech: "5b", BritishAdj: "E2", French: "6b", UIAA: "VII", Saxony: "VIIIa", Ewbank: "20", Ewbank2: "21", Finnish: "6", NOR: "6", Brazilian: "VI sup", Kurtyka: "VI.1"},
	{ID: 15, US: "5.10d", BritishTech: "5c", BritishAdj: "E3", French: "6b+", UIAA: "VII+", Saxony: "VIIIb", Ewbank: "20", Ewbank2: "22", Finnish: "6", NOR: "6+", Brazilian: "VI sup", Kurtyka: "VI.1+"},
	{ID: 16, US: "5.11a", BritishTech: "5c", BritishAdj: "E3", French: "6c", UIAA: "VII+", Saxony: "VIIIc", Ewbank: "21", Ewbank2: "22", Finnish: "6+", NOR: "6+", Brazilian: "7a", Kurtyka: "VI.2"},
	{ID: 17, US: "5.11b", BritishTech: "5c", BritishAdj: "E4", French: "6c", UIAA: "VIII-", Saxony: "VIIIc", Ewbank: "22", Ewbank2: "23", Finnish: "7-", NOR: "7-", Brazilian: "7b", Kurtyka: "VI.2+"},
	{ID: 18, US: "5.11c", BritishTech: "6a", BritishAdj: "E4", French: "6c+", UIAA: "VIII-", Saxony: "IXa", Ewbank: "23", Ewbank2: "24", Finnish: "7", NOR: "7", Brazilian: "7c", Kurtyka: "VI.2+"},
	{ID: 19, US: "5.11d", BritishTech: "6a", BritishAdj: "E5", French: "7a", UIAA: "VIII", Saxony: "IXb", Ewbank: "23", Ewbank2: "25", Finnish: "7", NOR: "7+", Brazilian: "7c", Kurtyka: "VI.3"},
	{ID: 20, US: "5.12a", BritishTech: "6a", BritishAdj: "E5", French: "7a+", UIAA: "VIII+", Saxony: "IXc", Ewbank: "24", Ewbank2: "26", Finnish: "7+", NOR: "7+/8-", Brazilian: "8a", Kurtyka: "VI.3+"},
	{ID: 21, US: "5.12b", BritishTech: "6a", BritishAdj: "E6", French: "7b", UIAA: "VIII+", Saxony: "IXc", Ewbank: "25", Ewbank2: "27", Finnish: "8-", NOR: "8-", Brazilian: "8b", Kurtyka: "VI.4"},
	{ID: 22, US: "5.12c", BritishTech: "6b", BritishAdj: "E6", French: "7b+", UIAA: "IX-", Saxony: "Xa", Ewbank: "26", Ewbank2: "28", Finnish: "8", NOR: "8", Brazilian: "8c", Kurtyka: "VI.5"},
	{ID: 23, US: "5.12d", BritishTech: "6b", BritishAdj: "E7", French: "7c", UIAA: "IX", Saxony: "Xb", Ewbank: "27", Ewbank2: "29", Finnish: "8+", NOR: "8/8+", Brazilian: "9a", Kurtyka: "VI.4+"},
	{ID: 24, US: "5.13a", BritishTech: "6b", BritishAdj: "E7", French: "7c+", UIAA: "IX+", Saxony: "Xc", Ewbank: "28", Ewbank2: "30", Finnish: "9-", NOR: "8+", Brazilian: "9b", Kurtyka: "VI.5"},
	{ID: 25, US: "5.13b", BritishTech: "6c", BritishAdj: "E8", French: "8a", UIAA: "IX+", Saxony: "Xc", Ewbank: "29", Ewbank2: "31", Finnish: "9", NOR: "9-", Brazilian: "9c", Kurtyka: "VI.5+"},
	{ID: 26, US: "5.13c", BritishTech: "6c", BritishAdj: "E8", French: "8a+", UIAA: "X-", Saxony: "XIa", Ewbank: "30", Ewbank2: "32", Finnish: "9+", NOR: "9-/9", Brazilian: "10a", Kurtyka: "VI.5+"},
	{ID: 27, US: "5.13d", BritishTech: "6c", BritishAdj: "E9", French: "8b", UIAA: "X", Saxony: "XIb", Ewbank: "31", Ewbank2: "33", Finnish: "10-", NOR: "9", Brazilian: "10b", Kurtyka: "VI.6"},
	{ID: 28, US: "5.14a", BritishTech: "7a", BritishAdj: "E10", French: "8b+", UIAA: "X+", Saxony: "XIc", Ewbank: "32", Ewbank2: "34", Finnish: "10", NOR: "9/9+", Brazilian: "10c", Kurtyka: "VI.6+"},
	{ID: 29, US: "5.14b", BritishTech: "7a", BritishAdj: "E11", French: "8c", UIAA: "X+", Saxony: "XIc", Ewbank: "33", Ewbank2: "35", Finnish: "10+", NOR: "9+", Brazilian: "11a", Kurtyka: "VI.7"},
	{ID: 30, US: "5.14c", BritishTech: "7b", BritishAdj: "E11", French: "8c+", UIAA: "XI-", Saxony: "XIIa", Ewbank: "34", Ewbank2: "36", Finnish: "11-", NOR: "9+/10-", Brazilian: "11b", Kurtyka: "VI.7+"},
	{ID: 31, US: "5.14d", BritishTech: "", BritishAdj: "", French: "9a", UIAA: "XI", Saxony: "XIIb", Ewbank: "35", Ewbank2: "37", Finnish: "11", NOR: "", Brazilian: "11c", Kurtyka: "VI.8"},
	{ID: 32, US: "5.15a", BritishTech: "", BritishAdj: "", French: "9a+", UIAA: "XI+", Saxony: "", Ewbank: "36", Ewbank2: "38", Finnish: "", NOR: "", Brazilian: "12a", Kurtyka: ""},
	{ID: 33, US: "5.15b", BritishTech: "", BritishAdj: "", French: "9b", UIAA: "XI+", Saxony: "", Ewbank: "37", Ewbank2: "39", Finnish: "", NOR: "", Brazilian: "12b", Kurtyka: ""},
	{ID: 34, US: "5.15c", BritishTech: "", BritishAdj: "", French: "9b+", UIAA: "XII-", Saxony: "", Ewbank: "38", Ewbank2: "40", Finnish: "", NOR: "", Brazilian: "12c", Kurtyka: ""},
	{ID: 35, US: "5.15d", BritishTech: "", BritishAdj: "", French: "9c", UIAA: "XII", Saxony: "", Ewbank: "39", Ewbank2: "", Finnish: "", NOR: "", Brazilian: "13a", Kurtyka: ""},
}

var GradesID = make(map[int]*Grade, 36)
var GradesSaxony = make(map[string]*Grade, 36)

func init() {
	sort.Slice(GradesList, func(i, j int) bool {
		return GradesList[i].ID < GradesList[j].ID
	})

	for _, grade := range GradesList {
		GradesID[grade.ID] = grade
	}

	for _, grade := range GradesList {
		if _, exists := GradesSaxony[grade.Saxony]; exists {
			continue
		}
		GradesSaxony[grade.Saxony] = grade
	}
}

//
//func printGrades() {
//	f, _ := gradesCsvFile.Open("grades.csv")
//	reader := csv.NewReader(f)
//	records, _ := reader.ReadAll()
//	for _, record := range records {
//		fmt.Printf(`{ID: %s, US: "%s", BritishTech: "%s", BritishAdj: "%s", French: "%s", UIAA: "%s", Saxony: "%s",  Ewbank: "%s", Ewbank2: "%s", Finnish: "%s", NOR: "%s", Brazilian: "%s", Kurtyka: "%s"},`,
//			record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12])
//		fmt.Println()
//	}
//}
