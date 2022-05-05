package pescriptlist

import (
	"privesc-toolkit-manager/types"
)

var PEScriptList = []types.PEScript{
	types.PEScript{Name: "Linpeas", Url: "https://www.linpeas.com/", System: types.Linux},
	types.PEScript{Name: "Winpeas", Url: "https://www.linpeas.com/", System: types.Windows},
	types.PEScript{Name: "Macpeas", Url: "https://www.linpeas.com/", System: types.Mac},
}
