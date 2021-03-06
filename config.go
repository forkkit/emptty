package main

import (
	"os"
	"strconv"
	"strings"
)

const (
	confTTYnumber     = "TTY_NUMBER"
	confSwitchTTY     = "SWITCH_TTY"
	confPrintIssue    = "PRINT_ISSUE"
	confDefaultUser   = "DEFAULT_USER"
	confAutologin     = "AUTOLOGIN"
	confLang          = "LANG"
	confDbusLaunch    = "DBUS_LAUNCH"
	confXinitrcLaunch = "XINITRC_LAUNCH"

	pathConfigFile = "/etc/emptty/conf"
)

// config defines structure of application configuration.
type config struct {
	defaultUser   string
	autologin     bool
	tty           int
	switchTTY     bool
	printIssue    bool
	lang          string
	dbusLaunch    bool
	xinitrcLaunch bool
}

// LoadConfig handles loading of application configuration.
func loadConfig() *config {
	c := config{tty: 0, switchTTY: true, printIssue: true, defaultUser: "", autologin: false, lang: "en_US.UTF-8", dbusLaunch: true}

	if fileExists(pathConfigFile) {
		err := readProperties(pathConfigFile, func(key string, value string) {
			switch key {
			case confTTYnumber:
				c.tty = parseTTY(value, "0")
			case confSwitchTTY:
				c.switchTTY = parseBool(value, "true")
			case confPrintIssue:
				c.printIssue = parseBool(value, "true")
			case confDefaultUser:
				c.defaultUser = sanitizeValue(value, "")
			case confAutologin:
				c.autologin = parseBool(value, "false")
			case confLang:
				c.lang = sanitizeValue(value, "en_US.UTF-8")
			case confDbusLaunch:
				c.dbusLaunch = parseBool(value, "true")
			case confXinitrcLaunch:
				c.xinitrcLaunch = parseBool(value, "false")
			}
		})
		handleErr(err)
	}

	os.Unsetenv(confTTYnumber)
	os.Unsetenv(confSwitchTTY)
	os.Unsetenv(confPrintIssue)
	os.Unsetenv(confDefaultUser)
	os.Unsetenv(confAutologin)
	os.Unsetenv(confDbusLaunch)

	return &c
}

// Parse TTY number.
func parseTTY(tty string, defaultValue string) int {
	val, err := strconv.ParseInt(sanitizeValue(tty, defaultValue), 10, 32)
	if err != nil {
		return 0
	}
	return int(val)
}

// Parse boolean values.
func parseBool(strBool string, defaultValue string) bool {
	val, err := strconv.ParseBool(sanitizeValue(strBool, defaultValue))
	if err != nil {
		return false
	}
	return val
}

// Sanitize value.
func sanitizeValue(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return strings.TrimSpace(value)
}

// Returns TTY number converted to string
func (c *config) strTTY() string {
	return strconv.Itoa(c.tty)
}
