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
    "net/http"
    "log"
    "html/template"
    "fmt"
    "strconv"
    "time"
)

func check_access(w http.ResponseWriter, r *http.Request) {
    _, err := r.Cookie(settings.CookieName)
    if err != nil {
        http.Redirect(w, r, "/", http.StatusFound)
    }
}

func login_handler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        _, err := r.Cookie(settings.CookieName)
        if err == nil {
            // cookie set
            http.Redirect(w, r, "/index", http.StatusFound)
        } else {
            t, err := template.New("login").Parse(login_tpl)
            if err != nil {
               log.Fatal(err)
            }
            t.Execute(w, loginPage)
        }
    } else {
        if r.FormValue("password") == settings.Password {
            // password right
            loginAttempts = 0
            loginPage.ErrorMessage = ""
            cookie := http.Cookie{
                Name: settings.CookieName,
                Value: settings.CookieId,
                MaxAge: settings.CookieMaxAge,
            }
            http.SetCookie(w, &cookie)
            http.Redirect(w, r, "/index", http.StatusFound)
        } else {
            // password wrong
            loginPage.ErrorMessage = "Wrong password!"
            if loginAttempts >= 5 {
                time.Sleep(20 * time.Second)
            } else {
                loginAttempts++
                http.Redirect(w, r, "/", http.StatusFound)
            }
        }
    }
}

func logout_handler(w http.ResponseWriter, r *http.Request) {
    cookie := http.Cookie{
        Name: settings.CookieName,
        Value: settings.CookieId,
        MaxAge: -1,
    }
    http.SetCookie(w, &cookie)
    http.Redirect(w, r, "/", http.StatusFound)
}

func index_handler(w http.ResponseWriter, r *http.Request) {
    check_access(w, r)
    logs := settings.Logs;
    updates := list_updates()
    uptime := get_uptime()
    free := get_free()
    date := get_date()
    w_val := get_w()
    df := get_df()
    sensors := get_sensors()
    disks := settings.Disks

    page := IndexPage{ appVersion, logs, updates, uptime, free, date, w_val, df, sensors, disks }
    
    t, err := template.New("index").Parse(index_tpl)
    if err != nil {
        log.Fatal(err)
    }
    err = t.Execute(w, page)
    if err != nil {
        log.Fatal(err)
    }
}

func log_view_handler(w http.ResponseWriter, r *http.Request) {
    check_access(w, r)
    
    logNum, err := strconv.Atoi(r.URL.Query().Get("log"))

    if (err != nil || logNum < 0 || logNum > len(settings.Logs)) {
        http.Redirect(w, r, "/index", http.StatusFound)
        return
    }
    
    fmt.Fprintf(w, "%v", query_log(settings.Logs[logNum]))
}

func dmesg_view_handler(w http.ResponseWriter, r *http.Request) {
    check_access(w, r)

    fmt.Fprintf(w, "%v", query_dmesg())
}

func smart_handler(w http.ResponseWriter, r *http.Request) {
    check_access(w, r)

    disk := r.URL.Query().Get("disk")

    if (disk != settings.Disks[0]) {
        http.Redirect(w, r, "/index", http.StatusFound)
		return
    }

    fmt.Fprintf(w, "%v", query_smart(disk))
}
