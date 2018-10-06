package main

/*================================= includes ======================*/

import (
	"bufio"
	"fmt"
	"io"
	"os"

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
		/* read from commandline */
		var str string

		for page_ctr = sa.start_page; page_ctr <= sa.end_page; page_ctr++ {
			for line_ctr = 0; line_ctr < sa.page_len; line_ctr++ {
				fmt.Scanln(&str)
				fmt.Println(str)
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
