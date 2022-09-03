package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("calc", func() {
	Title("Calculator Service")
	Description("Service for multiplying numbers, a Goa teaser")
	Server("calc", func() {
		Services("TestService")
	})
})

var _ = Service("TestService", func() {
	Description("The calc service performs operations on numbers.")

	Method("update", func() {
		Description("Update an existing workspace information")
		Payload(Workspace, func() {
			Required("id")
		})

		HTTP(func() {
			PATCH("/{id}")
		})
	})

	Method("set", func() {
		Description("Set the workspace to be used")
		Payload(Workspace)

		HTTP(func() {
			PUT("/{id}")
		})
	})

})
var Workspace = Type("workspace", func() {
	Attribute("id", String, "Workspace ID")
	Attribute("name", String)
	Attribute("description", String)
	Attribute("destination_ids", ArrayOf(String))
	Attribute("logo_url", String)
})
