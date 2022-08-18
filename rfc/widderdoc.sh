#!/bin/bash

# Requires widdershins
# Requires pandoc

mkdir -p out
rm -r out/*

widdershins \
  -o out/output.md \
  -c \
  -u rfc_template \
  ../spec/openapi.yaml

pandoc \
  -o out/output.docx \
  -t docx \
  --reference-doc rfc_template/custom-reference.docx \
  -f markdown \
  out/output.md 

open out/output.docx
