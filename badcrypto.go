package vuln

import "crypto/md5"


func main(){

	x := 0
	y := 2
    if x > y {
        h := md5.New()
    }

}