package utils

import (
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/url"
	"reflect"
	"swenson-he-task/models"
	"testing"
)

func TestParseFilters(t *testing.T) {
	t.Run("Test filtering string", func(t *testing.T) {
		r := &http.Request{}
		u := &url.URL{}
		q := u.Query()
		q.Set("productType", "large machine")
		r.URL = u
		r.URL.RawQuery = q.Encode()

		filters := ParseFilters(r, reflect.ValueOf(&models.CoffeeProduct{}).Elem())
		expectedFilters := []Filter{
			{Field: "productType", Value: "large machine", Operator: "=", Type: "string"},
		}
		assert.Equal(t, filters, &expectedFilters)
	})

	t.Run("Test filtering int", func(t *testing.T) {
		r := &http.Request{}
		u := &url.URL{}
		q := u.Query()
		q.Set("packSize", "3")
		r.URL = u
		r.URL.RawQuery = q.Encode()

		filters := ParseFilters(r, reflect.ValueOf(&models.CoffeeProduct{}).Elem())
		expectedFilters := []Filter{
			{Field: "packSize", Value: 3, Operator: "=", Type: "int"},
		}
		assert.Equal(t, filters, &expectedFilters)
	})

	t.Run("Test filtering string,int,bool", func(t *testing.T) {
		r := &http.Request{}
		u := &url.URL{}
		q := u.Query()
		q.Set("productType", "large machine")
		q.Set("packSize", "3")
		q.Set("waterLineCompatible", "true")
		r.URL = u
		r.URL.RawQuery = q.Encode()

		filters := ParseFilters(r, reflect.ValueOf(&models.CoffeeProduct{}).Elem())
		expectedFilters := []Filter{
			{Field: "productType", Value: "large machine", Operator: "=", Type: "string"},
			{Field: "waterLineCompatible", Value: true, Operator: "=", Type: "bool"},
			{Field: "packSize", Value: 3, Operator: "=", Type: "int"},
		}
		assert.Equal(t, filters, &expectedFilters)
	})

	t.Run("Test filtering string,int", func(t *testing.T) {
		r := &http.Request{}
		u := &url.URL{}
		q := u.Query()
		q.Set("productType", "large machine")
		q.Set("packSize", "3")
		r.URL = u
		r.URL.RawQuery = q.Encode()

		filters := ParseFilters(r, reflect.ValueOf(&models.CoffeeProduct{}).Elem())
		expectedFilters := []Filter{
			{Field: "productType", Value: "large machine", Operator: "=", Type: "string"},
			{Field: "packSize", Value: 3, Operator: "=", Type: "int"},
		}
		assert.Equal(t, filters, &expectedFilters)
	})

}
