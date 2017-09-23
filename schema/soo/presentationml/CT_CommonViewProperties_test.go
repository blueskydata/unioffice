// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package presentationml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/soo/presentationml"
)

func TestCT_CommonViewPropertiesConstructor(t *testing.T) {
	v := presentationml.NewCT_CommonViewProperties()
	if v == nil {
		t.Errorf("presentationml.NewCT_CommonViewProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed presentationml.CT_CommonViewProperties should validate: %s", err)
	}
}

func TestCT_CommonViewPropertiesMarshalUnmarshal(t *testing.T) {
	v := presentationml.NewCT_CommonViewProperties()
	buf, _ := xml.Marshal(v)
	v2 := presentationml.NewCT_CommonViewProperties()
	xml.Unmarshal(buf, v2)
}