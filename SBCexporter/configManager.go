package main

import (
	"fmt"
	//"os"
	"gopkg.in/yaml.v2"
    //"flag"
   // "log"
   "io/ioutil"
)
// Template used for struct and the functions NewConfig(), ValidateConfigPath() and ParseFlags() are copied from:
// https://dev.to/koddr/let-s-write-config-for-your-golang-web-app-on-right-way-yaml-5ggp
    type Config struct {
        Hosts []Host
        Authtimeout int `yaml:"authtimeout"`
    }
        type Host struct {
            HostName       string `yaml:"hostname"`
            Ipaddress      string `yaml:"ipaddress"`
            Username       string `yaml:"username"`
            Password       string `yaml:"password"`
            //exclude        string `yaml:"exclude"`
                Exclude struct {
                    // Server is the general server timeout to use
                    // for graceful shutdowns
                    SystemExporter bool `yaml:"systemstats"`
                    CallStats      bool `yaml:"callstats"`
                }`yaml:"exclude"`
            }


            func getConf(c *Config) *Config {

                yamlFile, err := ioutil.ReadFile("config.yml")
                if err != nil {
                    //log.Printf("yamlFile.Get err   #%v ", err)
                    fmt.Println("yamlFile.Get err   # ", err)
                }
                err = yaml.Unmarshal(yamlFile, c)
                if err != nil {
                   // log.Fatalf("Unmarshal: %v", err)
                    fmt.Println("yamlFile.Get err   # ", err)
                }
                return c
            }

func getIpAdrExp(exporterName string) []string{
    cfg := getConf(&Config{})

	var list []string
    switch exporterName {
        case "systemStats":
           for i := range cfg.Hosts {
            //for i := 0; i < len(cfg.Hosts); i++ {
                if (cfg.Hosts[i].Exclude.SystemExporter == false) {
                    list = append(list, cfg.Hosts[i].Ipaddress)
                }
            }
        case "callStats":
            for i:= range cfg.Hosts {
                if (cfg.Hosts[i].Exclude.CallStats == false) {
                    list = append(list, cfg.Hosts[i].Ipaddress)
                }
            }
            //INFO: have a switch case on all exporters made, NB!: must remember exact exporternames inside each exporter
        }
return list
}

func getAuth(ipadr string) (username string, password string) {
    var u, p string
    cfg := getConf(&Config{})

    for i:= range cfg.Hosts {
        if (cfg.Hosts[i].Ipaddress == ipadr) {
            u, p = cfg.Hosts[i].Username, cfg.Hosts[i].Password
        }
    }
   // return "test", "test"
    return u,p
}

func test() {
    ip := getIpAdrExp("systemStats")
    fmt.Println(ip)
}