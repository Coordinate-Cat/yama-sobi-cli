package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

type equipment struct {
	Type  string
	Name  string
	Price string
	Brand string
	Order string
	Url   string
}

func main() {
	equipments := []equipment{
		{Type: "Backpacks", Name: "THREE", Price: "¥34,100", Brand: "山と道", Order: "2021年7月30日", Url: "https://www.yamatomichi.com/products/three/"},
		{Type: "Pickel", Name: "Grivel Nepal SA Plus, 58cm", Price: "¥21,087", Brand: "Grivel"},
		{Type: "Shoes", Name: "コロンビア カラサワ ミスト オムニテック", Price: "¥18,150", Brand: "Columbia"},
		{Type: "Cap", Name: "New Era 9Twenty", Price: "¥4,730", Brand: "New Era"},
		{Type: "Sunglasses", Name: "Momentum", Price: "¥3,780", Brand: "Marsquest"},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   " {{ .Type | yellow }} ({{ .Price | red }})",
		Inactive: "  {{ .Type | blue }} ({{ .Price | red }})",
		Selected: "{{ .Type | red | blue }}",
		Details: `
── details ──────────────────────
{{ "Name :" | faint }}  {{ .Name | yellow }}
{{ "Type :" | faint }}  {{ .Type | blue }}
{{ "Brand:" | faint }}  {{ .Brand | blue }}
{{ "Order:" | faint }}  {{ .Order | blue }}
{{ "Price:" | faint }}  {{ .Price | yellow }}
─────────────────────────────────`,
	}

	searcher := func(input string, index int) bool {
		equipment := equipments[index]
		name := strings.Replace(strings.ToLower(equipment.Type), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Oka's Mt. climbing Equipments",
		Items:     equipments,
		Templates: templates,
		// 表示数
		Size:     8,
		Searcher: searcher,
	}

	i, _, err := prompt.Run()
	cmd, err := exec.Command("sh", "-c", "exit").CombinedOutput()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	// cmd := exec.Command("bash")
	fmt.Printf("You choose number %d: %s\n", i+1, equipments[i].Type)
	fmt.Printf(`
%s%v
	`, cmd, err)
	// hello ls:\n%s :Error:\n%v\n
}
