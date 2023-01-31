/**
 * @author mch
 */

package util

import (
	"flag"
	"fmt"
)

var (
	instance *Config
	conf = flag.String("conf","../etc/config.json","config.....")
)
func init(){
	flag.Parse()
	fmt.Println(*conf)
}

type Config struct {
	BasePath string `json:"base_path"`
	DataPath string `json:"data_path"`
}
