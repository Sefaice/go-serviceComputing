/*=================================================================

Program name:
	selpg (SELect PaGes)

Purpose:
	Sometimes one needs to extract only a specified range of
pages from an input text file. This program allows the user to do
that.

Author: sefaice

===================================================================*/

package main

/*================================= includes ======================*/

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

/*================================= types =========================*/

type selpg_args struct {
	start_page  int
	end_page    int
	in_filename string
	page_len    int /* default value, can be overriden by "-l number" on command line */
	page_type   int /* 'l' for lines-delimited, 'f' for form-feed-delimited */
	/* default is 'l' */
	print_dest string
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
	//flag.BoolVarP(&sa.page_type, "f", false, "page type")
	flag.StringVarP(&sa.print_dest, "d", "d", "", "print destination")

	// 改变默认的 Usage
	flag.Usage = usage

	flag.Parse()

	process_args()
	//process_input()
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

/*================================= usage() =======================*/

func usage() {
	fmt.Fprintf(os.Stderr, "\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n", progname)
	flag.PrintDefaults()
}

/*================================= EOF ===========================*/
