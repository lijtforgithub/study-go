package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto breakHere
			}
		}
	}

	// 标签
breakHere:
	{
		fmt.Println("done")
	}

	/**
		err := firstCheckError()
	    if err != nil {
	        goto onExit
	    }
	    err = secondCheckError()
	    if err != nil {
	        goto onExit
	    }
	    fmt.Println("done")
	    return
	onExit:
	    fmt.Println(err)
	    exitProcess()
	*/
}
