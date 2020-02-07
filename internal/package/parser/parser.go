package parser

import (
	"strings"

	"awesome/internal/package/node"
)

func ParseIndex(readme string) node.Node {
	var parseStatus = false
	var cname string  = ""
	var cobj  node.Node
	// May be used for logging
	var categoryCount int = 0
	var contentCount int  = 0

	var index node.Node

    for _, line := range strings.Split(strings.TrimSuffix(readme, "\n"), "\n") {
	    //fmt.Println(line)

	    if IsCategory(line) {
	    	if !IsCategoryIgnored(line) {
		    	parseStatus = true
		    	
		    	if (cname != "") {
		    		index.AddChild(cobj)
		    	}

		    	cname = LineToTitle(line)
		    	cobj  = node.New(cname,"", "")
		    	categoryCount++
		    	continue
		    } else {
		    	parseStatus = false
		    }
	    }

	    if (parseStatus == true) {
	    	if IsContent(line) {
	    		name, url, desc, _ := ParseContentFromLine(line)

	    		cobj.AddChild(node.New(name, url, desc))
	    		contentCount++
		    	continue
		    }
	    }
	}

	index.AddChild(cobj)

	return index
}

func IsCategory(line string) bool {
	return strings.HasPrefix(line, "## ") && !strings.HasPrefix(line, "### ")
}


func IsContent(line string) bool {
	return strings.HasPrefix(line, "- [") || strings.HasPrefix(line, "* [")
}

func IsCategoryIgnored(line string) bool {
	ignoreList := []string{"Table of Contents", "Contents", "Contributing", "TODO", "Introduction"}
	str        := LineToTitle(line)
	for _, s := range ignoreList {
        if (s == str) {
        	return true
        }
    }

    return false
}

func LineToTitle(line string) string {
	line = strings.Replace(line, "#", "", -1)
	line = strings.Trim(line, " ")

	return line
}

func ParseContentFromLine(line string) (string, string, string, error) {
	name := Split(line, "[", "]")
	url  := Split(line, "(", ")")
	line  = strings.Trim(line, "- ")
	desc :=  ""

	if (strings.HasSuffix(line, ".")) {
		desc = Split(line, " - ", "\n")	
	}

	return name, url, desc, nil
}

func Split(str, before, after string) string {
    a := strings.SplitAfterN(str, before, 2)
    b := strings.SplitAfterN(a[len(a)-1], after, 2)

    if 1 == len(b) {
        return b[0]
    }
    
    return b[0][0:len(b[0])-len(after)]
}
