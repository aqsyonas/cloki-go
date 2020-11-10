package model

func (TableSample) TableName() string {
	return "samples"
}

func (TableSample) TableEngine() string {
	return "MergeTree    PARTITION BY toDate(timestamp_ms / 1000)    ORDER BY (fingerprint, timestamp_ms);"
}

// swagger:model CreateUserStruct
type TableSample struct {
	FingerPrint uint64 `db:"fingerprint" clickhouse:"type:UInt64" json:"fingerprint"`
	// required: true
	TimestampMS int64 `db:"timestamp_ms" clickhouse:"type:Int64" json:"timestamp_ms"`
	//
	Value float64 `db:"value" clickhouse:"type:Float64" json:"value"`
	// example: 10
	// required: true
	String string `db:"string" clickhouse:"type:String" json:"string"`
}
