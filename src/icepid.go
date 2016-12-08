/*
 * Simple server dashboard
 *
 * Copyright 2016 by general-error
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
}

var appVersion string = "1.0.0"
var confVersion int = 1
var settings = Settings{}

func load_settings() Settings {
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

    log.Fatal("Error: load_settings\n", err)
    return configuration
}

func main() {
    settings = load_settings()
    http.HandleFunc("/", login_handler)
    http.HandleFunc("/index", index_handler)
    http.HandleFunc("/logout", logout_handler)
    http.HandleFunc("/log", log_view_handler)
    http.HandleFunc("/dmesg", dmesg_view_handler)
    log.Fatal(http.ListenAndServe(settings.Port, nil))
}
