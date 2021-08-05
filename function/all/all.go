package all

import (
	// blank import to ensure init() gets called for each package so it can
	// be properly registered.
	_ "transporter/function/gojajs"
	_ "transporter/function/omit"
	_ "transporter/function/opfilter"
	_ "transporter/function/ottojs"
	_ "transporter/function/pick"
	_ "transporter/function/pretty"
	_ "transporter/function/remap"
	_ "transporter/function/rename"
	_ "transporter/function/skip"
)
