package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"os/exec"
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

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
	GBK = Charset("gbk")
	HZGB2312 = Charset("HZGB2312")
)

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes,_=simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case GBK:
		var decodeBytes,_=simplifiedchinese.GBK.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case HZGB2312:
		var decodeBytes,_=simplifiedchinese.HZGB2312.NewDecoder().Bytes(byte)
		str= string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

type STDetail struct {
	worknumber		string  //工单编号
	porductname     string  //产品名字
	producttype     string  //产品类型
	productPN       string  //产品品号
	productserial   string  //产品序列号
	testresult      string  //测试结果1
	testresult2     string  //测试结果2
	testresult3     string  //测试结果3
	testTime        string  //测试时间
}

type  STDataWorkNum struct {
	worknumber		string //工单编号
	producttype		string //产品型号
	productname		string //产品名称
	productnum		string //生产数量
	completenumb	string //完成数量
	passrate		string //通过率
	workstate		string //工单状态
	startTime		string //开始时间
	endTime			string //结束时间
	remoteupdate	string //远程更新
}

func main() {

	//当前目录
	path := GetCurrentPath()

	//加载sqllite
	sqlpath := os.Getenv("LOCALAPPDATA")
	sqlpath = sqlpath + "\\tdcsrc\\C9D1576954E8B26E8BB19A42"
	//加载ini
	iniPath := path + "\\config.ini"
	cfg, err := ini.Load(iniPath)
	if err != nil {
		fmt.Println("====================")
		fmt.Println(err)
		fmt.Println("====================")
	}

	//sqlite
	db, err := sql.Open("sqlite3", sqlpath)
	if err != nil {
		fmt.Println("load sqlite3 error")
	}
	defer db.Close()

	selId := os.Args[1]
	//selId := "1"
	sql := "select worknumber,porductname,producttype,productpn,productserial,testresult,testresult2,testresult3,testtime FROM detailtb"

	sql = sql+ " where id =" + selId

	rows, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	var det STDetail
	for rows.Next() {
		err = rows.Scan(&det.worknumber, &det.porductname, &det.producttype, &det.productPN,&det.productserial, &det.testresult, &det.testresult2, &det.testresult3,&det.testTime)
	}

	timeFormat := det.testTime
	timeFormat = strings.Replace(timeFormat, "-", "", -1)
	timeFormat = strings.Replace(timeFormat, ":", "", -1)
	timeFormat = strings.Replace(timeFormat, " ", "", -1)
	pdfSavePath := cfg.Section("config").Key("pdfpath").String()
	if pdfSavePath == ""{
		pdfSavePath = GetCurrentPath() + "\\pdf"
	}
	os.Mkdir(pdfSavePath, os.ModePerm)  //创建目录

	filePath := pdfSavePath + "\\" + timeFormat + ".pdf"


	//deailwork
	sqlWork := "select worknumber,producttype,productname FROM worknumtb"
	sqlWork = sqlWork+ " where worknumber =" + det.worknumber

	rows, err = db.Query(sqlWork)
	if err != nil {
		log.Fatal(err)
	}
	var wktb STDataWorkNum
	for rows.Next() {
		err = rows.Scan(&wktb.worknumber, &wktb.producttype, &wktb.productname)
	}

	//test1
	decoded, err := base64.StdEncoding.DecodeString(det.testresult)
	json := string(decoded)

	//var data []byte = []byte(json)
	//json = ConvertByte2String(data,GB18030)

	//test2
	decoded2, err := base64.StdEncoding.DecodeString(det.testresult2)
	json2 := string(decoded2)
	//var data2 []byte = []byte(json2)
	//json2 = ConvertByte2String(data2,GB18030)

	//test3
	decoded3, err := base64.StdEncoding.DecodeString(det.testresult3)
	json3 := string(decoded3)
	//var data3 []byte = []byte(json3)
	//json3 = ConvertByte2String(data3,GB18030)

	const (
		colCount = 5
		colWd    = 40.0
		marginH  = 10.0
		lineHt   = 5.5
		cellGap  = 2.0
	)



	var pdf *gofpdf.Fpdf
	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(marginH, 15, marginH)

	//宋体
	simsunttf := path + "\\simsun.ttf"
	pdf.AddUTF8Font("chinafont","",simsunttf)
	//宋体粗体
	songblod := path + "\\simsunblod.ttf"
	pdf.AddUTF8Font("songblod","",songblod)

	//设置字号
	pdf.SetFont("chinafont", "", 20)
	pdf.AddPage()

	//标题
	titleStr := wktb.producttype + "型" + wktb.productname + "名称"
	wd := pdf.GetStringWidth(titleStr) + 6
	pdf.SetX((210 - wd) / 2)
	pdf.CellFormat(wd, 6, titleStr, "0", 0, "C", false, 0, "")
	pdf.Ln(-1)
	pdf.Ln(-1)
	pdf.Ln(-1)


	pdf.SetFont("chinafont", "", 14)
	//工号
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "工单编号", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, det.worknumber, "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	//版序号
	//pdf.SetX((210 - colWd*5) / 2)
	//pdf.CellFormat(colWd, 10, "版序号", "1", 0, "CM", false, 0, "")
	//pdf.CellFormat(colWd*4, 10, "3424334", "1", 0, "CM", false, 0, "")
	//pdf.Ln(-1)

	//产品型号
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "产品型号", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, det.productserial, "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	//产品图号
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "产品图号", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, det.productPN, "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)


	//测试员
	usrname := cfg.Section("config").Key("username").String()
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "测试员", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, usrname, "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)
	//测试时间
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "测试时间", "1", 0, "CM", false, 0, "")
	pdf.CellFormat(colWd*4, 10, det.testTime, "1", 0, "CM", false, 0, "")
	pdf.Ln(-1)

	titleTestResult :="测试失败"
	isSuccess1 := gjson.Get(string(json), "result")
	isSuccess2 := gjson.Get(string(json3), "result")
	if isSuccess1.String() =="true" && isSuccess2.String() == "true"{
		titleTestResult = "测试成功"
	}

	//测试结论
	pdf.SetX((210 - colWd*5) / 2)
	pdf.CellFormat(colWd, 10, "测试结论", "1", 0, "CM", false, 0, "")
	pdf.SetFillColor(0, 128, 0)
	pdf.CellFormat(colWd*4, 10, titleTestResult, "1", 0, "CM", true, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("songblod", "", 16)
	zukang := "电压测试"
	zukangwd := pdf.GetStringWidth(zukang) + 6
	pdf.SetX((210 - zukangwd) / 2)
	pdf.CellFormat(zukangwd, 10, zukang, "0", 0, "CM", false, 0, "")
	pdf.Ln(-1)


	pdf.SetFont("chinafont", "", 16)
	header := [colCount]string{"序号", "测试内容", "结果","测量值","正确范围"}
	// Headers
	pdf.SetX((210 - colWd*5) / 2)
	for colJ := 0; colJ < colCount; colJ++ {
		if colJ == 0 {
			pdf.CellFormat(colWd-15, 10, header[colJ], "1", 0, "CM", false, 0, "")
		} else if colJ == 1 {
			pdf.CellFormat(colWd +15, 10, header[colJ], "1", 0, "CM", false, 0, "")
		}else{
			pdf.CellFormat(colWd, 10, header[colJ], "1", 0, "CM", false, 0, "")
		}

	}
	pdf.Ln(-1)


	//json电压测试
	//==========================================
	pdf.SetFont("chinafont", "", 12)
	result := gjson.Get(string(json), "data")
	nRetRow1 :=0
	result.ForEach(func(key, value gjson.Result) bool {
		nRetRow1++
		serial := strconv.Itoa(nRetRow1)
		//计算行高
		curx, y := pdf.GetXY()
		x := curx-marginH/2
		fmt.Printf("GetY=%f\n",y)
		height := 0.
		_, lineHt := pdf.GetFontSize()

		testWidth :=colWd
		value.ForEach(func(key1, value1 gjson.Result) bool {
			lines := pdf.SplitLines([]byte(value1.String()), testWidth+15)
			h := float64(len(lines))*lineHt + cellGap*float64(len(lines))
			if h > height {
				height = h
			}
			return true
		})

		_, pageh := pdf.GetPageSize()
		_,_, _, mbottom := pdf.GetMargins()
		//如果大于页的底部，就增加一页
		if pdf.GetY()+height > pageh-mbottom {
			pdf.AddPage()
			y = pdf.GetY()
		}

		//填充
		pdf.SetX((210 - colWd*5) / 2)
		pdf.Rect(x, y, colWd-15, height, "")
		pdf.MultiCell(colWd-15, lineHt+cellGap, serial, "", "CM", false)
		//fmt.Printf("top=%f,bottom=%f\n",y,y+height)
		x = x+colWd-15
		nFormat :=1
		value.ForEach(func(key1, value1 gjson.Result) bool {
			if nFormat==1{
				pdf.SetXY(x, y)
				pdf.Rect(x, y, colWd+15, height, "")
				pdf.MultiCell(colWd+15, lineHt+cellGap, value1.String(), "", "CM", false)
			}
			if nFormat==2{
				pdf.SetFillColor(0, 128, 0)
				//pdf.SetLineWidth(0.1)
				x = x+15
				pdf.SetXY(x, y)
				pdf.Rect(x, y, colWd, height, "")
				pdf.CellFormat(colWd, height, value1.String(), "1", 0, "CM", true, 0, "")
				//pdf.MultiCell(colWd, lineHt+cellGap, value1.String(), "", "CM", true)
			}
			if nFormat==3{

				//pdf.SetLineWidth(1)
				//pdf.SetDrawColor(0, 128,0)
				pdf.SetXY(x, y)
				pdf.Rect(x, y, colWd, height, "")
				//pdf.CellFormat(colWd, height, "", "1", 0, "CM", true, 0, "")//结果
				pdf.MultiCell(colWd, lineHt+cellGap, value1.String(), "", "CM", false)
			}
			//fmt.Printf("left=%f,right=%f\n",x,x+colWd)
			nFormat++
			x += colWd
			return true
		})
		pdf.SetXY(x, y)
		pdf.Rect(x, y, colWd, height, "")
		pdf.MultiCell(colWd, lineHt+cellGap, "0.1-1.0", "", "CM", false)
		pdf.SetXY(curx, y+height)

		//fmt.Printf("curx=%f,curxy=%f\n",y,y+height)
		//pdf.Ln(-1)
		return true
	})

	if len(json2) >0{
		//程序烧录
		pdf.SetFont("songblod", "", 16)
		shaolu := "程序烧录"
		shaoluwd := pdf.GetStringWidth(shaolu) + 6
		pdf.SetX((210 - shaoluwd) / 2)
		pdf.CellFormat(shaoluwd, 10, shaolu, "0", 0, "CM", false, 0, "")
		pdf.Ln(-1)

		pdf.SetFont("chinafont", "", 16)
		pdf.SetX((210 - colWd*5) / 2)
		pdf.CellFormat(colWd, 10, "测试内容", "1", 0, "CM", false, 0, "")
		pdf.CellFormat(colWd*4, 10, "结果", "1", 0, "CM", false, 0, "")
		pdf.Ln(-1)

		pdf.SetFont("chinafont", "", 10)

		resulshaolu := gjson.Get(string(json2), "data")
		nShaolu :=1
		resulshaolu.ForEach(func(key, value gjson.Result) bool {
			pdf.SetX((210 - colWd*5) / 2)
			if nShaolu ==1{
				pdf.CellFormat(colWd, 10, "CPU1", "1", 0, "CM", false, 0, "")
			}
			if nShaolu ==2{
				pdf.CellFormat(colWd, 10, "CPU2", "1", 0, "CM", false, 0, "")
			}
			pdf.SetFillColor(0, 128, 0)
			pdf.CellFormat(colWd*4, 10, value.String(), "1", 0, "CM", true, 0, "")
			pdf.Ln(-1)
			nShaolu++
			return true
		})
	}


	if len(json3) >0{
		//功能测试========================================
		pdf.SetFont("songblod", "", 16)
		gntitle := "功能测试"
		gntitlewd := pdf.GetStringWidth(gntitle) + 6
		pdf.SetX((210 - gntitlewd) / 2)
		pdf.CellFormat(gntitlewd, 10, gntitle, "0", 0, "CM", false, 0, "")
		pdf.Ln(-1)

		//表头
		pdf.SetFont("chinafont", "", 8)
		resultHeader2 := gjson.Get(string(json3), "data")
		resultHeader2.ForEach(func(key, value gjson.Result) bool {
			//fmt.Println(key.String())
			var nCount float64
			value.ForEach(func(key1, value1 gjson.Result) bool {
				nCount++
				return true
			})

			var testWidth float64
			testWidth = 200/nCount
			pdf.SetX((210 - testWidth*nCount) / 2)

			value.ForEach(func(key1, value1 gjson.Result) bool {
				pdf.CellFormat(testWidth, 6, key1.String(), "1", 0, "CM", false, 0, "")
				return true
			})
			pdf.Ln(-1)
			return false //keep iterating
		})

		//内容
		result2 := gjson.Get(string(json3), "data")
		result2.ForEach(func(key, value gjson.Result) bool {
			var nCount float64
			value.ForEach(func(key1, value1 gjson.Result) bool {
				nCount++
				return true
			})

			//计算每列宽度
			var testWidth float64
			testWidth = 200/nCount


			//计算行高
			curx, y := pdf.GetXY()
			x := curx-marginH/2

			height := 0.
			_, lineHt := pdf.GetFontSize()


			value.ForEach(func(key1, value1 gjson.Result) bool {
				lines := pdf.SplitLines([]byte(value1.String()), testWidth)
				h := float64(len(lines))*lineHt + cellGap*float64(len(lines))
				if h > height {
					height = h
				}
				return true
			})

			_, pageh := pdf.GetPageSize()
			_,_, _, mbottom := pdf.GetMargins()
			//如果大于页的底部，就增加一页
			if pdf.GetY()+height > pageh-mbottom {
				pdf.AddPage()
				y = pdf.GetY()
			}

			//填充
			pdf.SetX((210 - testWidth*nCount) / 2)
			value.ForEach(func(key1, value1 gjson.Result) bool {
				pdf.Rect(x, y, testWidth, height, "")
				pdf.MultiCell(testWidth, lineHt+cellGap, value1.String(), "", "CM", false)
				x += testWidth
				pdf.SetXY(x, y)
				return true
			})
			pdf.SetXY(curx, y+height)
			return true //keep iterating
		})
	}

	os.Remove(filePath)
	pdf.OutputFileAndClose(filePath)
	cmd := exec.Command("cmd.exe", "/c", "start "+filePath)
	cmd.Run()

}
