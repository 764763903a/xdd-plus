package models

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/gorm"
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
			code = strings.Replace(code, `\n`, ``, -1)
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

func WriteHelpJS(acks []JdCookie) {
	cks := GetJdCookies(func(sb *gorm.DB) *gorm.DB {
		return sb.Where(fmt.Sprintf("%s = ?", Help), True)
	})
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
	for _, ck := range cks {
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
				if codes[k][len-1] == "undefined" || codes[k][len-1] == "" {
					codes[k] = codes[k][:len-1]
				}
			}
		}
	}
	var e = map[string][]string{
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
	var f = func(ss []string, s string) string {
		tss := []string{}
		for _, v := range ss {
			if v != s {
				tss = append(tss, v)
			}
		}
		return `'` + strings.Join(tss, "@") + `'`
	}

	for k := range codes {
		for _, ck := range acks {
			switch k {
			case "Fruit":
				e[k] = append(e[k], f(codes[k], ck.Fruit))
			case "Pet":
				e[k] = append(e[k], f(codes[k], ck.Pet))
			case "Bean":
				e[k] = append(e[k], f(codes[k], ck.Bean))
			case "JdFactory":
				e[k] = append(e[k], f(codes[k], ck.JdFactory))
			case "DreamFactory":
				e[k] = append(e[k], f(codes[k], ck.DreamFactory))
			case "Jxnc":
				e[k] = append(e[k], f(codes[k], ck.Jxnc))
			case "Jdzz":
				e[k] = append(e[k], f(codes[k], ck.Jdzz))
			case "Joy":
				e[k] = append(e[k], f(codes[k], ck.Joy))
			case "Sgmh":
				e[k] = append(e[k], f(codes[k], ck.Sgmh))
			case "Cfd":
				e[k] = append(e[k], f(codes[k], ck.Cfd))
			case "Cash":
				e[k] = append(e[k], f(codes[k], ck.Cash))
			}
			if len := len(codes[k]); len != 0 {
				if codes[k][len-1] == "undefined" || codes[k][len-1] == "" {
					codes[k] = codes[k][:len-1]
				}
			}
		}
	}
	tpl := `let codes = [%s];
for (let i = 0; i < codes.length; i++) {
	const index = (i + 1 === 1) ? '' : (i + 1);
	exports['%s' + index] = codes[i];
}`
	for k, codes := range e {
		switch k {
		case "Fruit":
			WriteToFile(
				ExecPath+"/scripts/jdFruitShareCodes.js",
				fmt.Sprintf(tpl, strings.Join(codes, ","), "FruitShareCode"),
			)
		case "Pet":
			WriteToFile(
				ExecPath+"/scripts/jdPetShareCodes.js",
				fmt.Sprintf(tpl, strings.Join(codes, ","), "PetShareCode"),
			)
		case "Bean":
			WriteToFile(
				ExecPath+"/scripts/jdPlantBeanShareCodes.js",
				fmt.Sprintf(tpl, strings.Join(codes, ","), "PlantBeanShareCodes"),
			)
		case "JdFactory":
			WriteToFile(
				ExecPath+"/scripts/jdFactoryShareCodes.js",
				fmt.Sprintf(tpl, strings.Join(codes, ","), "shareCodes.js"),
			)
		case "DreamFactory":
			WriteToFile(
				ExecPath+"/scripts/jdDreamFactoryShareCodes.js",
				fmt.Sprintf(tpl, strings.Join(codes, ","), "shareCodes.js"),
			)
		case "Jxnc":
			WriteToFile(
				ExecPath+"/scripts/jdJxncShareCodes.js",
				fmt.Sprintf(tpl, strings.Join(codes, ","), "JxncShareCode.js"),
			)
		case "Jdzz":

		case "Joy":

		case "Sgmh":

		case "Cfd":

		case "Cash":

		}
	}
}

func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt([]byte(content), n)
		// fmt.Println("write succeed!")
		defer f.Close()
	}
	return err
}
