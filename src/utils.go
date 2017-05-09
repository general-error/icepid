/*
 * Copyright 2016 2017 by general-error
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
    "io/ioutil"
    "os/exec"
    "bytes"
    "log"
    "strings"
    "time"
)

func query_log(path string) string {
    file, err := ioutil.ReadFile(path)
    
    if err != nil {
        log.Fatal(err)
    }

    return string(file);
}

func list_updates() string {
    cmd := exec.Command("apt",  "list", "--upgradable")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Start()
    err = cmd.Wait()

    var output string

    if err == nil {
        output = out.String()
    }

    if strings.Compare(output, "Listing...\n") == 0 {
        return "No updates.";
    } else {
        return output[len("Listing..."):len(output)];
    }

    log.Print("Error: list_updates\n", err)
    return "Error listing updates"
}

func get_uptime() string {
    out, err := exec.Command("uptime", "-p").Output()

    if err == nil {
        return string(out)
    }

    log.Print("Error: get_uptime\n", err)
    return ""
}

func query_dmesg() string {
    out, err := exec.Command("dmesg").Output()

    if err == nil {
        return string(out)
    }

    log.Print("Error: query_dmesg\n", err)
    return ""
}

func get_free() string {
    out, err := exec.Command("free", "-h").Output()

    if err == nil {
        return string(out)
    }

    log.Print("Error: get_free\n", err)
    return ""
}

func get_date() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

func get_w() string {
    out, err := exec.Command("w").Output()

    if err == nil {
        return string(out)
    }

    log.Print("Error: get_w\n", err)
    return ""
}

func get_df() string {
    out, err := exec.Command("df", "-h", "/", "/home").Output()

    if err == nil {
        return string(out)
    }

    log.Print("Error: get_df\n", err)
    return ""
}

func get_sensors() string {
    out, err := exec.Command("sensors").Output()

    if err == nil {
        return string(out)
    }

    log.Print("Error: get_sensors\n", err)
    return ""
}

func query_smart(disk string) string {
    out, err := exec.Command("./icepid-smart", "-d", disk).Output()

    if err == nil {
        return string(out)
    }

    log.Print("Error: query_smart\n", err)
    return ""
}
