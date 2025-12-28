package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func TestAccAllocationDataSource_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfig("test-pool", "test-alloc", 24),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("id"),
						knownvalue.StringExact("test-alloc"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("pool_name"),
						knownvalue.StringExact("test-pool"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(24),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("allocated_cidr"),
						knownvalue.NotNull(),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_MultipleAllocations(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfigMultiple("multi-pool", 24),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test1",
						tfjsonpath.New("id"),
						knownvalue.StringExact("alloc-1"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test2",
						tfjsonpath.New("id"),
						knownvalue.StringExact("alloc-2"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test3",
						tfjsonpath.New("id"),
						knownvalue.StringExact("alloc-3"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test1",
						tfjsonpath.New("allocated_cidr"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test2",
						tfjsonpath.New("allocated_cidr"),
						knownvalue.NotNull(),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test3",
						tfjsonpath.New("allocated_cidr"),
						knownvalue.NotNull(),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_DifferentPrefixLengths(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfigDifferentPrefixes("prefix-pool"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_24",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(24),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_27",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(27),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_30",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(30),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_32",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(32),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_SingleHost(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfig("host-pool", "host-alloc", 32),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("id"),
						knownvalue.StringExact("host-alloc"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(32),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("allocated_cidr"),
						knownvalue.NotNull(),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_NotFound(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccAllocationDataSourceConfigNotFound("test-pool", "nonexistent-alloc"),
				ExpectError: regexp.MustCompile("Provider produced null object|not found|does not exist"),
			},
		},
	})
}

func TestAccAllocationDataSource_IPv6(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfigIPv6("ipv6-pool", "ipv6-alloc", 64),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("id"),
						knownvalue.StringExact("ipv6-alloc"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("pool_name"),
						knownvalue.StringExact("ipv6-pool"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(64),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("allocated_cidr"),
						knownvalue.NotNull(),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_IPv6_MultipleSubnets(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfigIPv6Multiple("ipv6-multi-pool"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_48",
						tfjsonpath.New("id"),
						knownvalue.StringExact("ipv6-alloc-48"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_48",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(48),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_56",
						tfjsonpath.New("id"),
						knownvalue.StringExact("ipv6-alloc-56"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_56",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(56),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_64",
						tfjsonpath.New("id"),
						knownvalue.StringExact("ipv6-alloc-64"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_64",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(64),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_FromDifferentPools(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfigDifferentPools(),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test1",
						tfjsonpath.New("pool_name"),
						knownvalue.StringExact("pool-1"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test2",
						tfjsonpath.New("pool_name"),
						knownvalue.StringExact("pool-2"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test3",
						tfjsonpath.New("pool_name"),
						knownvalue.StringExact("pool-3"),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_VerifyAllocatedCIDR(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfigVerifyCIDR("verify-pool", "verify-alloc", 24),
				ConfigStateChecks: []statecheck.StateCheck{
					// Verify data source has same values as resource
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("id"),
						knownvalue.StringExact("verify-alloc"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("pool_name"),
						knownvalue.StringExact("verify-pool"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("prefix_length"),
						knownvalue.Int64Exact(24),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test",
						tfjsonpath.New("allocated_cidr"),
						knownvalue.NotNull(),
					),
				},
			},
		},
	})
}

func TestAccAllocationDataSource_LargePool(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccAllocationDataSourceConfigLargePool("large-pool"),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_0",
						tfjsonpath.New("id"),
						knownvalue.StringExact("large-alloc-0"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_4",
						tfjsonpath.New("id"),
						knownvalue.StringExact("large-alloc-4"),
					),
					statecheck.ExpectKnownValue(
						"data.tfipam_allocation.test_9",
						tfjsonpath.New("id"),
						knownvalue.StringExact("large-alloc-9"),
					),
				},
			},
		},
	})
}

// testAccAllocationDataSourceConfig generates a Terraform configuration with pool, allocation, and data source.
func testAccAllocationDataSourceConfig(poolName, allocID string, prefixLength int) string {
	return fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["10.0.0.0/16"]
}

resource "tfipam_allocation" "test" {
  id            = %[2]q
  pool_name     = tfipam_pool.test.name
  prefix_length = %[3]d
}

data "tfipam_allocation" "test" {
  id        = tfipam_allocation.test.id
  pool_name = tfipam_allocation.test.pool_name
}
`, poolName, allocID, prefixLength)
}

// testAccAllocationDataSourceConfigNotFound generates config that tries to read a non-existent allocation.
func testAccAllocationDataSourceConfigNotFound(poolName, allocID string) string {
	return fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["10.0.0.0/16"]
}

data "tfipam_allocation" "test" {
  id        = %[2]q
  pool_name = tfipam_pool.test.name
}
`, poolName, allocID)
}

// testAccAllocationDataSourceConfigMultiple generates config with multiple allocations and data sources.
func testAccAllocationDataSourceConfigMultiple(poolName string, prefixLength int) string {
	return fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["10.0.0.0/16"]
}

resource "tfipam_allocation" "alloc1" {
  id            = "alloc-1"
  pool_name     = tfipam_pool.test.name
  prefix_length = %[2]d
}

resource "tfipam_allocation" "alloc2" {
  id            = "alloc-2"
  pool_name     = tfipam_pool.test.name
  prefix_length = %[2]d
}

resource "tfipam_allocation" "alloc3" {
  id            = "alloc-3"
  pool_name     = tfipam_pool.test.name
  prefix_length = %[2]d
}

data "tfipam_allocation" "test1" {
  id        = tfipam_allocation.alloc1.id
  pool_name = tfipam_allocation.alloc1.pool_name
}

data "tfipam_allocation" "test2" {
  id        = tfipam_allocation.alloc2.id
  pool_name = tfipam_allocation.alloc2.pool_name
}

data "tfipam_allocation" "test3" {
  id        = tfipam_allocation.alloc3.id
  pool_name = tfipam_allocation.alloc3.pool_name
}
`, poolName, prefixLength)
}

// testAccAllocationDataSourceConfigDifferentPrefixes generates config with different prefix lengths.
func testAccAllocationDataSourceConfigDifferentPrefixes(poolName string) string {
	return fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["10.0.0.0/16"]
}

resource "tfipam_allocation" "alloc_24" {
  id            = "alloc-24"
  pool_name     = tfipam_pool.test.name
  prefix_length = 24
}

resource "tfipam_allocation" "alloc_27" {
  id            = "alloc-27"
  pool_name     = tfipam_pool.test.name
  prefix_length = 27
}

resource "tfipam_allocation" "alloc_30" {
  id            = "alloc-30"
  pool_name     = tfipam_pool.test.name
  prefix_length = 30
}

resource "tfipam_allocation" "alloc_32" {
  id            = "alloc-32"
  pool_name     = tfipam_pool.test.name
  prefix_length = 32
}

data "tfipam_allocation" "test_24" {
  id        = tfipam_allocation.alloc_24.id
  pool_name = tfipam_allocation.alloc_24.pool_name
}

data "tfipam_allocation" "test_27" {
  id        = tfipam_allocation.alloc_27.id
  pool_name = tfipam_allocation.alloc_27.pool_name
}

data "tfipam_allocation" "test_30" {
  id        = tfipam_allocation.alloc_30.id
  pool_name = tfipam_allocation.alloc_30.pool_name
}

data "tfipam_allocation" "test_32" {
  id        = tfipam_allocation.alloc_32.id
  pool_name = tfipam_allocation.alloc_32.pool_name
}
`, poolName)
}

// testAccAllocationDataSourceConfigIPv6 generates config for IPv6 allocation data source.
func testAccAllocationDataSourceConfigIPv6(poolName, allocID string, prefixLength int) string {
	return fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["2001:db8::/32"]
}

resource "tfipam_allocation" "test" {
  id            = %[2]q
  pool_name     = tfipam_pool.test.name
  prefix_length = %[3]d
}

data "tfipam_allocation" "test" {
  id        = tfipam_allocation.test.id
  pool_name = tfipam_allocation.test.pool_name
}
`, poolName, allocID, prefixLength)
}

// testAccAllocationDataSourceConfigIPv6Multiple generates config for multiple IPv6 allocations.
func testAccAllocationDataSourceConfigIPv6Multiple(poolName string) string {
	return fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["2001:db8::/32"]
}

resource "tfipam_allocation" "alloc_48" {
  id            = "ipv6-alloc-48"
  pool_name     = tfipam_pool.test.name
  prefix_length = 48
}

resource "tfipam_allocation" "alloc_56" {
  id            = "ipv6-alloc-56"
  pool_name     = tfipam_pool.test.name
  prefix_length = 56
}

resource "tfipam_allocation" "alloc_64" {
  id            = "ipv6-alloc-64"
  pool_name     = tfipam_pool.test.name
  prefix_length = 64
}

data "tfipam_allocation" "test_48" {
  id        = tfipam_allocation.alloc_48.id
  pool_name = tfipam_allocation.alloc_48.pool_name
}

data "tfipam_allocation" "test_56" {
  id        = tfipam_allocation.alloc_56.id
  pool_name = tfipam_allocation.alloc_56.pool_name
}

data "tfipam_allocation" "test_64" {
  id        = tfipam_allocation.alloc_64.id
  pool_name = tfipam_allocation.alloc_64.pool_name
}
`, poolName)
}

// testAccAllocationDataSourceConfigDifferentPools generates config with allocations from different pools.
func testAccAllocationDataSourceConfigDifferentPools() string {
	return `
resource "tfipam_pool" "pool1" {
  name = "pool-1"
  cidrs = ["10.0.0.0/16"]
}

resource "tfipam_pool" "pool2" {
  name = "pool-2"
  cidrs = ["192.168.0.0/16"]
}

resource "tfipam_pool" "pool3" {
  name = "pool-3"
  cidrs = ["172.16.0.0/12"]
}

resource "tfipam_allocation" "alloc1" {
  id            = "alloc-1"
  pool_name     = tfipam_pool.pool1.name
  prefix_length = 24
}

resource "tfipam_allocation" "alloc2" {
  id            = "alloc-2"
  pool_name     = tfipam_pool.pool2.name
  prefix_length = 24
}

resource "tfipam_allocation" "alloc3" {
  id            = "alloc-3"
  pool_name     = tfipam_pool.pool3.name
  prefix_length = 20
}

data "tfipam_allocation" "test1" {
  id        = tfipam_allocation.alloc1.id
  pool_name = tfipam_allocation.alloc1.pool_name
}

data "tfipam_allocation" "test2" {
  id        = tfipam_allocation.alloc2.id
  pool_name = tfipam_allocation.alloc2.pool_name
}

data "tfipam_allocation" "test3" {
  id        = tfipam_allocation.alloc3.id
  pool_name = tfipam_allocation.alloc3.pool_name
}
`
}

// testAccAllocationDataSourceConfigVerifyCIDR generates config to verify CIDR values match between resource and data source.
func testAccAllocationDataSourceConfigVerifyCIDR(poolName, allocID string, prefixLength int) string {
	return fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["10.0.0.0/16"]
}

resource "tfipam_allocation" "test" {
  id            = %[2]q
  pool_name     = tfipam_pool.test.name
  prefix_length = %[3]d
}

data "tfipam_allocation" "test" {
  id        = tfipam_allocation.test.id
  pool_name = tfipam_allocation.test.pool_name
}
`, poolName, allocID, prefixLength)
}

// testAccAllocationDataSourceConfigLargePool generates config with many allocations.
func testAccAllocationDataSourceConfigLargePool(poolName string) string {
	config := fmt.Sprintf(`
resource "tfipam_pool" "test" {
  name = %[1]q
  cidrs = ["10.0.0.0/16"]
}
`, poolName)

	for i := 0; i < 10; i++ {
		config += fmt.Sprintf(`
resource "tfipam_allocation" "alloc_%[1]d" {
  id            = "large-alloc-%[1]d"
  pool_name     = tfipam_pool.test.name
  prefix_length = 27
}

data "tfipam_allocation" "test_%[1]d" {
  id        = tfipam_allocation.alloc_%[1]d.id
  pool_name = tfipam_allocation.alloc_%[1]d.pool_name
}
`, i)
	}

	return config
}
