// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// +build ignore

package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

var (
	baseFlag = flag.String("base", ".", "base directory of the repo, relative to the working directory")
	outFlag  = flag.String("o", "Dockerfile-testing", "output file, relative to this directory")
	diffFlag = flag.Bool("d", false, "diff file against output file instead of writing")
)

func relPath(p string) string {
	if *baseFlag == "." {
		return "./" + p
	}
	rel, err := filepath.Rel(*baseFlag, p)
	if err != nil {
		panic(err)
	}
	return rel
}

var dockerfileTemplateFuncs = template.FuncMap{
	"join": path.Join,
}

// This generates a Dockerfile that copies all of the go.mod and go.sum
// files found and then runs "go mod download" for each module, before
// copying across the rest of the source code.
var dockerfileTemplate = template.Must(template.New("Dockerfile").Funcs(dockerfileTemplateFuncs).Parse(`
# Code generated by gendockerfile. DO NOT EDIT.
FROM golang:latest
ENV GO111MODULE=on
{{range .Dirs}}
COPY {{join . "go.mod"}} {{join . "go.sum"}} {{join "/go/src/go.elastic.co/apm" .}}/{{end}}

{{range .Dirs}}RUN cd {{join "/go/src/go.elastic.co/apm" .}} && go mod download
{{end}}
WORKDIR /go/src/go.elastic.co/apm
ADD . /go/src/go.elastic.co/apm
`[1:]))

func main() {
	flag.Parse()

	// Locate all go.mod files.
	var moduleDirs []string
	if err := filepath.Walk(*baseFlag, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if info.Name() == "go.mod" {
				moduleDirs = append(moduleDirs, relPath(filepath.Dir(path)))
			}
			return nil
		}
		name := info.Name()
		if name != *baseFlag && (name == "vendor" || strings.HasPrefix(name, ".")) {
			return filepath.SkipDir
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	sort.Strings(moduleDirs)

	var buf bytes.Buffer
	var out io.Writer = &buf
	outFile := filepath.Join(*baseFlag, "scripts", *outFlag)
	if !*diffFlag {
		f, err := os.Create(outFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		out = f
	}

	var data struct {
		Dirs []string
	}
	data.Dirs = moduleDirs

	if err := dockerfileTemplate.Execute(out, &data); err != nil {
		log.Fatal(err)
	}
	if *diffFlag {
		cmd := exec.Command("diff", "-c", outFile, "-")
		cmd.Stdin = &buf
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}