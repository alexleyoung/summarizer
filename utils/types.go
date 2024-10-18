package utils

type GenericPage struct {
	Titles []string
	Paragraphs []string
	Content string
}

type WikiPage struct {
	Title string
	Headings []string
	Subtitles []string
	Paragraphs []string
	Content string
}