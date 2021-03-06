package helpers

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	. "github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/onsi/gomega"
	. "github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-routing-test-helpers/schema"
	"github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/cf"
)

func GetOrgQuotaDefinitionUrl(orgGuid string, timeout time.Duration) (string, error) {
	orgGuid = strings.TrimSuffix(orgGuid, "\n")
	response := cf.Cf("curl", fmt.Sprintf("/v2/organizations/%s", string(orgGuid)))
	Expect(response.Wait(timeout)).To(Exit(0))

	var orgEntity schema.OrgResource
	err := json.Unmarshal(response.Out.Contents(), &orgEntity)
	if err != nil {
		return "", err
	}

	return orgEntity.Entity.QuotaDefinitionUrl, nil
}
