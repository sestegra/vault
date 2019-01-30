package pki

import "github.com/hashicorp/vault/logical/framework"

// addIssueAndSignCommonFields adds fields common to both CA and non-CA issuing
// and signing
func addIssueAndSignCommonFields(fields map[string]*framework.FieldSchema) map[string]*framework.FieldSchema {
	fields["exclude_cn_from_sans"] = &framework.FieldSchema{
		Type:    framework.TypeBool,
		Default: false,
		Description: `If true, the Common Name will not be
included in DNS or Email Subject Alternate Names.
Defaults to false (CN is included).`,
		DisplayName: "Exclude Common Name from Subject Alternative Names (SANs)",
	}

	fields["format"] = &framework.FieldSchema{
		Type:    framework.TypeString,
		Default: "pem",
		Description: `Format for returned data. Can be "pem", "der",
or "pem_bundle". If "pem_bundle" any private
key and issuing cert will be appended to the
certificate pem. Defaults to "pem".`,
		DisplayName: "Format",
		AllowedValues: []interface{}{"pem", "der", "pem_bundle"},
		DisplayValue: "pem",
	}

	fields["private_key_format"] = &framework.FieldSchema{
		Type:    framework.TypeString,
		Default: "der",
		Description: `Format for the returned private key.
Generally the default will be controlled by the "format"
parameter as either base64-encoded DER or PEM-encoded DER.
However, this can be set to "pkcs8" to have the returned
private key contain base64-encoded pkcs8 or PEM-encoded
pkcs8 instead. Defaults to "der".`,
		DisplayName: "Private Key Format",
		AllowedValues: []interface{}{"", "der", "pem", "pkcs8"},
		DisplayValue: "",
	}

	fields["ip_sans"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `The requested IP SANs, if any, in a
comma-delimited list`,
		DisplayName: "IP Subject Alternative Names (SANs)",
	}

	fields["uri_sans"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `The requested URI SANs, if any, in a
comma-delimited list.`,
		DisplayName: "URI Subject Alternative Names (SANs)",
	}

	fields["other_sans"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `Requested other SANs, in an array with the format
<oid>;UTF8:<utf8 string value> for each entry.`,
		DisplayName: "Other SANs",
	}

	return fields
}

// addNonCACommonFields adds fields with help text specific to non-CA
// certificate issuing and signing
func addNonCACommonFields(fields map[string]*framework.FieldSchema) map[string]*framework.FieldSchema {
	fields = addIssueAndSignCommonFields(fields)

	fields["role"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The desired role with configuration for this
request`,
	}

	fields["common_name"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested common name; if you want more than
one, specify the alternative names in the
alt_names map. If email protection is enabled
in the role, this may be an email address.`,
		DisplayName: "Common Name",
	}

	fields["alt_names"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested Subject Alternative Names, if any,
in a comma-delimited list. If email protection
is enabled for the role, this may contain
email addresses.`,
		DisplayName: "DNS/Email Subject Alternative Names (SANs)",
	}

	fields["serial_number"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested serial number, if any. If you want
more than one, specify alternative names in
the alt_names map using OID 2.5.4.5.`,
		DisplayName: "Serial Number",
	}

	fields["ttl"] = &framework.FieldSchema{
		Type: framework.TypeDurationSecond,
		Description: `The requested Time To Live for the certificate;
sets the expiration date. If not specified
the role default, backend default, or system
default TTL is used, in that order. Cannot
be larger than the role max TTL.`,
		DisplayName: "TTL",
	}

	return fields
}

// addCACommonFields adds fields with help text specific to CA
// certificate issuing and signing
func addCACommonFields(fields map[string]*framework.FieldSchema) map[string]*framework.FieldSchema {
	fields = addIssueAndSignCommonFields(fields)

	fields["alt_names"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested Subject Alternative Names, if any,
in a comma-delimited list. May contain both
DNS names and email addresses.`,
		DisplayName: "DNS/Email Subject Alternative Names (SANs)",
	}

	fields["common_name"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested common name; if you want more than
one, specify the alternative names in the alt_names
map. If not specified when signing, the common
name will be taken from the CSR; other names
must still be specified in alt_names or ip_sans.`,
		DisplayName: "Common Name",
	}

	fields["ttl"] = &framework.FieldSchema{
		Type: framework.TypeDurationSecond,
		Description: `The requested Time To Live for the certificate;
sets the expiration date. If not specified
the role default, backend default, or system
default TTL is used, in that order. Cannot
be larger than the mount max TTL. Note:
this only has an effect when generating
a CA cert or signing a CA cert, not when
generating a CSR for an intermediate CA.`,
		DisplayName: "TTL",
	}

	fields["ou"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `If set, OU (OrganizationalUnit) will be set to
this value.`,
	DisplayName: "Common Name",

	}

	fields["organization"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `If set, O (Organization) will be set to
this value.`,
	}

	fields["country"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `If set, Country will be set to
this value.`,
	}

	fields["locality"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `If set, Locality will be set to
this value.`,
		DisplayName: "Locality/City",
	}

	fields["province"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `If set, Province will be set to
this value.`,
		DisplayName: "Province/State",
	}

	fields["street_address"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `If set, Street Address will be set to
this value.`,
		DisplayName: "Street Address",
	}

	fields["postal_code"] = &framework.FieldSchema{
		Type: framework.TypeCommaStringSlice,
		Description: `If set, Postal Code will be set to
this value.`,
		DisplayName: "Postal Code",
	}

	fields["serial_number"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `The requested serial number, if any. If you want
more than one, specify alternative names in
the alt_names map using OID 2.5.4.5.`,
		DisplayName: "Serial Number",
	}

	return fields
}

// addCAKeyGenerationFields adds fields with help text specific to CA key
// generation and exporting
func addCAKeyGenerationFields(fields map[string]*framework.FieldSchema) map[string]*framework.FieldSchema {
	fields["exported"] = &framework.FieldSchema{
		Type: framework.TypeString,
		Description: `Must be "internal" or "exported". If set to
"exported", the generated private key will be
returned. This is your *only* chance to retrieve
the private key!`,
	}

	fields["key_bits"] = &framework.FieldSchema{
		Type:    framework.TypeInt,
		Default: 2048,
		Description: `The number of bits to use. You will almost
certainly want to change this if you adjust
the key_type.`,
		DisplayName: "Key Bits",
	}

	fields["key_type"] = &framework.FieldSchema{
		Type:    framework.TypeString,
		Default: "rsa",
		Description: `The type of key to use; defaults to RSA. "rsa"
and "ec" are the only valid values.`,
		DisplayName: "Key Type",
		AllowedValues: []interface{}{"rsa", "ec"},
		DisplayValue: "rsa",
	}

	return fields
}

// addCAIssueFields adds fields common to CA issuing, e.g. when returning
// an actual certificate
func addCAIssueFields(fields map[string]*framework.FieldSchema) map[string]*framework.FieldSchema {
	fields["max_path_length"] = &framework.FieldSchema{
		Type:        framework.TypeInt,
		Default:     -1,
		Description: "The maximum allowable path length",
		DisplayName: "Max Path Length",
		DisplayValue: -1,
	}

	fields["permitted_dns_domains"] = &framework.FieldSchema{
		Type:        framework.TypeCommaStringSlice,
		Description: `Domains for which this certificate is allowed to sign or issue child certificates. If set, all DNS names (subject and alt) on child certs must be exact matches or subsets of the given domains (see https://tools.ietf.org/html/rfc5280#section-4.2.1.10).`,
		DisplayName: "Permitted DNS Domains",
	}

	return fields
}
