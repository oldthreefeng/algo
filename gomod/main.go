/*
@Time : 2019/9/8 1:35
@Author : louis
@File : main
@Software: GoLand
*/

package main

import log "github.com/sirupsen/logrus"

func main() {
	log.Info("this is a log log")
	log.Warn("this is an another log log")
	log.Fatal("this a bad log")
}
