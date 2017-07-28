package house

const testVersion = 1

type verseMain struct {
	Index int
	Line  string
}

type verseOpener struct {
	Index int
	Line  string
}

var delta = []verseOpener{
	{1, "This is the house that Jack built."},
	{2, "This is the malt"},
	{3, "This is the rat"},
	{4, "This is the cat"},
	{5, "This is the dog"},
	{6, "This is the cow with the crumpled horn"},
	{7, "This is the maiden all forlorn"},
	{8, "This is the man all tattered and torn"},
	{9, "This is the priest all shaven and shorn"},
	{10, "This is the rooster that crowed in the morn"},
	{11, "This is the farmer sowing his corn"},
	{12, "This is the horse and the hound and the horn"},
}

var constant = []verseMain{
	{1, ""},
	{2, "that lay in the house that Jack built."},
	{3, "that ate the malt"},
	{4, "that killed the rat"},
	{5, "that worried the cat"},
	{6, "that tossed the dog"},
	{7, "that milked the cow with the crumpled horn"},
	{8, "that kissed the maiden all forlorn"},
	{9, "that married the man all tattered and torn"},
	{10, "that woke the priest all shaven and shorn"},
	{11, "that kept the rooster that crowed in the morn"},
	{12, "that belonged to the farmer sowing his corn"},
}

func Song() string {

	var song string
	for n := 0; n <= 11; n++ {
		if n == 0 {
			song += buildVerse(n)
		} else {
			song += "\n" + "\n" + buildVerse(n)
		}
	}

	return song
}

func Verse(n int) string {

	var output string
	output = buildVerse(n - 1)

	return output
}

func buildVerse(n int) string {

	var verses string

	if constant[n].Line != "" {
		verses = delta[n].Line + "\n" + constant[n].Line
	} else {
		verses = delta[n].Line + constant[n].Line
	}

	for l := n - 1; l >= 0; l-- {
		if constant[l].Line != "" {
			verses += "\n" + constant[l].Line
		}
	}

	return verses
}
