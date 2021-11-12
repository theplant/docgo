package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

//go:embed template
var bareBox embed.FS

func main() {

	f, err := os.Open("go.mod")
	if err != nil {
		fmt.Println("run inside go package where go.mod exists")
		os.Exit(1)
	}

	mod, _ := ioutil.ReadAll(f)
	lines := strings.Split(string(mod), "\n")
	var pkg string
	for _, line := range lines {
		if strings.HasPrefix(line, "module") {
			segs := strings.Split(line, " ")
			if len(segs) > 1 {
				pkg = segs[1]
			}
		}
	}
	if len(pkg) == 0 {
		fmt.Println("no package provide in go.mod")
		os.Exit(1)
	}

	validateFileExists := func(input string) error {
		dir := filepath.Base(input)
		_, err := os.Stat(dir)
		if err == nil {
			return fmt.Errorf("%s already exists, remove it first to generate.\n", err)
		}
		return nil
	}

	dir := "docsrc"
	if err = validateFileExists(dir); err != nil {
		sel := promptui.Select{
			Label: "Directory 'docsrc' exists, Overwrite?",
			Items: []string{
				"Yes",
				"No",
			},
		}
		result, _, _ := sel.Run()
		if result != 0 {
			return
		}
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}

	copyAndReplaceFiles(bareBox, ".", "template", pkg)
	fmt.Printf("\ncd %s && ./dev.sh to start your doc\n", dir)

}

func copyAndReplaceFiles(box embed.FS, dir string, template string, pkg string) {
	fs.WalkDir(box, template, func(path string, d fs.DirEntry, err1 error) error {
		if d != nil && d.IsDir() {
			return nil
		}
		newPath := strings.ReplaceAll(path, template+"/", "")
		fp := filepath.Join(dir, newPath)
		err := os.MkdirAll(filepath.Dir(fp), 0755)
		if err != nil {
			panic(err)
		}
		var f *os.File
		f, err = os.Create(fp)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		content, err := box.ReadFile(path)
		if err != nil {
			panic(err)
		}

		var fmod fs.FileMode = 0644
		fmt.Println(fmod)
		err = ioutil.WriteFile(fp, content, fmod)
		if err != nil {
			panic(err)
		}
		fmt.Println(fp, "generated")
		return err
	})

	runCmd(dir, "chmod", "+x", "./docsrc/dev.sh")

	fmt.Println("Done")

	replaceInFiles(dir, "github.com/theplant/docgo/x/docgo/template", pkg)
	replaceInFiles(dir, "docgoPackageName", filepath.Base(pkg))

	return
}

func runCmd(dir string, name string, args ...string) {
	cmdGet := exec.Command(name, args...)
	cmdGet.Dir = dir
	cmdGet.Stdout = os.Stdout
	cmdGet.Stderr = os.Stderr

	err := cmdGet.Run()
	if err != nil {
		panic(err)
	}
}

func replaceInFiles(baseDir string, from, to string) {
	var fileList []string
	err := filepath.Walk(baseDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range fileList {
		replaceInFile(file, from, to)
	}
}

func replaceInFile(filepath, from, to string) {
	read, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	newContents := strings.Replace(string(read), from, to, -1)

	// fmt.Println(newContents)

	err = ioutil.WriteFile(filepath, []byte(newContents), 0)
	if err != nil {
		panic(err)
	}
}
