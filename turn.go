package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/efekarakus/termcolor"
	"github.com/gookit/color"
	"github.com/hisamafahri/coco"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

const usage = `
Correct usage of turn:

turn [-inputType] <inputValue> [-out] {d|x|w|o|b|r|rgb|c|a|u} [-v]

Possible input type flags are:  
  -d, --decimal	decimal number	
  -x, --hex	hexadecimal number with or without preceding "0x" or "0X"  
  -o, --octal	octal number with or without preceding "0o" or "0O"  
  -bin, --bit	binary number as sequence of zeros and ones without whitespace  
  -rgb	rgb	values in a string. Whitespaces need doubles quotes. Examples: "-rgb r123g255b23" or -rgb "123 255 23" or -rgb "123, 255, 23"  
  -r -g -b	🟥 🟩 🟦 as 8-bit decimal value (0-255). Absent values will be considered as zero. Examples "-r 123 -g 255 -b23" or "-g 255"
  -rgbx		rgb values as hexadecimal triplet as a string. Example: "-rgbx 7bff17"    
  -a, --asc	a string beginning with an ASCII encoded character. Only the first character is considered.

Output flag:  
  -out, --output  indicates a list of desired output formats where applicable. A value of a default type is output if -out flag is omitted.

Possible output type arguments as are:  
  d		decimal number format  
  x		hexadecimal number format  
  b		binary number format  
  o		octal format  
  a		ASCII representation where applicable  
  c		character representation where applicable  
  u		UTF-8 representation where applicable  
  w, web	web formats
  
Other flags:  
  -v, --verbose		outputs additional information alternative versions/formats
  -h, --help		outputs this help text  
  -v, -ver, --version	outputs the version number if no other flags are present

Note:  
Order of the desired output types will be considered.  
Color output dependent of your terminal's abilities and settings.  

First version developed as a project for the CS50 course "Introduction to Computer Science" of the Harvard University 2022.  
Made with ❤️  and the Go programming language.

`

const version = `v1.0.23-beta.2`

func main() {
	// examplePtr := flag.String("example", "defaultValue", " Help text.")

	var verboseFlag, versionFlag bool
	var decimalValue, binaryDigits, redInt, greenInt, blueInt int
	var tempInt int64
	var red, green, blue float64
	var rgbSlice []string
	var hexValue, octValue, binValue, rgbValue, ascValue, rgbHexValue, outputTypes string
	var err, errRed, errGreen, errBlue error
	var rgbHexValueArray [3]uint8

	//decimal numbers
	flag.IntVar(&decimalValue, "d", 0, "Input of an decimal integer value.")
	flag.IntVar(&decimalValue, "decimal", 0, "Input of an decimal integer value.")

	flag.StringVar(&hexValue, "x", "", "Input of a hexadecimal integer value with or without preceding \"0x\" or \"0X\".")
	flag.StringVar(&hexValue, "hex", "", "Input of a hexadecimal integer value with or without preceding \"0x\" or \"0X\".")

	flag.StringVar(&octValue, "o", "", "Input of an octal integer value with or without preceding \"0o\" or \"0O\".")
	flag.StringVar(&octValue, "octal", "", "Input of an octal integer value with or without preceding \"0o\" or \"0O\".")

	flag.StringVar(&binValue, "bin", "", "Input of a binary integer value - no whitespace allowed.")
	flag.StringVar(&binValue, "bit", "", "Input of a binary integer value - no whitespace allowed.")

	flag.StringVar(&rgbValue, "rgb", "", "Input of three decimal integer value as a string. If values are separated with whitespaces double quotes around the whole expression are mandatory.\nExamples: -rgb r123g255b15 or -rgb \"123 255 15\"")
	flag.IntVar(&redInt, "r", 0, "Input of an decimal integer value between 0-255 (8-bit) for intensity of red color.")
	flag.IntVar(&greenInt, "g", 0, "Input of an decimal integer value between 0-255 (8-bit) for intensity of green color.")
	flag.IntVar(&blueInt, "b", 0, "Input of an decimal integer value between 0-255 (8-bit) for intensity of blue color.")
	flag.StringVar(&rgbHexValue, "rgbx", "", "Input of rgb values as hexadecimal triplet as a string. Example: -rgbx 7bff17  ")

	flag.StringVar(&ascValue, "a", "", "Input of a string beginning with an ASCII encoded character. Only the first character is considered")
	flag.StringVar(&ascValue, "asc", "", "Input of a string beginning with an ASCII encoded character. Only the first character is considered")

	flag.BoolVar(&verboseFlag, "v", false, "Turns on verbose output where applicable.")
	flag.BoolVar(&verboseFlag, "verbose", false, "Turns on verbose output where applicable.")

	flag.BoolVar(&versionFlag, "ver", false, "Outputs the version number")
	flag.BoolVar(&versionFlag, "version", false, "Outputs the version number")

	flag.StringVar(&outputTypes, "out", "", "Outputs desired output formats.")
	flag.StringVar(&outputTypes, "output", "", "Outputs desired output formats.")

	flag.Usage = func() { fmt.Print(usage) }

	flag.Parse()

	// version number if ver, version or v, but nothing else
	if (versionFlag || verboseFlag) && (!(isFlagPassed("out") || isFlagPassed("output") || isFlagPassed("d") || isFlagPassed("decimal") || isFlagPassed("x") || isFlagPassed("hex") || isFlagPassed("o") || isFlagPassed("octal") || (isFlagPassed("bin") || isFlagPassed("bit") || isFlagPassed("a") || isFlagPassed("asc") || isFlagPassed("rgb") || isFlagPassed("r") || isFlagPassed("g") || isFlagPassed("b") || isFlagPassed("rgbx")))) {
		fmt.Println(version)
	}

	// Check if the flags passed is consistent with the logic of the program's requirements, if not, we exit
	// short and long version not allowed at the same time
	if isFlagPassed("d") && isFlagPassed("decimal") {
		fmt.Println("Usage of the same type in short and long notation at the same time not supported.")
		os.Exit(1)
	}
	if isFlagPassed("o") && isFlagPassed("octal") {
		fmt.Println("Usage of the same type in short and long notation at the same time not supported.")
		os.Exit(1)
	}
	if isFlagPassed("bin") && isFlagPassed("bit") {
		fmt.Println("Usage of the same type in short and long notation at the same time not supported.")
		os.Exit(1)
	}
	if isFlagPassed("a") && isFlagPassed("asc") {
		fmt.Println("Usage of the same type in short and long notation at the same time not supported.")
		os.Exit(1)
	}
	if (isFlagPassed("rgb")) && (isFlagPassed("r") || isFlagPassed("g") || isFlagPassed("b")) {
		fmt.Println("Usage of the same type (in different notation) at the same time not supported.")
		os.Exit(1)
	}

	counterFlags := 0
	if isFlagPassed("d") || isFlagPassed("decimal") {
		counterFlags += 1
	}
	if isFlagPassed("x") || isFlagPassed("hex") {
		counterFlags += 1
	}
	if isFlagPassed("o") || isFlagPassed("octal") {
		counterFlags += 1
	}
	if isFlagPassed("bin") || isFlagPassed("bit") {
		counterFlags += 1
	}
	if isFlagPassed("a") || isFlagPassed("asc") {
		counterFlags += 1
	}
	if isFlagPassed("rgb") {
		counterFlags += 1
	}
	if isFlagPassed("r") || isFlagPassed("g") || isFlagPassed("b") {
		counterFlags += 1
	}

	if counterFlags > 1 {
		fmt.Printf("Only one input type at a time supported, found %v\n", counterFlags)
		os.Exit(1)
	}

	// handle all numeric types
	if isFlagPassed("x") || isFlagPassed("hex") {
		if !(strings.HasPrefix(hexValue, "0x") || strings.HasPrefix(hexValue, "0X")) {
			tempInt, err = strconv.ParseInt(hexValue, 16, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				decimalValue = int(tempInt)
			}
		} else {
			tempInt, err = strconv.ParseInt(hexValue, 0, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				decimalValue = int(tempInt)
			}
		}
	}

	if isFlagPassed("o") || isFlagPassed("octal") {
		if !(strings.HasPrefix(octValue, "0o") || strings.HasPrefix(octValue, "0O")) {
			tempInt, err = strconv.ParseInt(octValue, 8, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				decimalValue = int(tempInt)
			}
		} else {
			tempInt, err = strconv.ParseInt(octValue, 0, 0)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				decimalValue = int(tempInt)
			}
		}
	}

	if isFlagPassed("bin") || isFlagPassed("bit") {
		tempInt, err = strconv.ParseInt(binValue, 2, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			decimalValue = int(tempInt)
		}
	}

	if isFlagPassed("a") || isFlagPassed("asc") {
		ascValue = string(ascValue[0])
		character := []rune(ascValue)
		decimalValue = int(character[0])
	}

	// handles rgb values
	if isFlagPassed("rgb") {
		// strip from commata and semicolons, leave spaces as separators
		if strings.Contains(rgbValue, ",") {
			rgbValue = strings.ReplaceAll(rgbValue, ",", " ")
		}
		if strings.Contains(rgbValue, ";") {
			rgbValue = strings.ReplaceAll(rgbValue, ";", " ")
		}

		// if not beginning with a letter 'r' we consider the values fine for conversion
		//if rgbValue[0] != 114 && rgbValue[0] != 103 && rgbValue[0] != 98 {
		if strings.Contains(rgbValue, "r") {
			rgbValue = strings.ReplaceAll(rgbValue, "r", " ")
		}
		if strings.Contains(rgbValue, "g") {
			rgbValue = strings.ReplaceAll(rgbValue, "g", " ")
		}
		if strings.Contains(rgbValue, "b") {
			rgbValue = strings.ReplaceAll(rgbValue, "b", " ")
		}

		// as long we have double whitespace remove them
		for {
			if !(strings.Contains(rgbValue, "  ")) {
				break
			}
			rgbValue = strings.ReplaceAll(rgbValue, "  ", " ")
		}

		// trim whitespaces from begining and end
		rgbValue = strings.TrimSpace(rgbValue)

		// split string in substrings
		rgbSlice = strings.Split(rgbValue, " ")

		// check number of substrings, if more than 3 something went wrong
		if len(rgbSlice) != 3 {
			fmt.Println("Problem interpreting values. (Too many whitespaces or non-numerical characters included? too much or not enough values?)")
			os.Exit(1)
		}
		red, errRed = strconv.ParseFloat(rgbSlice[0], 64)
		green, errGreen = strconv.ParseFloat(rgbSlice[1], 64)
		blue, errBlue = strconv.ParseFloat(rgbSlice[2], 64)
		if errRed != nil {
			if verboseFlag {
				fmt.Println("Problem interpreting red value.")
				fmt.Println(errRed)
				fmt.Println("Red value set to 0")
				red = 0.0
			}
		} else if errGreen != nil {
			if verboseFlag {
				fmt.Println("Problem interpreting red value.")
				fmt.Println(errGreen)
				fmt.Println("Red value set to 0")
				green = 0.0
			}
		} else if errBlue != nil {
			if verboseFlag {
				fmt.Println("Problem interpreting red value.")
				fmt.Println(errBlue)
				fmt.Println("Red value set to 0")
				blue = 0.0
			}
		}

		if !((red >= 0 && red < 256) && (green >= 0 && green < 256) && (blue >= 0 && blue < 256)) {
			fmt.Println("Problem interpreting value (out of range 0-255).")
			os.Exit(1)
		}
	} else if isFlagPassed("r") || isFlagPassed("g") || isFlagPassed("b") {
		if (redInt >= 0 && redInt < 256) && (greenInt >= 0 && greenInt < 256) && (blueInt >= 0 && blueInt < 256) {
			red = float64(redInt)
			green = float64(greenInt)
			blue = float64(blueInt)
		} else {
			fmt.Println("Problem interpreting value (out of range 0-255).")
			os.Exit(1)
		}
	}

	if isFlagPassed("rgbx") {
		if strings.Contains(rgbHexValue, "#") {
			rgbHexValue = strings.ReplaceAll(rgbHexValue, "#", "")
		}
		if strings.Contains(rgbHexValue, " ") {
			rgbHexValue = strings.ReplaceAll(rgbHexValue, " ", "")
		}
		// only notation as with 6 digits allowed, no short forms
		if len(rgbHexValue) == 6 {
			rgbHexValueArray = coco.Hex2Rgb(rgbHexValue)
			red = float64(rgbHexValueArray[0])
			green = float64(rgbHexValueArray[1])
			blue = float64(rgbHexValueArray[2])
		} else {
			fmt.Println("Problem interpreting hexadecimal value.")
			os.Exit(1)
		}
	}

	// clean output types
	outputTypes = strings.ReplaceAll(outputTypes, "rgb", "r")

	// range over the output types to maintain order
	if isFlagPassed("out") || isFlagPassed("output") {
		for _, v := range outputTypes {

			// If input is not RGB or rgbx
			if !(isFlagPassed("rgb") || isFlagPassed("r") || isFlagPassed("g") || isFlagPassed("b") || isFlagPassed("rgbx")) {
				// output hex
				if v == 120 {
					if verboseFlag {
						fmt.Print("Hexadecimal: ")
					}
					fmt.Printf("%#x", decimalValue)
					if verboseFlag {
						fmt.Printf(" %x (without preceding \"0x\")\n", decimalValue)
					} else {
						fmt.Print("\n")
					}
				}

				// output dec
				if v == 100 {
					if verboseFlag {
						fmt.Print("Decimal: ")
					}
					fmt.Printf("%d", decimalValue)
					if verboseFlag {
						p := message.NewPrinter(language.English)
						p.Printf(" %d (English notation)\n", decimalValue)
					} else {
						fmt.Print("\n")
					}
				}

				// asked for binary output formatted in 8 digits wide words
				// if verbose is on the output type is mentioned before the output
				// and the number of significant digits is shown
				// output bin
				if v == 98 {
					if verboseFlag {
						fmt.Print("Binary: ")
					}

					if decimalValue > 0 {
						binaryDigits = int(math.Log2(float64(decimalValue)) + 1)
					} else {
						binaryDigits = int(math.Log2(float64(-1*decimalValue)) + 1)
					}

					// differing between different length of outputted binary numbers
					// no need to ouput 64 bit always
					switch {
					case binaryDigits > 0 && binaryDigits <= 8:
						fmt.Printf("%08b", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 8 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					case binaryDigits > 8 && binaryDigits <= 16:
						fmt.Printf("%016b", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 16 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					case binaryDigits > 16 && binaryDigits <= 24:
						fmt.Printf("%024b", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 24 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					case binaryDigits > 24 && binaryDigits <= 32:
						fmt.Printf("%032b", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 32 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					case binaryDigits > 32 && binaryDigits <= 40:
						fmt.Printf("%040b", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 40 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					case binaryDigits > 40 && binaryDigits <= 48:
						fmt.Printf("%048b\n", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 48 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					case binaryDigits > 48 && binaryDigits <= 56:
						fmt.Printf("%056b", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 56 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					case binaryDigits > 56:
						fmt.Printf("%064b", decimalValue)
						if verboseFlag {
							fmt.Printf(" %d significant digits of 64 shown (from right to left).\n", binaryDigits)
						} else {
							fmt.Print("\n")
						}
					}
				}

				// ouput octal
				if v == 111 {
					if verboseFlag {
						fmt.Print("Octal: ")
					}
					fmt.Printf("%O", decimalValue)
					if verboseFlag {
						fmt.Printf(" %o (without preceding \"0o\")\n", decimalValue)
					} else {
						fmt.Print("\n")
					}
				}

				// output char
				if v == 99 {
					if verboseFlag {
						fmt.Print("Char:  ")
					}
					fmt.Printf("%c", decimalValue)
					if verboseFlag {
						fmt.Printf(" %q a single-quoted character literal safely escaped (Go syntax)\n", decimalValue)
					} else {
						fmt.Print("\n")
					}
				}

				// output unicode respresentation
				if v == 117 {
					if verboseFlag {
						fmt.Print("Unicode: ")
					}
					if decimalValue >= 0 {
						fmt.Printf("%U", decimalValue)
					} else {
						fmt.Printf("N/A")
					}

					if verboseFlag {
						if decimalValue >= 0 {
							fmt.Printf(" %q a single-quoted character literal safely escaped (Go syntax)\n", decimalValue)
						} else {
							fmt.Print("\n")
						}
					} else {
						fmt.Print("\n")
					}
				}

				// ouput ASCII respresentation
				if v == 97 {
					asciiString := []string{"NUL", " SOH", " STX", " ETX", " EOT", " ENQ", "ACK", "BEL", "BS", "TAB", "LF", "VT", "FF", "CR", "SO", "SI", "DLE", "DC1", "DC2", "DC3", "DC4", "NAK", "SYN", "ETB", "CAN", "EM", "SUB", "ESC", "FS", "GS", "RS", "US", "SPACE", "DEL"}
					if verboseFlag {
						fmt.Print("ASCII: ")
					}
					if decimalValue >= 0 && decimalValue <= 32 {
						fmt.Printf("%s", asciiString[decimalValue])
					} else if decimalValue == 127 {
						fmt.Printf("%s", asciiString[len(asciiString)-1])
					} else if decimalValue > 32 && decimalValue < 127 {
						fmt.Printf("%c", decimalValue)
					} else {
						fmt.Printf("N/A")
					}
					if verboseFlag {
						if decimalValue >= 0 && decimalValue <= 32 {
							fmt.Printf("%s", asciiString[decimalValue])
						} else if decimalValue == 127 {
							fmt.Printf("%s", asciiString[len(asciiString)-1])
						} else if decimalValue > 32 && decimalValue < 127 {
							fmt.Printf(" %q a single-quoted character literal safely escaped (Go syntax)\n", decimalValue)
						} else {
							fmt.Printf(" N/A\n")
						}
					} else {
						fmt.Print("\n")
					}
				}
				// If input is RGB
			} else if isFlagPassed("rgb") || isFlagPassed("r") || isFlagPassed("g") || isFlagPassed("b") || isFlagPassed("rgbx") {

				// output x hex for rgb
				if v == 120 {
					if verboseFlag {
						fmt.Print("Web/hex: ")
					}
					fmt.Printf("#%s ", strings.ToLower(coco.Rgb2Hex(red, green, blue)))
					if verboseFlag {
						if termcolor.SupportsBasic(os.Stderr) {
							color.HEXStyle("000", strings.ToLower(coco.Rgb2Hex(red, green, blue))).Print("   ")
						}
						fmt.Printf(" rgb(%v,%v,%v) alternative version\n", red, green, blue)
					} else {
						fmt.Print("\n")
					}
				}
				// output w web for rgb
				if v == 119 {
					if verboseFlag {
						fmt.Print("RGB/web: ")
					}
					fmt.Printf("rgb(%v,%v,%v) ", red, green, blue)
					if verboseFlag {
						if termcolor.SupportsBasic(os.Stderr) {
							color.HEXStyle("000", strings.ToLower(coco.Rgb2Hex(red, green, blue))).Print("   ")
						}
						fmt.Printf(" #%s alternative hexadecimal version\n", strings.ToLower(coco.Rgb2Hex(red, green, blue)))
					} else {
						fmt.Print("\n")
					}
				}
			}
		}
	} else {
		// if "o" was omitted. This are the defaults.
		// fmt.Println("DEFAULT")

		// x, o, bin needs to be intepreted before outputting decimal below
		if isFlagPassed("x") || isFlagPassed("hex") {
			if !(strings.HasPrefix(hexValue, "0x") || strings.HasPrefix(hexValue, "0X")) {
				tempInt, err = strconv.ParseInt(hexValue, 16, 0)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else {
					decimalValue = int(tempInt)
				}
			} else {
				tempInt, err = strconv.ParseInt(hexValue, 0, 0)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else {
					decimalValue = int(tempInt)
				}
			}
		}
		if isFlagPassed("o") || isFlagPassed("octal") {
			if !(strings.HasPrefix(octValue, "0o") || strings.HasPrefix(octValue, "0O")) {
				tempInt, err = strconv.ParseInt(octValue, 8, 0)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else {
					decimalValue = int(tempInt)
				}
			} else {
				tempInt, err = strconv.ParseInt(octValue, 0, 0)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else {
					decimalValue = int(tempInt)
				}
			}
		}
		if isFlagPassed("bin") || isFlagPassed("bit") {
			tempInt, err = strconv.ParseInt(binValue, 2, 64)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				decimalValue = int(tempInt)
			}
		}

		// hex is default output for decimal integers
		if isFlagPassed("d") || isFlagPassed("decimal") {
			if verboseFlag {
				fmt.Print("Hexadecimal: ")
			}
			fmt.Printf("%#x", decimalValue)
			if verboseFlag {
				if isFlagPassed("d") || isFlagPassed("decimal") {
					fmt.Printf(" %x (without preceding \"0x\")\n", decimalValue)
				}
			} else {
				fmt.Print("\n")
			}
		}

		// dec is default output for bin, oct, hex
		if (isFlagPassed("bin") || isFlagPassed("bit")) || (isFlagPassed("o") || isFlagPassed("octal")) || (isFlagPassed("x") || isFlagPassed("hex")) {
			if verboseFlag {
				fmt.Print("Decimal: ")
			}
			fmt.Printf("%d", decimalValue)
			if verboseFlag {
				p := message.NewPrinter(language.English)
				p.Printf(" %d (English notation)\n", decimalValue)
			} else {
				fmt.Print("\n")
			}
		}

		// dec is default output for asc
		if isFlagPassed("a") || isFlagPassed("asc") {
			if verboseFlag {
				fmt.Print("Decimal: ")
			}
			fmt.Printf("%d", decimalValue)
			if verboseFlag {
				fmt.Printf(" %#x as hexadecimal value\n", decimalValue)
			} else {
				fmt.Print("\n")
			}
		}

		// x hex is default for rgb
		if isFlagPassed("rgb") || isFlagPassed("r") || isFlagPassed("g") || isFlagPassed("b") {
			if verboseFlag {
				fmt.Print("Web/hex: ")
			}
			fmt.Printf("#%s ", strings.ToLower(coco.Rgb2Hex(red, green, blue)))
			if verboseFlag {
				if termcolor.SupportsBasic(os.Stderr) {
					color.HEXStyle("000", strings.ToLower(coco.Rgb2Hex(red, green, blue))).Print("   ")
				}
				fmt.Printf(" rgb(%v,%v,%v) alternative version\n", red, green, blue)
			} else {
				fmt.Print("\n")
			}
		}

		// rgb web is default for rgbx
		if isFlagPassed("rgbx") {
			if verboseFlag {
				fmt.Print("RGB/web: ")
			}
			fmt.Printf("rgb(%v,%v,%v) ", red, green, blue)
			if verboseFlag {
				if termcolor.SupportsBasic(os.Stderr) {
					color.HEXStyle("000", strings.ToLower(coco.Rgb2Hex(red, green, blue))).Print("   ")
				}
				fmt.Printf(" #%s with preceding \"#\"\n", strings.ToLower(coco.Rgb2Hex(red, green, blue)))
			} else {
				fmt.Print("\n")
			}
		}
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
