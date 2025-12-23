# buid and run the scraper
br:
	go build && ./scraper

# run migrtions with goose
migrate:
	goose postgres postgres://ab01fazl:@localhost:5432/scraper up