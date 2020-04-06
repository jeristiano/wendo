package service

import (
	"github.com/jeristiano/wendo/pkg/export"
	"github.com/tealeg/xlsx"
	"strconv"
	time2 "time"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

func GetAll() {

}

func (t *Tag) Export() (string, error) {
	tags, err := t.GetAll()
	if err != nil {
		return "", err
	}

	xlsFile := xlsx.NewFile()
	sheet, err := xlsFile.AddSheet("报表导出")

	titles := []string{"ID", "名称", "创建人", "创建时间", "修改人", "修改时间"}
	row := sheet.AddRow()

	var cell *xlsx.Cell
	for _, title := range titles {
		cell = row.AddCell()
		cell.Value = title
	}

	for _, v := range tags {
		values := []string{
			strconv.Itoa(v.ID),
			v.Name,
			v.CreatedBy,
			v.CreatedBy,
			v.ModifiedBy,
			strconv.Itoa(v.ModifiedOn),
		}
		row = sheet.AddRow()
		for _, value := range values {
			cell = row.AddCell()
			cell.Value = value
		}
	}

	time := strconv.Itoa(int(time2.Now().Unix()))
	filename := "tags-" + time + export.EXT
	dirFullPath := export.GetExcelFullPath()
	err = file.IsNotExistMkDir(dirFullPath)
	if err != nil {
		return "", err
	}
	err = xlsFile.Save(dirFullPath + filename)
	if err != nil {
		return "", err
	}
	return filename, nil
}
