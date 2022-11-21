package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

type ValidatableAsUser interface {
	GetCreator() string
	GetUserId() string
	GetFirstName() string
	GetLastName() string
	GetCountryCode() string
	GetSubnationalEntity() string
	GetCity() string
	GetBio() string
	GetReferrer() string
	GetAccountAddress() string
}

func ValidateUserBasic(user ValidatableAsUser) (err error) {
	_, err = sdk.AccAddressFromBech32(user.GetCreator())
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	_, err = uuid.Parse(user.GetUserId())
	if err != nil {
		return ErrUserIdNotUUID
	}

	if user.GetFirstName() == "" {
		return ErrFirstNameBlank
	}

	if user.GetLastName() == "" {
		return ErrLastNameBlank
	}

	if user.GetCountryCode() == "" {
		return ErrCountryCodeBlank
	}
	if countriesByCode()[user.GetCountryCode()] == "" {
		return ErrCountryCodeInvalid
	}

	if user.GetSubnationalEntity() == "" {
		return ErrSubnationalEntityBlank
	}

	if user.GetCity() == "" {
		return ErrCityBlank
	}

	if user.GetBio() == "" {
		return ErrBioBlank
	}

	_, err = sdk.AccAddressFromBech32(user.GetReferrer())
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid referrer address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(user.GetAccountAddress())
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user account address (%s)", err)
	}

	return nil
}

func countriesByCode() map[string]string {
	return map[string]string{
		"ABW": "Aruba",
		"AFG": "Afghanistan",
		"AGO": "Angola",
		"AIA": "Anguilla",
		"ALA": "Aland Islands",
		"ALB": "Albania",
		"AND": "Andorra",
		"ANT": "Netherlands Antilles",
		"ARE": "United Arab Emirates",
		"ARG": "Argentina",
		"ARM": "Armenia",
		"ASM": "American Samoa",
		"ATA": "Antarctica",
		"ATF": "French Southern Territories",
		"ATG": "Antigua and Barbuda",
		"AUS": "Australia",
		"AUT": "Austria",
		"AZE": "Azerbaijan",
		"BDI": "Burundi",
		"BEL": "Belgium",
		"BEN": "Benin",
		"BFA": "Burkina Faso",
		"BGD": "Bangladesh",
		"BGR": "Bulgaria",
		"BHR": "Bahrain",
		"BHS": "Bahamas",
		"BIH": "Bosnia and Herzegovina",
		"BLM": "Saint-Barthélemy",
		"BLR": "Belarus",
		"BLZ": "Belize",
		"BMU": "Bermuda",
		"BOL": "Bolivia",
		"BRA": "Brazil",
		"BRB": "Barbados",
		"BRN": "Brunei Darussalam",
		"BTN": "Bhutan",
		"BVT": "Bouvet Island",
		"BWA": "Botswana",
		"CAF": "Central African Republic",
		"CAN": "Canada",
		"CCK": "Cocos (Keeling) Islands",
		"CHE": "Switzerland",
		"CHL": "Chile",
		"CHN": "China",
		"CIV": "Côte d'Ivoire",
		"CMR": "Cameroon",
		"COD": "Congo, (Kinshasa)",
		"COG": "Congo (Brazzaville)",
		"COK": "Cook Islands",
		"COL": "Colombia",
		"COM": "Comoros",
		"CPV": "Cape Verde",
		"CRI": "Costa Rica",
		"CUB": "Cuba",
		"CXR": "Christmas Island",
		"CYM": "Cayman Islands",
		"CYP": "Cyprus",
		"CZE": "Czech Republic",
		"DEU": "Germany",
		"DJI": "Djibouti",
		"DMA": "Dominica",
		"DNK": "Denmark",
		"DOM": "Dominican Republic",
		"DZA": "Algeria",
		"ECU": "Ecuador",
		"EGY": "Egypt",
		"ERI": "Eritrea",
		"ESH": "Western Sahara",
		"ESP": "Spain",
		"EST": "Estonia",
		"ETH": "Ethiopia",
		"FIN": "Finland",
		"FJI": "Fiji",
		"FLK": "Falkland Islands (Malvinas)",
		"FRA": "France",
		"FRO": "Faroe Islands",
		"FSM": "Micronesia, Federated States of",
		"GAB": "Gabon",
		"GBR": "United Kingdom",
		"GEO": "Georgia",
		"GGY": "Guernsey",
		"GHA": "Ghana",
		"GIB": "Gibraltar",
		"GIN": "Guinea",
		"GLP": "Guadeloupe",
		"GMB": "Gambia",
		"GNB": "Guinea-Bissau",
		"GNQ": "Equatorial Guinea",
		"GRC": "Greece",
		"GRD": "Grenada",
		"GRL": "Greenland",
		"GTM": "Guatemala",
		"GUF": "French Guiana",
		"GUM": "Guam",
		"GUY": "Guyana",
		"HKG": "Hong Kong, SAR China",
		"HMD": "Heard and Mcdonald Islands",
		"HND": "Honduras",
		"HRV": "Croatia",
		"HTI": "Haiti",
		"HUN": "Hungary",
		"IDN": "Indonesia",
		"IMN": "Isle of Man",
		"IND": "India",
		"IOT": "British Indian Ocean Territory",
		"IRL": "Ireland",
		"IRN": "Iran, Islamic Republic of",
		"IRQ": "Iraq",
		"ISL": "Iceland",
		"ISR": "Israel",
		"ITA": "Italy",
		"JAM": "Jamaica",
		"JEY": "Jersey",
		"JOR": "Jordan",
		"JPN": "Japan",
		"KAZ": "Kazakhstan",
		"KEN": "Kenya",
		"KGZ": "Kyrgyzstan",
		"KHM": "Cambodia",
		"KIR": "Kiribati",
		"KNA": "Saint Kitts and Nevis",
		"KOR": "Korea (South)",
		"KWT": "Kuwait",
		"LAO": "Lao PDR",
		"LBN": "Lebanon",
		"LBR": "Liberia",
		"LBY": "Libya",
		"LCA": "Saint Lucia",
		"LIE": "Liechtenstein",
		"LKA": "Sri Lanka",
		"LSO": "Lesotho",
		"LTU": "Lithuania",
		"LUX": "Luxembourg",
		"LVA": "Latvia",
		"MAC": "Macao, SAR China",
		"MAF": "Saint-Martin (French part)",
		"MAR": "Morocco",
		"MCO": "Monaco",
		"MDA": "Moldova",
		"MDG": "Madagascar",
		"MDV": "Maldives",
		"MEX": "Mexico",
		"MHL": "Marshall Islands",
		"MKD": "Macedonia, Republic of",
		"MLI": "Mali",
		"MLT": "Malta",
		"MMR": "Myanmar",
		"MNE": "Montenegro",
		"MNG": "Mongolia",
		"MNP": "Northern Mariana Islands",
		"MOZ": "Mozambique",
		"MRT": "Mauritania",
		"MSR": "Montserrat",
		"MTQ": "Martinique",
		"MUS": "Mauritius",
		"MWI": "Malawi",
		"MYS": "Malaysia",
		"MYT": "Mayotte",
		"NAM": "Namibia",
		"NCL": "New Caledonia",
		"NER": "Niger",
		"NFK": "Norfolk Island",
		"NGA": "Nigeria",
		"NIC": "Nicaragua",
		"NIU": "Niue",
		"NLD": "Netherlands",
		"NOR": "Norway",
		"NPL": "Nepal",
		"NRU": "Nauru",
		"NZL": "New Zealand",
		"OMN": "Oman",
		"PAK": "Pakistan",
		"PAN": "Panama",
		"PCN": "Pitcairn",
		"PER": "Peru",
		"PHL": "Philippines",
		"PLW": "Palau",
		"PNG": "Papua New Guinea",
		"POL": "Poland",
		"PRI": "Puerto Rico",
		"PRK": "Korea (North)",
		"PRT": "Portugal",
		"PRY": "Paraguay",
		"PSE": "Palestinian Territory",
		"PYF": "French Polynesia",
		"QAT": "Qatar",
		"REU": "Réunion",
		"ROU": "Romania",
		"RUS": "Russian Federation",
		"RWA": "Rwanda",
		"SAU": "Saudi Arabia",
		"SDN": "Sudan",
		"SEN": "Senegal",
		"SGP": "Singapore",
		"SGS": "South Georgia and the South Sandwich Islands",
		"SHN": "Saint Helena",
		"SJM": "Svalbard and Jan Mayen Islands",
		"SLB": "Solomon Islands",
		"SLE": "Sierra Leone",
		"SLV": "El Salvador",
		"SMR": "San Marino",
		"SOM": "Somalia",
		"SPM": "Saint Pierre and Miquelon",
		"SRB": "Serbia",
		"SSD": "South Sudan",
		"STP": "Sao Tome and Principe",
		"SUR": "Suriname",
		"SVK": "Slovakia",
		"SVN": "Slovenia",
		"SWE": "Sweden",
		"SWZ": "Swaziland",
		"SYC": "Seychelles",
		"SYR": "Syrian Arab Republic (Syria)",
		"TCA": "Turks and Caicos Islands",
		"TCD": "Chad",
		"TGO": "Togo",
		"THA": "Thailand",
		"TJK": "Tajikistan",
		"TKL": "Tokelau",
		"TKM": "Turkmenistan",
		"TLS": "Timor-Leste",
		"TON": "Tonga",
		"TTO": "Trinidad and Tobago",
		"TUN": "Tunisia",
		"TUR": "Turkey",
		"TUV": "Tuvalu",
		"TWN": "Taiwan, Republic of China",
		"TZA": "Tanzania, United Republic of",
		"UGA": "Uganda",
		"UKR": "Ukraine",
		"UMI": "US Minor Outlying Islands",
		"URY": "Uruguay",
		"USA": "United States of America",
		"UZB": "Uzbekistan",
		"VAT": "Holy See (Vatican City State)",
		"VCT": "Saint Vincent and Grenadines",
		"VEN": "Venezuela (Bolivarian Republic)",
		"VGB": "British Virgin Islands",
		"VIR": "Virgin Islands, US",
		"VNM": "Viet Nam",
		"VUT": "Vanuatu",
		"WLF": "Wallis and Futuna Islands",
		"WSM": "Samoa",
		"YEM": "Yemen",
		"ZAF": "South Africa",
		"ZMB": "Zambia",
		"ZWE": "Zimbabwe",
	}
}
