package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/afero"
	flag "github.com/spf13/pflag"
)

type ElfInventory struct {
	Calories int64
	Items    []int64
}

var inputFilename string

func init() {
	flag.StringVar(&inputFilename, "input", "", "Input filename")
}

func main() {
	var (
		err              error
		inputFs          afero.Fs        = afero.NewReadOnlyFs(afero.NewOsFs())
		inventories      []*ElfInventory = make([]*ElfInventory, 0)
		rawData          []byte
		lines            []string
		currentInventory *ElfInventory = new(ElfInventory)
	)

	flag.Parse()
	log.Printf("flags parsed: %#v\n", flag.Parsed())
	log.Printf("input filename: %#v\n", inputFilename)

	rawData, err = afero.ReadFile(inputFs, inputFilename)
	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(rawData), "\n")
	for _, eachLine := range lines {
		eachLine = strings.TrimSpace(eachLine)
		if len(eachLine) < 1 {
			// blank line (or EOF) means end of elf inventory
			inventories = append(inventories, currentInventory)
			currentInventory = new(ElfInventory)
		} else {
			cal, err := strconv.ParseInt(eachLine, 10, 0)
			if err != nil {
				panic(err)
			}
			currentInventory.Items = append(currentInventory.Items, cal)
		}
	}

	// there may be an EOF w/o newline which means an elf's inventory is not saved
	if len(currentInventory.Items) > 0 {
		inventories = append(inventories, currentInventory)
	}
	currentInventory = nil

	for _, inv := range inventories {
		var totalCalories int64 = 0
		for _, item := range inv.Items {
			totalCalories += item
		}
		inv.Calories = totalCalories
	}

	sort.Slice(inventories, func(i, j int) bool {
		return inventories[i].Calories > inventories[j].Calories
	})

	fmt.Printf("Top 1 Calories: %d (%#v)\n", inventories[0].Calories, inventories[0].Items)

	var topThreeCaloriesSum int64 = 0
	for i := 0; i < 3; i++ {
		topThreeCaloriesSum += inventories[i].Calories
	}
	fmt.Printf("Top 3 Calories total: %d\n", topThreeCaloriesSum)
}
