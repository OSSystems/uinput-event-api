/*
 * Copyright (C) 2019
 * O.S. Systems Sofware LTDA: contato@ossystems.com.br
 *
 * SPDX-License-Identifier: MIT
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/bendahl/uinput"
	"github.com/labstack/echo"
)

var keyMap = map[string]byte{
	"ESC":   0x01,
	"ENTER": 0x1C,
	"A":     0x1E,
	"R":     0x13,
	"Z":     0x2C,
	"HOME":  0x66,
	"UP":    0x67,
	"LEFT":  0x69,
	"RIGHT": 0x6A,
	"DOWN":  0x6C,
}

func main() {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("keyboard"))
	if err != nil {
		logrus.Fatal(err)
	}

	defer keyboard.Close()

	e := echo.New()
	e.HideBanner = true

	e.POST("/keypress", func(c echo.Context) error {
		var req struct {
			Key string `json:"key"`
		}

		if err = c.Bind(&req); err != nil {
			return err
		}

		key, ok := keyMap[strings.ToUpper(req.Key)]
		if !ok {
			return c.NoContent(http.StatusInternalServerError)
		}

		keyboard.KeyPress(int(key))

		return c.NoContent(http.StatusOK)
	})

	log.Fatal(e.Start(fmt.Sprintf(":%d", 8085)))
}
