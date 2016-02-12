package template

import (
	"fmt"
	"testing"

	"github.com/OpenNebula/goca/api"

	//. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/check.v1"
)

func Test(t *testing.T) {
	check.TestingT(t)
}

type S struct{}

var _ = check.Suite(&S{})

func (s *S) TestGetTemplateByName(c *check.C) {
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	flav := TemplateReqs{TemplateName: "newone", Client: client}
	res, error := flav.GetTemplateByName()
	fmt.Println(res[0].Id)
	c.Assert(error, check.IsNil)
}

/*
func (s *S) TestGetTemplate(c *check.C) {
	client, _ := api.NewRPCClient("http://localhost:2633/RPC2", "oneadmin", "RaifZuewjoc4")
	flav := TemplateReqs{TemplateId: 33, Client: client}
	_, error := flav.GetTemplate()
	c.Assert(error, check.IsNil)
}
*/
