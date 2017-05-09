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

func wrap(base string) string {
    const header =
`<html>
    <head><title>Icepid</title></head>
    <body>
    <style type="text/css">
    .container {
        display: flex;
        align-items: center;
        justify-content: center;
        height: 97vh; /* 100vh causes scrollbar in firefox */
        position: relative;
    }
    .title {
        align-self: flex-start;
        left: 0;
        top: 0;
        position: absolute;
        font-size: 300%;
    }
    .error {
        position: absolute;
        right: 0;
        top: 0;
        padding: 10px;
        font-size: 200%;
        color: red;
    }
    .letter::first-letter {
            color: #0000ff;
            font-size: 200%;
    }
    .header {
        display: flex;
        justify-content: space-between;
        background-color: #b6dcfc;
        margin: -8px;
        padding: 8px;
    }
    .cell {
        border: 1px solid black;
        display: inline-block;
        padding: 3px;
    }
    .bottom {
        position: fixed;
        font-size: small;
        background-color: #b6dcfc;
        bottom: 0;
        right: 0;
        left: 0;
        padding: 3px;
        height: 15px;
    }
    .main {
        margin-bottom: 30px;
    }
    </style>
    `
    const footer =
`</body></html>`

    return header + base + footer
}

var loginTpl string = wrap(
    `<div class="container">
        <div class="title">
            <p class="letter">Icepid</p>
        </div>
        <div>
            <form action="/" method="post">
                <input name="password" type="password" autofocus required >
                <input type="submit" value="Login">
            </form>
        </div>
        <div class="error">
            {{.ErrorMessage}}
        </div>
    </div>
    `)

var indexTpl string = wrap(
    `<div class="header">
        <form action="/logout" method="post">
            <input type="submit" value="Logout">
        </form>
        <div>{{.Uptime}}</div>
        <div>{{.Date}}</div>
    </div>
    <div class="main">
        <p>Updates:</p>
        <pre>{{.Updates}}</pre>
        <hr>
        <pre class="cell">{{.W}}</pre><br>
        <pre class="cell">{{.Free}}</pre><br>
        <pre class="cell">{{.Df}}</pre><br>
        <pre class="cell">{{.Sensors}}</pre><br>
        <hr>
        <p>SMART:</p>
        <ul>
            {{range $key, $val := .Disks}}
            <li><a href="/disk?disk={{$val}}">{{$val}}</a></li>
            {{end}}
        </ul>
        <p>Log files:</p>
        <ul>
            <li><a href="/dmesg">dmesg</a></li>
            {{range $key, $val := .Logs}}
            <li><a href="/log?log={{$key}}">{{$val}}</a></li>
            {{end}}
        </ul>
    </div>
    <div class="bottom">v{{.Version}}</div>
     `)
