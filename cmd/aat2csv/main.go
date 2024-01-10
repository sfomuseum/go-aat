package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/sfomuseum/go-aat"
	"github.com/sfomuseum/go-csvdict"
)

func main() {

	var terms string

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Parse Getty AAT XML data and emit as CSV data to STDOUT with the following columns: id, term, preferred, languages.\n")
		fmt.Fprintf(os.Stderr, "Usage:\n\t %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	
	flag.StringVar(&terms, "terms", "", "The path to a local file on disk containing Getty AAT XML vocabulary data. If empty that data will be retrieved from the Getty servers.")
	flag.Parse()

	r, err := aat.TermsReader(terms)

	if err != nil {
		log.Fatalf("Failed to open %s for reading, %v", terms, err)
	}

	defer r.Close()

	var vocab *aat.Vocabulary

	dec := xml.NewDecoder(r)
	err = dec.Decode(&vocab)

	if err != nil {
		log.Fatalf("Failed to decode %s, %v", terms, err)
	}

	var csv_wr *csvdict.Writer

	mu := new(sync.RWMutex)

	write_rows := func(rows []map[string]string) error {

		mu.Lock()
		defer mu.Unlock()

		for _, row := range rows {

			if csv_wr == nil {

				fieldnames := make([]string, 0)

				for k, _ := range row {
					fieldnames = append(fieldnames, k)
				}

				wr, err := csvdict.NewWriter(os.Stdout, fieldnames)

				if err != nil {
					return fmt.Errorf("Failed to create CSV writer, %v", err)
				}

				wr.WriteHeader()
				csv_wr = wr
			}

			err = csv_wr.WriteRow(row)

			if err != nil {
				return fmt.Errorf("Failed to write row, %v", err)
			}
		}

		return nil
	}

	for _, s := range vocab.Subject {

		out := make([]map[string]string, 0)

		pref := s.Terms.PreferredTerm

		t_langs := pref.TermLanguages.TermLanguage
		langs := make([]string, len(t_langs))

		for idx, l := range t_langs {
			langs[idx] = l.Language
		}

		row := map[string]string{
			"id":        strconv.Itoa(s.SubjectID),
			"term":      pref.TermText,
			"preferred": "1",
			"languages": strings.Join(langs, ","),
		}

		out = append(out, row)

		for _, t := range s.Terms.NonPreferredTerm {

			row := map[string]string{
				"id":        strconv.Itoa(s.SubjectID),
				"term":      t.TermText,
				"preferred": "0",
				"languages": "",
			}

			if t.TermLanguages != nil {

				t_langs := t.TermLanguages.TermLanguage
				langs := make([]string, len(t_langs))

				for idx, l := range t_langs {
					langs[idx] = l.Language
				}

				row["languages"] = strings.Join(langs, ",")
			}

			out = append(out, row)
		}

		err = write_rows(out)

		if err != nil {
			log.Fatalf("Failed to write rows, %v", err)
		}

	}

	csv_wr.Flush()
}
