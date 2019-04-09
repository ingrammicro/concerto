package format

import (
	"fmt"
	"github.com/ingrammicro/concerto/utils"
	"io"
	"reflect"
	"strings"
	"text/tabwriter"

	log "github.com/Sirupsen/logrus"
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
	log.Debug("printItemAux")

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
				fmt.Fprintln(w, fmt.Sprintf("%s:\t%s", it.Type().Field(i).Tag.Get("header"), it.Field(i).Elem()))
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
	log.Debug("printListHeadersAux")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Type.Kind() == reflect.Struct {
			f.printListHeadersAux(w, field.Type)
		}

		showTags := strings.Split(field.Tag.Get("show"), ",")
		if !utils.Contains(showTags, "nolist") {
			fmt.Fprint(w, fmt.Sprintf("%+v\t", field.Tag.Get("header")))
		}
	}
}

func (f *TextFormatter) printListBodyAux(w *tabwriter.Writer, t reflect.Value) {
	log.Debug("printListBodyAux")

	for i := 0; i < t.NumField(); i++ {
		showTags := strings.Split(t.Type().Field(i).Tag.Get("show"), ",")
		if !utils.Contains(showTags, "nolist") {
			field := t.Field(i)
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
					f.printListBodyAux(w, field)
				} else {
					fmt.Fprint(w, fmt.Sprintf("%+v\t", field.Interface()))
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

	// Headers
	f.printListHeadersAux(w, reflect.TypeOf(items).Elem())
	fmt.Fprintln(w)

	// Body
	for pos := 0; pos < its.Len(); pos++ {
		f.printListBodyAux(w, its.Index(pos))
		fmt.Fprintln(w)
	}

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
