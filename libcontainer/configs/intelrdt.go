package configs

type IntelRdt struct {
	// The identity for RDT Class of Service
	ClosID string `json:"closID,omitempty"`
	// The schema for L3 cache id and capacity bitmask (CBM)
	// Format: "L3:<cache_id0>=<cbm0>;<cache_id1>=<cbm1>;..."
	L3CacheSchema string `json:"l3_cache_schema,omitempty"`

	// The schema of memory bandwidth percentage per L3 cache id
	// Format: "MB:<cache_id0>=bandwidth0;<cache_id1>=bandwidth1;..."
	MemBwSchema string `json:"memBwSchema,omitempty"`
}
