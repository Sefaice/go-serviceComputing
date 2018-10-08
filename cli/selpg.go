package main

/*================================= includes ======================*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	flag "github.com/spf13/pflag"
)

/*================================= types =========================*/

type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int  /* default value, can be overriden by "-l number" on command line */
	page_type   bool /* 'l' for lines-delimited, 'f' for form-feed-delimited default is 'l' */
	print_dest  string
}

/*================================= globals =======================*/

var progname = "selpgGO"

/*================================= prototypes ====================*/

/*================================= main()=== =====================*/

var sa = new(selpg_args)

func main() {
	flag.IntVarP(&sa.start_page, "s", "s", -1, "start page")
	flag.IntVarP(&sa.end_page, "e", "e", -1, "end page")
	flag.IntVarP(&sa.page_len, "l", "l", 72, "page length")
	flag.BoolVarP(&sa.page_type, "f", "f", false, "page type")
	flag.StringVarP(&sa.print_dest, "d", "d", "", "print destination")

	// 改变默认的 Usage
	flag.Usage = usage

	flag.Parse()

	process_args()
	process_input()
}

/*================================= process_args() ================*/

func process_args() {

	/* check the command-line arguments for validity */
	if flag.NFlag() < 2 {
		fmt.Fprintf(os.Stderr, "%s: not enough arguments\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	/* handle 1st arg - start page */
	//not required the first argument must be -s
	//not check the start_page is an int
	if sa.start_page < 1 {
		fmt.Fprintf(os.Stderr, "%s: invalid start page %d\n", progname, sa.start_page)
		flag.Usage()
		os.Exit(3)
	}

	/* handle 2nd arg - end page */
	if sa.end_page < 1 || sa.end_page < sa.start_page {
		fmt.Fprintf(os.Stderr, "%s: invalid end page %d\n", progname, sa.end_page)
		flag.Usage()
		os.Exit(5)
	}

	/* no need loop for opt args because using flag */

	/* handle last optional arg - filename */
	if flag.NArg() == 1 {
		//exist
		sa.in_filename = flag.Arg(0)
		_, err := os.Stat(sa.in_filename)
		if err != nil && os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s: input file \"%s\" does not exist\n", progname, sa.in_filename)
			flag.Usage()
			os.Exit(6)
		}

		//openable
		_, err = os.Open(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n", progname, sa.in_filename)
			flag.Usage()
			os.Exit(7)
		}

	}

}

/*================================= process_input() ===============*/

func process_input() {

	line_ctr := 0
	page_ctr := 1

	//selpg内容通过管道输入给 grep, grep从中搜出带有keyword文件的内容,copy from blog
	fout := os.Stdout
	if sa.print_dest != "" {
		cmd := exec.Command("grep", "-nf", "keyword")
		inpipe, err := cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
			os.Exit(8)
		}
		defer inpipe.Close()
		cmd.Stdout = fout
		cmd.Start()
	}

	/**another code of -d with pipe, copy from github
	cmd_grep := exec.Command("./" + sa.print_dest)
	stdin_grep, grep_error := cmd_grep.StdinPipe()
	if grep_error != nil {
		fmt.Println("Error happened about standard input pipe ", grep_error)
		os.Exit(30)
	}
	writer := stdin_grep
	if grep_error := cmd_grep.Start(); grep_error != nil {
		fmt.Println("Error happened in execution ", grep_error)
		os.Exit(30)
	}
	if sa.page == true { //-d type
		process_input_f_d(reader, writer, args, &page_ctr)
	} else { //-l type
		process_input_l_d(reader, writer, sa, &page_ctr, &line_ctr)
	}
	stdin_grep.Close()
	//make sure all the infor in the buffer could be read
	if err := cmd_grep.Wait(); err != nil {
		fmt.Println("Error happened in Wait process")
		os.Exit(30)
	}
	*/

	if flag.NArg() == 1 {
		/* read from file input */
		fin, err := os.Open(sa.in_filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: could not open input file \"%s\"\n", progname, sa.in_filename)
			flag.Usage()
			os.Exit(7)
		}
		//defer calls when this func returns
		defer fin.Close()

		if sa.page_type == false {
			//read by line
			br := bufio.NewReader(fin)

			file_end := false

			for page_ctr = sa.start_page; page_ctr <= sa.end_page && file_end == false; page_ctr++ {
				for line_ctr = 0; line_ctr < sa.page_len && file_end == false; line_ctr++ {
					a, _, c := br.ReadLine()
					if c == io.EOF {
						file_end = true
					}
					fmt.Println(string(a))
				}
			}
		} else {
			//read by char and stop a page when '\f'
			//readString can split string by the param char you input
			br := bufio.NewReader(fin)
			file_end := false

			for page_ctr = sa.start_page; page_ctr <= sa.end_page && file_end == false; page_ctr++ {
				for !file_end {
					line, _ := br.ReadString('\f')
					if line == "" {
						file_end = true
						break
					}
					fmt.Println(line)
				}
			}
		}
	} else {
		/* read from commandline */
		reader := bufio.NewReader(os.Stdin)

		for page_ctr = sa.start_page; page_ctr <= sa.end_page; page_ctr++ {
			for line_ctr = 0; line_ctr < sa.page_len; line_ctr++ {
				//fmt.Scanln(&str) cannot read space in a line, so use below instead
				str, _, _ := reader.ReadLine()
				fmt.Println(string(str))
			}
		}
	}

	/* end main loop */

	if page_ctr < sa.start_page {
		fmt.Fprintf(os.Stderr, "%s: start_page (%d) greater than total pages (%d), no output written\n", progname, sa.start_page, page_ctr)
	} else if page_ctr < sa.end_page {
		fmt.Fprintf(os.Stderr, "%s: end_page (%d) greater than total pages (%d), less output than expected\n", progname, sa.end_page, page_ctr)
	}

}

/*================================= usage() =======================*/

func usage() {
	fmt.Fprintf(os.Stderr, "\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n", progname)
	flag.PrintDefaults()
}

/*================================= EOF ===========================*/
