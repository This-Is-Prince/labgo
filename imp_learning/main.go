package main

import (
	"fmt"

	"github.com/imp_learning/bitmanipulation"
	"github.com/imp_learning/copierpackage"
	"github.com/imp_learning/encoderdecoder"
	"github.com/imp_learning/errorunderstanding"
	"github.com/imp_learning/overflow"
	"github.com/imp_learning/reflectpackage"
	"github.com/imp_learning/unsafepackage"
)

func main() {
	fmt.Println("Important learning while understanding open source Go repos and their structure.")

	bitmanipulation.Learn(false)
	reflectpackage.Learn(false)
	copierpackage.Learn(false)
	unsafepackage.Learn(false)
	overflow.Learn(false)
	errorunderstanding.Learn(false)
	encoderdecoder.Learn(true)
}
