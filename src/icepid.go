/*
 * Simple server dashboard
 *
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
    "net/http"
    "log"
    "encoding/json"
    "os"
)

type Settings struct {
    Version int
    Port string
    CookieMaxAge int    // in seconds
    CookieId string
    CookieName string
    Password string
    Logs []string
    Disks []string
}

type IndexPage struct {
    Version string
    Logs []string
    Updates string
    Uptime string
    Free string
    Date string
    W string
    Df string
    Sensors string
    Disks []string
}

type LoginPage struct {
    ErrorMessage string
}

const appVersion string = "1.1.1"
const confVersion int = 2
var loginAttempts int
var settings = Settings{}
var loginPage = LoginPage{}

func loadSettings() Settings {
    file, _ := os.Open("icepid.json")
    decoder := json.NewDecoder(file)
    configuration := Settings{}
    err := decoder.Decode(&configuration)

    if configuration.Version != confVersion {
        log.Fatal("Error: configuration file version mismatch\n")
    }

    if err == nil {
        return configuration
    }

    log.Fatal("Error: loadSettings\n", err)
    return configuration
}

func main() {
    settings = loadSettings()
    http.HandleFunc("/", loginHandler)
    http.HandleFunc("/index", indexHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/log", logHandler)
    http.HandleFunc("/dmesg", dmesgHandler)
    http.HandleFunc("/disk", smartHandler)
    log.Fatal(http.ListenAndServe(settings.Port, nil))
}
