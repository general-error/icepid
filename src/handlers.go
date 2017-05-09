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
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func checkAccess(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie(settings.CookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := r.Cookie(settings.CookieName)
		if err == nil {
			// cookie set
			http.Redirect(w, r, "/index", http.StatusFound)
		} else {
			t, err := template.New("login").Parse(loginTpl)
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
				Name:   settings.CookieName,
				Value:  settings.CookieId,
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

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   settings.CookieName,
		Value:  settings.CookieId,
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusFound)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	checkAccess(w, r)
	logs := settings.Logs
	updates := getUpdates()
	uptime := getUptime()
	free := getFree()
	date := getDate()
	wVal := getW()
	df := getDf()
	sensors := getSensors()
	disks := settings.Disks

	page := IndexPage{appVersion, logs, updates, uptime, free, date, wVal, df, sensors, disks}

	t, err := template.New("index").Parse(indexTpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(w, page)
	if err != nil {
		log.Fatal(err)
	}
}

func logHandler(w http.ResponseWriter, r *http.Request) {
	checkAccess(w, r)

	logNum, err := strconv.Atoi(r.URL.Query().Get("log"))

	if err != nil || logNum < 0 || logNum > len(settings.Logs) {
		http.Redirect(w, r, "/index", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "%v", getLog(settings.Logs[logNum]))
}

func dmesgHandler(w http.ResponseWriter, r *http.Request) {
	checkAccess(w, r)

	fmt.Fprintf(w, "%v", getDmesg())
}

func smartHandler(w http.ResponseWriter, r *http.Request) {
	checkAccess(w, r)

	disk := r.URL.Query().Get("disk")

	if disk != settings.Disks[0] {
		http.Redirect(w, r, "/index", http.StatusFound)
		return
	}

	fmt.Fprintf(w, "%v", getSmart(disk))
}
