/**
 * @Author pibing
 * @create 2020/11/15 10:41 AM
 */

package excel

import (
	"fmt"
	"strconv"
	"testing"
)

func TestExportExcel(t *testing.T) {
	header := []string{"name", "age"}
	headerKV := map[string]string{
		"name": "姓名",
		"age":  "年龄",
	}
	name := "test" //文件名称
	path := "./"   //文件路径
	list := make([]map[string]interface{}, 0)
	for i := 0; i < 10; i++ {
		item := map[string]interface{}{
			"name": "name" + strconv.Itoa(i),
			"age":  10 + i,
		}
		list = append(list, item)
	} //假设10条数据

	fileName, err := ExportExcel(name, path, header, headerKV, list)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("文件path：", fileName)
}

func TestUploadExcel(t *testing.T) {

	head := []string{"name", "age"} //表头所对应的字段名称
	filePath := "./test.xlsx"       //文件路径
	list, err := ReadExcel(filePath, head)
	fmt.Println(err)
	for i := 0; i < len(list); i++ {
		fmt.Print("name:", list[i]["name"])
		fmt.Println("  age:", list[i]["age"])
	}
}
