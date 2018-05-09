package fizz

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Index struct {
	Name    string
	Columns []string
	Unique  bool
	Options Options
}

func (f fizzer) AddIndex() interface{} {
	return func(table string, columns interface{}, options Options) {
		i := Index{}
		switch t := columns.(type) {
		default:
			log.WithField("table", table).WithField("type", fmt.Sprintf("%T", t)).Warn("unexpected type when adding index")
		case string:
			i.Columns = []string{t}
		case []interface{}:
			cl := make([]string, len(t))
			for i, c := range t {
				cl[i] = c.(string)
			}
			i.Columns = cl
		}

		if options["name"] != nil {
			i.Name = options["name"].(string)
		} else {
			i.Name = fmt.Sprintf("%s_%s_idx", table, strings.Join(i.Columns, "_"))
		}
		i.Unique = options["unique"] != nil
		f.add(f.Bubbler.AddIndex(Table{
			Name:    table,
			Indexes: []Index{i},
		}))
	}
}

func (f fizzer) DropIndex() interface{} {
	return func(table, name string) {
		f.add(f.Bubbler.DropIndex(Table{
			Name: table,
			Indexes: []Index{
				{Name: name},
			},
		}))
	}
}

func (f fizzer) RenameIndex() interface{} {
	return func(table, old, new string) {
		f.add(f.Bubbler.RenameIndex(Table{
			Name: table,
			Indexes: []Index{
				{Name: old},
				{Name: new},
			},
		}))
	}
}
