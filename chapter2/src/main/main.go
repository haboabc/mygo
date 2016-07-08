package main

import (
	"fmt"
	"hehe"
	"haha"
	"strconv"
	"os"
	"flag"
	"bufio"
	"io"
	"time"
)

var (
	/* 注册命令行选项和存储解析结果的变量
	 * -i 选项 指定infile 默认 "infile"
	 * -o 选项 指定outfile 默认 "outfile"
	 * -a 选项 指定sort  默认"qsort"
	 **/
	 // 存储结果变量 = 选项， 默认， 说明
	 infile	*string = flag.String("i", "", "File contains values for sort")
	outfile	*string	= flag.String("o", "outfile", "File to receive sorted values")
	algorithm *string = flag.String("a", "qsort", "Sort algorithm")
)

func Usage() {
	fmt.Println("Usage:...\n\t./main -i <inputfile> -o <outpurfile> -a <sortmethod>")
}

func readValues(infile string) (values []int, err error) {
	/* os.Open 打开文件 */
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Fail to open inputfile")
	} else {
		defer file.Close()
	}

	/* 创建读句柄 */
	br := bufio.NewReader(file)

	values  = make([]int, 0)

	for {
		/* 读取一行 */
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
			err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line")
			return
		}

		str :=  string(line)

		value, err1 := strconv.Atoi(str)
		//value, err1 := strconv.Atoi(line)

		if err1 != nil {
			fmt.Println("strconv.Atoi(line)")
			err = err1
			return
		}

		values = append(values, value)
	}


	return

}

func writeValues(values []int, outfile string) error {
	/* 创建文件 */
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Create file error")
		return err
	}
	defer file.Close()


	/* 将切片内容写到文件 */
	for _, v := range(values) {
		str := strconv.Itoa(v)
		file.WriteString(str + "\n")
	}
	return nil
}

func main() {
	/* 解析命令行 */
	//args := os.Args
	flag.Parse() //解些命令行, 解析的结果会

	if len(*infile) != 0 {
		fmt.Println("infile = ", *infile,
		" outfile = ", *outfile,
		" algorithm = ", *algorithm)
	} else {
		Usage()
	}

	/* 读取文件内容 */
	values, err := readValues(*infile)
	t1 := time.Now()
	if err != nil {
		fmt.Println("read error")
	} else {
		switch *algorithm {
		case "qsort":
			hehe.Asort(values)
		case "bubble":
			haha.Bubble(values)
		default:
			fmt.Println("not support this methon")
		}
	}

	t2 := time.Now()
	fmt.Println("the sorting cost ", t2.Sub(t1), " to complete")

	ret := writeValues(values, *outfile)
	if ret != nil {
		fmt.Println("write err")
	}
}
