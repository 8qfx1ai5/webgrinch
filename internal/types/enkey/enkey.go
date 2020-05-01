package enkey

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

// Key holds all information to make an encryption
type Key struct {
	from    string
	to      string
	subsets [][]rune
}

// UseFrom generates new key based on from
func (k *Key) UseFrom(in string) (err error) {
	k.from = in
	k.subsets = make([][]rune, 1)
	k.subsets[0] = []rune(in)
	k.shuffle()
	return nil
}

// UseRegex generates new key based on a regex
func (k *Key) UseRegex(regex string) (err error) {
	subsets, err := createSubsetsFromRegex(regex)
	k.UseSubsets(subsets)
	return nil
}

// UseSubsets generates new key based on a slice of slices
func (k *Key) UseSubsets(subsets [][]rune) (err error) {
	k.subsets = subsets
	k.from = ""
	for _, subset := range k.subsets {
		k.from += string(subset)
	}
	k.shuffle()
	return nil
}

// UseExistingKey takes the values and does not shuffle
func (k *Key) UseExistingKey(from string, to string) (err error) {
	k.from = from
	k.to = to
	k.subsets = make([][]rune, 1)
	k.subsets[0] = []rune(from)
	return nil
}

// shuffle the subsets and generate 'to'
func (k *Key) shuffle() (err error) {
	k.to = ""
	for _, subset := range k.subsets {
		newSubset := make([]rune, len(subset))
		copy(newSubset, subset)
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(newSubset), func(i, j int) { newSubset[i], newSubset[j] = newSubset[j], newSubset[i] })
		k.to += string(newSubset)
	}
	return nil
}

// GetTo returns the current 'to' value
func (k *Key) GetTo() (to string) {
	return k.to
}

// GetFrom returns the current 'from' value
func (k *Key) GetFrom() (to string) {
	return k.from
}

// analyse a regex and try to figure out the contained  characters
// function is unit tested
func createSubsetsFromRegex(regex string) (out [][]rune, err error) {
	r := regexp.MustCompile(`\[([^\[\]]*)\]`)
	submatchall := r.FindAllString(regex, -1)
	if submatchall == nil {
		out = make([][]rune, 1)
		out[0], err = processSingleRegexSubset(regex)
		return
	}
	out = make([][]rune, 0)
	for _, submatch := range submatchall {
		sub, err := processSingleRegexSubset(submatch[1 : len(submatch)-1])
		if err != nil {
			return out, err
		}
		out = append(out, sub)
	}
	return out, nil
}

type regexCase struct {
	regex string
	all   string
}

var ranges = []regexCase{
	regexCase{`[A-Z]-[A-Z]`, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"},
	regexCase{`\d-\d`, "0123456789"},
	regexCase{`[a-z]-[a-z]`, "abcdefghijklmnopqrstuvwxyz"},
}

// process single subset
func processSingleRegexSubset(regex string) (out []rune, err error) {

	for _, matchCandidate := range ranges {
		r := regexp.MustCompile(matchCandidate.regex)
		submatch := r.FindString(regex)
		if submatch == "" {
			continue
		}
		r2 := regexp.MustCompile("[" + submatch + "]*")
		thisMatch := r2.FindString(matchCandidate.all)
		if thisMatch == "" {
			return out, fmt.Errorf("regex '" + submatch + "' not implemented")
		}
		out = []rune(thisMatch)
		r3 := regexp.MustCompile(submatch)
		otherMatches := r3.Split(regex, 2)
		for _, otherMatch := range otherMatches {
			if otherMatch == "" {
				continue
			}
			tmp, err := processSingleRegexSubset(otherMatch)
			if err != nil {
				return out, err
			}
			out = append(out, tmp...)
		}
		return out, err
	}

	// no regex found, return original
	return []rune(regex), nil
}
