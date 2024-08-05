package main

import (
	"fmt"
	"log"
	"net/http"
)

const banner = `
             .__   ________                              
  ___________|  |  \_____  \ ___  ______________  _____  
 /  ___/ ____/  |   /  ____/ \  \/  /  _ \_  __ \/     \ 
 \___ < <_|  |  |__/       \  >    <  <_> )  | \/  Y Y  \
/____  >__   |____/\_______ \/__/\_ \____/|__|  |__|_|  /
     \/   |__|             \/      \/                 \/ 
`

func main() {
	http.HandleFunc("/gen", generateHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println(banner)
	fmt.Println("listening on :7892")
	log.Fatal(http.ListenAndServe(":7892", nil))
}
