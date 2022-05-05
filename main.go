package main

import (
	"fmt"

	"privesc-toolkit-manager/pescriptlist"
	"privesc-toolkit-manager/syscheck"
	"privesc-toolkit-manager/types"

	"github.com/manifoldco/promptui"
)

func main() {
	fmt.Println("=================================")
	banner()
	fmt.Println("=================================")
	fmt.Println("Welcome to Privesc Toolkit Manager")
	fmt.Println("This tool is designed to help you find and download Privesc scripts")
	fmt.Println("_________________________________")
	currentSystem := syscheck.GetSystem()
	fmt.Println("Your system is: ", types.GetSystemName(currentSystem))
	ToolsSelect := promptui.Select{
		Label: "Select Scripts to Download",
		Items: types.GetNames(types.FilterBySystem(pescriptlist.PEScriptList, currentSystem)),
	}
	_, Tools, err := ToolsSelect.Run()
	if err == nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Tools: ", Tools)
}

func banner() {
	fmt.Println(`	/$$$$$$$           /$$                                                  
    | $$__  $$         |__/                                                  
    | $$  \ $$ /$$$$$$  /$$ /$$    /$$ /$$$$$$   /$$$$$$$  /$$$$$$$          
    | $$$$$$$//$$__  $$| $$|  $$  /$$//$$__  $$ /$$_____/ /$$_____/          
    | $$____/| $$  \__/| $$ \  $$/$$/| $$$$$$$$|  $$$$$$ | $$                
    | $$     | $$      | $$  \  $$$/ | $$_____/ \____  $$| $$                
    | $$     | $$      | $$   \  $/  |  $$$$$$$ /$$$$$$$/|  $$$$$$$          
    |__/     |__/      |__/    \_/    \_______/|_______/  \_______/          
                                                                             
                                                                             
                                                                             
     /$$$$$$$$                  /$$ /$$       /$$   /$$                      
    |__  $$__/                 | $$| $$      |__/  | $$                      
       | $$  /$$$$$$   /$$$$$$ | $$| $$   /$$ /$$ /$$$$$$                    
       | $$ /$$__  $$ /$$__  $$| $$| $$  /$$/| $$|_  $$_/                    
       | $$| $$  \ $$| $$  \ $$| $$| $$$$$$/ | $$  | $$                      
       | $$| $$  | $$| $$  | $$| $$| $$_  $$ | $$  | $$ /$$                  
       | $$|  $$$$$$/|  $$$$$$/| $$| $$ \  $$| $$  |  $$$$/                  
       |__/ \______/  \______/ |__/|__/  \__/|__/   \___/                    
                                                                             
                                                                             
                                                                             
     /$$      /$$                                                            
    | $$$    /$$$                                                            
    | $$$$  /$$$$  /$$$$$$  /$$$$$$$   /$$$$$$   /$$$$$$   /$$$$$$   /$$$$$$ 
    | $$ $$/$$ $$ |____  $$| $$__  $$ |____  $$ /$$__  $$ /$$__  $$ /$$__  $$
    | $$  $$$| $$  /$$$$$$$| $$  \ $$  /$$$$$$$| $$  \ $$| $$$$$$$$| $$  \__/
    | $$\  $ | $$ /$$__  $$| $$  | $$ /$$__  $$| $$  | $$| $$_____/| $$      
    | $$ \/  | $$|  $$$$$$$| $$  | $$|  $$$$$$$|  $$$$$$$|  $$$$$$$| $$      
    |__/     |__/ \_______/|__/  |__/ \_______/ \____  $$ \_______/|__/      
                                                /$$  \ $$                    
                                               |  $$$$$$/                    
                                                \______/                     `)
}
