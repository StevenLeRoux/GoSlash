package redirect

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Apply(target string, list []string) string {

	// check list len, if 1 => only alias, no more args
	re := regexp.MustCompile(`\${([0-9]+:[^}]+)}`)
	m := make(map[int]string)
	for i := 1; i < len(list); i++ {
		m[i] = list[i]
	}
	fmt.Println(m)

	plist := re.FindAllStringSubmatch(target, -1)

	var key string
	var defvalue string
	var altvalue bool
	var onlyifnotset bool

	for i := range plist {
		key = ""
		defvalue = ""
		expr := plist[i][1]

		if strings.Contains(expr, ":") {
			key = expr[:strings.Index(expr, ":")]
			defvalue = expr[strings.Index(expr, ":"):][1:]

			if strings.HasPrefix(defvalue, "+") {
				altvalue = true
				defvalue = defvalue[1:]
			} else if strings.HasPrefix(defvalue, "-") {
				altvalue = false
				defvalue = defvalue[1:]
			} else if strings.HasPrefix(defvalue, "!") {
				onlyifnotset = true
				altvalue = false
				defvalue = defvalue[1:]
			} else {
				altvalue = false
			}

		} else {
			key = expr
		}

		// Support the following syntaxes
		//
		// ${X:-DEFAULT}  If parameter X is not set, use 'DEFAULT', otherwise substitute parameter X
		// ${X:!DEFAULT}  If parameter X is not set, use 'DEFAULT', otherwise use the empty string ''
		// ${X:+DEFAULT}  If parameter X is set, use 'DEFAULT' where $X was replaced by value of X
		// ${X} is equivalent to ${X:-}

		keyint, _ := strconv.Atoi(key)
		if v, ok := m[keyint]; ok {
			//  Positional parameter was set, replace occurences of ${X[:...]} with its value.
			if defvalue != "" && altvalue {

				defvalue = strings.Replace(defvalue, "$"+key, v, -1)
				target = strings.Replace(target, "${"+expr+"}", defvalue, -1)
			} else {
				if !onlyifnotset {
					target = strings.Replace(target, "${"+expr+"}", v, -1)
				} else {
					target = strings.Replace(target, "${"+expr+"}", "", -1)
				}

			}
		} else {
			if defvalue != "" && !altvalue {
				target = strings.Replace(target, "${"+expr+"}", defvalue, -1)
			} else {
				target = strings.Replace(target, "${"+expr+"}", "", -1)
			}

		}
	}

	return target
}
