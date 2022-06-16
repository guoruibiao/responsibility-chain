package main

import (
	"context"
	"fmt"

	"github.com/guoruibiao/responsibility-chain/params"
	"github.com/guoruibiao/responsibility-chain/process"
)

func main() {
	// bootstrap
	process.Init()

	// mock request parameters
	reqParams := params.CommonParams{
		Name: "tiger",
		Age: 26,
		Address: "大连",
	}

	// register custom functions
	process.Register("upper_name", process.UpperName)
	process.Register("decrease_age", process.DecreaseAge)
	process.Register("modify_address", process.ModifyAddress)
	process.Register("decrease_age", process.DecreaseAge) // this will not be run because duplicated key name
	process.Register("decrease_age_2", process.DecreaseAge) // this will be run

	// run after all functions registered
	ctx := context.Background()
	if _, err := process.Run(&ctx, &reqParams); err != nil {
		fmt.Println(err)
		panic(err)
	}

	// print final result
	fmt.Println("-------------------")
	fmt.Printf("name=%s\n", ctx.Value("name"))
	fmt.Printf("age=%d\n", ctx.Value("age"))
	fmt.Printf("address=%s\n", ctx.Value("address"))
	fmt.Println("-------------------")
}
