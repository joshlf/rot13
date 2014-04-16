// Copyright 2014 The Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os/exec"
	"reflect"
	"testing"
)

const (
	input = `bTYOEYc5I1LjFXMqmCEtZr7poi1Css5jYr/452QyhSvZVmmkuZnnkvYNZV5Na2v44r0flh5nwPVx
ciw+pjhouS9LasXKiiq2H5/VIHlMVTauO3tpdLpEuzxEmI3KDOzATHT3EINnh3hJQIXJ9p3zwEjY
YwMt643HrXq1rWP5ikLC9sl0gPJrExP+CLZ97H4pEglo0U0gZ+NRNDTkysstHIwu8n/dLPBAPVoq
JrglOI9/+ugLEfl+ZxaothgC9mQUy3H9Ln+qlZzF1VGpZJiphu/9bc9VOT+Q5SsuzRKaO9ShRHfH
Wv1zr0f9NLcs3ZYXiAng9kakvU9ANTdn/LWIPg34GmD7h9MdCh6Co1iZIg+GlnFdaDyrBXLG63uS
fQo5gvHTcRIVxoy1ZNHVV+I+DfjF5SeOsoOIe/tzmna1PWLJkckxpi9AZXZI5h8eJ1nkCtvD9zMS
UcXYwmfuqY0zPOrZxEdwTsyK8+zBPYNS9uTAenXeLLAQMyn3TvnzOfjGnCa9ANSDRwHzue1m9Tix
AnvT3rKWALR0ld8rq810Rf6mBDc1JH14xQR+iCuIqkVsA6ttqtZZkkeYKWUu5bKEYHhWwEXGEsgm
wZl1RcVEzYrTgKvzPHnJ+7b7y+mgipI8PXg8rjhVDOzhxsNVGX2yazvTCRvzW7S9GfQG59f7i/YF
yrx+Sxj1UT4QfMkmydFqA4ZQF23vXBadKYqalY5tZPfvLnJexj4E0Eew9SZkDBn5uAeVU1j4Y4fn
L0vRMibOYx81RW88yGR1e9Jw/ZXU4WyQZLkPF7Lq5xPr1UR1H7ZAg0s5ddd5hEUoU+lnP8sWiD93
gMDn6jA3lsNhoZgctzzCF/EksA6FVC75T8v7LbesuRoCs0fDZeiGD7ScgzlPtAStTUduE8k6bk/H
UFR1zGcu7vyQTXwr2qvnnrUPZaQ8Bafw84v+UdEIgrXMKJRk9pMjLdVCmtTaohKRvqJcYNioFRou
xAcBeYsRPMyYYxt33MxjKUXpZNkF83PR/VyNzZZ3QPwqJ6MBndhUpyz0ETAEXHaNtguy2g/Ueg5p
hcO7ykFF9q3IIyliBFwao1/r2IgGZR0l5VfYQKnf5TYywMO8rkO6nMGM235NEq/nMSbb675xDnwx
yzwY1lMOwIFX+6mq8ux6ITqmpPIUKXyk3H4ksiT5V3bDz6Sh9njkgMVk1GjQRDjBYwvc8ygpt7kA
DkI1g/otbwb2+bWCMdbwizN6CQB7RyY0onz2zGfgCHHkujFbp7JSRNkxvhT3hwqdwKXQtyUUuIAL
K/DlKlmJev58g30vNxRgxaFDmYBdR3WkOO1H8H07IP4qw9SRRtd28ojdz46Oq5QyqpNtzeMidQ==`
	output = `oGLBRLp5V1YwSKZdzPRgMe7cbv1Pff5wLe/452DluFiMIzzxhMaaxiLAMI5An2i44e0syu5ajCIk
pvj+cwubhF9YnfKXvvd2U5/IVUyZIGnhB3gcqYcRhmkRzV3XQBmNGUG3RVAau3uWDVKW9c3mjRwL
LjZg643UeKd1eJC5vxYP9fy0tCWeRkC+PYM97U4cRtyb0H0tM+AEAQGxlffgUVjh8a/qYCONCIbd
WetyBV9/+htYRsy+MknbgutP9zDHl3U9Ya+dyMmS1ITcMWvcuh/9op9IBG+D5FfhmEXnB9FuEUsU
Ji1me0s9AYpf3MLKvNat9xnxiH9NAGqa/YJVCt34TzQ7u9ZqPu6Pb1vMVt+TyaSqnQleOKYT63hF
sDb5tiUGpEVIkbl1MAUII+V+QswS5FrBfbBVr/gmzan1CJYWxpxkcv9NMKMV5u8rW1axPgiQ9mZF
HpKLjzshdL0mCBeMkRqjGflX8+mOCLAF9hGNraKrYYNDZla3GiamBswTaPn9NAFQEjUmhr1z9Gvk
NaiG3eXJNYE0yq8ed810Es6zOQp1WU14kDE+vPhVdxIfN6ggdgMMxxrLXJHh5oXRLUuJjRKTRftz
jMy1EpIRmLeGtXimCUaW+7o7l+ztvcV8CKt8ewuIQBmukfAITK2lnmiGPEimJ7F9TsDT59s7v/LS
lek+Fkw1HG4DsZxzlqSdN4MDS23iKOnqXLdnyL5gMCsiYaWrkw4R0Rrj9FMxQOa5hNrIH1w4L4sa
Y0iEZvoBLk81EJ88lTE1r9Wj/MKH4JlDMYxCS7Yd5kCe1HE1U7MNt0f5qqq5uRHbH+yaC8fJvQ93
tZQa6wN3yfAubMtpgmmPS/RxfN6SIP75G8i7YorfhEbPf0sQMrvTQ7FptmyCgNFgGHqhR8x6ox/U
HSE1mTph7ilDGKje2diaaeHCMnD8Onsj84i+HqRVteKZXWEx9cZwYqIPzgGnbuXEidWpLAvbSEbh
kNpOrLfECZlLLkg33ZkwXHKcMAxS83CE/IlAmMM3DCjdW6ZOaquHclm0RGNRKUnAgthl2t/Hrt5c
upB7lxSS9d3VVlyvOSjnb1/e2VtTME0y5IsLDXas5GLljZB8exB6aZTZ235ARd/aZFoo675kQajk
lmjL1yZBjVSK+6zd8hk6VGdzcCVHXKlx3U4xfvG5I3oQm6Fu9awxtZIx1TwDEQwOLjip8ltcg7xN
QxV1t/bgojo2+oJPZqojvmA6PDO7ElL0bam2mTstPUUxhwSoc7WFEAxkiuG3ujdqjXKDglHHhVNY
X/QyXyzWri58t30iAkEtknSQzLOqE3JxBB1U8U07VC4dj9FEEgq28bwqm46Bd5DldcAgmrZvqD==`
)

func TestRot13(t *testing.T) {
	c := exec.Command("go", "run", "rot13.go")

	ibuf := []byte(input)
	obuf := []byte(output)
	c.Stdin = bytes.NewBuffer(ibuf)
	stdout := &bytes.Buffer{}
	c.Stdout = stdout

	if err := c.Run(); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	result := stdout.Bytes()

	if !reflect.DeepEqual(obuf, result) {
		t.Errorf("Expected \"%s\"; got \"%s\"", output, string(result))
	}
}
