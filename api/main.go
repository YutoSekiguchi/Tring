package main

import (
	"github.com/YutoSekiguchi/tring/router"
	"github.com/YutoSekiguchi/tring/util"
)

func main() {
	db := util.InitDb()
	router.InitRouter(db)
}