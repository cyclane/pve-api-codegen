# Proxmox VE API codegen

This is a tool to generate API libraries from the Proxmox VE api-docs
`apidata.js` file.

## Why JavaScript

The code generator doesn't more performance than JS can provide, and we need to use to
parse the `apidata.js` file anyway (which currently can be easily converted to JSON,
 but this could change in the future considering it is provided as a `.js` file and not
 `.json`).
