package utils

import (
	"fmt"
	"strings"
)
//Like sql查询字符串前后加%
func Like(q string)  string{
	q = strings.TrimSpace(q)
	if q == ""{
		return ""
	}
	q = strings.ReplaceAll(q,"/","//")
	q = strings.ReplaceAll(q,"%","/%")
	q = strings.Replace(q,"_","/_",-1)
	return fmt.Sprintf("%%%s%%",q)
}
