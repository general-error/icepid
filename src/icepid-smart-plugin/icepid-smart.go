/*
 * SMART data reader
 *
 * Copyright 2017 by general-error
 *
 * This file is part of Icepid.
 *
 * Icepid is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Icepid is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with Icepid.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
)

const appVersion string = "1.0.0"

var reg = regexp.MustCompile(`^[[:alnum:]]+$`)

func main() {
	var v = flag.Bool("v", false, "show program version and exit")
	var s = flag.String("d", "", "device file name like sda")
	flag.Parse()

	if *v {
		fmt.Println("icepid-smart version " + appVersion)
		os.Exit(0)
	}

	if *s == "" {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if !reg.MatchString(*s) {
		log.Print("Error: wrong drive identifier " + *s)
		os.Exit(1)
	}

	out, err := exec.Command("smartctl", "-a", "/dev/"+*s).Output()

	if err != nil {
		log.Print("Error: cannot read SMART data\n", err)
		os.Exit(1)
	}

	fmt.Print(string(out))
}
