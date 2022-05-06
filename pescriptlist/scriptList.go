package pescriptlist

import (
	"privesc-toolkit-manager/types"
)

var PEScriptList = []types.PEScript{
	types.PEScript{Name: "Linpeas", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/linpeas_linux_amd64", System: types.Linux, Arch: types.Amd64},
	types.PEScript{Name: "Linpeas", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/linpeas_linux_arm64", System: types.Linux, Arch: types.Arm64},
	types.PEScript{Name: "Winpeas", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/winPEASx64.exe", System: types.Windows, Arch: types.Amd64},
	types.PEScript{Name: "Winpeas", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/winPEASx86.exe", System: types.Windows, Arch: types.I386},
	types.PEScript{Name: "Macpeas", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/linpeas_darwin_amd64", System: types.Mac, Arch: types.Amd64},
	types.PEScript{Name: "Macpeas", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/linpeas_darwin_arm64", System: types.Mac, Arch: types.Arm64},
	types.PEScript{Name: "Macpeas sh", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/linpeas.sh", System: types.Mac, Arch: types.NoArch},
	types.PEScript{Name: "Linpeas sh", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/linpeas.sh", System: types.Mac, Arch: types.NoArch},
	types.PEScript{Name: "Winpeas bat", Url: "https://github.com/carlospolop/PEASS-ng/releases/download/20220504/winPEAS.bat", System: types.Windows, Arch: types.NoArch},
	types.PEScript{Name: "Linux Smart Enumeration", Url: "https://github.com/diego-treitos/linux-smart-enumeration/releases/download/4.5nw/lse.sh", System: types.Linux, Arch: types.NoArch},
	types.PEScript{Name: "Linux Exploit Suggester", Url: "https://raw.githubusercontent.com/mzet-/linux-exploit-suggester/master/linux-exploit-suggester.sh", System: types.Linux, Arch: types.NoArch},
	types.PEScript{Name: "LinEnum", Url: "https://raw.githubusercontent.com/rebootuser/LinEnum/master/LinEnum.sh", System: types.Linux, Arch: types.NoArch},
	types.PEScript{Name: "Linux PrivChecker", Url: "https://raw.githubusercontent.com/linted/linuxprivchecker/master/linuxprivchecker.sh", System: types.Linux, Arch: types.NoArch},
}

func GetToolId(name string) int {
	for i, v := range PEScriptList {
		if v.Name == name {
			return i
		}
	}
	panic("Unknown Tool")
}
