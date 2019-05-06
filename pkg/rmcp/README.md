# RMCP

This package implements the subset of the [Alert Standard Format](https://www.dmtf.org/standards/asf) [Specification v2.0](https://www.dmtf.org/sites/default/files/standards/documents/DSP0136.pdf) required for IPMI v2.0.

Note that although IPMIv2 refers to "RMCP+", this is not a modification to RMCP: it merely leverages the `0x7` IPMI message class in the RMCP header to allow a non-ASF-specified payload.
An RMCPv1 implementation is seemingly sufficient for use with IPMIv2.
