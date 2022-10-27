package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"time"
)

func main() {
	f := excelize.NewFile()
	sheetName := "成绩单"
	f.SetSheetName("Sheet1", sheetName)
	data := [][]interface{}{
		{"考试成绩表"},
		{"考试名称：其中考试", nil, nil, nil, "基础科目", nil, nil, "理科科目"},
		{"序号", "学号", "姓名", "斑级", "数学", "英语", "语文", "化学", "生物", "物理", "总分", "日期"},
		{1, 10001, "学生1", "1斑", 93, 80, 89, 86, 58, 77, 0, time.Now()},
		{2, 10002, "学生2", "2斑", 93, 80, 89, 86, 58, 77, 0, time.Now()},
		{3, 10003, "学生3", "1斑", 93, 80, 89, 86, 58, 77, 0, time.Now()},
		{4, 10004, "学生4", "2斑", 93, 80, 89, 86, 58, 77, 0, time.Now()},
		{5, 10005, "学生5", "2斑", 93, 80, 89, 86, 58, 77, 0, time.Now()},
		{6, 10006, "学生6", "1斑", 93, 80, 89, 86, 58, 77, 0, time.Now()},
		{7, 10007, "学生7", "1斑", 93, 80, 89, 86, 58, 77, 0, time.Now()},
	}
	for i, row := range data {
		starCell, err := excelize.JoinCellName("A", i+1)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		if err := f.SetSheetRow(sheetName, starCell, &row); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	}

	formulaType, ref := excelize.STCellFormulaTypeShared, "K4:K10"
	if err := f.SetCellFormula(sheetName, "K4", "=sum(E4,J4)",
		excelize.FormulaOpts{Ref: &ref, Type: &formulaType}); err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	mergeCellRanges := [][]string{{"a1", "k1"}, {"a2", "d2"}, {"e2", "g2"}, {"h2", "j2"}}

	for _, ranges := range mergeCellRanges {
		if err := f.MergeCell(sheetName, ranges[0], ranges[1]); err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
	}
	if err := f.SaveAs("book1.xlsx"); err != nil {
		fmt.Printf("err: %v\n", err)
	}
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	/* // 创建一个工作表
		index := newFunction1(f)
		// 设置单元格的值
		f.SetCellValue("Sheet2", "A2", "Hello world.")
		f.SetCellValue("Sheet1", "B2", 100)
		// 设置工作簿的默认工作表
		f.SetActiveSheet(index)
		// 根据指定路径保存文件
		if err := f.SaveAs("Book1.xlsx"); err != nil {
			fmt.Println(err)
		}
	}

	func newFunction1(f *excelize.File) int {
		index := newFunction(f)
		return index
	}

	func newFunction(f *excelize.File) int {
		index := f.NewSheet("Sheet2")
		return index
	}
	*/
}
