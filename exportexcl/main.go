package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)


	func test1(path string) {
		//
		var file *xlsx.File
		var sheet *xlsx.Sheet
		var row *xlsx.Row
		var cell *xlsx.Cell
		file = xlsx.NewFile()
		sheet ,_= file.AddSheet("Sheet1")
		row = sheet.AddRow()
		//添加表头
		cell = row.AddCell()
		cell.Value = "工单号"
		cell = row.AddCell()
		cell.Value = "生产数量"
		cell = row.AddCell()
		cell.Value = "完成数量"
		//添加内容
		for i:=0;i<8 ;i++  {
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = "11"
			cell = row.AddCell()
			cell.Value = "22"
			cell = row.AddCell()
			cell.Value = "33"
		}

		err := file.Save(path)
		//err = file.Save("d:\\MyXLSXFile.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

	func test2() {
		var file *xlsx.File
		var sheet *xlsx.Sheet
		var row *xlsx.Row
		var cell *xlsx.Cell
		var err error

		file, _ = xlsx.OpenFile("MyXLSXFile.xlsx")
		sheet = file.Sheet["Sheet1"]
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = "000101"
		cell = row.AddCell()
		cell.Value = "中文1"
		err = file.Save("MyXLSXFile1.xlsx")
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

	func generateExcl(exclpath,dbpath string){
		//excel
		var file *xlsx.File
		var sheet *xlsx.Sheet
		var row *xlsx.Row
		var cell *xlsx.Cell
		file = xlsx.NewFile()
		sheet ,_= file.AddSheet("Sheet1")
		row = sheet.AddRow()
		//添加表头
		cell = row.AddCell()
		cell.Value = "工单编号"
		cell = row.AddCell()
		cell.Value = "产品型号"
		cell = row.AddCell()
		cell.Value = "产品名称"

		cell = row.AddCell()
		cell.Value = "生产数量"
		cell = row.AddCell()
		cell.Value = "完成数量"
		cell = row.AddCell()
		cell.Value = "通过率"

		cell = row.AddCell()
		cell.Value = "工单状态"
		cell = row.AddCell()
		cell.Value = "开始时间"
		cell = row.AddCell()
		cell.Value = "结束时间"


		//sqlite
		db, err := sql.Open("sqlite3", dbpath)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		rows, err := db.Query("select worknumber,producttype,productname,productnum,completenumb,passrate,workstate,startTime,endTime FROM worknumtb")
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var worknumber		string 	//工单编号
			var producttype		string 	//产品型号
			var productname		string 	//产品名称
			var productnum		string 	//生产数量
			var completenumb 	string 	//完成数量
			var passrate		string 	//通过率
			var workstate 		string 	//工单状态
			var startTime 		string 	//开始时间
			var endTime			string 	//结束时间
			err = rows.Scan(&worknumber, &producttype, &productname, &productnum,&completenumb, &passrate, &workstate, &startTime,&endTime)

			//excel 添加内容
			row = sheet.AddRow()

			cell = row.AddCell()
			cell.Value = worknumber
			cell = row.AddCell()
			cell.Value = producttype
			cell = row.AddCell()
			cell.Value = productname

			cell = row.AddCell()
			cell.Value = productnum
			cell = row.AddCell()
			cell.Value = completenumb
			cell = row.AddCell()
			cell.Value = passrate

			cell = row.AddCell()
			cell.Value = workstate
			cell = row.AddCell()
			cell.Value = startTime
			cell = row.AddCell()
			cell.Value = endTime
		}

		err = file.Save(exclpath)
		if err != nil {
			fmt.Printf(err.Error())
		}


	}

	func main() {
		//fmt.Println("命令行参数数量:",len(os.Args))

		//var exclpath string
		exclpath := os.Args[1]
		dbPath := os.Args[2]
		//for _,v:= range os.Args{
		//	//fmt.Printf("args[%v]=[%v]\n",k,v)
		//	exclpath = v
		//}

		generateExcl(exclpath,dbPath)
		//test1(exclpath)
	}

