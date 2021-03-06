// mkpass - A simple password generator.
package main

import (
	"bufio"
	cRand "crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	mRand "math/rand"
	"os"
	"regexp"
	"strings"
)

const Version = "0.0.2"

/* Reads a file of words, filters the words and returns an array of words.
 *
 * path: Path to the file to read (each line is a different word)
 * minLen: The minumum length of the words. Any words shorter than the minimum
 *         length are filtered out.
 */
func readDict(path string, minLen int) ([]string, error) {
	// Build the regexp to use as the filter
	filterRegexp := fmt.Sprintf("^[[:word:]]{%d,}$", minLen)
	validWord := regexp.MustCompile(filterRegexp)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var line = scanner.Text()
		// Filter out invalid lines
		if validWord.MatchString(line) {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

/* Returns an array of the passed length containing randoms words from the
 * passed array.
 *
 * dict: The array of words.
 * numberOfWords: THe number of strings to return.
 */
func genPassword(dict []string, numberOfWords int) ([]string, error) {
	var words []string

	if numberOfWords < 1 {
		return []string{}, errors.New("Need a positive number of words to generate")
	}

	max := len(dict)
	if max < 1 {
		return []string{}, errors.New("No words in the dict")
	}

	for i := 0; i < numberOfWords; i++ {
		rand_index := mRand.Intn(max)

		word := dict[rand_index]

		words = append(words, word)
	}
	return words, nil
}

func init() {
	// Seed math/rand using crypto/rand
	max := *big.NewInt(9223372036854775807) // largest int64

	seed, err := cRand.Int(cRand.Reader, &max)
	if err != nil {
		fmt.Println("Unable to seed the random number generator")
		fmt.Println(err)
		os.Exit(1)
	}

	mRand.Seed(seed.Int64())
}

func main() {
	// Parse Arguments
	var numberOfWords, minLenOfWords int
	var wordFile string
	var showVersion bool

	flag.IntVar(&numberOfWords, "n", 4, "The number of words to generate")
	flag.IntVar(&minLenOfWords, "l", 6, "The minimum lenght of each word")
	flag.StringVar(&wordFile, "f", "/usr/share/dict/words", "A file that contains possible words (one per line)")
	flag.BoolVar(&showVersion, "v", false, "Display the version")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s version %s\n", os.Args[0], Version)
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nOptions\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if showVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	// Go to Work
	dict, err := readDict(wordFile, minLenOfWords)
	if err != nil {
		fmt.Printf("Unable to read words from '%s'\n", wordFile)
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Choosing from %d words\n", len(dict))

	words, err := genPassword(dict, numberOfWords)
	if err != nil {
		fmt.Println("Unable to generate a password")
		fmt.Println(err)
		os.Exit(1)
	}

	// Return the result
	fmt.Println("Your words are:")
	fmt.Println(strings.Join(words, " "))
}
