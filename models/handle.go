package models

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"gorm.io/gorm"
)

func initHandle() {
	//获取路径
	Save = make(chan *JdCookie)
	go func() {
		init := true
		for {
			get := <-Save
			if get.Pool == "s" {
				initCookie()
				continue
			}
			cks := GetJdCookies(func(sb *gorm.DB) *gorm.DB {
				return sb.Where(fmt.Sprintf("%s >= ? and %s != ?", Priority, Hack), 0, True)
			})
			tmp := []JdCookie{}
			for _, ck := range cks {
				if ck.Priority >= 0 && ck.Hack != True {
					tmp = append(tmp, ck)
				}
			}
			cks = tmp
			cookies := "{"
			hh := []string{}
			for i, ck := range cks {
				hh = append(hh,
					fmt.Sprintf("CookieJD%d:'pt_key=%s;pt_pin=%s;'", i+1, ck.PtKey, ck.PtPin),
				)
			}
			cookies += strings.Join(hh, ",")
			cookies += "}"
			f, err := os.OpenFile(ExecPath+"/scripts/jdCookie.js", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
			if err != nil {
				logs.Warn("创建jdCookie.js失败，", err)
			}

			f.WriteString(fmt.Sprintf(`
var cookies = %s
var pins = process.env.pins
if(pins){
	pins = pins.split("&")
	for (var key in cookies) {
	    c = false
	    for (var pin of pins) {
		   if (pin && cookies[key].indexOf(pin) != -1) {
			  c = true
			  break
		   }
	    }
	    if (!c) {
		   delete cookies[key]
	    }
	}
}
module.exports = cookies`, cookies))
			f.Close()
			WriteHelpJS(cks)
			go CopyConfigAll()
			// tmp = []JdCookie{}
			// for _, ck := range cks {
			// 	if ck.Hack != True {
			// 		tmp = append(tmp, ck)
			// 	}
			// }
			// cks = tmp
			if Config.Mode == Parallel {
				for i := range Config.Containers {
					(&Config.Containers[i]).read()
				}
				for i := range Config.Containers {
					(&Config.Containers[i]).write(cks)
				}
			} else {
				resident := []JdCookie{}
				if Config.Resident != "" {
					tmp := cks
					cks = []JdCookie{}
					for _, ck := range tmp {
						if strings.Contains(Config.Resident, ck.PtPin) {
							resident = append(resident, ck)
						} else {
							cks = append(cks, ck)
						}
					}
				}
				type balance struct {
					Container Container
					Weigth    float64
					Ready     []JdCookie
					Should    int
				}
				availables := []Container{}
				parallels := []Container{}
				bs := []balance{}
				for i := range Config.Containers {
					(&Config.Containers[i]).read()
					if Config.Containers[i].Available {
						if Config.Containers[i].Mode == Parallel {
							parallels = append(parallels, Config.Containers[i])
						} else {
							availables = append(availables, Config.Containers[i])
							bs = append(bs, balance{
								Container: Config.Containers[i],
								Weigth:    float64(Config.Containers[i].Weigth),
							})
						}
					}
				}
				bat := cks
				for {
					left := []JdCookie{}
					l := len(cks)
					total := 0.0
					for i := range bs {
						total += float64(bs[i].Weigth)
					}
					for i := range bs {
						if bs[i].Weigth == 0 {
							bs[i].Should = 0
						} else {
							bs[i].Should = int(math.Ceil(bs[i].Weigth / total * float64(l)))
						}

					}
					a := 0
					for i := range bs {
						j := bs[i].Should
						if j == 0 {
							continue
						}
						s := 0
						if bs[i].Container.Limit > 0 && j > bs[i].Container.Limit {
							s = a + bs[i].Container.Limit
							left = append(left, cks[s:a+j]...)
							bs[i].Weigth = 0
						} else {
							s = a + j
						}
						if s > l {
							s = l
						}
						bs[i].Ready = append(bs[i].Ready, cks[a:s]...)
						a += j
						if a >= l-1 {
							break
						}
					}
					if len(left) != 0 {
						cks = left
						continue
					}
					break
				}
				for i := range bs {
					bs[i].Container.write(append(resident, bs[i].Ready...))
				}
				for i := range parallels {
					parallels[i].write(append(resident, bat...))
				}
			}
			if init {
				go func() {
					for {
						Save <- &JdCookie{
							Pool: "s",
						}

						time.Sleep(time.Minute * 30)
						// time.Sleep(time.Second * 1)
					}
				}()
				init = false
			}
		}
	}()
}
