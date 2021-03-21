package main

import (
	"net/http"
	"os"
	"path"
	"strings"

	types "github.com/rancher/wrangler/pkg/schemas"

	"github.com/futuretea/go-harvester/codegen/generator"
)

var (
	outputDir  = "./pkg/generated"
	baseCattle = "../client/generated"
)

func getVersion(schemas *types.Schemas) (version string) {
	for _, schema := range schemas.Schemas() {
		if version == "" {
			version = schema.Attributes["version"].(string)
			continue
		}
		if version != schema.Attributes["version"].(string) {
			panic("schema set contains two APIVersions")
		}
	}

	return
}

func GenerateClient(schemas *types.Schemas, group string, backendTypes map[string]bool) {
	version := getVersion(schemas)
	groupBase := strings.Split(group, ".")[0]

	cattleOutputPackage := path.Join(baseCattle, groupBase, version)

	if err := generator.GenerateClient(schemas, backendTypes, outputDir, cattleOutputPackage); err != nil {
		panic(err)
	}
}

var (
	userSchema = types.Schema{
		ID:         "user",
		CodeName:   "User",
		PluralName: "harvester.cattle.io.users",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "harvester.cattle.io",
			"importPackage": "github.com/rancher/harvester/pkg/apis/harvester.cattle.io/v1alpha1",
			"importAlias":   "harv1",
			"importType":    "User",
			"namespaced":    false,
		},
	}
	imageSchema = types.Schema{
		ID:         "image",
		CodeName:   "Image",
		PluralName: "harvester.cattle.io.virtualmachineimages",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "harvester.cattle.io",
			"importPackage": "github.com/rancher/harvester/pkg/apis/harvester.cattle.io/v1alpha1",
			"importAlias":   "harv1",
			"importType":    "VirtualMachineImage",
			"namespaced":    true,
		},
	}
	keypairSchema = types.Schema{
		ID:         "keypair",
		CodeName:   "Keypair",
		PluralName: "harvester.cattle.io.keypairs",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "harvester.cattle.io",
			"importPackage": "github.com/rancher/harvester/pkg/apis/harvester.cattle.io/v1alpha1",
			"importAlias":   "harv1",
			"importType":    "KeyPair",
			"namespaced":    true,
		},
	}
	settingSchema = types.Schema{
		ID:         "setting",
		CodeName:   "Setting",
		PluralName: "harvester.cattle.io.settings",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "harvester.cattle.io",
			"importPackage": "github.com/rancher/harvester/pkg/apis/harvester.cattle.io/v1alpha1",
			"importAlias":   "harv1",
			"importType":    "Setting",
			"namespaced":    false,
		},
	}
	networkSchema = types.Schema{
		ID:         "network",
		CodeName:   "Network",
		PluralName: "k8s.cni.cncf.io.network-attachment-definitions",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "k8s.cni.cncf.io",
			"importPackage": "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1",
			"importAlias":   "cniv1",
			"importType":    "NetworkAttachmentDefinition",
			"namespaced":    true,
		},
	}
	vmSchema = types.Schema{
		ID:         "virtualmachine",
		CodeName:   "VirtualMachine",
		PluralName: "kubevirt.io.virtualmachines",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		ResourceActions: map[string]types.Action{
			"abortMigration": {},
			"migrate":        {},
			"pause":          {},
			"unpause":        {},
			"restart":        {},
			"start":          {},
			"stop":           {},
			"restore": {
				Input: "restoreInput",
			},
			"backup": {
				Input: "backupInput",
			},
			"ejectCdRom": {
				Input: "ejectCdRomActionInput",
			},
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "kubevirt.io",
			"importPackage": "kubevirt.io/client-go/api/v1",
			"importAlias":   "kubevirtv1",
			"importType":    "VirtualMachine",
			"namespaced":    true,
		},
	}
	vmiSchema = types.Schema{
		ID:         "virtualMachineInstance",
		CodeName:   "VirtualMachineInstance",
		PluralName: "kubevirt.io.virtualmachineinstance",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "kubevirt.io",
			"importPackage": "kubevirt.io/client-go/api/v1",
			"importAlias":   "kubevirtv1",
			"importType":    "VirtualMachineInstance",
			"namespaced":    true,
		},
	}
	volumeSchema = types.Schema{
		ID:         "volume",
		CodeName:   "Volume",
		PluralName: "cdi.kubevirt.io.datavolumes",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "cdi.kubevirt.io",
			"importPackage": "kubevirt.io/containerized-data-importer/pkg/apis/core/v1beta1",
			"importAlias":   "cdiv1beta1",
			"importType":    "DataVolume",
			"namespaced":    true,
		},
	}
	serviceSchema = types.Schema{
		ID:         "service",
		CodeName:   "Service",
		PluralName: "services",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "",
			"importPackage": "k8s.io/api/core/v1",
			"importAlias":   "corev1",
			"importType":    "Service",
			"namespaced":    true,
		},
	}
	nodeSchema = types.Schema{
		ID:         "node",
		CodeName:   "Node",
		PluralName: "nodes",
		CollectionMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		Attributes: map[string]interface{}{
			"version":       "v1",
			"group":         "",
			"importPackage": "k8s.io/api/core/v1",
			"importAlias":   "corev1",
			"importType":    "Node",
			"namespaced":    false,
		},
	}
)

func main() {
	_ = os.Unsetenv("GOPATH")

	v1Schemas, err := types.NewSchemas()
	if err != nil {
		panic(err)
	}
	v1Schemas.
		MustAddSchema(nodeSchema).
		MustAddSchema(serviceSchema).
		MustAddSchema(userSchema).
		MustAddSchema(vmSchema).
		MustAddSchema(imageSchema).
		MustAddSchema(keypairSchema).
		MustAddSchema(settingSchema).
		MustAddSchema(networkSchema).
		MustAddSchema(vmiSchema).
		MustAddSchema(volumeSchema)

	GenerateClient(v1Schemas, "", map[string]bool{})
}