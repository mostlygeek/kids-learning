package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/mostlygeek/kids-learning/lib"
)

var wordlist = []string{
	"bad", "bag", "bam", "ban", "bar", "bat", "bed", "beg",
	"ben", "bet", "bib", "bid", "big", "bin", "bit", "bot",
	"bow", "bug", "bum", "bun", "bus", "but", "cab", "cam",
	"can", "cap", "car", "cat", "caw", "cod", "cog", "con",
	"cop", "cot", "cow", "cub", "cup", "cut", "dad", "dag",
	"dan", "den", "dew", "did", "dig", "dim", "dip", "dog",
	"dot", "dub", "dug", "fad", "fan", "far", "fat", "fed",
	"fib", "fig", "fit", "fix", "fog", "for", "fox", "fun",
	"fur", "gag", "gas", "gem", "get", "gig", "gim", "gin",
	"god", "got", "gum", "gun", "had", "ham", "hat", "hem",
	"hen", "her", "hew", "hex", "hic", "him", "hip", "his",
	"hit", "hog", "hop", "hot", "how", "hub", "hug", "hut",
	"jab", "jam", "jar", "jaw", "jet", "job", "jog", "keg",
	"ken", "kim", "kit", "lad", "lag", "lan", "lap", "lax",
	"leg", "let", "lid", "lip", "log", "lot", "lug", "mac",
	"mad", "man", "mat", "men", "met", "min", "mix", "mob",
	"mud", "mug", "mum", "nab", "nag", "net", "new", "nip",
	"nit", "not", "now", "nut", "pad", "pal", "pat", "paw",
	"pen", "pet", "pic", "pig", "pin", "pip", "pit", "pod",
	"pop", "pow", "pub", "pug", "pus", "put", "rad", "rag",
	"ram", "ran", "rat", "red", "rep", "ret", "rex", "rib",
	"rid", "rig", "rim", "rip", "rob", "rod", "rot", "row",
	"rub", "rug", "rum", "run", "rut", "sad", "sag", "sal",
	"sam", "sap", "sar", "sat", "saw", "sen", "set", "sew",
	"sin", "sip", "sit", "six", "sob", "sod", "sog", "son",
	"sot", "sov", "sow", "sub", "sum", "sun", "sur", "tab",
	"tad", "tag", "tap", "tav", "ted", "til", "tin", "tip",
	"tod", "tub", "tug", "van", "vet", "vex", "vim", "wad",
	"wag", "wah", "war", "was", "wax", "web", "wed", "wet",
	"wig", "win", "wit", "wiz", "won", "wow", "zag", "zak",
	"zig", "zip",
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// generate 5 pages of random cvc words
	pages := 5
	numwords := len(wordlist)
	listSize := pages * 48 /* 48 cells in the table */
	list := make([]string, listSize, listSize)
	for i := 0; i < listSize; i++ {
		list[i] = wordlist[rand.Intn(numwords)]
	}

	pdf, err := lib.TableLayout(pages, list)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	if err := pdf.OutputFileAndClose("cvc.pdf"); err != nil {
		fmt.Println(err.Error())
	}
}
