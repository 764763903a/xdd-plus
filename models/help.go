package models

import (
	"fmt"
	"strings"
)

func getVhelpRule(num int) string {
	rules := ""
	var codes = map[string][]string{
		"Fruit":        {},
		"Pet":          {},
		"Bean":         {},
		"JdFactory":    {},
		"DreamFactory": {},
		"Jxnc":         {},
		"Jdzz":         {},
		"Joy":          {},
		"Sgmh":         {},
		"Cfd":          {},
		"Cash":         {},
	}
	cks := GetJdCookies()
	for _, ck := range cks {
		if ck.Help == True || Cdle {
			for k := range codes {
				switch k {
				case "Fruit":
					codes[k] = append(codes[k], ck.Fruit)
				case "Pet":
					codes[k] = append(codes[k], ck.Pet)
				case "Bean":
					codes[k] = append(codes[k], ck.Bean)
				case "JdFactory":
					codes[k] = append(codes[k], ck.JdFactory)
				case "DreamFactory":
					codes[k] = append(codes[k], ck.DreamFactory)
				case "Jxnc":
					codes[k] = append(codes[k], ck.Jxnc)
				case "Jdzz":
					codes[k] = append(codes[k], ck.Jdzz)
				case "Joy":
					codes[k] = append(codes[k], ck.Joy)
				case "Sgmh":
					codes[k] = append(codes[k], ck.Sgmh)
				case "Cfd":
					codes[k] = append(codes[k], ck.Cfd)
				case "Cash":
					codes[k] = append(codes[k], ck.Cash)
				}
				if len := len(codes[k]); len != 0 {
					if codes[k][len-1] == "undefined" || codes[k][len-1] == "" || codes[k][len-1] == "--" {
						codes[k] = codes[k][:len-1]
					}
				}
			}
		}
	}
	for k := range codes {
		for i, code := range codes[k] {
			code = strings.Replace(code, `"`, `\"`, -1)
			rules += fmt.Sprintf("My%s%d=\"%s\"\n", k, i+1, code)
			codes[k][i] = fmt.Sprintf("${My%s%d}", k, i+1)
		}
	}
	for k := range codes {
		for i := 0; i < num; i++ {
			if len(codes[k]) > 0 {
				rules += fmt.Sprintf("ForOther"+k+"%d=\"%s\"\n", i+1, strings.Join(codes[k], "@"))
			}
		}
	}
	return rules
}

func getQLHelp(num int) map[string]string {
	var codes = map[string][]string{
		"Fruit":        {},
		"Pet":          {},
		"Bean":         {},
		"JdFactory":    {},
		"DreamFactory": {},
		"Jxnc":         {},
		"Jdzz":         {},
		"Joy":          {},
		"Sgmh":         {},
		"Cfd":          {},
		"Cash":         {},
	}
	cks := GetJdCookies()
	for _, ck := range cks {
		if ck.Help == True || Cdle {
			for k := range codes {
				switch k {
				case "Fruit":
					codes[k] = append(codes[k], ck.Fruit)
				case "Pet":
					codes[k] = append(codes[k], ck.Pet)
				case "Bean":
					codes[k] = append(codes[k], ck.Bean)
				case "JdFactory":
					codes[k] = append(codes[k], ck.JdFactory)
				case "DreamFactory":
					codes[k] = append(codes[k], ck.DreamFactory)
				case "Jxnc":
					codes[k] = append(codes[k], ck.Jxnc)
				case "Jdzz":
					codes[k] = append(codes[k], ck.Jdzz)
				case "Joy":
					codes[k] = append(codes[k], ck.Joy)
				case "Sgmh":
					codes[k] = append(codes[k], ck.Sgmh)
				case "Cfd":
					codes[k] = append(codes[k], ck.Cfd)
				case "Cash":
					codes[k] = append(codes[k], ck.Cash)
				}
				if len := len(codes[k]); len != 0 {
					if codes[k][len-1] == "undefined" || codes[k][len-1] == "" || codes[k][len-1] == "--" {
						codes[k] = codes[k][:len-1]
					}
				}
			}
		}
	}
	var e = map[string]string{
		"Fruit":        "",
		"Pet":          "",
		"Bean":         "",
		"JdFactory":    "",
		"DreamFactory": "",
		"Jxnc":         "",
		"Jdzz":         "",
		"Joy":          "",
		"Sgmh":         "",
		"Cfd":          "",
		"Cash":         "",
	}
	for k := range codes {
		vv := codes[k]
		for i := range vv {
			vv[i] = strings.Replace(vv[i], `"`, `\"`, -1)

		}
		e[k] += strings.Join(vv, "@")
	}
	for k := range e {
		n := []string{}
		for i := 0; i < num; i++ {
			n = append(n, e[k])
		}
		e[k] = strings.Join(n, "&")
	}
	var f = map[string]string{}
	for k := range e {
		switch k {
		case "Fruit":
			f["FRUITSHARECODES"] = e[k]
		case "Pet":
			f["PETSHARECODES"] = e[k]
		case "Bean":
			f["PLANT_BEAN_SHARECODES"] = e[k]
		case "JdFactory":
			f["DDFACTORY_SHARECODES"] = e[k]
		case "DreamFactory":
			f["DREAM_FACTORY_SHARE_CODES"] = e[k]
		case "Jxnc":
			f["JXNC_SHARECODES"] = e[k]
		// case "Jdzz":
		// 	f[k] = e[k]
		// case "Joy":
		// 	f[k] = e[k]
		case "Sgmh":
			f["JDSGMH_SHARECODES"] = e[k]
		// case "Cfd":
		// 	f[k] = e[k]
		case "Cash":
			f["JD_CASH_SHARECODES"] = e[k]
		}
	}
	return f
}
