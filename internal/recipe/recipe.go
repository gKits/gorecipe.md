package recipe

import (
	"errors"
	"io"
	"os"
	"text/template"

	"github.com/kkyr/go-recipe/pkg/recipe"
)

const recipeTemplate = `# {{.title}}                                                                                    
                                                                                                                        
{{.description}}                                                                                                        
                                                                                                                        
## Ingredients                                                                                                          
                                                                                                                        
{{range .ingredients}}{{println "-" .}}{{end}}                                                                          
                                                                                                                        
## Instructions                                                                                                         
                                                                                                                        
{{range $index, $instruction := .instructions}}{{len (printf "a%*s" $index "")}}{{println "." $instruction}}{{end}}     
`

type MDScraper struct {
	w    io.Writer
	tmpl *template.Template
	cfg
}

type cfg struct {
	tmplStr         string
	withHugoHeaders bool
}

var defaultCfg = cfg{
	tmplStr:         recipeTemplate,
	withHugoHeaders: false,
}

type recipeData struct {
	Title, Description        string
	Ingredients, Instructions []string
}

func NewMDScraper(w io.Writer, opts ...Option) (*MDScraper, error) {
	cfg := defaultCfg
	for _, opt := range opts {
		opt(&cfg)
	}

	tmpl, err := template.New("").Parse(cfg.tmplStr)
	if err != nil {
		return nil, err
	}

	return &MDScraper{
		w:    w,
		tmpl: tmpl,
		cfg:  cfg,
	}, nil
}

func MDScrape(w io.Writer, url string, opts ...Option) error {
	conv, err := NewMDScraper(w, opts...)
	if err != nil {
		return err
	}
	return conv.MDScrape(url)
}

func (conv *MDScraper) MDScrape(url string) error {
	rec, err := recipe.ScrapeURL(url)
	if err != nil {
		panic(err)
	}

	var (
		data recipeData
		ok   bool
	)
	if data.Title, ok = rec.Name(); !ok {
		return errors.New("scrape does not contain a recipe name")
	}

	if data.Description, ok = rec.Description(); !ok {
		return errors.New("scrape does not contain a recipe description")
	}

	if data.Ingredients, ok = rec.Ingredients(); !ok {
		return errors.New("scrape does not contain recipe ingredients")
	}

	if data.Instructions, ok = rec.Instructions(); !ok {
		return errors.New("scrape does not contain recipe instructions")
	}
	return conv.tmpl.Execute(os.Stdout, data)
}

type Option func(*cfg)

func WithTemplate(tmpl string) Option {
	return func(c *cfg) {
		c.tmplStr = tmpl
	}
}

func WithHugoHeaders(b bool) Option {
	return func(c *cfg) {
		c.withHugoHeaders = b
	}
}
