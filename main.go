package main

import (
	"fmt"

	"privesc-toolkit-manager/downloader"
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
	currentArch := syscheck.GetArch()
	fmt.Println("Your system is: ", types.GetSystemName(currentSystem))
	ToolsSelect := promptui.Select{
		Label: "Select Scripts to Download",
		Items: types.GetNames(types.FilterBySystem(pescriptlist.PEScriptList, currentSystem, currentArch)),
	}
	_, Tool, err := ToolsSelect.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("You selected: ", Tool)
	fmt.Println("Downloading...")
	err = downloader.Download(pescriptlist.PEScriptList[pescriptlist.GetToolId(Tool)].Url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Download complete!")
}

func banner() {
	fmt.Println(`    /$$$$$$$           /$$                                                  
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
