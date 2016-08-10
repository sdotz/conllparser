package conllparser_test

import (
	"testing"
	"github.com/sdotz/slask/conllparser"
	"fmt"
)

func TestParse(t *testing.T) {
	conllOutput :=
		`1	Sally	_	NOUN	NNP	_	2	nsubj	_	_
		2	went	_	VERB	VBD	_	0	ROOT	_	_
		3	ham	_	NOUN	NN	_	2	dobj	_	_
		4	on	_	ADP	IN	_	2	prep	_	_
		5	a	_	DET	DT	_	7	det	_	_
		6	peanut	_	NOUN	NN	_	7	nn	_	_
		7	butter	_	NOUN	NN	_	4	pobj	_	_
		8	and	_	CONJ	CC	_	7	cc	_	_
		9	bacon	_	NOUN	NN	_	10	nn	_	_
		10	sandwich	_	NOUN	NN	_	7	conj	_	_
		11	.	_	.	.	_	2	punct	_	_`

	conllBytes := []byte(conllOutput)

	conll := conllparser.Parse(conllBytes)

	fmt.Println(conll)

	if conll == nil {
		t.Error("Nope")
	}

}