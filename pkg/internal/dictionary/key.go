package dictionary

type DictObj struct {
	MapName string
	MaxLen  int
	Dict    []string
}

var the = DictObj{MapName: "the"}
var in = DictObj{MapName: "in"}
var and = DictObj{MapName: "and"}

var Hipku4 = []DictObj{
	the,
	AnimalAdjectives,
	AnimalColors,
	AnimalNouns,
	AnimalVerbs,
	in,
	the,
	NatureAdjectives,
	NatureNouns,
	PlantNouns,
	PlantVerbs,
}

var Hipku6 = []DictObj{
	Adjectives,
	Nouns,
	and,
	Adjectives,
	Nouns,
	Verbs,
	Adjectives,
	Adjectives,
	Adjectives,
	Adjectives,
	Adjectives,
	Nouns,
	Adjectives,
	Nouns,
	Verbs,
	Adjectives,
	Nouns,
}
