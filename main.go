package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return  dir
	//return strings.Replace(dir, "\\", "/", -1)
}

func loremList() []string {
	return []string{
		"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod " +
			"tempor incididunt ut labore et dolore magna aliqua.",
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut " +
			"aliquip ex ea commodo consequat.",
		"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum " +
			"dolore eu fugiat nulla pariatur.",
		"Excepteur sint occaecat cupidatat non proident, sunt in culpa qui " +
			"officia deserunt mollit anim id est laborum.",
	}
}

func getPdfName() string{
	path := GetCurrentPath()
	pdfPath := path + "\\hello.pdf"
	pdfPath ="g:\\2222222.pdf"
	os.Remove(pdfPath)
	return pdfPath
}
//func ExampleFpdf_CellFormat_tables() {
//	pdf := gofpdf.New("P", "mm", "A4", "")
//	type countryType struct {
//		nameStr, capitalStr, areaStr, popStr string
//	}
//	countryList := make([]countryType, 0, 8)
//	header := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}
//	loadData := func(fileStr string) {
//		fl, err := os.Open(fileStr)
//		if err == nil {
//			scanner := bufio.NewScanner(fl)
//			var c countryType
//			for scanner.Scan() {
//				// Austria;Vienna;83859;8075
//				lineStr := scanner.Text()
//				list := strings.Split(lineStr, ";")
//				if len(list) == 4 {
//					c.nameStr = list[0]
//					c.capitalStr = list[1]
//					c.areaStr = list[2]
//					c.popStr = list[3]
//					countryList = append(countryList, c)
//				} else {
//					err = fmt.Errorf("error tokenizing %s", lineStr)
//				}
//			}
//			fl.Close()
//			if len(countryList) == 0 {
//				err = fmt.Errorf("error loading data from %s", fileStr)
//			}
//		}
//		if err != nil {
//			pdf.SetError(err)
//		}
//	}
//	// Simple table
//	basicTable := func() {
//		left := (210.0 - 4*40) / 2
//		pdf.SetX(left)
//		for _, str := range header {
//			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
//		}
//		pdf.Ln(-1)
//		for _, c := range countryList {
//			pdf.SetX(left)
//			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
//			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
//			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
//			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
//			pdf.Ln(-1)
//		}
//	}
//	// Better table
//	improvedTable := func() {
//		// Column widths
//		w := []float64{40.0, 35.0, 40.0, 45.0}
//		wSum := 0.0
//		for _, v := range w {
//			wSum += v
//		}
//		left := (210 - wSum) / 2
//		// 	Header
//		pdf.SetX(left)
//		for j, str := range header {
//			pdf.CellFormat(w[j], 7, str, "1", 0, "C", false, 0, "")
//		}
//		pdf.Ln(-1)
//		// Data
//		for _, c := range countryList {
//			pdf.SetX(left)
//			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", false, 0, "")
//			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", false, 0, "")
//			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
//				"LR", 0, "R", false, 0, "")
//			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
//				"LR", 0, "R", false, 0, "")
//			pdf.Ln(-1)
//		}
//		pdf.SetX(left)
//		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
//	}
//	// Colored table
//	fancyTable := func() {
//		// Colors, line width and bold font
//		pdf.SetFillColor(255, 0, 0)
//		pdf.SetTextColor(255, 255, 255)
//		pdf.SetDrawColor(128, 0, 0)
//		pdf.SetLineWidth(.3)
//		pdf.SetFont("", "B", 0)
//		// 	Header
//		w := []float64{40, 35, 40, 45}
//		wSum := 0.0
//		for _, v := range w {
//			wSum += v
//		}
//		left := (210 - wSum) / 2
//		pdf.SetX(left)
//		for j, str := range header {
//			pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
//		}
//		pdf.Ln(-1)
//		// Color and font restoration
//		pdf.SetFillColor(224, 235, 255)
//		pdf.SetTextColor(0, 0, 0)
//		pdf.SetFont("", "", 0)
//		// 	Data
//		fill := false
//		for _, c := range countryList {
//			pdf.SetX(left)
//			pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", fill, 0, "")
//			pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", fill, 0, "")
//			pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
//				"LR", 0, "R", fill, 0, "")
//			pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
//				"LR", 0, "R", fill, 0, "")
//			pdf.Ln(-1)
//			fill = !fill
//		}
//		pdf.SetX(left)
//		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
//	}
//	loadData(example.TextFile("countries.txt"))
//	pdf.SetFont("Arial", "", 14)
//	pdf.AddPage()
//	basicTable()
//	pdf.AddPage()
//	improvedTable()
//	pdf.AddPage()
//	fancyTable()
//	filePath := getPdfName()
//	pdf.OutputFileAndClose(filePath)
//
//	// Output:
//	// Successfully generated pdf/Fpdf_CellFormat_tables.pdf
//}



func ExampleFpdf_SplitLines_tables() {
	const (
		colCount = 3
		colWd    = 60.0
		marginH  = 15.0
		lineHt   = 5.5
		cellGap  = 2.0
	)
	// var colStrList [colCount]string
	type cellType struct {
		str  string
		list [][]byte
		ht   float64
	}
	var (
		cellList [colCount]cellType
		cell     cellType
	)

	pdf := gofpdf.New("P", "mm", "A4", "") // 210 x 297
	header := [colCount]string{"Column A", "Column B", "Column C"}
	alignList := [colCount]string{"L", "C", "R"}
	strList := loremList()
	pdf.SetMargins(marginH, 15, marginH)
	pdf.SetFont("Arial", "", 14)
	pdf.AddPage()

	// Headers
	pdf.SetTextColor(224, 224, 224)
	pdf.SetFillColor(64, 64, 64)
	for colJ := 0; colJ < colCount; colJ++ {
		pdf.CellFormat(colWd, 10, header[colJ], "1", 0, "CM", true, 0, "")
	}
	pdf.Ln(-1)
	pdf.SetTextColor(24, 24, 24)
	pdf.SetFillColor(255, 255, 255)

	// Rows
	y := pdf.GetY()
	count := 0
	for rowJ := 0; rowJ < 2; rowJ++ {
		maxHt := lineHt
		// Cell height calculation loop
		for colJ := 0; colJ < colCount; colJ++ {
			count++
			if count > len(strList) {
				count = 1
			}
			cell.str = strings.Join(strList[0:count], " ")
			cell.list = pdf.SplitLines([]byte(cell.str), colWd-cellGap-cellGap)
			cell.ht = float64(len(cell.list)) * lineHt
			if cell.ht > maxHt {
				maxHt = cell.ht
			}
			cellList[colJ] = cell
		}
		// Cell render loop
		x := marginH
		for colJ := 0; colJ < colCount; colJ++ {
			pdf.Rect(x, y, colWd, maxHt+cellGap+cellGap, "D")
			cell = cellList[colJ]
			cellY := y + cellGap + (maxHt-cell.ht)/2
			for splitJ := 0; splitJ < len(cell.list); splitJ++ {
				pdf.SetXY(x+cellGap, cellY)
				pdf.CellFormat(colWd-cellGap-cellGap, lineHt, string(cell.list[splitJ]), "", 0,
					alignList[colJ], false, 0, "")
				cellY += lineHt
			}
			x += colWd
		}
		y += maxHt + cellGap + cellGap
	}


	//////////////////
	type countryType struct {
		nameStr, capitalStr, areaStr, popStr string
	}
	countryList := make([]countryType, 0, 8)
	header1 := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}
	// Simple table
	basicTable := func() {
		left := (210.0 - 4*40) / 2
		pdf.SetX(left)
		for _, str := range header1 {
			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
	}
	pdf.SetFont("Arial", "", 14)
	basicTable()
	pdf.AddPage()

	filePath := getPdfName()
	pdf.OutputFileAndClose(filePath)

	// Output:
	// Successfully generated pdf/Fpdf_SplitLines_tables.pdf
}

func ExampleFpdf_MultiCell() {
	filePath := "g:\\hello111.pdf"
	os.Remove(filePath)
	pdf := gofpdf.New("P", "mm", "A4", "")
	titleStr := "20000 Leagues Under the Seas"
	pdf.SetTitle(titleStr, false)
	pdf.SetAuthor("Jules Verne", false)
	pdf.SetHeaderFunc(func() {
		// Arial bold 15
		pdf.SetFont("Arial", "B", 15)
		// Calculate width of title and position
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetX((210 - wd) / 2)
		// Colors of frame, background and text
		pdf.SetDrawColor(0, 80, 180)
		pdf.SetFillColor(230, 230, 0)
		pdf.SetTextColor(220, 50, 50)
		// Thickness of frame (1 mm)
		pdf.SetLineWidth(1)
		// Title
		pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
		// Line break
		pdf.Ln(10)
	})
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("Arial", "I", 8)
		// Text color in gray
		pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})


	pdf.OutputFileAndClose(filePath)

}



func main() {

	const (
		colCount = 5
		colWd    = 30.0
		marginH  = 15.0
		lineHt   = 5.5
		cellGap  = 2.0
	)


	//filePath := getPdfName()
	filePath := "g:\\hello.pdf"
	os.Remove(filePath)
	path := GetCurrentPath()

	var pdf *gofpdf.Fpdf
	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginH, 15, marginH)

	//宋体
	simsunttf := path + "\\simsun.ttf"
	pdf.AddUTF8Font("chinafont","",simsunttf)
	//宋体粗体
	songblod := path + "\\simsunblod.ttf"
	pdf.AddUTF8Font("songblod","",songblod)

	pdf.SetFont("chinafont", "", 20)
	pdf.AddPage()

	//标题
	titleStr :="CJK1型号地板卡测试报告"
	wd := pdf.GetStringWidth(titleStr) + 6
	pdf.SetX((210 - wd) / 2)
	pdf.CellFormat(wd, 6, titleStr, "0", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)


	pdf.SetFont("chinafont", "", 14)
	//工号
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "工号", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, "3424334", "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	//版序号
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "版序号", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, "3424334", "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	//产品型号
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "产品型号", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	//产品图号
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "产品图号", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd, 10, "", "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	//测试员
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "测试员", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, "3424334", "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	//测试时间
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "测试时间", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, "3424334", "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	//测试结论
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "测试结论", "1", 0, "CM", false, 0, "")
	pdf.SetFillColor(0, 128, 0)
	pdf.CellFormat(colWd*4, 10, "3424334", "1", 0, "CM", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("songblod", "", 16)
	zukang := "电压测试"
	zukangwd := pdf.GetStringWidth(zukang) + 6
	pdf.SetX((210 - zukangwd) / 2)
	pdf.CellFormat(zukangwd, 10, zukang, "0", 0, "CM", false, 0, "")
	pdf.Ln(-1)


	pdf.SetFont("chinafont", "", 16)
	header := [colCount]string{"序号", "测试内容", "正确范围", "测量值", "结果"}
	// Headers
	pdf.SetX((210 - colWd*5) / 2)
	for colJ := 0; colJ < colCount; colJ++ {
		pdf.CellFormat(colWd, 10, header[colJ], "1", 0, "CM", false, 0, "")
	}
	pdf.Ln(-1)

	//内容
	for colJ :=0; colJ<10; colJ++{
		serial := strconv.Itoa(colJ+1)
		pdf.SetX((210 - colWd*5) / 2)
		pdf.CellFormat(colWd, 6, serial, "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWd, 6, "aa2", "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWd, 6, "aa3", "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWd, 6, "aa4", "1", 0, "CM", false, 0, "")

		pdf.SetFillColor(0, 128, 0)
		pdf.CellFormat(colWd, 6, "aa4", "1", 0, "CM", true, 0, "")
		pdf.Ln(-1)
	}

	pdf.Ln(-1)

	var lineNumber float64
	lineNumber = 7.0
    //测试 7行
    var testWidth float64
	testWidth = 150/lineNumber
	pdf.SetX((210 - testWidth*lineNumber) / 2)
	pdf.CellFormat(testWidth, 6, "aa2", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(testWidth, 6, "aa2", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(testWidth, 6, "aa3", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(testWidth, 6, "aa4", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(testWidth, 6, "aa4", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(testWidth, 6, "aa4", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(testWidth, 6, "aa4", "1", 0, "CM", false, 0, "")


	pdf.OutputFileAndClose(filePath)

}
