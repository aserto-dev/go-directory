module github.com/aserto-dev/go-directory/magefiles

go 1.17

// !!! required to workaround mage induces a clui dependency update which is 1.18 dependent and breaks the build !!!
replace github.com/rivo/uniseg => github.com/rivo/uniseg v0.2.0

replace github.com/zricethezav/gitleaks/v8 => github.com/zricethezav/gitleaks/v8 v8.3.0

replace github.com/gitleaks/go-gitdiff => github.com/gitleaks/go-gitdiff v0.7.4

require (
	github.com/aserto-dev/clui v0.8.1
	github.com/aserto-dev/mage-loot v0.8.10
	github.com/magefile/mage v1.14.0
)

require (
	github.com/allegro/bigcache/v3 v3.0.2 // indirect
	github.com/fatih/color v1.13.0 // indirect
	github.com/go-test/deep v1.0.8 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kyokomi/emoji v2.2.4+incompatible // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.2 // indirect
	github.com/tidwall/gjson v1.14.3 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/ulikunitz/xz v0.5.10 // indirect
	golang.org/x/sys v0.0.0-20220915200043-7b5979e65e41 // indirect
	golang.org/x/term v0.0.0-20220722155259-a9ba230a4035 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
