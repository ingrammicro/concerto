package format

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/ingrammicro/concerto/utils"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"
)

// TextFormatter prints items and lists
type TextFormatter struct {
	output io.Writer
}

// NewTextFormatter creates a new TextFormatter
func NewTextFormatter(out io.Writer) *TextFormatter {
	log.Debug("Creating Text formatter")

	return &TextFormatter{
		output: out,
	}
}

func (f *TextFormatter) printItemAux(w *tabwriter.Writer, item interface{}) error {
	it := reflect.ValueOf(item)
	for i := 0; i < it.NumField(); i++ {
		showTags := strings.Split(it.Type().Field(i).Tag.Get("show"), ",")
		if !utils.Contains(showTags, "noshow") {
			switch it.Field(i).Type().String() {
			case "time.Time":
				fmt.Fprintln(w, fmt.Sprintf("%s:\t%+v", it.Type().Field(i).Tag.Get("header"), it.Field(i).Interface()))
			case "json.RawMessage":
				fmt.Fprintln(w, fmt.Sprintf("%s:\t%s", it.Type().Field(i).Tag.Get("header"), it.Field(i).Interface()))
			case "*json.RawMessage":
				if it.Field(i).IsNil() {
					fmt.Fprintln(w, fmt.Sprintf("%s:\t", it.Type().Field(i).Tag.Get("header")))
				} else {
					fmt.Fprintln(w, fmt.Sprintf("%s:\t%s", it.Type().Field(i).Tag.Get("header"), it.Field(i).Elem()))
				}
			default:
				if it.Field(i).Kind() == reflect.Struct {
					f.printItemAux(w, it.Field(i).Interface())
				} else {
					fmt.Fprintln(w, fmt.Sprintf("%s:\t%+v", it.Type().Field(i).Tag.Get("header"), it.Field(i).Interface()))
				}
			}
		}
	}
	return nil
}

// PrintItem prints item
func (f *TextFormatter) PrintItem(item interface{}) error {
	log.Debug("PrintItem")

	w := tabwriter.NewWriter(f.output, 15, 1, 3, ' ', 0)
	f.printItemAux(w, item)
	w.Flush()

	return nil
}

func (f *TextFormatter) printListHeadersAux(w *tabwriter.Writer, t reflect.Type) {
	n := 0
	if t.Kind() != reflect.Struct {
		n = t.Elem().NumField()
	} else {
		n = t.NumField()
	}
	var field reflect.StructField
	for i := 0; i < n; i++ {
		if t.Kind() != reflect.Struct {
			field = t.Elem().Field(i)
		} else {
			field = t.Field(i)
		}
		if field.Type.Kind() == reflect.Struct {
			f.printListHeadersAux(w, field.Type)
		}
		showTags := strings.Split(field.Tag.Get("show"), ",")
		if !utils.Contains(showTags, "nolist") && field.Tag.Get("header") != "" {
			fmt.Fprint(w, fmt.Sprintf("%+v\t", field.Tag.Get("header")))
		}
	}
}

func (f *TextFormatter) printListBodyAux(w *tabwriter.Writer, rv reflect.Value, depth int) {
	switch rv.Kind() {
	//case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
	case reflect.Slice:
		for i := 0; i < rv.Len(); i++ {
			if rv.Index(i).Kind() == reflect.Ptr {
				if !rv.Index(i).IsNil() {
					f.printListBodyAux(w, rv.Index(i).Elem(), depth+1)
					fmt.Fprintln(w)
				}
			} else {
				///usr/local/go/bin/go run main.go firewall list
				f.printListBodyAux(w, rv.Index(i), depth+1)
				fmt.Fprintln(w)
			}
		}
	case reflect.Struct:
		for i := 0; i < rv.NumField(); i++ {
			showTags := strings.Split(rv.Type().Field(i).Tag.Get("show"), ",")
			if !utils.Contains(showTags, "nolist") {
				field := rv.Field(i)
				switch field.Type().String() {
				case "time.Time":
					fmt.Fprint(w, fmt.Sprintf("%+v\t", field.Interface()))
				case "json.RawMessage":
					fmt.Fprint(w, fmt.Sprintf("%s\t", field.Interface()))
				case "*json.RawMessage":
					if field.IsNil() {
						fmt.Fprint(w, fmt.Sprintf(" \t"))
					} else {
						fmt.Fprint(w, fmt.Sprintf("%s\t", field.Elem()))
					}
				default:
					if field.Kind() == reflect.Struct {
						f.printListBodyAux(w, field, depth+1)
					} else {
						fmt.Fprint(w, fmt.Sprintf("%+v\t", field))
					}
				}
			}
		}
	}
}

// PrintList prints item list
func (f *TextFormatter) PrintList(items interface{}) error {
	log.Debug("PrintList")

	// should be an array
	its := reflect.ValueOf(items)
	t := its.Type().Kind()
	if t != reflect.Slice {
		return fmt.Errorf("couldn't print list. Expected slice, but received %s", t.String())
	}

	w := tabwriter.NewWriter(f.output, 15, 1, 3, ' ', 0)
	f.printListHeadersAux(w, reflect.TypeOf(items).Elem())
	fmt.Fprintln(w)

	f.printListBodyAux(w, reflect.ValueOf(items), 0)
	fmt.Fprintln(w)

	w.Flush()

	return nil
}

// PrintError prints an error
func (f *TextFormatter) PrintError(context string, err error) {
	log.Debug("PrintError")

	f.output.Write([]byte(fmt.Sprintf("ERROR: %s\n -> %s\n", context, err)))
}

// PrintFatal prints an error and exists
func (f *TextFormatter) PrintFatal(context string, err error) {
	log.Debug("PrintFatal")

	f.PrintError(context, err)
	osExit(1)
}
