package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Options is struct of options.
type Options struct {
	config string
}

func main() {

	var opts Options

	flag.StringVar(&opts.config, "c", "config.yaml", "Path of config file")
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
	paths := buildSiteMap(config.ProxyRoot, config.Path)
	log.Printf("Path is found in %v", paths)

	// Create source file.
	log.Println("Creating source file...")
	codeTmpl, err := loadTemplateFromBinary("templates/code_template.go.tmpl")
	if err != nil {
		log.Fatal("Failed to load code template")
	}
	baseURL := strings.Replace(config.ProxyRoot, config.WebRoot, "", 1)
	server := fmt.Sprintf("%s:%d", config.Server, config.Port)
	tmplModel := NewTemplateModel(baseURL, server, config.LogFile)
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
	if err = exec.Command("go", "build", "-o", "index.cgi").Run(); err != nil {
		log.Fatalf("Failed to build source file %s", err)
	}
	defer os.Remove("index.cgi")

	// Open executable
	bin, err := os.Open("index.cgi")
	if err != nil {
		panic(err)
	}
	defer bin.Close()

	// Copy executable for each directory.
	log.Println("Installing executable")
	for _, path := range paths {
		target := path + "/index.cgi"
		fmt.Println("Installing to", target)
		if err = os.MkdirAll(path, 0755); err != nil {
			log.Fatalf("Failed to make directory: %s", path)
		}
		dst, err := os.OpenFile(target, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			panic(err)
		}
		defer dst.Close()
		if _, err := io.Copy(dst, bin); err != nil {
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
