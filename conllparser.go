package conllparser

import (
	"strings"
	"strconv"
)

/*
CoNLL-U fields:
ID: Word index, integer starting at 1 for each new sentence; may be a range for tokens with multiple words.
FORM: Word form or punctuation symbol.
LEMMA: Lemma or stem of word form.
UPOSTAG: Universal part-of-speech tag drawn from our revised version of the Google universal POS tags.
XPOSTAG: Language-specific part-of-speech tag; underscore if not available.
FEATS: List of morphological features from the universal feature inventory or from a defined language-specific extension; underscore if not available.
HEAD: Head of the current token, which is either a value of ID or zero (0).
DEPREL: Universal Stanford dependency relation to the HEAD (root iff HEAD = 0) or a defined language-specific subtype of one.
DEPS: List of secondary dependencies (head-deprel pairs).
MISC: Any other annotation.
*/
type ConllWord struct {
	id      int
	form    string
	lemma   string
	upostag string
	xpostag string
	feats   string
	head    int
	deprel  string
	deps    string
	misc    string
}

func Parse(conllBytes []byte) []ConllWord {
	conllString := string(conllBytes[:])
	conll := []ConllWord{}

	for _, v := range strings.Split(conllString, "\n") {
		parts := strings.Split(v, "\t")

		id, _ := strconv.ParseInt(parts[0], 10, 0)

		head, _ := strconv.ParseInt(parts[6], 10, 0)

		conllWord := ConllWord{
			int(id),
			parts[1],
			parts[2],
			parts[3],
			parts[4],
			parts[5],
			int(head),
			parts[7],
			parts[8],
			parts[9],
		}

		conll = append(conll, conllWord)
	}

	return conll
}
