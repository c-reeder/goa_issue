package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("TestAPI", func() {
	Title("My Test API")
	Description("API for testing goa stuff")
	Server("TestServer", func() {
		Services("TestService")
	})
})

var _ = Service("TestService", func() {
	Description("The calc service performs operations on numbers.")

	Method("create", func() {
		Description("Create a new thing")
		Payload(FirstType)

		HTTP(func() {
			POST("")
		})
	})

})
var FirstType = Type("FirstType", func() {
	Meta("struct:pkg:path", "types")
	Attribute("id", String)
	Attribute("name", String)
	Attribute("description", String)
	Attribute("thing", SecondType)
})

var SecondType = Type("SecondType", func() {
	Attribute("Description", String)
})
