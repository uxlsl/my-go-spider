package pipelines

import "fmt"

type HttpBinPipeline struct{
	In chan string
}

func (h *HttpBinPipeline) Process(){
	for item:= range h.In{
		fmt.Println("Process item")
		fmt.Println(item)
	}
}
