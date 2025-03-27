package resp

import (
	"encoding/json"
	"fmt"
)

func PrintResp(model any) {
	queryResp, _ := json.MarshalIndent(model, "", "  ")
	fmt.Println(string(queryResp))
}
