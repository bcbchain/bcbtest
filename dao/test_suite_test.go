package dao

import (
	"bcbtest/config"
	"fmt"
	"testing"
)

func TestGetFirstSuiteName(t *testing.T) {
	config.LoadConfig("/home/deploy/bcbtest/conf/config.yaml")
	Init()
	name := GetFirstSuiteNames(10001)
	fmt.Println(name)
}

func TestGetSecondSuiteNames(t *testing.T) {
	config.LoadConfig("/home/deploy/bcbtest/conf/config.yaml")
	Init()
	name := GetSecondSuiteNames("checktx流程")
	fmt.Println(name)
}

func TestGetThirdSuiteNames(t *testing.T) {
	config.LoadConfig("/home/deploy/bcbtest/conf/config.yaml")
	Init()
	names := GetThirdSuiteNames("交易tx解析")
	for _, name := range names {
		fmt.Println(name)
	}
}
