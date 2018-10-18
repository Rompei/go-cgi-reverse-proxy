package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Options is struct of options.
type Options struct {
	config  string
	rootDir string
}

func main() {

	var opts Options

	flag.StringVar(&opts.config, "c", "config.yaml", "Path of config file")
	flag.StringVar(&opts.rootDir, "r", ".", "Path to proxy root directory")
	flag.Parse()

	path, err := exec.LookPath("go")
	if err != nil {
		log.Fatal("go command was not found.")
	}
	log.Printf("go command is located in %s", path)

	log.Printf("Loading config file in %s", opts.config)
	config, err := loadConfig(opts.config)
	if err != nil {
		log.Fatal("Failed to load config file")
	}

	// Load path.
	paths := buildSiteMap(opts.rootDir, config.Path)
	log.Printf("Path is found in %v", paths)

	// Create source file.
	log.Println("Creating source file...")
	codeTmpl, err := loadTemplateFromBinary("code_template.go.tmpl")
	if err != nil {
		log.Fatal("Failed to load code template")
	}
	server := fmt.Sprintf("%s:%d", config.Server, config.Port)
	tmplModel := NewTemplateModel(config.ProxyRoot, server)
	var w bytes.Buffer
	if err = codeTmpl.Execute(&w, tmplModel); err != nil {
		log.Fatal("Failed to execute template.")
	}

	log.Println("Source code was generated!")
	fmt.Println(w.String())

	if err = ioutil.WriteFile("cgi.go", w.Bytes(), 0644); err != nil {
		log.Fatal("Failed to create source file.")
	}
	defer os.Remove("cgi.go")

	// Build executable.
	log.Println("Building executable...")
	if err = exec.Command("go", "build", "-o", "tmp.cgi").Run(); err != nil {
		log.Fatalf("Failed to build source file %s", err)
	}
	defer os.Remove("tmp.cgi")

	// Open executable
	bin, err := ioutil.ReadFile("tmp.cgi")
	if err != nil {
		log.Fatal("Failed to open executable.")
	}

	// Copy executable for each directory.
	log.Println("Installing executable")
	for _, path := range paths {
		target := path + "/index.cgi"
		fmt.Println("Installing to", target)
		if err = os.MkdirAll(path, 0755); err != nil {
			log.Fatalf("Failed to make directory: %s", path)
		}
		if err := ioutil.WriteFile(target, bin, 0755); err != nil {
			log.Fatal("Failed to copy executable.")
		}
	}
	log.Println("Completed!")
}

func buildSiteMap(root string, paths []interface{}) (mp []string) {
	mp = append(mp, root)
	parsePaths(root, paths, &mp)
	return
}

func parsePaths(root string, paths []interface{}, mp *[]string) {
	current := ""
	for i := range paths {
		switch p := paths[i].(type) {
		case string:
			current = p
			pp := root + "/" + p
			*mp = append(*mp, pp)
		case []interface{}:
			parsePaths(root+"/"+current, p, mp)
		}
	}
}
