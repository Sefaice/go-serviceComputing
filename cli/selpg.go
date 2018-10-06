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

var (
	start_page  int
	end_page    int
	in_filename string
	page_len    int /* default value, can be overriden by "-l number" on command line */
	page_type   int /* 'l' for lines-delimited, 'f' for form-feed-delimited */
	/* default is 'l' */
	print_dest string
)

/*================================= globals =======================*/

/*================================= prototypes ====================*/

/*================================= main()=== =====================*/

func init() {

	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	flag.Parse()

	flag.Usage()
}

/*================================= usage() =======================*/

func usage() {
	fmt.Fprintf(os.Stderr, "\nUSAGE: %s -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n")
	flag.PrintDefaults()
}

/*================================= EOF ===========================*/
