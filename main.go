package main

import (
	"errors"
	"github.com/axiom-data-science/usher"
	"log"
	"strings"
	"time"
)

type IfcbFileMapper struct{}

func (fm *IfcbFileMapper) GetFileDestPath(relSrcFile string, absSrcFile string,
	baseSrcFile string, mappedRootSrcPath string, mappedRootDestPath string) (string, error) {

	//all data other than beads should be put in a separate path
	//so a single directory contains the real dataset
	const dataDir = "data"
	const beadsDir = "beads"

	//allow root level metadata.csv files
	if relSrcFile == "metadata.csv" {
		return dataDir + "/" + relSrcFile, nil
	}

	//example filename: D20230525T192231_IFCB162.adc
	if baseSrcFile[0] != 'D' {
		return "", errors.New("file " + relSrcFile + " does not start with D prefix, ignoring")
	} else if len(baseSrcFile) < 16 {
		return "", errors.New("file " + relSrcFile + " has a base filename less than 16 characters, ignoring")
	}

	fileTime, err := time.Parse("20060102T150405", baseSrcFile[1:16])
	if err != nil {
		log.Println(err)
		return "", errors.New("couldn't parse date from file " + relSrcFile + ", ignoring")
	}
	destPath := fileTime.Format("2006/D20060102/") + baseSrcFile

	//put beads in beads directory if any directory component is "beads"
	pathComponents := strings.Split(relSrcFile, "/")
	pathComponents = pathComponents[:len(pathComponents)-1]
	var isBeads bool = false
	for _, pathComponent := range pathComponents {
		if pathComponent == "beads" {
			isBeads = true
			break
		}
	}
	if isBeads {
		destPath = beadsDir + "/" + destPath
	} else {
		destPath = dataDir + "/" + destPath
	}

	return destPath, nil
}

func main() {
  //run usher with a single ifcb mapper
	usher.Run(map[string]usher.FileMapper{
		"ifcb": &IfcbFileMapper{},
	})
}
