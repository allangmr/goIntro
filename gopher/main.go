package main

import (
	"fmt"
	"time"
)

func main(){
  now := time.Unix(123456789, 0).UTC() 
  fmt.Println("                    ")
  fmt.Println("    `.-::::::-.`    ")
  fmt.Println(".:-::::::::::::::-:.")
  fmt.Println(" `_:::    ::    :::_`")
  fmt.Println(" .:( ^   :: ^   ):. ")
  fmt.Println(" `:::   (..)   :::. ")
  fmt.Println(" `:::::::UU:::::::` ")
  fmt.Println(" .::::::::::::::::. ")
  fmt.Println(" O::::::::::::::::O ")
  fmt.Println(" -::::::::::::::::- ")
  fmt.Println(" `::::::::::::::::` ")
  fmt.Println("  .::::::::::::::.  ")
  fmt.Println("    oO:::::::Oo     ")
  fmt.Printf("%v %v\n", "Actual time:", now)
}