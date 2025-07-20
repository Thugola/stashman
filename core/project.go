package core

import (
	"fmt"
	"encoding/json"
	"os"
	"strings"
	"path/filepath"
)

const stashFileSuffix = ".stash.json"
var stashFiles []string
var StashFilePath string

func IsStashFile(filename string) bool {
		return len(filename) >= 11 &&
			filename[len(filename)-11:] == stashFileSuffix
}

func CreateStashFile() {
	var title string
	var path string
	var snippets []Snippet

	for {
		fmt.Printf("[>] Create a title for the project: ")
		n, err := fmt.Scanln(&title)
		if n == 0 || err != nil {
			fmt.Println("[!] Title cannot be empty.")
			title = ""
		} else {
			break
		}
	}

	for {
		fmt.Printf("[>] Enter a root path for the project: ")
		n, err := fmt.Scanln(&path)
		if n == 0 || err != nil {
			fmt.Println("[!] Path cannot be empty.")
			path = ""
		} else {
			break
		}
	}

	path = strings.TrimSpace(path)
	path = filepath.Clean(path)
	if !strings.HasSuffix(path, string(os.PathSeparator)) {
		path += string(os.PathSeparator)
	}
	
	StashFilePath = path + title + stashFileSuffix
	
	data := map[string]interface{}{
		"Title":    title,
		"Snippets": snippets,
	}

	file, err := os.Create(StashFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	json.NewEncoder(file).Encode(data)
	os.Exit(0)
}

func ValidateStashFileContent() {
	data, err := os.ReadFile(StashFilePath)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}

	var obj map[string]interface{}
	if err := json.Unmarshal(data, &obj); err != nil {
		fmt.Println("Error: " + err.Error())
		os.Exit(1)
	}
	
	if len(obj) != 2 {
		fmt.Println("A stash file was found, but it has more entries than it should: " + StashFilePath)
		os.Exit(1)
	}

	_, hasTitle := obj["Title"]
	_, hasSnippets := obj["Snippets"]

	if hasTitle && hasSnippets {
		return
	} else {
		fmt.Println("A stash file was found, but its entries are invalid: " + StashFilePath)
		os.Exit(1)
	}
}

func CheckStashFileCount() {
	if len(stashFiles) > 1 {
		fmt.Println("Multiple stash files encountered:")
		for i, stashFile := range stashFiles {
			fmt.Println(fmt.Sprintf("%d: %s", i+1, stashFile))
		}
		fmt.Println("Stashman cannot proceed.")
		os.Exit(1)
	}

	if len(stashFiles) == 0 {
		var confirmStashCreation string
		confirmationKey := "y"
		denialKey := "n"

		fmt.Println("No stash file encountered for the project.")
		for {
			fmt.Printf("[>] Do you want to create one? (Y/n): ")
			_, err := fmt.Scanln(&confirmStashCreation)
			if strings.ToLower(confirmStashCreation) == confirmationKey || err != nil {
				CreateStashFile()
			} else if strings.ToLower(confirmStashCreation) == denialKey {
				os.Exit(0)
			} else {
				continue
			}
		}
	}
}

func LoadOrInitProject() {
	dir, _ := os.Getwd()
	fsRoot := "/"

	for {
		files, err := os.ReadDir(dir)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			os.Exit(1)
		}
	
		for _, file := range files {
			filename := file.Name()
			if IsStashFile(filename) {
				StashFilePath = filepath.Join(dir, filename)
				stashFiles = append(stashFiles, StashFilePath)
			}
		}

		if len(stashFiles) >= 1 {
			break
		}

		if dir != fsRoot {
			parent := filepath.Dir(dir)
			dir = parent
		} else {
			break
		}
	}

	CheckStashFileCount()
	ValidateStashFileContent()
}
