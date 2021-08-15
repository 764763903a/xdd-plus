package models

// import (
// 	"reflect"
// 	"strconv"
// 	"strings"

// 	"github.com/boltdb/bolt"
// )

// type Env struct {
// 	ID    int
// 	Name  string
// 	Value string
// }

// type Bd struct {
// 	name  string
// 	model interface{}
// }

// func (sf *Bd) Name(name string) error {
// 	sf.name = name
// 	return nil
// }

// func (sf *Bd) GetAll(is []interface{}) error {
// 	db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(sf.name))
// 		b.ForEach(func(_, v []byte) error {
// 			i := is[0]
// 			var _v = reflect.ValueOf(&ck).Elem()
// 			for _, vv := range strings.Split(string(v), ";") {
// 				v := strings.Split(vv, "=")
// 				if len(v) == 2 {
// 					t := _v.FieldByName(v[0])
// 					if t.CanSet() {
// 						switch t.Kind() {
// 						case reflect.Int:
// 							i, _ := strconv.Atoi(v[1])
// 							t.SetInt(int64(i))
// 						case reflect.String:
// 							t.SetString(v[1])
// 						}
// 					}
// 				}
// 			}
// 			is = append(is, i)
// 			return nil
// 		})
// 		return nil
// 	})
// 	return nil
// }
